package manifests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
)

func TestRulesConfigMap_ReturnsDataEntriesPerRule(t *testing.T) {
	cm_shards, err := RulesConfigMapShards(testOptions())
	require.NoError(t, err)
	require.NotNil(t, cm_shards)

	cm := cm_shards[0]

	require.Len(t, cm.Data, 4)
	require.Contains(t, cm.Data, "dev-alerting-rules-alerts1.yaml")
	require.Contains(t, cm.Data, "dev-recording-rules-recs1.yaml")
	require.Contains(t, cm.Data, "prod-alerting-rules-alerts2.yaml")
	require.Contains(t, cm.Data, "prod-recording-rules-recs2.yaml")
}

func TestRulesConfigMap_ReturnsTenantMapPerRule(t *testing.T) {
	opts := testOptions()
	cm_shards, err := RulesConfigMapShards(opts)
	require.NoError(t, err)
	require.NotNil(t, cm_shards)

	cm := cm_shards[0]
	require.Len(t, cm.Data, 4)
	require.Contains(t, opts.Tenants.Configs["tenant-a"].RuleFiles, "-rules-0___tenant-a___dev-alerting-rules-alerts1.yaml")
	require.Contains(t, opts.Tenants.Configs["tenant-a"].RuleFiles, "-rules-0___tenant-a___prod-alerting-rules-alerts2.yaml")
	require.Contains(t, opts.Tenants.Configs["tenant-b"].RuleFiles, "-rules-0___tenant-b___dev-recording-rules-recs1.yaml")
	require.Contains(t, opts.Tenants.Configs["tenant-b"].RuleFiles, "-rules-0___tenant-b___prod-recording-rules-recs2.yaml")
}

func TestRulesConfigMapSharding(t *testing.T) {
	cm_shards, err := RulesConfigMapShards(testOptions_withSharding())
	require.NoError(t, err)
	require.NotNil(t, cm_shards)
	require.Len(t, cm_shards, 2)
}

func testOptions() *Options {
	return &Options{
		Tenants: Tenants{
			Configs: map[string]TenantConfig{
				"tenant-a": {},
				"tenant-b": {},
			},
		},
		AlertingRules: []logstorev1.AlertingRule{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rules",
					Namespace: "dev",
					UID:       types.UID("alerts1"),
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "tenant-a",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "rule-a",
						},
						{
							Name: "rule-b",
						},
					},
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rules",
					Namespace: "prod",
					UID:       types.UID("alerts2"),
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "tenant-a",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "rule-c",
						},
						{
							Name: "rule-d",
						},
					},
				},
			},
		},
		RecordingRules: []logstorev1.RecordingRule{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rules",
					Namespace: "dev",
					UID:       types.UID("recs1"),
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "tenant-b",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "rule-a",
						},
						{
							Name: "rule-b",
						},
					},
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rules",
					Namespace: "prod",
					UID:       types.UID("recs2"),
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "tenant-b",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "rule-c",
						},
						{
							Name: "rule-d",
						},
					},
				},
			},
		},
	}
}

func testOptions_withSharding() *Options {
	// Generate a list of dummy rules to create a large amount of data
	// that should result in sharding the rules ConfigMap
	// In this case, each Alerting rule amounts to 598 bytes of ConfigMap data
	// and 2000 of them will be split into 2 shards
	var alertingRules []logstorev1.AlertingRule

	for i := range 2000 {
		alertingRules = append(alertingRules, logstorev1.AlertingRule{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "alerting-rules",
				Namespace: "dev",
				UID:       types.UID(fmt.Sprintf("alerts%d", i)),
			},
			Spec: logstorev1.AlertingRuleSpec{
				TenantID: "tenant-a",
				Groups: []*logstorev1.AlertingRuleGroup{
					{
						Name: "rule-a",
						Rules: []*logstorev1.AlertingRuleGroupSpec{
							{
								Alert: fmt.Sprintf("test%d", i),
								Expr: `|
									sum(rate({kubernetes_namespace_name="openshift-operators-redhat", kubernetes_pod_name=~"logstore-operator-controller-manager.*"} |= "error" [1m])) by (job)
									/
									sum(rate({kubernetes_namespace_name="openshift-operators-redhat", kubernetes_pod_name=~"logstore-operator-controller-manager.*"}[1m])) by (job)
									> 0.01
									`,
							},
						},
					},
				},
			},
		})
	}

	return &Options{
		Name:      "sharding-test",
		Namespace: "namespace",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Ruler: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
		Tenants: Tenants{
			Configs: map[string]TenantConfig{
				"tenant-a": {},
				"tenant-b": {},
			},
		},
		AlertingRules: alertingRules,
	}
}
