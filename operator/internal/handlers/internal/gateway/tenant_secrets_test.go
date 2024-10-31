package gateway

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s/k8sfakes"
	"example.com/acme/logstore/operator/internal/manifests"
)

func TestGetTenantSecrets(t *testing.T) {
	for _, mode := range []logstorev1.ModeType{logstorev1.Static, logstorev1.Dynamic} {
		for _, tc := range []struct {
			name      string
			authNSpec []logstorev1.AuthenticationSpec
			object    client.Object
			expected  []*manifests.TenantSecrets
		}{
			{
				name: "oidc",
				authNSpec: []logstorev1.AuthenticationSpec{
					{
						TenantName: "test",
						TenantID:   "test",
						OIDC: &logstorev1.OIDCSpec{
							Secret: &logstorev1.TenantSecretSpec{
								Name: "test",
							},
						},
					},
				},
				object: &corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test",
						Namespace: "some-ns",
					},
					Data: map[string][]byte{
						"clientID":     []byte("test"),
						"clientSecret": []byte("test"),
					},
				},
				expected: []*manifests.TenantSecrets{
					{
						TenantName: "test",
						OIDCSecret: &manifests.OIDCSecret{
							ClientID:     "test",
							ClientSecret: "test",
						},
					},
				},
			},
			{
				name: "mTLS",
				authNSpec: []logstorev1.AuthenticationSpec{
					{
						TenantName: "test",
						TenantID:   "test",
						MTLS: &logstorev1.MTLSSpec{
							CA: &logstorev1.CASpec{
								CA:    "test",
								CAKey: "special-ca.crt",
							},
						},
					},
				},
				object: &corev1.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test",
						Namespace: "some-ns",
					},
					Data: map[string]string{
						"special-ca.crt": "my-specila-ca",
					},
				},
				expected: []*manifests.TenantSecrets{
					{
						TenantName: "test",
						MTLSSecret: &manifests.MTLSSecret{
							CAPath: "/var/run/tenants-ca/test/special-ca.crt",
						},
					},
				},
			},
		} {
			t.Run(strings.Join([]string{string(mode), tc.name}, "_"), func(t *testing.T) {
				k := &k8sfakes.FakeClient{}
				s := &logstorev1.LogstoreStack{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "mystack",
						Namespace: "some-ns",
					},
					Spec: logstorev1.LogstoreStackSpec{
						Tenants: &logstorev1.TenantsSpec{
							Mode:           mode,
							Authentication: tc.authNSpec,
						},
					},
				}

				k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
					if name.Name == "test" && name.Namespace == "some-ns" {
						k.SetClientObject(object, tc.object)
					}
					return nil
				}
				ts, err := getTenantSecrets(context.TODO(), k, s)
				require.NoError(t, err)
				require.ElementsMatch(t, ts, tc.expected)
			})
		}
	}
}

func TestExtractOIDCSecret(t *testing.T) {
	type test struct {
		name       string
		tenantName string
		secret     *corev1.Secret
		wantErr    bool
	}
	table := []test{
		{
			name:       "missing clientID",
			tenantName: "tenant-a",
			secret:     &corev1.Secret{},
			wantErr:    true,
		},
		{
			name:       "all set",
			tenantName: "tenant-a",
			secret: &corev1.Secret{
				Data: map[string][]byte{
					"clientID":     []byte("test"),
					"clientSecret": []byte("test"),
				},
			},
		},
	}
	for _, tst := range table {
		t.Run(tst.name, func(t *testing.T) {
			t.Parallel()

			_, err := extractOIDCSecret(tst.secret)
			if !tst.wantErr {
				require.NoError(t, err)
			}
			if tst.wantErr {
				require.NotNil(t, err)
			}
		})
	}
}

func TestCheckKeyIsPresent(t *testing.T) {
	type test struct {
		name       string
		tenantName string
		configMap  *corev1.ConfigMap
		wantErr    bool
	}
	table := []test{
		{
			name:       "missing key",
			tenantName: "tenant-a",
			configMap:  &corev1.ConfigMap{},
			wantErr:    true,
		},
		{
			name:       "all set",
			tenantName: "tenant-a",
			configMap: &corev1.ConfigMap{
				Data: map[string]string{
					"test": "test",
				},
			},
		},
	}
	for _, tst := range table {
		t.Run(tst.name, func(t *testing.T) {
			t.Parallel()

			err := checkKeyIsPresent(tst.configMap, "test")
			if !tst.wantErr {
				require.NoError(t, err)
			}
			if tst.wantErr {
				require.NotNil(t, err)
			}
		})
	}
}
