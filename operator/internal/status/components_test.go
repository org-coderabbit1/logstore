package status

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s/k8sfakes"
	"example.com/acme/logstore/operator/internal/manifests"
)

func createPodList(baseName string, ready bool, phases ...corev1.PodPhase) *corev1.PodList {
	items := []corev1.Pod{}
	for i, p := range phases {
		items = append(items, corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("%s-pod-%d", baseName, i),
			},
			Status: corev1.PodStatus{
				Phase: p,
				ContainerStatuses: []corev1.ContainerStatus{
					{
						Ready: ready,
					},
				},
			},
		})
	}

	return &corev1.PodList{
		Items: items,
	}
}

func setupListClient(t *testing.T, stack *logstorev1.LogstoreStack, componentPods map[string]*corev1.PodList) (*k8sfakes.FakeClient, *k8sfakes.FakeStatusWriter) {
	k, sw := setupFakesNoError(t, stack)
	k.ListStub = func(_ context.Context, list client.ObjectList, options ...client.ListOption) error {
		componentLabel := ""
		for _, o := range options {
			if m, ok := o.(client.MatchingLabels); ok {
				componentLabel = m["app.kubernetes.io/component"]
			}
		}

		if componentLabel == "" {
			t.Fatalf("no component label on list call: %s", options)
		}

		podList, ok := componentPods[componentLabel]
		if !ok {
			t.Fatalf("no pods found for label: %s", componentLabel)
		}

		k.SetClientObjectList(list, podList)
		return nil
	}

	return k, sw
}

func TestGenerateComponentStatus(t *testing.T) {
	empty := logstorev1.PodStatusMap{
		logstorev1.PodFailed:  []string{},
		logstorev1.PodPending: []string{},
		logstorev1.PodRunning: []string{},
		logstorev1.PodReady:   []string{},
	}

	tt := []struct {
		desc                string
		componentPods       map[string]*corev1.PodList
		wantComponentStatus *logstorev1.LogstoreStackComponentStatus
	}{
		{
			desc: "no pods",
			componentPods: map[string]*corev1.PodList{
				manifests.LabelCompactorComponent:     {},
				manifests.LabelDistributorComponent:   {},
				manifests.LabelIngesterComponent:      {},
				manifests.LabelQuerierComponent:       {},
				manifests.LabelQueryFrontendComponent: {},
				manifests.LabelIndexGatewayComponent:  {},
				manifests.LabelRulerComponent:         {},
				manifests.LabelGatewayComponent:       {},
			},
			wantComponentStatus: &logstorev1.LogstoreStackComponentStatus{
				Compactor:     empty,
				Distributor:   empty,
				IndexGateway:  empty,
				Ingester:      empty,
				Querier:       empty,
				QueryFrontend: empty,
				Gateway:       empty,
				Ruler:         empty,
			},
		},
		{
			desc: "all one pod running",
			componentPods: map[string]*corev1.PodList{
				manifests.LabelCompactorComponent:     createPodList(manifests.LabelCompactorComponent, false, corev1.PodRunning),
				manifests.LabelDistributorComponent:   createPodList(manifests.LabelDistributorComponent, false, corev1.PodRunning),
				manifests.LabelIngesterComponent:      createPodList(manifests.LabelIngesterComponent, false, corev1.PodRunning),
				manifests.LabelQuerierComponent:       createPodList(manifests.LabelQuerierComponent, false, corev1.PodRunning),
				manifests.LabelQueryFrontendComponent: createPodList(manifests.LabelQueryFrontendComponent, false, corev1.PodRunning),
				manifests.LabelIndexGatewayComponent:  createPodList(manifests.LabelIndexGatewayComponent, false, corev1.PodRunning),
				manifests.LabelRulerComponent:         createPodList(manifests.LabelRulerComponent, false, corev1.PodRunning),
				manifests.LabelGatewayComponent:       createPodList(manifests.LabelGatewayComponent, false, corev1.PodRunning),
			},
			wantComponentStatus: &logstorev1.LogstoreStackComponentStatus{
				Compactor:     logstorev1.PodStatusMap{logstorev1.PodRunning: {"compactor-pod-0"}},
				Distributor:   logstorev1.PodStatusMap{logstorev1.PodRunning: {"distributor-pod-0"}},
				IndexGateway:  logstorev1.PodStatusMap{logstorev1.PodRunning: {"index-gateway-pod-0"}},
				Ingester:      logstorev1.PodStatusMap{logstorev1.PodRunning: {"ingester-pod-0"}},
				Querier:       logstorev1.PodStatusMap{logstorev1.PodRunning: {"querier-pod-0"}},
				QueryFrontend: logstorev1.PodStatusMap{logstorev1.PodRunning: {"query-frontend-pod-0"}},
				Gateway:       logstorev1.PodStatusMap{logstorev1.PodRunning: {"logstorestack-gateway-pod-0"}},
				Ruler:         logstorev1.PodStatusMap{logstorev1.PodRunning: {"ruler-pod-0"}},
			},
		},
		{
			desc: "all pods without ruler",
			componentPods: map[string]*corev1.PodList{
				manifests.LabelCompactorComponent:     createPodList(manifests.LabelCompactorComponent, false, corev1.PodRunning),
				manifests.LabelDistributorComponent:   createPodList(manifests.LabelDistributorComponent, false, corev1.PodRunning),
				manifests.LabelIngesterComponent:      createPodList(manifests.LabelIngesterComponent, false, corev1.PodRunning),
				manifests.LabelQuerierComponent:       createPodList(manifests.LabelQuerierComponent, false, corev1.PodRunning),
				manifests.LabelQueryFrontendComponent: createPodList(manifests.LabelQueryFrontendComponent, false, corev1.PodRunning),
				manifests.LabelIndexGatewayComponent:  createPodList(manifests.LabelIndexGatewayComponent, false, corev1.PodRunning),
				manifests.LabelRulerComponent:         {},
				manifests.LabelGatewayComponent:       createPodList(manifests.LabelGatewayComponent, false, corev1.PodRunning),
			},
			wantComponentStatus: &logstorev1.LogstoreStackComponentStatus{
				Compactor:     logstorev1.PodStatusMap{logstorev1.PodRunning: {"compactor-pod-0"}},
				Distributor:   logstorev1.PodStatusMap{logstorev1.PodRunning: {"distributor-pod-0"}},
				IndexGateway:  logstorev1.PodStatusMap{logstorev1.PodRunning: {"index-gateway-pod-0"}},
				Ingester:      logstorev1.PodStatusMap{logstorev1.PodRunning: {"ingester-pod-0"}},
				Querier:       logstorev1.PodStatusMap{logstorev1.PodRunning: {"querier-pod-0"}},
				QueryFrontend: logstorev1.PodStatusMap{logstorev1.PodRunning: {"query-frontend-pod-0"}},
				Gateway:       logstorev1.PodStatusMap{logstorev1.PodRunning: {"logstorestack-gateway-pod-0"}},
				Ruler:         empty,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			stack := &logstorev1.LogstoreStack{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-stack",
					Namespace: "some-ns",
				},
			}

			k, _ := setupListClient(t, stack, tc.componentPods)

			componentStatus, err := generateComponentStatus(context.Background(), k, stack)
			require.NoError(t, err)
			require.Equal(t, tc.wantComponentStatus, componentStatus)

			// one list call for each component
			require.Equal(t, 8, k.ListCallCount())
		})
	}
}
