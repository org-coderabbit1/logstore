package status

import (
	"context"
	"testing"
	"time"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/manifests"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
)

func TestRefreshSuccess(t *testing.T) {
	now := time.Now()
	stack := &logstorev1.LogstoreStack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	req := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	componentPods := map[string]*corev1.PodList{
		manifests.LabelCompactorComponent:     createPodList(manifests.LabelCompactorComponent, corev1.PodRunning),
		manifests.LabelDistributorComponent:   createPodList(manifests.LabelDistributorComponent, corev1.PodRunning),
		manifests.LabelIngesterComponent:      createPodList(manifests.LabelIngesterComponent, corev1.PodRunning),
		manifests.LabelQuerierComponent:       createPodList(manifests.LabelQuerierComponent, corev1.PodRunning),
		manifests.LabelQueryFrontendComponent: createPodList(manifests.LabelQueryFrontendComponent, corev1.PodRunning),
		manifests.LabelIndexGatewayComponent:  createPodList(manifests.LabelIndexGatewayComponent, corev1.PodRunning),
		manifests.LabelRulerComponent:         createPodList(manifests.LabelRulerComponent, corev1.PodRunning),
		manifests.LabelGatewayComponent:       createPodList(manifests.LabelGatewayComponent, corev1.PodRunning),
	}

	wantStatus := logstorev1.LogstoreStackStatus{
		Components: logstorev1.LogstoreStackComponentStatus{
			Compactor:     map[corev1.PodPhase][]string{corev1.PodRunning: {"compactor-pod-0"}},
			Distributor:   map[corev1.PodPhase][]string{corev1.PodRunning: {"distributor-pod-0"}},
			IndexGateway:  map[corev1.PodPhase][]string{corev1.PodRunning: {"index-gateway-pod-0"}},
			Ingester:      map[corev1.PodPhase][]string{corev1.PodRunning: {"ingester-pod-0"}},
			Querier:       map[corev1.PodPhase][]string{corev1.PodRunning: {"querier-pod-0"}},
			QueryFrontend: map[corev1.PodPhase][]string{corev1.PodRunning: {"query-frontend-pod-0"}},
			Gateway:       map[corev1.PodPhase][]string{corev1.PodRunning: {"logstorestack-gateway-pod-0"}},
			Ruler:         map[corev1.PodPhase][]string{corev1.PodRunning: {"ruler-pod-0"}},
		},
		Storage: logstorev1.LogstoreStackStorageStatus{},
		Conditions: []metav1.Condition{
			{
				Type:               string(logstorev1.ConditionReady),
				Reason:             string(logstorev1.ReasonReadyComponents),
				Message:            messageReady,
				Status:             metav1.ConditionTrue,
				LastTransitionTime: metav1.NewTime(now),
			},
		},
	}

	k, sw := setupListClient(t, stack, componentPods)

	err := Refresh(context.Background(), k, req, now)

	require.NoError(t, err)
	require.Equal(t, 1, k.GetCallCount())
	require.Equal(t, 8, k.ListCallCount())

	require.Equal(t, 1, sw.UpdateCallCount())
	_, updated, _ := sw.UpdateArgsForCall(0)
	updatedStack, ok := updated.(*logstorev1.LogstoreStack)
	if !ok {
		t.Fatalf("not a LogstoreStack: %T", updatedStack)
	}

	require.Equal(t, wantStatus, updatedStack.Status)
}
