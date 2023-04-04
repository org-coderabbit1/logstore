package manifests

import (
	"testing"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/manifests/storage"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func TestTolerationsAreSetForEachComponent(t *testing.T) {
	tolerations := []corev1.Toleration{{
		Key:      "type",
		Operator: corev1.TolerationOpEqual,
		Value:    "storage",
		Effect:   corev1.TaintEffectNoSchedule,
	}}
	optsWithTolerations := Options{
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				Compactor: &logstorev1.LogstoreComponentSpec{
					Tolerations: tolerations,
					Replicas:    1,
				},
				Distributor: &logstorev1.LogstoreComponentSpec{
					Tolerations: tolerations,
					Replicas:    1,
				},
				Ingester: &logstorev1.LogstoreComponentSpec{
					Tolerations: tolerations,
					Replicas:    1,
				},
				Querier: &logstorev1.LogstoreComponentSpec{
					Tolerations: tolerations,
					Replicas:    1,
				},
				QueryFrontend: &logstorev1.LogstoreComponentSpec{
					Tolerations: tolerations,
					Replicas:    1,
				},
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Tolerations: tolerations,
					Replicas:    1,
				},
				Ruler: &logstorev1.LogstoreComponentSpec{
					Tolerations: tolerations,
					Replicas:    1,
				},
			},
		},
		ObjectStorage: storage.Options{},
	}

	optsWithoutTolerations := Options{
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				Compactor: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				Distributor: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				Ingester: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				Querier: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				QueryFrontend: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				Ruler: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
		ObjectStorage: storage.Options{},
	}

	t.Run("distributor", func(t *testing.T) {
		assert.Equal(t, tolerations, NewDistributorDeployment(optsWithTolerations).Spec.Template.Spec.Tolerations)
		assert.Empty(t, NewDistributorDeployment(optsWithoutTolerations).Spec.Template.Spec.Tolerations)
	})

	t.Run("query_frontend", func(t *testing.T) {
		assert.Equal(t, tolerations, NewQueryFrontendDeployment(optsWithTolerations).Spec.Template.Spec.Tolerations)
		assert.Empty(t, NewQueryFrontendDeployment(optsWithoutTolerations).Spec.Template.Spec.Tolerations)
	})

	t.Run("querier", func(t *testing.T) {
		assert.Equal(t, tolerations, NewQuerierDeployment(optsWithTolerations).Spec.Template.Spec.Tolerations)
		assert.Empty(t, NewQuerierDeployment(optsWithoutTolerations).Spec.Template.Spec.Tolerations)
	})

	t.Run("ingester", func(t *testing.T) {
		assert.Equal(t, tolerations, NewIngesterStatefulSet(optsWithTolerations).Spec.Template.Spec.Tolerations)
		assert.Empty(t, NewIngesterStatefulSet(optsWithoutTolerations).Spec.Template.Spec.Tolerations)
	})

	t.Run("compactor", func(t *testing.T) {
		assert.Equal(t, tolerations, NewCompactorStatefulSet(optsWithTolerations).Spec.Template.Spec.Tolerations)
		assert.Empty(t, NewCompactorStatefulSet(optsWithoutTolerations).Spec.Template.Spec.Tolerations)
	})

	t.Run("index_gateway", func(t *testing.T) {
		assert.Equal(t, tolerations, NewIndexGatewayStatefulSet(optsWithTolerations).Spec.Template.Spec.Tolerations)
		assert.Empty(t, NewIndexGatewayStatefulSet(optsWithoutTolerations).Spec.Template.Spec.Tolerations)
	})

	t.Run("ruler", func(t *testing.T) {
		assert.Equal(t, tolerations, NewRulerStatefulSet(optsWithTolerations).Spec.Template.Spec.Tolerations)
		assert.Empty(t, NewRulerStatefulSet(optsWithoutTolerations).Spec.Template.Spec.Tolerations)
	})
}

func TestNodeSelectorsAreSetForEachComponent(t *testing.T) {
	nodeSelectors := map[string]string{"type": "storage"}
	optsWithNodeSelectors := Options{
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				Compactor: &logstorev1.LogstoreComponentSpec{
					NodeSelector: nodeSelectors,
					Replicas:     1,
				},
				Distributor: &logstorev1.LogstoreComponentSpec{
					NodeSelector: nodeSelectors,
					Replicas:     1,
				},
				Ingester: &logstorev1.LogstoreComponentSpec{
					NodeSelector: nodeSelectors,
					Replicas:     1,
				},
				Querier: &logstorev1.LogstoreComponentSpec{
					NodeSelector: nodeSelectors,
					Replicas:     1,
				},
				QueryFrontend: &logstorev1.LogstoreComponentSpec{
					NodeSelector: nodeSelectors,
					Replicas:     1,
				},
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					NodeSelector: nodeSelectors,
					Replicas:     1,
				},
				Ruler: &logstorev1.LogstoreComponentSpec{
					NodeSelector: nodeSelectors,
					Replicas:     1,
				},
			},
		},
		ObjectStorage: storage.Options{},
	}

	optsWithoutNodeSelectors := Options{
		Stack: logstorev1.LogstoreStackSpec{
			Template: &logstorev1.LogstoreTemplateSpec{
				Compactor: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				Distributor: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				Ingester: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				Querier: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				QueryFrontend: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
				Ruler: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
				},
			},
		},
		ObjectStorage: storage.Options{},
	}

	t.Run("distributor", func(t *testing.T) {
		assert.Equal(t, nodeSelectors, NewDistributorDeployment(optsWithNodeSelectors).Spec.Template.Spec.NodeSelector)
		assert.Empty(t, NewDistributorDeployment(optsWithoutNodeSelectors).Spec.Template.Spec.NodeSelector)
	})

	t.Run("query_frontend", func(t *testing.T) {
		assert.Equal(t, nodeSelectors, NewQueryFrontendDeployment(optsWithNodeSelectors).Spec.Template.Spec.NodeSelector)
		assert.Empty(t, NewQueryFrontendDeployment(optsWithoutNodeSelectors).Spec.Template.Spec.NodeSelector)
	})

	t.Run("querier", func(t *testing.T) {
		assert.Equal(t, nodeSelectors, NewQuerierDeployment(optsWithNodeSelectors).Spec.Template.Spec.NodeSelector)
		assert.Empty(t, NewQuerierDeployment(optsWithoutNodeSelectors).Spec.Template.Spec.NodeSelector)
	})

	t.Run("ingester", func(t *testing.T) {
		assert.Equal(t, nodeSelectors, NewIngesterStatefulSet(optsWithNodeSelectors).Spec.Template.Spec.NodeSelector)
		assert.Empty(t, NewIngesterStatefulSet(optsWithoutNodeSelectors).Spec.Template.Spec.NodeSelector)
	})

	t.Run("compactor", func(t *testing.T) {
		assert.Equal(t, nodeSelectors, NewCompactorStatefulSet(optsWithNodeSelectors).Spec.Template.Spec.NodeSelector)
		assert.Empty(t, NewCompactorStatefulSet(optsWithoutNodeSelectors).Spec.Template.Spec.NodeSelector)
	})

	t.Run("index_gateway", func(t *testing.T) {
		assert.Equal(t, nodeSelectors, NewIndexGatewayStatefulSet(optsWithNodeSelectors).Spec.Template.Spec.NodeSelector)
		assert.Empty(t, NewIndexGatewayStatefulSet(optsWithoutNodeSelectors).Spec.Template.Spec.NodeSelector)
	})

	t.Run("ruler", func(t *testing.T) {
		assert.Equal(t, nodeSelectors, NewRulerStatefulSet(optsWithNodeSelectors).Spec.Template.Spec.NodeSelector)
		assert.Empty(t, NewRulerStatefulSet(optsWithoutNodeSelectors).Spec.Template.Spec.NodeSelector)
	})
}
