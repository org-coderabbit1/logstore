package gateway

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"

	configv1 "example.com/acme/logstore/operator/api/config/v1"
	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"
	"example.com/acme/logstore/operator/internal/handlers/internal/openshift"
	"example.com/acme/logstore/operator/internal/manifests"
	"example.com/acme/logstore/operator/internal/status"
)

// BuildOptions returns the options needed to generate Kubernetes resource
// manifests for the logstorestack-gateway.
// The returned error can be a status.DegradedError in the following cases:
//   - The tenants spec is missing.
//   - The tenants spec is invalid.
func BuildOptions(ctx context.Context, log logr.Logger, k k8s.Client, stack *logstorev1.LogstoreStack, fg configv1.FeatureGates) (string, manifests.Tenants, error) {
	var (
		err        error
		baseDomain string
		secrets    []*manifests.TenantSecrets
		configs    map[string]manifests.TenantConfig
		tenants    manifests.Tenants
	)

	if !fg.LogstoreStackGateway {
		return "", tenants, nil
	}

	if stack.Spec.Tenants == nil {
		return "", tenants, &status.DegradedError{
			Message: "Invalid tenants configuration: TenantsSpec cannot be nil when gateway flag is enabled",
			Reason:  logstorev1.ReasonInvalidTenantsConfiguration,
			Requeue: false,
		}
	}

	if err = validateModes(stack); err != nil {
		return "", tenants, &status.DegradedError{
			Message: fmt.Sprintf("Invalid tenants configuration: %s", err),
			Reason:  logstorev1.ReasonInvalidTenantsConfiguration,
			Requeue: false,
		}
	}

	switch stack.Spec.Tenants.Mode {
	case logstorev1.OpenshiftLogging, logstorev1.OpenshiftNetwork:
		baseDomain, err = getOpenShiftBaseDomain(ctx, k)
		if err != nil {
			return "", tenants, err
		}

		if stack.Spec.Proxy == nil {
			// If the LogstoreStack has no proxy set but there is a cluster-wide proxy setting,
			// set the LogstoreStack proxy to that.
			ocpProxy, proxyErr := openshift.GetProxy(ctx, k)
			if proxyErr != nil {
				return "", tenants, proxyErr
			}

			stack.Spec.Proxy = ocpProxy
		}
	default:
		secrets, err = getTenantSecrets(ctx, k, stack)
		if err != nil {
			return "", tenants, err
		}
	}

	// extract the existing tenant's id, cookieSecret if exists, otherwise create new.
	configs, err = getTenantConfigFromSecret(ctx, k, stack)
	if err != nil {
		log.Error(err, "error in getting tenant secret data")
	}

	tenants = manifests.Tenants{
		Secrets: secrets,
		Configs: configs,
	}

	return baseDomain, tenants, nil
}
