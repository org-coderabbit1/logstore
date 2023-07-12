package manifests

import (
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
)

func TestNewIndexGatewayStatefulSet_HasTemplateConfigHashAnnotation(t *testing.T) {
	ss := NewIndexGatewayStatefulSet(Options{
		Name:       "abcd",
		Namespace:  "efgh",
		ConfigSHA1: "deadbeef",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	})

	expected := "logstore.acme.com/config-hash"
	annotations := ss.Spec.Template.Annotations
	require.Contains(t, annotations, expected)
	require.Equal(t, annotations[expected], "deadbeef")
}

func TestNewIndexGatewayStatefulSet_HasTemplateCertRotationRequiredAtAnnotation(t *testing.T) {
	ss := NewIndexGatewayStatefulSet(Options{
		Name:                   "abcd",
		Namespace:              "efgh",
		CertRotationRequiredAt: "deadbeef",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	})
	expected := "logstore.acme.com/certRotationRequiredAt"
	annotations := ss.Spec.Template.Annotations
	require.Contains(t, annotations, expected)
	require.Equal(t, annotations[expected], "deadbeef")
}

func TestNewIndexGatewayStatefulSet_SelectorMatchesLabels(t *testing.T) {
	// You must set the .spec.selector field of a StatefulSet to match the labels of
	// its .spec.template.metadata.labels. Prior to Kubernetes 1.8, the
	// .spec.selector field was defaulted when omitted. In 1.8 and later versions,
	// failing to specify a matching Pod Selector will result in a validation error
	// during StatefulSet creation.
	// See https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#pod-selector
	ss := NewIndexGatewayStatefulSet(Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	})

	l := ss.Spec.Template.GetObjectMeta().GetLabels()
	for key, value := range ss.Spec.Selector.MatchLabels {
		require.Contains(t, l, key)
		require.Equal(t, l[key], value)
	}
}

func TestBuildIndexGateway_PodDisruptionBudget(t *testing.T) {
	opts := Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	}
	objs, err := BuildIndexGateway(opts)

	require.NoError(t, err)
	require.Len(t, objs, 4)

	pdb := objs[3].(*policyv1.PodDisruptionBudget)
	require.NotNil(t, pdb)
	require.Equal(t, "abcd-index-gateway", pdb.Name)
	require.Equal(t, "efgh", pdb.Namespace)
	require.NotNil(t, pdb.Spec.MinAvailable.IntVal)
	require.Equal(t, int32(1), pdb.Spec.MinAvailable.IntVal)
	require.EqualValues(t, ComponentLabels(LabelIndexGatewayComponent, opts.Name),
		pdb.Spec.Selector.MatchLabels)
}

func TestNewIndexGatewayStatefulSet_TopologySpreadConstraints(t *testing.T) {
	depl := NewIndexGatewayStatefulSet(Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
			Replication: &logstorev1.ReplicationSpec{
				Zones: []logstorev1.ZoneSpec{
					{
						TopologyKey: "zone",
						MaxSkew:     3,
					},
					{
						TopologyKey: "region",
						MaxSkew:     2,
					},
				},
				Factor: 1,
			},
		},
	})

	require.Equal(t, []corev1.TopologySpreadConstraint{
		{
			MaxSkew:           3,
			TopologyKey:       "zone",
			WhenUnsatisfiable: "DoNotSchedule",
			LabelSelector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/component": "index-gateway",
					"app.kubernetes.io/instance":  "abcd",
				},
			},
		},
		{
			MaxSkew:           2,
			TopologyKey:       "region",
			WhenUnsatisfiable: "DoNotSchedule",
			LabelSelector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/component": "index-gateway",
					"app.kubernetes.io/instance":  "abcd",
				},
			},
		},
	}, depl.Spec.Template.Spec.TopologySpreadConstraints)
}
