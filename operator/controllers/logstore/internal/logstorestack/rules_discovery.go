package logstorestack

import (
	"context"
	"time"

	"github.com/ViaQ/logerr/v2/kverrors"
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	annotationRulesDiscoveredAt = "logstore.acme.com/rulesDiscoveredAt"
)

// AnnotateForDiscoveredRules adds/updates the `logstore.acme.com/rulesDiscoveredAt` annotation
// to all instance of LogstoreStack on all namespaces to trigger the reconciliation loop.
func AnnotateForDiscoveredRules(ctx context.Context, k k8s.Client) error {
	timeStamp := time.Now().UTC().Format(time.RFC3339)

	var stacks logstorev1.LogstoreStackList
	err := k.List(ctx, &stacks, client.MatchingLabelsSelector{Selector: labels.Everything()})
	if err != nil {
		return kverrors.Wrap(err, "failed to list any logstorestack instances", "req")
	}

	for _, s := range stacks.Items {
		ss := s.DeepCopy()
		if err := updateAnnotation(ctx, k, ss, annotationRulesDiscoveredAt, timeStamp); err != nil {
			return kverrors.Wrap(err, "failed to update logstorestack `rulesDiscoveredAt` annotation", "name", ss.Name, "namespace", ss.Namespace)
		}
	}

	return nil
}
