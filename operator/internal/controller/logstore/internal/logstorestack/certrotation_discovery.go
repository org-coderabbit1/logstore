package logstorestack

import (
	"context"
	"fmt"
	"time"

	"github.com/ViaQ/logerr/v2/kverrors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"
)

const certRotationRequiredAtKey = "logstore.acme.com/certRotationRequiredAt"

// AnnotateForRequiredCertRotation adds/updates the `logstore.acme.com/certRotationRequiredAt` annotation
// to the named Logstorestack if any of the managed client/serving/ca certificates expired. If no LogstoreStack
// is found, then skip reconciliation.
func AnnotateForRequiredCertRotation(ctx context.Context, k k8s.Client, name, namespace string) error {
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
	timeStamp := time.Now().UTC().Format(time.RFC3339)
	if err := updateAnnotation(ctx, k, ss, certRotationRequiredAtKey, timeStamp); err != nil {
		return kverrors.Wrap(err, fmt.Sprintf("failed to update logstorestack `%s` annotation", certRotationRequiredAtKey), "key", key)
	}

	return nil
}
