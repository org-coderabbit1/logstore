package manifests

import (
	"testing"

	"github.com/stretchr/testify/require"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/manifests/storage"
)

func TestNewCompactorStatefulSet_SelectorMatchesLabels(t *testing.T) {
	// You must set the .spec.selector field of a StatefulSet to match the labels of
	// its .spec.template.metadata.labels. Prior to Kubernetes 1.8, the
	// .spec.selector field was defaulted when omitted. In 1.8 and later versions,
	// failing to specify a matching Pod Selector will result in a validation error
	// during StatefulSet creation.
	// See https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#pod-selector
	sts := NewCompactorStatefulSet(Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Compactor: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	})

	l := sts.Spec.Template.GetObjectMeta().GetLabels()
	for key, value := range sts.Spec.Selector.MatchLabels {
		require.Contains(t, l, key)
		require.Equal(t, l[key], value)
	}
}

func TestNewCompactorStatefulSet_HasTemplateConfigHashAnnotation(t *testing.T) {
	ss := NewCompactorStatefulSet(Options{
		Name:       "abcd",
		Namespace:  "efgh",
		ConfigSHA1: "deadbeef",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Compactor: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	})

	annotations := ss.Spec.Template.Annotations
	require.Contains(t, annotations, AnnotationLogstoreConfigHash)
	require.Equal(t, annotations[AnnotationLogstoreConfigHash], "deadbeef")
}

func TestNewCompactorStatefulSet_HasTemplateObjectStorageHashAnnotation(t *testing.T) {
	ss := NewCompactorStatefulSet(Options{
		Name:      "abcd",
		Namespace: "efgh",
		ObjectStorage: storage.Options{
			SecretSHA1: "deadbeef",
		},
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Compactor: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	})

	annotations := ss.Spec.Template.Annotations
	require.Contains(t, annotations, AnnotationLogstoreObjectStoreHash)
	require.Equal(t, annotations[AnnotationLogstoreObjectStoreHash], "deadbeef")
}

func TestNewCompactorStatefulSet_HasTemplateCertRotationRequiredAtAnnotation(t *testing.T) {
	ss := NewCompactorStatefulSet(Options{
		Name:                   "abcd",
		Namespace:              "efgh",
		CertRotationRequiredAt: "deadbeef",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Compactor: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	})

	annotations := ss.Spec.Template.Annotations
	require.Contains(t, annotations, AnnotationCertRotationRequiredAt)
	require.Equal(t, annotations[AnnotationCertRotationRequiredAt], "deadbeef")
}
