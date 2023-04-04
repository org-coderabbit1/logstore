package handlers

import (
	"context"

	"github.com/ViaQ/logerr/v2/kverrors"
	"github.com/go-logr/logr"
	configv1 "example.com/acme/logstore/operator/apis/config/v1"
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/certrotation"
	"example.com/acme/logstore/operator/internal/external/k8s"
	"example.com/acme/logstore/operator/internal/handlers/internal/certificates"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

// CheckCertExpiry handles the case if the LogstoreStack managed signing CA, client and/or serving
// certificates expired. Returns true if any of those expired and an error representing the reason
// of expiry.
func CheckCertExpiry(ctx context.Context, log logr.Logger, req ctrl.Request, k k8s.Client, fg configv1.FeatureGates) error {
	ll := log.WithValues("logstorestack", req.String(), "event", "checkCertExpiry")

	var stack logstorev1.LogstoreStack
	if err := k.Get(ctx, req.NamespacedName, &stack); err != nil {
		if apierrors.IsNotFound(err) {
			// maybe the user deleted it before we could react? Either way this isn't an issue
			ll.Error(err, "could not find the requested logstore stack", "name", req.String())
			return nil
		}
		return kverrors.Wrap(err, "failed to lookup logstorestack", "name", req.String())
	}

	var mode logstorev1.ModeType
	if stack.Spec.Tenants != nil {
		mode = stack.Spec.Tenants.Mode
	}

	opts, err := certificates.GetOptions(ctx, k, req, mode)
	if err != nil {
		return kverrors.Wrap(err, "failed to lookup certificates secrets", "name", req.String())
	}

	if optErr := certrotation.ApplyDefaultSettings(&opts, fg.BuiltInCertManagement); optErr != nil {
		ll.Error(optErr, "failed to conform options to build settings")
		return optErr
	}

	if err := certrotation.SigningCAExpired(opts); err != nil {
		return err
	}

	if err := certrotation.CertificatesExpired(opts); err != nil {
		return err
	}

	return nil
}
