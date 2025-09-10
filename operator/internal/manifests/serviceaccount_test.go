package manifests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/manifests/storage"
)

func TestServiceAccountName_MatchesPodSpecServiceAccountName(t *testing.T) {
	opts := Options{
		Name:      "logstorestack",
		Namespace: "ns",
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

	sa := BuildServiceAccount(opts)

	t.Run("distributor", func(t *testing.T) {
		assert.Equal(t, sa.GetName(), NewDistributorDeployment(opts).Spec.Template.Spec.ServiceAccountName)
	})

	t.Run("query_frontend", func(t *testing.T) {
		assert.Equal(t, sa.GetName(), NewQueryFrontendDeployment(opts).Spec.Template.Spec.ServiceAccountName)
	})

	t.Run("querier", func(t *testing.T) {
		assert.Equal(t, sa.GetName(), NewQuerierDeployment(opts).Spec.Template.Spec.ServiceAccountName)
	})

	t.Run("ingester", func(t *testing.T) {
		assert.Equal(t, sa.GetName(), NewIngesterStatefulSet(opts).Spec.Template.Spec.ServiceAccountName)
	})

	t.Run("compactor", func(t *testing.T) {
		assert.Equal(t, sa.GetName(), NewCompactorStatefulSet(opts).Spec.Template.Spec.ServiceAccountName)
	})

	t.Run("index_gateway", func(t *testing.T) {
		assert.Equal(t, sa.GetName(), NewIndexGatewayStatefulSet(opts).Spec.Template.Spec.ServiceAccountName)
	})
}
