package metrics

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/ViaQ/logerr/v2/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s/k8sfakes"
)

func TestRegisterLogstoreStackMetrics(t *testing.T) {
	logger := log.NewLogger("test", log.WithOutput(io.Discard))
	client := &k8sfakes.FakeClient{}
	registry := prometheus.NewPedanticRegistry()

	err := RegisterLogstoreStackCollector(logger, client, registry)
	require.NoError(t, err)
}

func TestLogstoreStackMetricsCollect(t *testing.T) {
	tt := []struct {
		desc        string
		k8sError    error
		stacks      *logstorev1.LogstoreStackList
		wantMetrics string
	}{
		{
			desc:        "no stacks",
			k8sError:    nil,
			stacks:      &logstorev1.LogstoreStackList{},
			wantMetrics: "",
		},
		{
			desc:     "one demo",
			k8sError: nil,
			stacks: &logstorev1.LogstoreStackList{
				Items: []logstorev1.LogstoreStack{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-stack",
							Namespace: "test-namespace",
						},
						Spec: logstorev1.LogstoreStackSpec{
							Size: logstorev1.SizeOneXDemo,
						},
					},
				},
			},
			wantMetrics: `# HELP logstorestack_info Information about deployed LogstoreStack instances. Value is always 1.
# TYPE logstorestack_info gauge
logstorestack_info{size="1x.demo",stack_name="test-stack",stack_namespace="test-namespace"} 1
`,
		},
		{
			desc:     "one small with warning",
			k8sError: nil,
			stacks: &logstorev1.LogstoreStackList{
				Items: []logstorev1.LogstoreStack{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-stack",
							Namespace: "test-namespace",
						},
						Spec: logstorev1.LogstoreStackSpec{
							Size: logstorev1.SizeOneXSmall,
						},
						Status: logstorev1.LogstoreStackStatus{
							Conditions: []metav1.Condition{
								{
									Type:   string(logstorev1.ConditionWarning),
									Status: metav1.ConditionTrue,
									Reason: string(logstorev1.ReasonStorageNeedsSchemaUpdate),
								},
							},
						},
					},
				},
			},
			wantMetrics: `# HELP logstorestack_info Information about deployed LogstoreStack instances. Value is always 1.
# TYPE logstorestack_info gauge
logstorestack_info{size="1x.small",stack_name="test-stack",stack_namespace="test-namespace"} 1
# HELP logstorestack_status_condition Counts the current status conditions of the LogstoreStack.
# TYPE logstorestack_status_condition gauge
logstorestack_status_condition{condition="Warning",reason="StorageNeedsSchemaUpdate",size="1x.small",stack_name="test-stack",stack_namespace="test-namespace",status="false"} 0
logstorestack_status_condition{condition="Warning",reason="StorageNeedsSchemaUpdate",size="1x.small",stack_name="test-stack",stack_namespace="test-namespace",status="true"} 1
`,
		},
		{
			desc:     "multiple conditions, inactive warning",
			k8sError: nil,
			stacks: &logstorev1.LogstoreStackList{
				Items: []logstorev1.LogstoreStack{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-stack",
							Namespace: "test-namespace",
						},
						Spec: logstorev1.LogstoreStackSpec{
							Size: logstorev1.SizeOneXSmall,
						},
						Status: logstorev1.LogstoreStackStatus{
							Conditions: []metav1.Condition{
								{
									Type:   string(logstorev1.ConditionReady),
									Status: metav1.ConditionTrue,
									Reason: string(logstorev1.ReasonReadyComponents),
								},
								{
									Type:   string(logstorev1.ConditionPending),
									Status: metav1.ConditionFalse,
									Reason: string(logstorev1.ReasonPendingComponents),
								},
								{
									Type:   string(logstorev1.ConditionWarning),
									Status: metav1.ConditionFalse,
									Reason: string(logstorev1.ReasonStorageNeedsSchemaUpdate),
								},
							},
						},
					},
				},
			},
			wantMetrics: `# HELP logstorestack_info Information about deployed LogstoreStack instances. Value is always 1.
# TYPE logstorestack_info gauge
logstorestack_info{size="1x.small",stack_name="test-stack",stack_namespace="test-namespace"} 1
# HELP logstorestack_status_condition Counts the current status conditions of the LogstoreStack.
# TYPE logstorestack_status_condition gauge
logstorestack_status_condition{condition="Pending",reason="PendingComponents",size="1x.small",stack_name="test-stack",stack_namespace="test-namespace",status="false"} 1
logstorestack_status_condition{condition="Pending",reason="PendingComponents",size="1x.small",stack_name="test-stack",stack_namespace="test-namespace",status="true"} 0
logstorestack_status_condition{condition="Ready",reason="ReadyComponents",size="1x.small",stack_name="test-stack",stack_namespace="test-namespace",status="false"} 0
logstorestack_status_condition{condition="Ready",reason="ReadyComponents",size="1x.small",stack_name="test-stack",stack_namespace="test-namespace",status="true"} 1
logstorestack_status_condition{condition="Warning",reason="StorageNeedsSchemaUpdate",size="1x.small",stack_name="test-stack",stack_namespace="test-namespace",status="false"} 1
logstorestack_status_condition{condition="Warning",reason="StorageNeedsSchemaUpdate",size="1x.small",stack_name="test-stack",stack_namespace="test-namespace",status="true"} 0
`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			logger := log.NewLogger("test", log.WithOutput(io.Discard))
			k := &k8sfakes.FakeClient{}
			k.ListStub = func(_ context.Context, list client.ObjectList, _ ...client.ListOption) error {
				if tc.k8sError != nil {
					return tc.k8sError
				}

				k.SetClientObjectList(list, tc.stacks)
				return nil
			}

			expected := strings.NewReader(tc.wantMetrics)

			c := &logstoreStackCollector{
				log:       logger,
				k8sClient: k,
			}

			if err := testutil.CollectAndCompare(c, expected); err != nil {
				t.Error(err)
			}
		})
	}
}
