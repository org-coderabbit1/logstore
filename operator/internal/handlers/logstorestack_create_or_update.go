package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/ViaQ/logerr/v2/kverrors"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrlutil "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	configv1 "example.com/acme/logstore/operator/api/config/v1"
	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"
	"example.com/acme/logstore/operator/internal/handlers/internal/gateway"
	"example.com/acme/logstore/operator/internal/handlers/internal/rules"
	"example.com/acme/logstore/operator/internal/handlers/internal/serviceaccounts"
	"example.com/acme/logstore/operator/internal/handlers/internal/storage"
	"example.com/acme/logstore/operator/internal/handlers/internal/tlsprofile"
	"example.com/acme/logstore/operator/internal/manifests"
	"example.com/acme/logstore/operator/internal/status"
)

// CreateOrUpdateLogstoreStack handles LogstoreStack create and update events.
func CreateOrUpdateLogstoreStack(
	ctx context.Context,
	log logr.Logger,
	req ctrl.Request,
	k k8s.Client,
	s *runtime.Scheme,
	fg configv1.FeatureGates,
) (logstorev1.CredentialMode, error) {
	ll := log.WithValues("logstorestack", req.NamespacedName, "event", "createOrUpdate")

	var stack logstorev1.LogstoreStack
	if err := k.Get(ctx, req.NamespacedName, &stack); err != nil {
		if apierrors.IsNotFound(err) {
			// maybe the user deleted it before we could react? Either way this isn't an issue
			ll.Error(err, "could not find the requested logstore stack", "name", req.NamespacedName)
			return "", nil
		}
		return "", kverrors.Wrap(err, "failed to lookup logstorestack", "name", req.NamespacedName)
	}

	img := os.Getenv(manifests.EnvRelatedImageLogstore)
	if img == "" {
		img = manifests.DefaultContainerImage
	}

	gwImg := os.Getenv(manifests.EnvRelatedImageGateway)
	if gwImg == "" {
		gwImg = manifests.DefaultLogstoreStackGatewayImage
	}

	objStore, err := storage.BuildOptions(ctx, k, &stack, fg)
	if err != nil {
		return "", err
	}

	baseDomain, tenants, err := gateway.BuildOptions(ctx, ll, k, &stack, fg)
	if err != nil {
		return "", err
	}

	if err = rules.Cleanup(ctx, ll, k, &stack); err != nil {
		return "", err
	}

	alertingRules, recordingRules, ruler, ocpOptions, err := rules.BuildOptions(ctx, ll, k, &stack)
	if err != nil {
		return "", err
	}

	certRotationRequiredAt := ""
	if stack.Annotations != nil {
		certRotationRequiredAt = stack.Annotations[manifests.AnnotationCertRotationRequiredAt]
	}

	timeoutConfig, err := manifests.NewTimeoutConfig(stack.Spec.Limits)
	if err != nil {
		ll.Error(err, "failed to parse query timeout")
		return "", &status.DegradedError{
			Message: fmt.Sprintf("Error parsing query timeout: %s", err),
			Reason:  logstorev1.ReasonQueryTimeoutInvalid,
			Requeue: false,
		}
	}

	// Here we will translate the logstorev1.LogstoreStack options into manifest options
	opts := manifests.Options{
		Name:                   req.Name,
		Namespace:              req.Namespace,
		Image:                  img,
		GatewayImage:           gwImg,
		GatewayBaseDomain:      baseDomain,
		Stack:                  stack.Spec,
		Gates:                  fg,
		ObjectStorage:          objStore,
		CertRotationRequiredAt: certRotationRequiredAt,
		AlertingRules:          alertingRules,
		RecordingRules:         recordingRules,
		Ruler:                  ruler,
		Timeouts:               timeoutConfig,
		Tenants:                tenants,
		OpenShiftOptions:       ocpOptions,
	}

	ll.Info("begin building manifests")

	if optErr := manifests.ApplyDefaultSettings(&opts); optErr != nil {
		ll.Error(optErr, "failed to conform options to build settings")
		return "", optErr
	}

	if fg.LogstoreStackGateway {
		if optErr := manifests.ApplyGatewayDefaultOptions(&opts); optErr != nil {
			ll.Error(optErr, "failed to apply defaults options to gateway settings")
			return "", optErr
		}
	}

	tlsProfileType := configv1.TLSProfileType(fg.TLSProfile)
	// Overwrite the profile from the flags and use the profile from the apiserver instead
	if fg.OpenShift.ClusterTLSPolicy {
		tlsProfileType = configv1.TLSProfileType("")
	}

	tlsProfile, err := tlsprofile.GetTLSSecurityProfile(ctx, k, tlsProfileType)
	if err != nil {
		// The API server is not guaranteed to be there nor have a result.
		ll.Error(err, "failed to get security profile. will use default tls profile.")
	}

	if optErr := manifests.ApplyTLSSettings(&opts, tlsProfile); optErr != nil {
		ll.Error(optErr, "failed to conform options to tls profile settings")
		return "", optErr
	}

	objects, err := manifests.BuildAll(opts)
	if err != nil {
		ll.Error(err, "failed to build manifests")
		return "", err
	}

	ll.Info("manifests built", "count", len(objects))

	// The status is updated before the objects are actually created to
	// avoid the scenario in which the configmap is successfully created or
	// updated and another resource is not. This would cause the status to
	// be possibly misaligned with the configmap, which could lead to
	// a user possibly being unable to read logs.
	if err := status.SetStorageSchemaStatus(ctx, k, req, objStore.Schemas); err != nil {
		ll.Error(err, "failed to set storage schema status")
		return "", err
	}

	var errCount int32

	for _, obj := range objects {
		l := ll.WithValues(
			"object_name", obj.GetName(),
			"object_kind", obj.GetObjectKind(),
		)

		if isNamespacedResource(obj) {
			obj.SetNamespace(req.Namespace)

			if err := ctrl.SetControllerReference(&stack, obj, s); err != nil {
				l.Error(err, "failed to set controller owner reference to resource")
				errCount++
				continue
			}
		}

		depAnnotations, err := dependentAnnotations(ctx, k, obj)
		if err != nil {
			l.Error(err, "failed to set dependent annotations")
			return "", err
		}

		desired := obj.DeepCopyObject().(client.Object)
		mutateFn := manifests.MutateFuncFor(obj, desired, depAnnotations)

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
		return "", kverrors.New("failed to configure logstorestack resources", "name", req.NamespacedName)
	}

	return objStore.CredentialMode, nil
}

func dependentAnnotations(ctx context.Context, k k8s.Client, obj client.Object) (map[string]string, error) {
	a := obj.GetAnnotations()
	saName, ok := a[corev1.ServiceAccountNameKey]
	if !ok || saName == "" {
		return nil, nil
	}

	key := client.ObjectKey{Name: saName, Namespace: obj.GetNamespace()}
	uid, err := serviceaccounts.GetUID(ctx, k, key)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		corev1.ServiceAccountUIDKey: uid,
	}, nil
}

// isNamespacedResource determines if an object should be managed or not by a LogstoreStack
func isNamespacedResource(obj client.Object) bool {
	switch obj.(type) {
	case *rbacv1.ClusterRole, *rbacv1.ClusterRoleBinding:
		return false
	default:
		return true
	}
}
