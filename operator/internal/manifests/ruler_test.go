package manifests_test

import (
	"math/rand"
	"testing"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/manifests"
	"example.com/acme/logstore/operator/internal/manifests/openshift"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
)

func TestNewRulerStatefulSet_HasTemplateConfigHashAnnotation(t *testing.T) {
	ss := manifests.NewRulerStatefulSet(manifests.Options{
		Name:       "abcd",
		Namespace:  "efgh",
		ConfigSHA1: "deadbeef",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Ruler: &logstorev1.LogstoreComponentSpec{
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

func TestNewRulerStatefulSet_HasTemplateCertRotationRequiredAtAnnotation(t *testing.T) {
	ss := manifests.NewRulerStatefulSet(manifests.Options{
		Name:                   "abcd",
		Namespace:              "efgh",
		CertRotationRequiredAt: "deadbeef",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Ruler: &logstorev1.LogstoreComponentSpec{
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

func TestBuildRuler_HasExtraObjectsForTenantMode(t *testing.T) {
	objs, err := manifests.BuildRuler(manifests.Options{
		Name:      "abcd",
		Namespace: "efgh",
		OpenShiftOptions: openshift.Options{
			BuildOpts: openshift.BuildOptions{
				LogstoreStackName:      "abc",
				LogstoreStackNamespace: "efgh",
			},
		},
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				Ruler: &logstorev1.LogstoreComponentSpec{
					Replicas: rand.Int31(),
				},
			},
			Rules: &logstorev1.RulesSpec{
				Enabled: true,
			},
			Tenants: &logstorev1.TenantsSpec{
				Mode: logstorev1.OpenshiftLogging,
			},
		},
	})

	require.NoError(t, err)
	require.Len(t, objs, 7)
}

func TestNewRulerStatefulSet_SelectorMatchesLabels(t *testing.T) {
	// You must set the .spec.selector field of a StatefulSet to match the labels of
	// its .spec.template.metadata.labels. Prior to Kubernetes 1.8, the
	// .spec.selector field was defaulted when omitted. In 1.8 and later versions,
	// failing to specify a matching Pod Selector will result in a validation error
	// during StatefulSet creation.
	// See https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#pod-selector
	sts := manifests.NewRulerStatefulSet(manifests.Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Ruler: &logstorev1.LogstoreComponentSpec{
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

func TestNewRulerStatefulSet_MountsRulesInPerTenantIDSubDirectories(t *testing.T) {
	opts := manifests.Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: logstorev1.LogstoreStackSpec{
			StorageClassName: "standard",
			Template: &logstorev1.LogstoreTemplateSpec{
				Ruler: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
		Tenants: manifests.Tenants{
			Configs: map[string]manifests.TenantConfig{
				"tenant-a": {RuleFiles: []string{"test-rules-0___tenant-a___rule-a-alerts.yaml", "test-rules-0___tenant-a___rule-b-recs.yaml"}},
				"tenant-b": {RuleFiles: []string{"test-rules-0___tenant-b___rule-a-alerts.yaml", "test-rules-0___tenant-b___rule-b-recs.yaml"}},
			},
		},
		RulesConfigMapNames: []string{"config"},
	}
	sts := manifests.NewRulerStatefulSet(opts)

	vs := sts.Spec.Template.Spec.Volumes

	var volumeNames []string
	for _, v := range vs {
		volumeNames = append(volumeNames, v.Name)
	}

	require.NotEmpty(t, volumeNames)
}

func TestNewRulerStatefulSet_ShardedRulesConfigMap(t *testing.T) {
	// Create a large config map which will be split into 2 shards
	opts := testOptions_withSharding()
	rulesCMShards, err := manifests.RulesConfigMapShards(opts)
	require.NoError(t, err)
	require.NotNil(t, rulesCMShards)
	require.Len(t, rulesCMShards, 2)

	for _, shard := range rulesCMShards {
		opts.RulesConfigMapNames = append(opts.RulesConfigMapNames, shard.Name)
	}

	// Create the Ruler StatefulSet and mount the ConfigMap shards into the Ruler pod
	sts := manifests.NewRulerStatefulSet(*opts)

	vs := sts.Spec.Template.Spec.Volumes

	var volumeNames []string
	var volumeProjections []corev1.VolumeProjection
	for _, v := range vs {
		volumeNames = append(volumeNames, v.Name)
		if v.Name == manifests.RulesStorageVolumeName() {
			volumeProjections = append(volumeProjections, v.Projected.Sources...)
		}
	}

	require.Len(t, volumeNames, 2)
	require.Len(t, volumeProjections, 2)
}
