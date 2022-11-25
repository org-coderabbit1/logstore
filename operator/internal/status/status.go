package status

import (
	"context"

	"github.com/ViaQ/logerr/v2/kverrors"
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Refresh executes an aggregate update of the LogstoreStack Status struct, i.e.
// - It recreates the Status.Components pod status map per component.
// - It sets the appropriate Status.Condition to true that matches the pod status maps.
func Refresh(ctx context.Context, k k8s.Client, req ctrl.Request) error {
	if err := SetComponentsStatus(ctx, k, req); err != nil {
		return err
	}

	var s logstorev1.LogstoreStack
	if err := k.Get(ctx, req.NamespacedName, &s); err != nil {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return kverrors.Wrap(err, "failed to lookup logstorestack", "name", req.NamespacedName)
	}

	cs := s.Status.Components

	// Check for failed pods first
	failed := len(cs.Compactor[corev1.PodFailed]) +
		len(cs.Distributor[corev1.PodFailed]) +
		len(cs.Ingester[corev1.PodFailed]) +
		len(cs.Querier[corev1.PodFailed]) +
		len(cs.QueryFrontend[corev1.PodFailed]) +
		len(cs.Gateway[corev1.PodFailed]) +
		len(cs.IndexGateway[corev1.PodFailed]) +
		len(cs.Ruler[corev1.PodFailed])

	unknown := len(cs.Compactor[corev1.PodUnknown]) +
		len(cs.Distributor[corev1.PodUnknown]) +
		len(cs.Ingester[corev1.PodUnknown]) +
		len(cs.Querier[corev1.PodUnknown]) +
		len(cs.QueryFrontend[corev1.PodUnknown]) +
		len(cs.Gateway[corev1.PodUnknown]) +
		len(cs.IndexGateway[corev1.PodUnknown]) +
		len(cs.Ruler[corev1.PodUnknown])

	if failed != 0 || unknown != 0 {
		return SetFailedCondition(ctx, k, req)
	}

	// Check for pending pods
	pending := len(cs.Compactor[corev1.PodPending]) +
		len(cs.Distributor[corev1.PodPending]) +
		len(cs.Ingester[corev1.PodPending]) +
		len(cs.Querier[corev1.PodPending]) +
		len(cs.QueryFrontend[corev1.PodPending]) +
		len(cs.Gateway[corev1.PodPending]) +
		len(cs.IndexGateway[corev1.PodPending]) +
		len(cs.Ruler[corev1.PodPending])

	if pending != 0 {
		return SetPendingCondition(ctx, k, req)
	}
	return SetReadyCondition(ctx, k, req)
}
