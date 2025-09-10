package gateway

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s/k8sfakes"
	"example.com/acme/logstore/operator/internal/manifests"
)

var tenantConfigData = []byte(`
tenants:
- name: application
  id: test-123
  openshift:
    serviceAccount: logstorestack-dev-gateway
    cookieSecret: test123
- name: infrastructure
  id: test-456
  openshift:
    serviceAccount: logstorestack-dev-gateway
    cookieSecret: test456
- name: audit
  id: test-789
  openshift:
    serviceAccount: logstorestack-dev-gateway
    cookieSecret: test789
`)

func TestGetTenantConfigSecretData_SecretExist(t *testing.T) {
	k := &k8sfakes.FakeClient{}
	s := &logstorev1.LogstoreStack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "logstorestack-dev",
			Namespace: "some-ns",
		},
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if name.Name == "logstorestack-dev-gateway" && name.Namespace == "some-ns" {
			k.SetClientObject(object, &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "logstorestack-dev-gateway",
					Namespace: "some-ns",
				},
				Data: map[string][]byte{
					"tenants.yaml": tenantConfigData,
				},
			})
		}
		return nil
	}

	ts, err := getTenantConfigFromSecret(context.TODO(), k, s)
	require.NotNil(t, ts)
	require.NoError(t, err)

	expected := map[string]manifests.TenantConfig{
		"application": {
			OpenShift: &manifests.TenantOpenShiftSpec{
				CookieSecret: "test123",
			},
		},
		"infrastructure": {
			OpenShift: &manifests.TenantOpenShiftSpec{
				CookieSecret: "test456",
			},
		},
		"audit": {
			OpenShift: &manifests.TenantOpenShiftSpec{
				CookieSecret: "test789",
			},
		},
	}
	require.Equal(t, expected, ts)
}

func TestGetTenantConfigSecretData_SecretNotExist(t *testing.T) {
	k := &k8sfakes.FakeClient{}
	s := &logstorev1.LogstoreStack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "logstorestack-dev",
			Namespace: "some-ns",
		},
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	ts, err := getTenantConfigFromSecret(context.TODO(), k, s)
	require.Nil(t, ts)
	require.Error(t, err)
}
