package status

import (
	"context"

	"github.com/ViaQ/logerr/v2/kverrors"
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

// SetStorageSchemaStatus updates the storage status component
func SetStorageSchemaStatus(ctx context.Context, k k8s.Client, req ctrl.Request, schemas []logstorev1.ObjectStorageSchema) error {
	var s logstorev1.LogstoreStack
	if err := k.Get(ctx, req.NamespacedName, &s); err != nil {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return kverrors.Wrap(err, "failed to lookup logstorestack", "name", req.NamespacedName)
	}

	s.Status.Storage = logstorev1.LogstoreStackStorageStatus{
		Schemas: schemas,
	}

	return k.Status().Update(ctx, &s)
}
