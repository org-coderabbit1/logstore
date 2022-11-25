package handlers

import (
	"context"
	"fmt"

	"github.com/ViaQ/logerr/v2/kverrors"
	"github.com/go-logr/logr"

	configv1 "example.com/acme/logstore/operator/apis/config/v1"
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/certrotation"
	"example.com/acme/logstore/operator/internal/external/k8s"
	"example.com/acme/logstore/operator/internal/handlers/internal/certificates"
	"example.com/acme/logstore/operator/internal/manifests"
	"example.com/acme/logstore/operator/internal/status"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrlutil "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// CreateOrRotateCertificates handles the LogstoreStack client and serving certificate creation and rotation
// including the signing CA and a ca bundle or else returns an error. It returns only a degrade-condition-worthy
// error if building the manifests fails for any reason.
func CreateOrRotateCertificates(ctx context.Context, log logr.Logger, req ctrl.Request, k k8s.Client, s *runtime.Scheme, fg configv1.FeatureGates) error {
	ll := log.WithValues("logstorestack", req.String(), "event", "createOrRotateCerts")

	var stack logstorev1.LogstoreStack
	if err := k.Get(ctx, req.NamespacedName, &stack); err != nil {
		if apierrors.IsNotFound(err) {
			// maybe the user deleted it before we could react? Either way this isn't an issue
			ll.Error(err, "could not find the requested LogstoreStack", "name", req.String())
			return nil
		}
		return kverrors.Wrap(err, "failed to lookup LogstoreStack", "name", req.String())
	}

	var mode logstorev1.ModeType
	if stack.Spec.Tenants != nil {
		mode = stack.Spec.Tenants.Mode
	}

	opts, err := certificates.GetOptions(ctx, k, req, mode)
	if err != nil {
		return kverrors.Wrap(err, "failed to lookup certificates secrets", "name", req.String())
	}

	ll.Info("begin building certificate manifests")

	if optErr := certrotation.ApplyDefaultSettings(&opts, fg.BuiltInCertManagement); optErr != nil {
		ll.Error(optErr, "failed to conform options to build settings")
		return optErr
	}

	objects, err := certrotation.BuildAll(opts)
	if err != nil {
		ll.Error(err, "failed to build certificate manifests")
		return &status.DegradedError{
			Message: "Failed to rotate TLS certificates",
			Reason:  logstorev1.ReasonFailedCertificateRotation,
			Requeue: true,
		}
	}

	ll.Info("certificate manifests built", "count", len(objects))

	var errCount int32

	for _, obj := range objects {
		l := ll.WithValues(
			"object_name", obj.GetName(),
			"object_kind", obj.GetObjectKind(),
		)

		obj.SetNamespace(req.Namespace)

		if err := ctrl.SetControllerReference(&stack, obj, s); err != nil {
			l.Error(err, "failed to set controller owner reference to resource")
			errCount++
			continue
		}

		desired := obj.DeepCopyObject().(client.Object)
		mutateFn := manifests.MutateFuncFor(obj, desired, nil)

		op, err := ctrl.CreateOrUpdate(ctx, k, obj, mutateFn)
		if err != nil {
			l.Error(err, "failed to configure resource")
			errCount++
			continue
		}

		msg := fmt.Sprintf("Resource has been %s", op)
		switch op {
		case ctrlutil.OperationResultNone:
			l.V(1).Info(msg)
		default:
			l.Info(msg)
		}
	}

	if errCount > 0 {
		return kverrors.New("failed to create or rotate LogstoreStack certificates", "name", req.String())
	}

	return nil
}
