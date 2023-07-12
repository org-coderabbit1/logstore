package manifests

import (
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
)

func TestNewQueryFrontendDeployment_SelectorMatchesLabels(t *testing.T) {
	ss := NewQueryFrontendDeployment(Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				QueryFrontend: &logstorev1.LogstoreComponentSpec{
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

func TestNewQueryFrontendDeployment_HasTemplateConfigHashAnnotation(t *testing.T) {
	ss := NewQueryFrontendDeployment(Options{
		Name:       "abcd",
		Namespace:  "efgh",
		ConfigSHA1: "deadbeef",
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				QueryFrontend: &logstorev1.LogstoreComponentSpec{
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

func TestNewQueryFrontendDeployment_HasTemplateCertRotationRequiredAtAnnotation(t *testing.T) {
	ss := NewQueryFrontendDeployment(Options{
		Name:                   "abcd",
		Namespace:              "efgh",
		CertRotationRequiredAt: "deadbeef",
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				QueryFrontend: &logstorev1.LogstoreComponentSpec{
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

func TestBuildQueryFrontend_PodDisruptionBudget(t *testing.T) {
	opts := Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				QueryFrontend: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
	}
	objs, err := BuildQueryFrontend(opts)

	require.NoError(t, err)
	require.Len(t, objs, 4)

	pdb := objs[3].(*policyv1.PodDisruptionBudget)
	require.NotNil(t, pdb)
	require.Equal(t, "abcd-query-frontend", pdb.Name)
	require.Equal(t, "efgh", pdb.Namespace)
	require.NotNil(t, pdb.Spec.MinAvailable.IntVal)
	require.Equal(t, int32(1), pdb.Spec.MinAvailable.IntVal)
	require.EqualValues(t, ComponentLabels(LabelQueryFrontendComponent, opts.Name), pdb.Spec.Selector.MatchLabels)
}

func TestNewQueryFrontendDeployment_TopologySpreadConstraints(t *testing.T) {
	depl := NewQueryFrontendDeployment(Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				QueryFrontend: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
			Replication: &logstorev1.ReplicationSpec{
				Zones: []logstorev1.ZoneSpec{
					{
						TopologyKey: "zone",
						MaxSkew:     1,
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
			MaxSkew:           1,
			TopologyKey:       "zone",
			WhenUnsatisfiable: "DoNotSchedule",
			LabelSelector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/component": "query-frontend",
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
					"app.kubernetes.io/component": "query-frontend",
					"app.kubernetes.io/instance":  "abcd",
				},
			},
		},
	}, depl.Spec.Template.Spec.TopologySpreadConstraints)
}
