package storage

import (
	"context"
	"fmt"
	"time"

	configv1 "example.com/acme/logstore/operator/api/config/v1"
	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s"
	"example.com/acme/logstore/operator/internal/manifests/storage"
	"example.com/acme/logstore/operator/internal/status"
)

// BuildOptions returns the object storage options to generate Kubernetes resource manifests
// which require access to object storage buckets.
// The returned error can be a status.DegradedError in the following cases:
//   - The user-provided object storage secret is missing.
//   - The object storage Secret data is invalid.
//   - The object storage schema config is invalid.
//   - The object storage CA ConfigMap is missing if one referenced.
//   - The object storage CA ConfigMap data is invalid.
//   - The object storage token cco auth secret is missing (Only on OpenShift STS-clusters)
func BuildOptions(ctx context.Context, k k8s.Client, stack *logstorev1.LogstoreStack, fg configv1.FeatureGates) (storage.Options, error) {
	storageSecret, tokenCCOAuthSecret, err := getSecrets(ctx, k, stack, fg)
	if err != nil {
		return storage.Options{}, err
	}

	objStore, err := extractSecrets(stack.Spec.Storage.Secret, storageSecret, tokenCCOAuthSecret, fg)
	if err != nil {
		return storage.Options{}, &status.DegradedError{
			Message: fmt.Sprintf("Invalid object storage secret contents: %s", err),
			Reason:  logstorev1.ReasonInvalidObjectStorageSecret,
			Requeue: false,
		}
	}

	if objStore.CredentialMode == logstorev1.CredentialModeTokenCCO && tokenCCOAuthSecret == nil {
		// If we have no token cco auth secret at this point, it is an error
		return storage.Options{}, &status.DegradedError{
			Message: "Missing OpenShift cloud credentials secret",
			Reason:  logstorev1.ReasonMissingTokenCCOAuthSecret,
			Requeue: true,
		}
	}

	now := time.Now().UTC()
	storageSchemas, err := storage.BuildSchemaConfig(
		now,
		stack.Spec.Storage,
		stack.Status.Storage,
	)
	if err != nil {
		return storage.Options{}, &status.DegradedError{
			Message: fmt.Sprintf("Invalid object storage schema contents: %s", err),
			Reason:  logstorev1.ReasonInvalidObjectStorageSchema,
			Requeue: false,
		}
	}

	objStore.Schemas = storageSchemas
	objStore.AllowStructuredMetadata = allowStructuredMetadata(storageSchemas, now)

	if stack.Spec.Storage.TLS == nil {
		return objStore, nil
	}

	tlsConfig := stack.Spec.Storage.TLS
	if tlsConfig.CA == "" {
		return storage.Options{}, &status.DegradedError{
			Message: "Missing object storage CA config map",
			Reason:  logstorev1.ReasonMissingObjectStorageCAConfigMap,
			Requeue: false,
		}
	}

	cm, err := getCAConfigMap(ctx, k, stack, tlsConfig.CA)
	if err != nil {
		return storage.Options{}, err
	}

	caKey := defaultCAKey
	if tlsConfig.CAKey != "" {
		caKey = tlsConfig.CAKey
	}

	var caHash string
	caHash, err = checkCAConfigMap(cm, caKey)
	if err != nil {
		return storage.Options{}, &status.DegradedError{
			Message: fmt.Sprintf("Invalid object storage CA configmap contents: %s", err),
			Reason:  logstorev1.ReasonInvalidObjectStorageCAConfigMap,
			Requeue: false,
		}
	}

	objStore.SecretSHA1 = fmt.Sprintf("%s;%s", objStore.SecretSHA1, caHash)
	objStore.TLS = &storage.TLSConfig{CA: cm.Name, Key: caKey}

	return objStore, nil
}

func allowStructuredMetadata(schemas []logstorev1.ObjectStorageSchema, now time.Time) bool {
	activeVersion := logstorev1.ObjectStorageSchemaV11
	for _, s := range schemas {
		time, _ := s.EffectiveDate.UTCTime()
		if time.Before(now) {
			activeVersion = s.Version
		}
	}

	return activeVersion != logstorev1.ObjectStorageSchemaV11 &&
		activeVersion != logstorev1.ObjectStorageSchemaV12
}
