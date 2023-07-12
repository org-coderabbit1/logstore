package manifests

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "example.com/acme/logstore/operator/apis/config/v1"
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/manifests/internal"
)

func TestNewIngesterStatefulSet_HasTemplateConfigHashAnnotation(t *testing.T) {
	ss := NewIngesterStatefulSet(Options{
		Name:       "abcd",
		Namespace:  "efgh",
		ConfigSHA1: "deadbeef",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Ingester: &logstorev1.LogstoreComponentSpec{
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

func TestNewIngesterStatefulSet_HasTemplateCertRotationRequiredAtAnnotation(t *testing.T) {
	ss := NewIngesterStatefulSet(Options{
		Name:                   "abcd",
		Namespace:              "efgh",
		CertRotationRequiredAt: "deadbeef",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Ingester: &logstorev1.LogstoreComponentSpec{
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

func TestNewIngesterStatefulSet_SelectorMatchesLabels(t *testing.T) {
	// You must set the .spec.selector field of a StatefulSet to match the labels of
	// its .spec.template.metadata.labels. Prior to Kubernetes 1.8, the
	// .spec.selector field was defaulted when omitted. In 1.8 and later versions,
	// failing to specify a matching Pod Selector will result in a validation error
	// during StatefulSet creation.
	// See https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#pod-selector
	sts := NewIngesterStatefulSet(Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Ingester: &logstorev1.LogstoreComponentSpec{
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

func TestBuildIngester_PodDisruptionBudget(t *testing.T) {
	for _, tc := range []struct {
		Name                 string
		PDBMinAvailable      int
		ExpectedMinAvailable int
	}{
		{
			Name:                 "Small stack",
			PDBMinAvailable:      1,
			ExpectedMinAvailable: 1,
		},
		{
			Name:                 "Medium stack",
			PDBMinAvailable:      2,
			ExpectedMinAvailable: 2,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			opts := Options{
				Name:      "abcd",
				Namespace: "efgh",
				Gates:     v1.FeatureGates{},
				ResourceRequirements: internal.ComponentResources{
					Ingester: internal.ResourceRequirements{
						PDBMinAvailable: tc.PDBMinAvailable,
					},
				},
				Stack: logstorev1.LogstoreStackSpec{
					Template: &logstorev1.LogstoreTemplateSpec{
						Ingester: &logstorev1.LogstoreComponentSpec{
							Replicas: rand.Int31(),
						},
					},
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
				},
			}
			objs, err := BuildIngester(opts)
			require.NoError(t, err)
			require.Len(t, objs, 4)

			pdb := objs[3].(*policyv1.PodDisruptionBudget)
			require.NotNil(t, pdb)
			require.Equal(t, "abcd-ingester", pdb.Name)
			require.Equal(t, "efgh", pdb.Namespace)
			require.NotNil(t, pdb.Spec.MinAvailable.IntVal)
			require.Equal(t, int32(tc.ExpectedMinAvailable), pdb.Spec.MinAvailable.IntVal)
			require.EqualValues(t, ComponentLabels(LabelIngesterComponent, opts.Name), pdb.Spec.Selector.MatchLabels)
		})
	}
}

func TestNewIngesterStatefulSet_TopologySpreadConstraints(t *testing.T) {
	ss := NewIngesterStatefulSet(Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				Ingester: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
			Replication: &logstorev1.ReplicationSpec{
				Zones: []logstorev1.ZoneSpec{
					{
						TopologyKey: "zone",
						MaxSkew:     2,
					},
					{
						TopologyKey: "region",
						MaxSkew:     1,
					},
				},
				Factor: 1,
			},
		},
	})

	require.Equal(t, []corev1.TopologySpreadConstraint{
		{
			MaxSkew:           2,
			TopologyKey:       "zone",
			WhenUnsatisfiable: "DoNotSchedule",
			LabelSelector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/component": "ingester",
					"app.kubernetes.io/instance":  "abcd",
				},
			},
		},
		{
			MaxSkew:           1,
			TopologyKey:       "region",
			WhenUnsatisfiable: "DoNotSchedule",
			LabelSelector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/component": "ingester",
					"app.kubernetes.io/instance":  "abcd",
				},
			},
		},
	}, ss.Spec.Template.Spec.TopologySpreadConstraints)
}
