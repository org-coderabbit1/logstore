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

// AnnotateForRulerConfig adds/updates the `logstore.acme.com/rulerConfigDiscoveredAt` annotation
// to the named Logstorestack in the same namespace of the RulerConfig. If no LogstoreStack is found, then
// skip reconciliation.
func AnnotateForRulerConfig(ctx context.Context, k k8s.Client, name, namespace string) error {
	var s logstorev1.LogstoreStack
	key := client.ObjectKey{Name: name, Namespace: namespace}

	if err := k.Get(ctx, key, &s); err != nil {
		if apierrors.IsNotFound(err) {
			// Do nothing
			return nil
		}

		return kverrors.Wrap(err, "failed to get logstorestack", "key", key)
	}

	ss := s.DeepCopy()
	if ss.Annotations == nil {
		ss.Annotations = make(map[string]string)
	}

	ss.Annotations["logstore.acme.com/rulerConfigDiscoveredAt"] = time.Now().UTC().Format(time.RFC3339)

	if err := k.Update(ctx, ss); err != nil {
		return kverrors.Wrap(err, "failed to update logstorestack `rulerConfigDiscoveredAt` annotation", "key", key)
	}

	return nil
}
