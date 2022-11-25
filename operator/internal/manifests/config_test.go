package manifests_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/apis/logstore/v1beta1"
	"example.com/acme/logstore/operator/internal/manifests"
	"example.com/acme/logstore/operator/internal/manifests/internal/config"
	"example.com/acme/logstore/operator/internal/manifests/openshift"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
)

func TestConfigMap_ReturnsSHA1OfBinaryContents(t *testing.T) {
	opts := randomConfigOptions()

	_, sha1C, err := manifests.LogstoreConfigMap(opts)
	require.NoError(t, err)
	require.NotEmpty(t, sha1C)
}

func TestConfigOptions_UserOptionsTakePrecedence(t *testing.T) {
	// regardless of what is provided by the default sizing parameters we should always prefer
	// the user-defined values. This creates an all-inclusive manifests.Options and then checks
	// that every value is present in the result
	opts := randomConfigOptions()
	res := manifests.ConfigOptions(opts)

	expected, err := json.Marshal(opts.Stack)
	require.NoError(t, err)

	actual, err := json.Marshal(res.Stack)
	require.NoError(t, err)

	assert.JSONEq(t, string(expected), string(actual))
}

func randomConfigOptions() manifests.Options {
	return manifests.Options{
		Name:      uuid.New().String(),
		Namespace: uuid.New().String(),
		Image:     uuid.New().String(),
		Stack: logstorev1.LogstoreStackSpec{
			Size:              logstorev1.SizeOneXExtraSmall,
			Storage:           logstorev1.ObjectStorageSpec{},
			StorageClassName:  uuid.New().String(),
			ReplicationFactor: rand.Int31(),
			Limits: &logstorev1.LimitsSpec{
				Global: &logstorev1.LimitsTemplateSpec{
					IngestionLimits: &logstorev1.IngestionLimitSpec{
						IngestionRate:             rand.Int31(),
						IngestionBurstSize:        rand.Int31(),
						MaxLabelNameLength:        rand.Int31(),
						MaxLabelValueLength:       rand.Int31(),
						MaxLabelNamesPerSeries:    rand.Int31(),
						MaxGlobalStreamsPerTenant: rand.Int31(),
						MaxLineSize:               rand.Int31(),
					},
					QueryLimits: &logstorev1.QueryLimitSpec{
						MaxEntriesLimitPerQuery: rand.Int31(),
						MaxChunksPerQuery:       rand.Int31(),
						MaxQuerySeries:          rand.Int31(),
					},
				},
				Tenants: map[string]logstorev1.LimitsTemplateSpec{
					uuid.New().String(): {
						IngestionLimits: &logstorev1.IngestionLimitSpec{
							IngestionRate:             rand.Int31(),
							IngestionBurstSize:        rand.Int31(),
							MaxLabelNameLength:        rand.Int31(),
							MaxLabelValueLength:       rand.Int31(),
							MaxLabelNamesPerSeries:    rand.Int31(),
							MaxGlobalStreamsPerTenant: rand.Int31(),
							MaxLineSize:               rand.Int31(),
						},
						QueryLimits: &logstorev1.QueryLimitSpec{
							MaxEntriesLimitPerQuery: rand.Int31(),
							MaxChunksPerQuery:       rand.Int31(),
							MaxQuerySeries:          rand.Int31(),
						},
					},
				},
			},
			Template: &logstorev1.LogstoreTemplateSpec{
				Compactor: &logstorev1.LogstoreComponentSpec{
					Replicas: 1,
					NodeSelector: map[string]string{
						uuid.New().String(): uuid.New().String(),
					},
					Tolerations: []corev1.Toleration{
						{
							Key:               uuid.New().String(),
							Operator:          corev1.TolerationOpEqual,
							Value:             uuid.New().String(),
							Effect:            corev1.TaintEffectNoExecute,
							TolerationSeconds: pointer.Int64Ptr(rand.Int63()),
						},
					},
				},
				Distributor: &logstorev1.LogstoreComponentSpec{
					Replicas: rand.Int31(),
					NodeSelector: map[string]string{
						uuid.New().String(): uuid.New().String(),
					},
					Tolerations: []corev1.Toleration{
						{
							Key:               uuid.New().String(),
							Operator:          corev1.TolerationOpEqual,
							Value:             uuid.New().String(),
							Effect:            corev1.TaintEffectNoExecute,
							TolerationSeconds: pointer.Int64Ptr(rand.Int63()),
						},
					},
				},
				Ingester: &logstorev1.LogstoreComponentSpec{
					Replicas: rand.Int31(),
					NodeSelector: map[string]string{
						uuid.New().String(): uuid.New().String(),
					},
					Tolerations: []corev1.Toleration{
						{
							Key:               uuid.New().String(),
							Operator:          corev1.TolerationOpEqual,
							Value:             uuid.New().String(),
							Effect:            corev1.TaintEffectNoExecute,
							TolerationSeconds: pointer.Int64Ptr(rand.Int63()),
						},
					},
				},
				Querier: &logstorev1.LogstoreComponentSpec{
					Replicas: rand.Int31(),
					NodeSelector: map[string]string{
						uuid.New().String(): uuid.New().String(),
					},
					Tolerations: []corev1.Toleration{
						{
							Key:               uuid.New().String(),
							Operator:          corev1.TolerationOpEqual,
							Value:             uuid.New().String(),
							Effect:            corev1.TaintEffectNoExecute,
							TolerationSeconds: pointer.Int64Ptr(rand.Int63()),
						},
					},
				},
				QueryFrontend: &logstorev1.LogstoreComponentSpec{
					Replicas: rand.Int31(),
					NodeSelector: map[string]string{
						uuid.New().String(): uuid.New().String(),
					},
					Tolerations: []corev1.Toleration{
						{
							Key:               uuid.New().String(),
							Operator:          corev1.TolerationOpEqual,
							Value:             uuid.New().String(),
							Effect:            corev1.TaintEffectNoExecute,
							TolerationSeconds: pointer.Int64Ptr(rand.Int63()),
						},
					},
				},
				IndexGateway: &logstorev1.LogstoreComponentSpec{
					Replicas: rand.Int31(),
					NodeSelector: map[string]string{
						uuid.New().String(): uuid.New().String(),
					},
					Tolerations: []corev1.Toleration{
						{
							Key:               uuid.New().String(),
							Operator:          corev1.TolerationOpEqual,
							Value:             uuid.New().String(),
							Effect:            corev1.TaintEffectNoExecute,
							TolerationSeconds: pointer.Int64Ptr(rand.Int63()),
						},
					},
				},
			},
		},
	}
}

func TestConfigOptions_RetentionConfig(t *testing.T) {
	tt := []struct {
		desc        string
		spec        logstorev1.LogstoreStackSpec
		wantOptions config.RetentionOptions
	}{
		{
			desc: "no retention",
			spec: logstorev1.LogstoreStackSpec{},
			wantOptions: config.RetentionOptions{
				Enabled: false,
			},
		},
		{
			desc: "global retention, extra small",
			spec: logstorev1.LogstoreStackSpec{
				Size: logstorev1.SizeOneXExtraSmall,
				Limits: &logstorev1.LimitsSpec{
					Global: &logstorev1.LimitsTemplateSpec{
						Retention: &logstorev1.RetentionLimitSpec{
							Days: 14,
						},
					},
				},
			},
			wantOptions: config.RetentionOptions{
				Enabled:           true,
				DeleteWorkerCount: 10,
			},
		},
		{
			desc: "global and tenant retention, extra small",
			spec: logstorev1.LogstoreStackSpec{
				Size: logstorev1.SizeOneXExtraSmall,
				Limits: &logstorev1.LimitsSpec{
					Global: &logstorev1.LimitsTemplateSpec{
						Retention: &logstorev1.RetentionLimitSpec{
							Days: 14,
						},
					},
					Tenants: map[string]logstorev1.LimitsTemplateSpec{
						"development": {
							Retention: &logstorev1.RetentionLimitSpec{
								Days: 3,
							},
						},
					},
				},
			},
			wantOptions: config.RetentionOptions{
				Enabled:           true,
				DeleteWorkerCount: 10,
			},
		},
		{
			desc: "tenant retention, extra small",
			spec: logstorev1.LogstoreStackSpec{
				Size: logstorev1.SizeOneXExtraSmall,
				Limits: &logstorev1.LimitsSpec{
					Tenants: map[string]logstorev1.LimitsTemplateSpec{
						"development": {
							Retention: &logstorev1.RetentionLimitSpec{
								Days: 3,
							},
						},
					},
				},
			},
			wantOptions: config.RetentionOptions{
				Enabled:           true,
				DeleteWorkerCount: 10,
			},
		},
		{
			desc: "global retention, medium",
			spec: logstorev1.LogstoreStackSpec{
				Size: logstorev1.SizeOneXMedium,
				Limits: &logstorev1.LimitsSpec{
					Global: &logstorev1.LimitsTemplateSpec{
						Retention: &logstorev1.RetentionLimitSpec{
							Days: 14,
						},
					},
				},
			},
			wantOptions: config.RetentionOptions{
				Enabled:           true,
				DeleteWorkerCount: 150,
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			inOpt := manifests.Options{
				Stack: tc.spec,
			}
			options := manifests.ConfigOptions(inOpt)
			require.Equal(t, tc.wantOptions, options.Retention)
		})
	}
}

func TestConfigOptions_RulerAlertManager(t *testing.T) {
	tt := []struct {
		desc        string
		opts        manifests.Options
		wantOptions *config.AlertManagerConfig
	}{
		{
			desc: "static mode",
			opts: manifests.Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Static,
					},
				},
			},
			wantOptions: nil,
		},
		{
			desc: "dynamic mode",
			opts: manifests.Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Dynamic,
					},
				},
			},
			wantOptions: nil,
		},
		{
			desc: "openshift-logging mode",
			opts: manifests.Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
				},
				OpenShiftOptions: openshift.Options{
					BuildOpts: openshift.BuildOptions{
						AlertManagerEnabled: true,
					},
				},
			},
			wantOptions: &config.AlertManagerConfig{
				EnableV2:        true,
				EnableDiscovery: true,
				RefreshInterval: "1m",
				Hosts:           "https://_web._tcp.alertmanager-operated.openshift-monitoring.svc",
			},
		},
		{
			desc: "openshift-network mode",
			opts: manifests.Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftNetwork,
					},
				},
				OpenShiftOptions: openshift.Options{
					BuildOpts: openshift.BuildOptions{
						AlertManagerEnabled: true,
					},
				},
			},
			wantOptions: &config.AlertManagerConfig{
				EnableV2:        true,
				EnableDiscovery: true,
				RefreshInterval: "1m",
				Hosts:           "https://_web._tcp.alertmanager-operated.openshift-monitoring.svc",
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			cfg := manifests.ConfigOptions(tc.opts)
			err := manifests.ConfigureOptionsForMode(&cfg, tc.opts)

			require.Nil(t, err)
			require.Equal(t, tc.wantOptions, cfg.Ruler.AlertManager)
		})
	}
}

func TestConfigOptions_RulerAlertManager_UserOverride(t *testing.T) {
	tt := []struct {
		desc        string
		opts        manifests.Options
		wantOptions *config.AlertManagerConfig
	}{
		{
			desc: "static mode",
			opts: manifests.Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Static,
					},
				},
			},
			wantOptions: nil,
		},
		{
			desc: "dynamic mode",
			opts: manifests.Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Dynamic,
					},
				},
			},
			wantOptions: nil,
		},
		{
			desc: "openshift-logging mode",
			opts: manifests.Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
					Rules: &logstorev1.RulesSpec{
						Enabled: true,
					},
				},
				Ruler: manifests.Ruler{
					Spec: &v1beta1.RulerConfigSpec{
						AlertManagerSpec: &v1beta1.AlertManagerSpec{
							EnableV2: false,
							DiscoverySpec: &v1beta1.AlertManagerDiscoverySpec{
								EnableSRV:       false,
								RefreshInterval: "2m",
							},
							Endpoints: []string{"http://my-alertmanager"},
						},
					},
				},
				OpenShiftOptions: openshift.Options{
					BuildOpts: openshift.BuildOptions{
						AlertManagerEnabled: true,
					},
				},
			},
			wantOptions: &config.AlertManagerConfig{
				EnableV2:        false,
				EnableDiscovery: false,
				RefreshInterval: "2m",
				Hosts:           "http://my-alertmanager",
			},
		},
		{
			desc: "openshift-network mode",
			opts: manifests.Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftNetwork,
					},
					Rules: &logstorev1.RulesSpec{
						Enabled: true,
					},
				},
				Ruler: manifests.Ruler{
					Spec: &v1beta1.RulerConfigSpec{
						AlertManagerSpec: &v1beta1.AlertManagerSpec{
							EnableV2: false,
							DiscoverySpec: &v1beta1.AlertManagerDiscoverySpec{
								EnableSRV:       false,
								RefreshInterval: "2m",
							},
							Endpoints: []string{"http://my-alertmanager"},
						},
					},
				},
				OpenShiftOptions: openshift.Options{
					BuildOpts: openshift.BuildOptions{
						AlertManagerEnabled: true,
					},
				},
			},
			wantOptions: &config.AlertManagerConfig{
				EnableV2:        false,
				EnableDiscovery: false,
				RefreshInterval: "2m",
				Hosts:           "http://my-alertmanager",
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			cfg := manifests.ConfigOptions(tc.opts)
			err := manifests.ConfigureOptionsForMode(&cfg, tc.opts)
			require.Nil(t, err)
			require.Equal(t, tc.wantOptions, cfg.Ruler.AlertManager)
		})
	}
}
