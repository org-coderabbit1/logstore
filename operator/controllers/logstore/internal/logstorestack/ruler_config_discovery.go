package logstorestack

import (
	"context"
	"time"

	"github.com/ViaQ/logerr/v2/kverrors"
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	annotationRulerConfigDiscoveredAt = "logstore.acme.com/rulerConfigDiscoveredAt"
)

// AnnotateForRulerConfig adds/updates the `logstore.acme.com/rulerConfigDiscoveredAt` annotation
// to the named Logstorestack in the same namespace of the RulerConfig. If no LogstoreStack is found, then
// skip reconciliation.
func AnnotateForRulerConfig(ctx context.Context, k k8s.Client, name, namespace string) error {
	key := client.ObjectKey{Name: name, Namespace: namespace}
	ss, err := getLogstoreStack(ctx, k, key)
	if ss == nil || err != nil {
		return err
	}

	timeStamp := time.Now().UTC().Format(time.RFC3339)
	if err := updateAnnotation(ctx, k, ss, annotationRulerConfigDiscoveredAt, timeStamp); err != nil {
		return kverrors.Wrap(err, "failed to update logstorestack `rulerConfigDiscoveredAt` annotation", "key", key)
	}

	return nil
}

// RemoveRulerConfigAnnotation removes the `logstore.acme.com/rulerConfigDiscoveredAt` annotation
// from the named Logstorestack in the same namespace of the RulerConfig. If no LogstoreStack is found, then
// skip reconciliation.
func RemoveRulerConfigAnnotation(ctx context.Context, k k8s.Client, name, namespace string) error {
	key := client.ObjectKey{Name: name, Namespace: namespace}
	ss, err := getLogstoreStack(ctx, k, key)
	if ss == nil || err != nil {
		return err
	}

	if err := removeAnnotation(ctx, k, ss, annotationRulerConfigDiscoveredAt); err != nil {
		return kverrors.Wrap(err, "failed to update logstorestack `rulerConfigDiscoveredAt` annotation", "key", key)
	}

	return nil
}

func getLogstoreStack(ctx context.Context, k k8s.Client, key client.ObjectKey) (*logstorev1.LogstoreStack, error) {
	var s logstorev1.LogstoreStack

	if err := k.Get(ctx, key, &s); err != nil {
		if apierrors.IsNotFound(err) {
			// Do nothing
			return nil, nil
		}

		return nil, kverrors.Wrap(err, "failed to get logstorestack", "key", key)
	}

	return s.DeepCopy(), nil
}
