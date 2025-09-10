package status

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/manifests"
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
		manifests.LabelCompactorComponent:     createPodList(manifests.LabelCompactorComponent, true, corev1.PodRunning),
		manifests.LabelDistributorComponent:   createPodList(manifests.LabelDistributorComponent, true, corev1.PodRunning),
		manifests.LabelIngesterComponent:      createPodList(manifests.LabelIngesterComponent, true, corev1.PodRunning),
		manifests.LabelQuerierComponent:       createPodList(manifests.LabelQuerierComponent, true, corev1.PodRunning),
		manifests.LabelQueryFrontendComponent: createPodList(manifests.LabelQueryFrontendComponent, true, corev1.PodRunning),
		manifests.LabelIndexGatewayComponent:  createPodList(manifests.LabelIndexGatewayComponent, true, corev1.PodRunning),
		manifests.LabelRulerComponent:         createPodList(manifests.LabelRulerComponent, true, corev1.PodRunning),
		manifests.LabelGatewayComponent:       createPodList(manifests.LabelGatewayComponent, true, corev1.PodRunning),
	}

	wantStatus := logstorev1.LogstoreStackStatus{
		Components: logstorev1.LogstoreStackComponentStatus{
			Compactor:     logstorev1.PodStatusMap{logstorev1.PodReady: {"compactor-pod-0"}},
			Distributor:   logstorev1.PodStatusMap{logstorev1.PodReady: {"distributor-pod-0"}},
			IndexGateway:  logstorev1.PodStatusMap{logstorev1.PodReady: {"index-gateway-pod-0"}},
			Ingester:      logstorev1.PodStatusMap{logstorev1.PodReady: {"ingester-pod-0"}},
			Querier:       logstorev1.PodStatusMap{logstorev1.PodReady: {"querier-pod-0"}},
			QueryFrontend: logstorev1.PodStatusMap{logstorev1.PodReady: {"query-frontend-pod-0"}},
			Gateway:       logstorev1.PodStatusMap{logstorev1.PodReady: {"logstorestack-gateway-pod-0"}},
			Ruler:         logstorev1.PodStatusMap{logstorev1.PodReady: {"ruler-pod-0"}},
		},
		Storage: logstorev1.LogstoreStackStorageStatus{
			CredentialMode: logstorev1.CredentialModeStatic,
		},
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

	err := Refresh(context.Background(), k, req, now, logstorev1.CredentialModeStatic, nil)

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

func TestRefreshSuccess_ZoneAwarePendingPod(t *testing.T) {
	now := time.Now()
	req := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "test-stack",
			Namespace: "test-ns",
		},
	}
	stack := logstorev1.LogstoreStack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-stack",
			Namespace: "test-ns",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Replication: &logstorev1.ReplicationSpec{
				Zones: []logstorev1.ZoneSpec{
					{
						TopologyKey: corev1.LabelTopologyZone,
					},
				},
			},
		},
	}
	testPod := corev1.Pod{
		Status: corev1.PodStatus{
			Phase: corev1.PodPending,
		},
	}

	k, sw := setupFakesNoError(t, &stack)
	k.ListStub = func(ctx context.Context, ol client.ObjectList, _ ...client.ListOption) error {
		switch ol.(type) {
		case *corev1.PodList:
			k.SetClientObjectList(ol, &corev1.PodList{
				Items: []corev1.Pod{
					testPod,
				},
			})
		case *corev1.NodeList:
			k.SetClientObjectList(ol, &corev1.NodeList{
				Items: []corev1.Node{},
			})
		}
		return nil
	}

	err := Refresh(context.Background(), k, req, now, logstorev1.CredentialModeStatic, nil)

	require.NoError(t, err)
	require.Equal(t, 1, k.GetCallCount())
	require.Equal(t, 9, k.ListCallCount())
	require.Equal(t, 1, sw.UpdateCallCount())
	_, updated, _ := sw.UpdateArgsForCall(0)
	updatedStack, ok := updated.(*logstorev1.LogstoreStack)
	if !ok {
		t.Fatalf("not a LogstoreStack: %T", updatedStack)
	}

	require.Len(t, updatedStack.Status.Conditions, 1)
	condition := updatedStack.Status.Conditions[0]
	require.Equal(t, conditionDegradedNodeLabels.Reason, condition.Reason)
	require.Equal(t, conditionDegradedNodeLabels.Type, condition.Type)
}
