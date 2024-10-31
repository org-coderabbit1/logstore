package status

import (
	"context"

	"github.com/ViaQ/logerr/v2/kverrors"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"
	"example.com/acme/logstore/operator/internal/manifests"
)

// generateComponentStatus updates the pod status map component
func generateComponentStatus(ctx context.Context, k k8s.Client, s *logstorev1.LogstoreStack) (*logstorev1.LogstoreStackComponentStatus, error) {
	var err error
	result := &logstorev1.LogstoreStackComponentStatus{}
	result.Compactor, err = appendPodStatus(ctx, k, manifests.LabelCompactorComponent, s.Name, s.Namespace)
	if err != nil {
		return nil, kverrors.Wrap(err, "failed lookup LogstoreStack component pods status", "name", manifests.LabelCompactorComponent)
	}

	result.Querier, err = appendPodStatus(ctx, k, manifests.LabelQuerierComponent, s.Name, s.Namespace)
	if err != nil {
		return nil, kverrors.Wrap(err, "failed lookup LogstoreStack component pods status", "name", manifests.LabelQuerierComponent)
	}

	result.Distributor, err = appendPodStatus(ctx, k, manifests.LabelDistributorComponent, s.Name, s.Namespace)
	if err != nil {
		return nil, kverrors.Wrap(err, "failed lookup LogstoreStack component pods status", "name", manifests.LabelDistributorComponent)
	}

	result.QueryFrontend, err = appendPodStatus(ctx, k, manifests.LabelQueryFrontendComponent, s.Name, s.Namespace)
	if err != nil {
		return nil, kverrors.Wrap(err, "failed lookup LogstoreStack component pods status", "name", manifests.LabelQueryFrontendComponent)
	}

	result.IndexGateway, err = appendPodStatus(ctx, k, manifests.LabelIndexGatewayComponent, s.Name, s.Namespace)
	if err != nil {
		return nil, kverrors.Wrap(err, "failed lookup LogstoreStack component pods status", "name", manifests.LabelIngesterComponent)
	}

	result.Ingester, err = appendPodStatus(ctx, k, manifests.LabelIngesterComponent, s.Name, s.Namespace)
	if err != nil {
		return nil, kverrors.Wrap(err, "failed lookup LogstoreStack component pods status", "name", manifests.LabelIndexGatewayComponent)
	}

	result.Gateway, err = appendPodStatus(ctx, k, manifests.LabelGatewayComponent, s.Name, s.Namespace)
	if err != nil {
		return nil, kverrors.Wrap(err, "failed lookup LogstoreStack component pods status", "name", manifests.LabelGatewayComponent)
	}

	result.Ruler, err = appendPodStatus(ctx, k, manifests.LabelRulerComponent, s.Name, s.Namespace)
	if err != nil {
		return nil, kverrors.Wrap(err, "failed lookup LogstoreStack component pods status", "name", manifests.LabelRulerComponent)
	}

	return result, nil
}

func appendPodStatus(ctx context.Context, k k8s.Client, component, stack, ns string) (logstorev1.PodStatusMap, error) {
	psm := logstorev1.PodStatusMap{}
	pods := &corev1.PodList{}
	opts := []client.ListOption{
		client.MatchingLabels(manifests.ComponentLabels(component, stack)),
		client.InNamespace(ns),
	}
	if err := k.List(ctx, pods, opts...); err != nil {
		return nil, kverrors.Wrap(err, "failed to list pods for LogstoreStack component", "name", stack, "component", component)
	}
	for _, pod := range pods.Items {
		status := podStatus(&pod)
		psm[status] = append(psm[status], pod.Name)
	}

	if len(psm) == 0 {
		psm = logstorev1.PodStatusMap{
			logstorev1.PodFailed:  []string{},
			logstorev1.PodPending: []string{},
			logstorev1.PodRunning: []string{},
			logstorev1.PodReady:   []string{},
		}
	}
	return psm, nil
}

func podStatus(pod *corev1.Pod) logstorev1.PodStatus {
	status := pod.Status
	switch status.Phase {
	case corev1.PodFailed:
		return logstorev1.PodFailed
	case corev1.PodPending:
		return logstorev1.PodPending
	case corev1.PodRunning:
	default:
		return logstorev1.PodStatusUnknown
	}

	for _, c := range status.ContainerStatuses {
		if !c.Ready {
			return logstorev1.PodRunning
		}
	}

	return logstorev1.PodReady
}
