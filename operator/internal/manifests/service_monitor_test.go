package manifests

import (
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"

	configv1 "example.com/acme/logstore/operator/apis/config/v1"
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test that all serviceMonitor match the labels of their services so that we know all serviceMonitor
// will work when deployed.
func TestServiceMonitorMatchLabels(t *testing.T) {
	type test struct {
		Service        *corev1.Service
		ServiceMonitor *monitoringv1.ServiceMonitor
	}

	featureGates := configv1.FeatureGates{
		BuiltInCertManagement:      configv1.BuiltInCertManagement{Enabled: true},
		ServiceMonitors:            true,
		ServiceMonitorTLSEndpoints: true,
	}

	opt := Options{
		Name:      "test",
		Namespace: "test",
		Image:     "test",
		Gates:     featureGates,
		Stack: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
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
				Gateway: &logstorev1.LogstoreComponentSpec{
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
	}

	table := []test{
		{
			Service:        NewDistributorHTTPService(opt),
			ServiceMonitor: NewDistributorServiceMonitor(opt),
		},
		{
			Service:        NewIngesterHTTPService(opt),
			ServiceMonitor: NewIngesterServiceMonitor(opt),
		},
		{
			Service:        NewQuerierHTTPService(opt),
			ServiceMonitor: NewQuerierServiceMonitor(opt),
		},
		{
			Service:        NewQueryFrontendHTTPService(opt),
			ServiceMonitor: NewQueryFrontendServiceMonitor(opt),
		},
		{
			Service:        NewCompactorHTTPService(opt),
			ServiceMonitor: NewCompactorServiceMonitor(opt),
		},
		{
			Service:        NewGatewayHTTPService(opt),
			ServiceMonitor: NewGatewayServiceMonitor(opt),
		},
		{
			Service:        NewIndexGatewayHTTPService(opt),
			ServiceMonitor: NewIndexGatewayServiceMonitor(opt),
		},
		{
			Service:        NewRulerHTTPService(opt),
			ServiceMonitor: NewRulerServiceMonitor(opt),
		},
	}

	for _, tst := range table {
		testName := fmt.Sprintf("%s_%s", tst.Service.GetName(), tst.ServiceMonitor.GetName())
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			for k, v := range tst.ServiceMonitor.Spec.Selector.MatchLabels {
				if assert.Contains(t, tst.Service.Spec.Selector, k) {
					// only assert Equal if the previous assertion is successful or this will panic
					assert.Equal(t, v, tst.Service.Spec.Selector[k])
				}
			}
		})
	}
}

func TestServiceMonitorEndpoints_ForBuiltInCertRotation(t *testing.T) {
	type test struct {
		Service        *corev1.Service
		ServiceMonitor *monitoringv1.ServiceMonitor
	}

	featureGates := configv1.FeatureGates{
		BuiltInCertManagement:      configv1.BuiltInCertManagement{Enabled: true},
		ServiceMonitors:            true,
		ServiceMonitorTLSEndpoints: true,
	}

	opt := Options{
		Name:      "test",
		Namespace: "test",
		Image:     "test",
		Gates:     featureGates,
		Stack: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
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
	}

	table := []test{
		{
			Service:        NewDistributorHTTPService(opt),
			ServiceMonitor: NewDistributorServiceMonitor(opt),
		},
		{
			Service:        NewIngesterHTTPService(opt),
			ServiceMonitor: NewIngesterServiceMonitor(opt),
		},
		{
			Service:        NewQuerierHTTPService(opt),
			ServiceMonitor: NewQuerierServiceMonitor(opt),
		},
		{
			Service:        NewQueryFrontendHTTPService(opt),
			ServiceMonitor: NewQueryFrontendServiceMonitor(opt),
		},
		{
			Service:        NewCompactorHTTPService(opt),
			ServiceMonitor: NewCompactorServiceMonitor(opt),
		},
		{
			Service:        NewIndexGatewayHTTPService(opt),
			ServiceMonitor: NewIndexGatewayServiceMonitor(opt),
		},
		{
			Service:        NewRulerHTTPService(opt),
			ServiceMonitor: NewRulerServiceMonitor(opt),
		},
	}

	for _, tst := range table {
		testName := fmt.Sprintf("%s_%s", tst.Service.GetName(), tst.ServiceMonitor.GetName())
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			require.NotNil(t, tst.ServiceMonitor.Spec.Endpoints)
			require.NotNil(t, tst.ServiceMonitor.Spec.Endpoints[0].TLSConfig)

			// Do not use bearer authentication for logstore endpoints
			require.Empty(t, tst.ServiceMonitor.Spec.Endpoints[0].BearerTokenFile)
			require.Empty(t, tst.ServiceMonitor.Spec.Endpoints[0].BearerTokenSecret)

			// Check using built-in PKI
			c := tst.ServiceMonitor.Spec.Endpoints[0].TLSConfig
			require.Equal(t, c.CA.ConfigMap.LocalObjectReference.Name, signingCABundleName(opt.Name))
			require.Equal(t, c.Cert.Secret.LocalObjectReference.Name, tst.Service.Name)
			require.Equal(t, c.KeySecret.LocalObjectReference.Name, tst.Service.Name)
		})
	}
}

func TestServiceMonitorEndpoints_ForGatewayServiceMonitor(t *testing.T) {
	tt := []struct {
		desc  string
		opts  Options
		total int
		want  []monitoringv1.Endpoint
	}{
		{
			desc: "default",
			opts: Options{
				Name:      "test",
				Namespace: "test",
				Image:     "test",
				Stack: logstorev1.LogstoreStackSpec{
					Size: logstorev1.SizeOneXExtraSmall,
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Static,
					},
					Template: &logstorev1.LogstoreTemplateSpec{
						Gateway: &logstorev1.LogstoreComponentSpec{
							Replicas: 1,
						},
					},
				},
			},
			total: 1,
			want: []monitoringv1.Endpoint{
				{
					Port:   gatewayInternalPortName,
					Path:   "/metrics",
					Scheme: "http",
				},
			},
		},
		{
			desc: "with http encryption",
			opts: Options{
				Name:      "test",
				Namespace: "test",
				Image:     "test",
				Gates: configv1.FeatureGates{
					LogstoreStackGateway:           true,
					BuiltInCertManagement:      configv1.BuiltInCertManagement{Enabled: true},
					ServiceMonitors:            true,
					ServiceMonitorTLSEndpoints: true,
				},
				Stack: logstorev1.LogstoreStackSpec{
					Size: logstorev1.SizeOneXExtraSmall,
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Static,
					},
					Template: &logstorev1.LogstoreTemplateSpec{
						Gateway: &logstorev1.LogstoreComponentSpec{
							Replicas: 1,
						},
					},
				},
			},
			total: 1,
			want: []monitoringv1.Endpoint{
				{
					Port:   gatewayInternalPortName,
					Path:   "/metrics",
					Scheme: "https",
					BearerTokenSecret: corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "test-gateway-token",
						},
						Key: corev1.ServiceAccountTokenKey,
					},
					TLSConfig: &monitoringv1.TLSConfig{
						SafeTLSConfig: monitoringv1.SafeTLSConfig{
							CA: monitoringv1.SecretOrConfigMap{
								ConfigMap: &corev1.ConfigMapKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: gatewaySigningCABundleName("test-gateway"),
									},
									Key: caFile,
								},
							},
							ServerName: "test-gateway-http.test.svc.cluster.local",
						},
					},
				},
			},
		},
		{
			desc: "openshift-logging",
			opts: Options{
				Name:      "test",
				Namespace: "test",
				Image:     "test",
				Stack: logstorev1.LogstoreStackSpec{
					Size: logstorev1.SizeOneXExtraSmall,
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
					Template: &logstorev1.LogstoreTemplateSpec{
						Gateway: &logstorev1.LogstoreComponentSpec{
							Replicas: 1,
						},
					},
				},
			},
			total: 2,
			want: []monitoringv1.Endpoint{
				{
					Port:   gatewayInternalPortName,
					Path:   "/metrics",
					Scheme: "http",
				},
				{
					Port:   "opa-metrics",
					Path:   "/metrics",
					Scheme: "http",
				},
			},
		},
		{
			desc: "openshift-logging with http encryption",
			opts: Options{
				Name:      "test",
				Namespace: "test",
				Image:     "test",
				Gates: configv1.FeatureGates{
					LogstoreStackGateway:           true,
					BuiltInCertManagement:      configv1.BuiltInCertManagement{Enabled: true},
					ServiceMonitors:            true,
					ServiceMonitorTLSEndpoints: true,
				},
				Stack: logstorev1.LogstoreStackSpec{
					Size: logstorev1.SizeOneXExtraSmall,
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
					Template: &logstorev1.LogstoreTemplateSpec{
						Gateway: &logstorev1.LogstoreComponentSpec{
							Replicas: 1,
						},
					},
				},
			},
			total: 2,
			want: []monitoringv1.Endpoint{
				{
					Port:   gatewayInternalPortName,
					Path:   "/metrics",
					Scheme: "https",
					BearerTokenSecret: corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "test-gateway-token",
						},
						Key: corev1.ServiceAccountTokenKey,
					},
					TLSConfig: &monitoringv1.TLSConfig{
						SafeTLSConfig: monitoringv1.SafeTLSConfig{
							CA: monitoringv1.SecretOrConfigMap{
								ConfigMap: &corev1.ConfigMapKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: gatewaySigningCABundleName("test-gateway"),
									},
									Key: caFile,
								},
							},
							ServerName: "test-gateway-http.test.svc.cluster.local",
						},
					},
				},
				{
					Port:   "opa-metrics",
					Path:   "/metrics",
					Scheme: "https",
					BearerTokenSecret: corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "test-gateway-token",
						},
						Key: corev1.ServiceAccountTokenKey,
					},
					TLSConfig: &monitoringv1.TLSConfig{
						SafeTLSConfig: monitoringv1.SafeTLSConfig{
							CA: monitoringv1.SecretOrConfigMap{
								ConfigMap: &corev1.ConfigMapKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: gatewaySigningCABundleName("test-gateway"),
									},
									Key: caFile,
								},
							},
							ServerName: "test-gateway-http.test.svc.cluster.local",
						},
					},
				},
			},
		},
		{
			desc: "openshift-network",
			opts: Options{
				Name:      "test",
				Namespace: "test",
				Image:     "test",
				Stack: logstorev1.LogstoreStackSpec{
					Size: logstorev1.SizeOneXExtraSmall,
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftNetwork,
					},
					Template: &logstorev1.LogstoreTemplateSpec{
						Gateway: &logstorev1.LogstoreComponentSpec{
							Replicas: 1,
						},
					},
				},
			},
			total: 2,
			want: []monitoringv1.Endpoint{
				{
					Port:   gatewayInternalPortName,
					Path:   "/metrics",
					Scheme: "http",
				},
				{
					Port:   "opa-metrics",
					Path:   "/metrics",
					Scheme: "http",
				},
			},
		},
		{
			desc: "openshift-network with http encryption",
			opts: Options{
				Name:      "test",
				Namespace: "test",
				Image:     "test",
				Gates: configv1.FeatureGates{
					LogstoreStackGateway:           true,
					BuiltInCertManagement:      configv1.BuiltInCertManagement{Enabled: true},
					ServiceMonitors:            true,
					ServiceMonitorTLSEndpoints: true,
				},
				Stack: logstorev1.LogstoreStackSpec{
					Size: logstorev1.SizeOneXExtraSmall,
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftNetwork,
					},
					Template: &logstorev1.LogstoreTemplateSpec{
						Gateway: &logstorev1.LogstoreComponentSpec{
							Replicas: 1,
						},
					},
				},
			},
			total: 2,
			want: []monitoringv1.Endpoint{
				{
					Port:   gatewayInternalPortName,
					Path:   "/metrics",
					Scheme: "https",
					BearerTokenSecret: corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "test-gateway-token",
						},
						Key: corev1.ServiceAccountTokenKey,
					},
					TLSConfig: &monitoringv1.TLSConfig{
						SafeTLSConfig: monitoringv1.SafeTLSConfig{
							CA: monitoringv1.SecretOrConfigMap{
								ConfigMap: &corev1.ConfigMapKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: gatewaySigningCABundleName("test-gateway"),
									},
									Key: caFile,
								},
							},
							ServerName: "test-gateway-http.test.svc.cluster.local",
						},
					},
				},
				{
					Port:   "opa-metrics",
					Path:   "/metrics",
					Scheme: "https",
					BearerTokenSecret: corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "test-gateway-token",
						},
						Key: corev1.ServiceAccountTokenKey,
					},
					TLSConfig: &monitoringv1.TLSConfig{
						SafeTLSConfig: monitoringv1.SafeTLSConfig{
							CA: monitoringv1.SecretOrConfigMap{
								ConfigMap: &corev1.ConfigMapKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: gatewaySigningCABundleName("test-gateway"),
									},
									Key: caFile,
								},
							},
							ServerName: "test-gateway-http.test.svc.cluster.local",
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			sm := NewGatewayServiceMonitor(tc.opts)
			require.Len(t, sm.Spec.Endpoints, tc.total)

			for _, endpoint := range tc.want {
				require.Contains(t, sm.Spec.Endpoints, endpoint)
			}
		})
	}
}
