package logstore

import (
	"context"

	"github.com/ViaQ/logerr/v2/kverrors"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"
	"example.com/acme/logstore/operator/internal/handlers"
)

var createOrUpdatePodWithLabelPred = builder.WithPredicates(predicate.Funcs{
	UpdateFunc:  func(e event.UpdateEvent) bool { return eventPodHasLabel(e.ObjectNew) },
	CreateFunc:  func(e event.CreateEvent) bool { return eventPodHasLabel(e.Object) },
	DeleteFunc:  func(e event.DeleteEvent) bool { return false },
	GenericFunc: func(e event.GenericEvent) bool { return false },
})

// LogstoreStackZoneAwarePodReconciler watches all the logstore component pods and updates the pod annotations with the topology node labels.
type LogstoreStackZoneAwarePodReconciler struct {
	client.Client
	Log logr.Logger
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *LogstoreStackZoneAwarePodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logstorePod := &corev1.Pod{}
	if err := r.Client.Get(ctx, req.NamespacedName, logstorePod); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, kverrors.Wrap(err, "failed to get pod", "name", req.NamespacedName)
	}

	err := handlers.AnnotatePodWithAvailabilityZone(ctx, r.Log, r.Client, logstorePod)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LogstoreStackZoneAwarePodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	b := ctrl.NewControllerManagedBy(mgr)
	return r.buildController(k8s.NewCtrlBuilder(b))
}

func (r *LogstoreStackZoneAwarePodReconciler) buildController(bld k8s.Builder) error {
	return bld.
		Named("ZoneAwarePod").
		Watches(&corev1.Pod{}, &handler.EnqueueRequestForObject{}, createOrUpdatePodWithLabelPred).
		Complete(r)
}

func eventPodHasLabel(object client.Object) bool {
	pod, isPod := object.(*corev1.Pod)
	if !isPod {
		return false
	}

	_, hasLabel := pod.Labels[logstorev1.LabelZoneAwarePod]
	return hasLabel
}
