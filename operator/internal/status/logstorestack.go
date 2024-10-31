package status

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"
)

const (
	messageReady                           = "All components ready"
	messageFailed                          = "One or more LogstoreStack components failed"
	messagePending                         = "One or more LogstoreStack components pending on dependencies"
	messageRunning                         = "All components are running, but some readiness checks are failing"
	messageDegradedMissingNodes            = "Cluster contains no nodes matching the labels used for zone-awareness"
	messageDegradedEmptyNodeLabel          = "No value for the labels used for zone-awareness"
	messageWarningNeedsSchemaVersionUpdate = "The schema configuration does not contain the most recent schema version and needs an update"
)

var (
	conditionFailed = metav1.Condition{
		Type:    string(logstorev1.ConditionFailed),
		Message: messageFailed,
		Reason:  string(logstorev1.ReasonFailedComponents),
	}
	conditionPending = metav1.Condition{
		Type:    string(logstorev1.ConditionPending),
		Message: messagePending,
		Reason:  string(logstorev1.ReasonPendingComponents),
	}
	conditionRunning = metav1.Condition{
		Type:    string(logstorev1.ConditionPending),
		Message: messageRunning,
		Reason:  string(logstorev1.ReasonPendingComponents),
	}
	conditionReady = metav1.Condition{
		Type:    string(logstorev1.ConditionReady),
		Message: messageReady,
		Reason:  string(logstorev1.ReasonReadyComponents),
	}
	conditionDegradedNodeLabels = metav1.Condition{
		Type:    string(logstorev1.ConditionDegraded),
		Message: messageDegradedMissingNodes,
		Reason:  string(logstorev1.ReasonZoneAwareNodesMissing),
	}
	conditionDegradedEmptyNodeLabel = metav1.Condition{
		Type:    string(logstorev1.ConditionDegraded),
		Message: messageDegradedEmptyNodeLabel,
		Reason:  string(logstorev1.ReasonZoneAwareEmptyLabel),
	}
)

// DegradedError contains information about why the managed LogstoreStack has an invalid configuration.
type DegradedError struct {
	Message string
	Reason  logstorev1.LogstoreStackConditionReason
	Requeue bool
}

func (e *DegradedError) Error() string {
	return fmt.Sprintf("cluster degraded: %s", e.Message)
}

func generateConditions(ctx context.Context, cs *logstorev1.LogstoreStackComponentStatus, k k8s.Client, stack *logstorev1.LogstoreStack, degradedErr *DegradedError) ([]metav1.Condition, error) {
	conditions := generateWarnings(stack.Status.Storage.Schemas)

	mainCondition, err := generateCondition(ctx, cs, k, stack, degradedErr)
	if err != nil {
		return nil, err
	}

	conditions = append(conditions, mainCondition)
	return conditions, nil
}

func generateCondition(ctx context.Context, cs *logstorev1.LogstoreStackComponentStatus, k k8s.Client, stack *logstorev1.LogstoreStack, degradedErr *DegradedError) (metav1.Condition, error) {
	if degradedErr != nil {
		return metav1.Condition{
			Type:    string(logstorev1.ConditionDegraded),
			Message: degradedErr.Message,
			Reason:  string(degradedErr.Reason),
		}, nil
	}

	// Check for failed pods first
	failed := len(cs.Compactor[logstorev1.PodFailed]) +
		len(cs.Distributor[logstorev1.PodFailed]) +
		len(cs.Ingester[logstorev1.PodFailed]) +
		len(cs.Querier[logstorev1.PodFailed]) +
		len(cs.QueryFrontend[logstorev1.PodFailed]) +
		len(cs.Gateway[logstorev1.PodFailed]) +
		len(cs.IndexGateway[logstorev1.PodFailed]) +
		len(cs.Ruler[logstorev1.PodFailed])

	if failed != 0 {
		return conditionFailed, nil
	}

	// Check for pending pods
	pending := len(cs.Compactor[logstorev1.PodPending]) +
		len(cs.Distributor[logstorev1.PodPending]) +
		len(cs.Ingester[logstorev1.PodPending]) +
		len(cs.Querier[logstorev1.PodPending]) +
		len(cs.QueryFrontend[logstorev1.PodPending]) +
		len(cs.Gateway[logstorev1.PodPending]) +
		len(cs.IndexGateway[logstorev1.PodPending]) +
		len(cs.Ruler[logstorev1.PodPending])

	if pending != 0 {
		if stack.Spec.Replication != nil && len(stack.Spec.Replication.Zones) > 0 {
			// When there are pending pods and zone-awareness is enabled check if there are any nodes
			// that can satisfy the constraints and emit a condition if not.
			nodesOk, labelsOk, err := checkForZoneawareNodes(ctx, k, stack.Spec.Replication.Zones)
			if err != nil {
				return metav1.Condition{}, err
			}

			if !nodesOk {
				return conditionDegradedNodeLabels, nil
			}

			if !labelsOk {
				return conditionDegradedEmptyNodeLabel, nil
			}
		}

		return conditionPending, nil
	}

	// Check if there are pods that are running but not ready
	running := len(cs.Compactor[logstorev1.PodRunning]) +
		len(cs.Distributor[logstorev1.PodRunning]) +
		len(cs.Ingester[logstorev1.PodRunning]) +
		len(cs.Querier[logstorev1.PodRunning]) +
		len(cs.QueryFrontend[logstorev1.PodRunning]) +
		len(cs.Gateway[logstorev1.PodRunning]) +
		len(cs.IndexGateway[logstorev1.PodRunning]) +
		len(cs.Ruler[logstorev1.PodRunning])

	if running > 0 {
		return conditionRunning, nil
	}

	return conditionReady, nil
}

func checkForZoneawareNodes(ctx context.Context, k client.Client, zones []logstorev1.ZoneSpec) (nodesOk bool, labelsOk bool, err error) {
	nodeLabels := client.HasLabels{}
	for _, z := range zones {
		nodeLabels = append(nodeLabels, z.TopologyKey)
	}

	nodeList := &corev1.NodeList{}
	if err := k.List(ctx, nodeList, nodeLabels); err != nil {
		return false, false, err
	}

	if len(nodeList.Items) == 0 {
		return false, false, nil
	}

	for _, node := range nodeList.Items {
		for _, nodeLabel := range nodeLabels {
			if node.Labels[nodeLabel] == "" {
				return true, false, nil
			}
		}
	}

	return true, true, nil
}

func generateWarnings(schemas []logstorev1.ObjectStorageSchema) []metav1.Condition {
	warnings := make([]metav1.Condition, 0, 2)

	if len(schemas) > 0 && schemas[len(schemas)-1].Version != logstorev1.ObjectStorageSchemaV13 {
		warnings = append(warnings, metav1.Condition{
			Type:    string(logstorev1.ConditionWarning),
			Reason:  string(logstorev1.ReasonStorageNeedsSchemaUpdate),
			Message: messageWarningNeedsSchemaVersionUpdate,
		})
	}

	return warnings
}
