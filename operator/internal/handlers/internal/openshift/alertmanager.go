package openshift

import (
	"context"

	"github.com/ViaQ/logerr/v2/kverrors"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"
	"example.com/acme/logstore/operator/internal/manifests/openshift"
)

// AlertManagerSVCExists returns true if the Openshift AlertManager is present in the cluster.
func AlertManagerSVCExists(ctx context.Context, stack logstorev1.LogstoreStackSpec, k k8s.Client) (bool, error) {
	if stack.Tenants == nil || (stack.Tenants.Mode != logstorev1.OpenshiftLogging && stack.Tenants.Mode != logstorev1.OpenshiftNetwork) {
		return false, nil
	}

	var svc corev1.Service
	key := client.ObjectKey{Name: openshift.MonitoringSVCOperated, Namespace: openshift.MonitoringNS}

	err := k.Get(ctx, key, &svc)
	if err != nil && !apierrors.IsNotFound(err) {
		return false, kverrors.Wrap(err, "failed to lookup alertmanager service", "name", key)
	}

	return err == nil, nil
}

// UserWorkloadAlertManagerSVCExists returns true if the Openshift User Workload AlertManager is present in the cluster.
func UserWorkloadAlertManagerSVCExists(ctx context.Context, stack logstorev1.LogstoreStackSpec, k k8s.Client) (bool, error) {
	if stack.Tenants == nil || stack.Tenants.Mode != logstorev1.OpenshiftLogging {
		return false, nil
	}

	var svc corev1.Service
	key := client.ObjectKey{Name: openshift.MonitoringSVCOperated, Namespace: openshift.MonitoringUserWorkloadNS}

	err := k.Get(ctx, key, &svc)
	if err != nil && !apierrors.IsNotFound(err) {
		return false, kverrors.Wrap(err, "failed to lookup user workload alertmanager service", "name", key)
	}

	return err == nil, nil
}
