package logstore

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/controller/logstore/internal/logstorestack"
)

// RulerConfigReconciler reconciles a RulerConfig object
type RulerConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=logstore.acme.com,resources=rulerconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=logstore.acme.com,resources=rulerconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=logstore.acme.com,resources=rulerconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *RulerConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var rc logstorev1.RulerConfig
	key := client.ObjectKey{Name: req.Name, Namespace: req.Namespace}
	if err := r.Get(ctx, key, &rc); err != nil {
		if errors.IsNotFound(err) {
			// RulerConfig not found, remove annotation from LogstoreStack.
			err = logstorestack.RemoveRulerConfigAnnotation(ctx, r.Client, req.Name, req.Namespace)
			if err != nil {
				return ctrl.Result{}, err
			}

			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	err := logstorestack.AnnotateForRulerConfig(ctx, r.Client, req.Name, req.Namespace)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RulerConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&logstorev1.RulerConfig{}).
		Complete(r)
}
