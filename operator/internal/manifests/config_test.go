package manifests

import (
	"encoding/json"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/manifests/internal/config"
	"example.com/acme/logstore/operator/internal/manifests/openshift"
)

func TestConfigMap_ReturnsSHA1OfBinaryContents(t *testing.T) {
	opts := randomConfigOptions()

	_, sha1C, err := LogstoreConfigMap(opts)
	require.NoError(t, err)
	require.NotEmpty(t, sha1C)
}

func TestConfigOptions_UserOptionsTakePrecedence(t *testing.T) {
	// regardless of what is provided by the default sizing parameters we should always prefer
	// the user-defined values. This creates an all-inclusive Options and then checks
	// that every value is present in the result
	opts := randomConfigOptions()
	res := ConfigOptions(opts)

	expected, err := json.Marshal(opts.Stack)
	require.NoError(t, err)

	actual, err := json.Marshal(res.Stack)
	require.NoError(t, err)

	assert.JSONEq(t, string(expected), string(actual))
}

func testTimeoutConfig() TimeoutConfig {
	return TimeoutConfig{
		Logstore: config.HTTPTimeoutConfig{
			IdleTimeout:  1 * time.Second,
			ReadTimeout:  1 * time.Minute,
			WriteTimeout: 10 * time.Minute,
		},
	}
}

func randomConfigOptions() Options {
	return Options{
		Name:      uuid.New().String(),
		Namespace: uuid.New().String(),
		Image:     uuid.New().String(),
		Timeouts:  testTimeoutConfig(),
		Stack: logstorev1.LogstoreStackSpec{
			Size:             logstorev1.SizeOneXExtraSmall,
			Storage:          logstorev1.ObjectStorageSpec{},
			StorageClassName: uuid.New().String(),
			Replication: &logstorev1.ReplicationSpec{
				Factor: rand.Int31(),
			},
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
						PerStreamRateLimit:        rand.Int31(),
						PerStreamRateLimitBurst:   rand.Int31(),
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
							PerStreamRateLimit:        rand.Int31(),
							PerStreamRateLimitBurst:   rand.Int31(),
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

func TestConfigOptions_GossipRingConfig(t *testing.T) {
	tt := []struct {
		desc        string
		spec        logstorev1.LogstoreStackSpec
		wantOptions config.GossipRing
	}{
		{
			desc: "defaults",
			spec: logstorev1.LogstoreStackSpec{},
			wantOptions: config.GossipRing{
				InstancePort:         9095,
				BindPort:             7946,
				MembersDiscoveryAddr: "my-stack-gossip-ring.my-ns.svc.cluster.local",
			},
		},
		{
			desc: "defaults with empty config",
			spec: logstorev1.LogstoreStackSpec{
				HashRing: &logstorev1.HashRingSpec{
					Type: logstorev1.HashRingMemberList,
				},
			},
			wantOptions: config.GossipRing{
				InstancePort:         9095,
				BindPort:             7946,
				MembersDiscoveryAddr: "my-stack-gossip-ring.my-ns.svc.cluster.local",
			},
		},
		{
			desc: "user selected any instance addr",
			spec: logstorev1.LogstoreStackSpec{
				HashRing: &logstorev1.HashRingSpec{
					Type: logstorev1.HashRingMemberList,
					MemberList: &logstorev1.MemberListSpec{
						InstanceAddrType: logstorev1.InstanceAddrDefault,
					},
				},
			},
			wantOptions: config.GossipRing{
				InstancePort:         9095,
				BindPort:             7946,
				MembersDiscoveryAddr: "my-stack-gossip-ring.my-ns.svc.cluster.local",
			},
		},
		{
			desc: "user selected podIP instance addr",
			spec: logstorev1.LogstoreStackSpec{
				HashRing: &logstorev1.HashRingSpec{
					Type: logstorev1.HashRingMemberList,
					MemberList: &logstorev1.MemberListSpec{
						InstanceAddrType: logstorev1.InstanceAddrPodIP,
					},
				},
			},
			wantOptions: config.GossipRing{
				InstanceAddr:         "${HASH_RING_INSTANCE_ADDR}",
				InstancePort:         9095,
				BindPort:             7946,
				MembersDiscoveryAddr: "my-stack-gossip-ring.my-ns.svc.cluster.local",
			},
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			inOpt := Options{
				Name:      "my-stack",
				Namespace: "my-ns",
				Stack:     tc.spec,
				Timeouts:  testTimeoutConfig(),
			}
			options := ConfigOptions(inOpt)
			require.Equal(t, tc.wantOptions, options.GossipRing)
		})
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

			inOpt := Options{
				Stack:    tc.spec,
				Timeouts: testTimeoutConfig(),
			}
			options := ConfigOptions(inOpt)
			require.Equal(t, tc.wantOptions, options.Retention)
		})
	}
}

func TestConfigOptions_RulerAlertManager(t *testing.T) {
	tt := []struct {
		desc        string
		opts        Options
		wantOptions *config.AlertManagerConfig
	}{
		{
			desc: "static mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Static,
					},
				},
				Timeouts: testTimeoutConfig(),
			},
			wantOptions: nil,
		},
		{
			desc: "dynamic mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Dynamic,
					},
				},
				Timeouts: testTimeoutConfig(),
			},
			wantOptions: nil,
		},
		{
			desc: "openshift-logging mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
				},
				Timeouts: testTimeoutConfig(),
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
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftNetwork,
					},
				},
				Timeouts: testTimeoutConfig(),
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

			cfg := ConfigOptions(tc.opts)
			err := ConfigureOptionsForMode(&cfg, tc.opts)

			require.Nil(t, err)
			require.Equal(t, tc.wantOptions, cfg.Ruler.AlertManager)
		})
	}
}

func TestConfigOptions_RulerAlertManager_UserOverride(t *testing.T) {
	tt := []struct {
		desc        string
		opts        Options
		wantOptions *config.AlertManagerConfig
	}{
		{
			desc: "static mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Static,
					},
				},
				Timeouts: testTimeoutConfig(),
			},
			wantOptions: nil,
		},
		{
			desc: "dynamic mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Dynamic,
					},
				},
				Timeouts: testTimeoutConfig(),
			},
			wantOptions: nil,
		},
		{
			desc: "openshift-logging mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
					Rules: &logstorev1.RulesSpec{
						Enabled: true,
					},
				},
				Timeouts: testTimeoutConfig(),
				Ruler: Ruler{
					Spec: &logstorev1.RulerConfigSpec{
						AlertManagerSpec: &logstorev1.AlertManagerSpec{
							EnableV2: false,
							DiscoverySpec: &logstorev1.AlertManagerDiscoverySpec{
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
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftNetwork,
					},
					Rules: &logstorev1.RulesSpec{
						Enabled: true,
					},
				},
				Timeouts: testTimeoutConfig(),
				Ruler: Ruler{
					Spec: &logstorev1.RulerConfigSpec{
						AlertManagerSpec: &logstorev1.AlertManagerSpec{
							EnableV2: false,
							DiscoverySpec: &logstorev1.AlertManagerDiscoverySpec{
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

			cfg := ConfigOptions(tc.opts)
			err := ConfigureOptionsForMode(&cfg, tc.opts)
			require.Nil(t, err)
			require.Equal(t, tc.wantOptions, cfg.Ruler.AlertManager)
		})
	}
}

func TestConfigOptions_RulerOverrides_OCPApplicationTenant(t *testing.T) {
	tt := []struct {
		desc        string
		opts        Options
		wantOptions map[string]config.LogstoreOverrides
	}{
		{
			desc: "static mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Static,
					},
				},
				Timeouts: testTimeoutConfig(),
			},
			wantOptions: nil,
		},
		{
			desc: "dynamic mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Dynamic,
					},
				},
				Timeouts: testTimeoutConfig(),
			},
			wantOptions: nil,
		},
		{
			desc: "openshift-logging mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
					Rules: &logstorev1.RulesSpec{
						Enabled: true,
					},
				},
				Timeouts: testTimeoutConfig(),
				Ruler: Ruler{
					Spec: &logstorev1.RulerConfigSpec{
						AlertManagerSpec: &logstorev1.AlertManagerSpec{
							EnableV2: false,
							DiscoverySpec: &logstorev1.AlertManagerDiscoverySpec{
								EnableSRV:       false,
								RefreshInterval: "2m",
							},
							Endpoints: []string{"http://my-alertmanager"},
						},
					},
				},
				OpenShiftOptions: openshift.Options{
					BuildOpts: openshift.BuildOptions{
						AlertManagerEnabled:             true,
						UserWorkloadAlertManagerEnabled: true,
					},
				},
			},
			wantOptions: map[string]config.LogstoreOverrides{
				"application": {
					Ruler: config.RulerOverrides{
						AlertManager: &config.AlertManagerConfig{
							Hosts:           "https://_web._tcp.alertmanager-operated.openshift-user-workload-monitoring.svc",
							EnableV2:        true,
							EnableDiscovery: true,
							RefreshInterval: "1m",
							Notifier: &config.NotifierConfig{
								TLS: config.TLSConfig{
									ServerName: pointer.String("alertmanager-user-workload.openshift-user-workload-monitoring.svc.cluster.local"),
									CAPath:     pointer.String("/var/run/ca/alertmanager/service-ca.crt"),
								},
								HeaderAuth: config.HeaderAuth{
									Type:            pointer.String("Bearer"),
									CredentialsFile: pointer.String("/var/run/secrets/kubernetes.io/serviceaccount/token"),
								},
							},
						},
					},
				},
			},
		},
		{
			desc: "openshift-network mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftNetwork,
					},
					Rules: &logstorev1.RulesSpec{
						Enabled: true,
					},
				},
				Timeouts: testTimeoutConfig(),
				Ruler: Ruler{
					Spec: &logstorev1.RulerConfigSpec{
						AlertManagerSpec: &logstorev1.AlertManagerSpec{
							EnableV2: false,
							DiscoverySpec: &logstorev1.AlertManagerDiscoverySpec{
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
			wantOptions: nil,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			cfg := ConfigOptions(tc.opts)
			err := ConfigureOptionsForMode(&cfg, tc.opts)
			require.Nil(t, err)
			require.EqualValues(t, tc.wantOptions, cfg.Overrides)
		})
	}
}

func TestConfigOptions_RulerOverrides(t *testing.T) {
	tt := []struct {
		desc        string
		opts        Options
		wantOptions map[string]config.LogstoreOverrides
	}{
		{
			desc: "static mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Static,
					},
				},
				Timeouts: testTimeoutConfig(),
			},
			wantOptions: nil,
		},
		{
			desc: "dynamic mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Dynamic,
					},
				},
				Timeouts: testTimeoutConfig(),
			},
			wantOptions: nil,
		},
		{
			desc: "openshift-logging mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
					Rules: &logstorev1.RulesSpec{
						Enabled: true,
					},
				},
				Timeouts: testTimeoutConfig(),
				Ruler: Ruler{
					Spec: &logstorev1.RulerConfigSpec{
						AlertManagerSpec: &logstorev1.AlertManagerSpec{
							EnableV2: false,
							DiscoverySpec: &logstorev1.AlertManagerDiscoverySpec{
								EnableSRV:       false,
								RefreshInterval: "2m",
							},
							Endpoints: []string{"http://my-alertmanager"},
						},
						Overrides: map[string]logstorev1.RulerOverrides{
							"application": {
								AlertManagerOverrides: &logstorev1.AlertManagerSpec{
									ExternalURL:    "external",
									ExternalLabels: map[string]string{"external": "label"},
									EnableV2:       false,
									Endpoints:      []string{"http://application-alertmanager"},
									DiscoverySpec: &logstorev1.AlertManagerDiscoverySpec{
										EnableSRV:       false,
										RefreshInterval: "3m",
									},
									Client: &logstorev1.AlertManagerClientConfig{
										TLS: &logstorev1.AlertManagerClientTLSConfig{
											ServerName: pointer.String("application.svc"),
											CAPath:     pointer.String("/tenant/application/alertmanager/ca.crt"),
											CertPath:   pointer.String("/tenant/application/alertmanager/cert.crt"),
											KeyPath:    pointer.String("/tenant/application/alertmanager/cert.key"),
										},
										HeaderAuth: &logstorev1.AlertManagerClientHeaderAuth{
											Type:        pointer.String("Bearer"),
											Credentials: pointer.String("letmeinplz"),
										},
									},
								},
							},
							"other-tenant": {
								AlertManagerOverrides: &logstorev1.AlertManagerSpec{
									ExternalURL:    "external1",
									ExternalLabels: map[string]string{"external1": "label1"},
									EnableV2:       false,
									Endpoints:      []string{"http://other-alertmanager"},
									DiscoverySpec: &logstorev1.AlertManagerDiscoverySpec{
										EnableSRV:       true,
										RefreshInterval: "5m",
									},
									Client: &logstorev1.AlertManagerClientConfig{
										TLS: &logstorev1.AlertManagerClientTLSConfig{
											ServerName: pointer.String("other.svc"),
											CAPath:     pointer.String("/tenant/other/alertmanager/ca.crt"),
											CertPath:   pointer.String("/tenant/other/alertmanager/cert.crt"),
											KeyPath:    pointer.String("/tenant/other/alertmanager/cert.key"),
										},
										BasicAuth: &logstorev1.AlertManagerClientBasicAuth{
											Username: pointer.String("user"),
											Password: pointer.String("pass"),
										},
									},
								},
							},
						},
					},
				},
				OpenShiftOptions: openshift.Options{
					BuildOpts: openshift.BuildOptions{
						AlertManagerEnabled:             true,
						UserWorkloadAlertManagerEnabled: true,
					},
				},
			},
			wantOptions: map[string]config.LogstoreOverrides{
				"application": {
					Ruler: config.RulerOverrides{
						AlertManager: &config.AlertManagerConfig{
							ExternalURL:     "external",
							Hosts:           "http://application-alertmanager",
							EnableV2:        false,
							EnableDiscovery: false,
							RefreshInterval: "3m",
							ExternalLabels:  map[string]string{"external": "label"},
							Notifier: &config.NotifierConfig{
								TLS: config.TLSConfig{
									ServerName: pointer.String("application.svc"),
									CAPath:     pointer.String("/tenant/application/alertmanager/ca.crt"),
									CertPath:   pointer.String("/tenant/application/alertmanager/cert.crt"),
									KeyPath:    pointer.String("/tenant/application/alertmanager/cert.key"),
								},
								HeaderAuth: config.HeaderAuth{
									Type:        pointer.String("Bearer"),
									Credentials: pointer.String("letmeinplz"),
								},
							},
						},
					},
				},
				"other-tenant": {
					Ruler: config.RulerOverrides{
						AlertManager: &config.AlertManagerConfig{
							ExternalURL:     "external1",
							Hosts:           "http://other-alertmanager",
							EnableV2:        false,
							EnableDiscovery: true,
							RefreshInterval: "5m",
							ExternalLabels:  map[string]string{"external1": "label1"},
							Notifier: &config.NotifierConfig{
								TLS: config.TLSConfig{
									ServerName: pointer.String("other.svc"),
									CAPath:     pointer.String("/tenant/other/alertmanager/ca.crt"),
									CertPath:   pointer.String("/tenant/other/alertmanager/cert.crt"),
									KeyPath:    pointer.String("/tenant/other/alertmanager/cert.key"),
								},
								BasicAuth: config.BasicAuth{
									Username: pointer.String("user"),
									Password: pointer.String("pass"),
								},
							},
						},
					},
				},
			},
		},
		{
			desc: "openshift-network mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
					Rules: &logstorev1.RulesSpec{
						Enabled: true,
					},
				},
				Timeouts: testTimeoutConfig(),
				Ruler: Ruler{
					Spec: &logstorev1.RulerConfigSpec{
						AlertManagerSpec: &logstorev1.AlertManagerSpec{
							EnableV2: false,
							DiscoverySpec: &logstorev1.AlertManagerDiscoverySpec{
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
			wantOptions: nil,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			cfg := ConfigOptions(tc.opts)
			err := ConfigureOptionsForMode(&cfg, tc.opts)
			require.Nil(t, err)
			require.EqualValues(t, tc.wantOptions, cfg.Overrides)
		})
	}
}

func TestConfigOptions_RulerOverrides_OCPUserWorkloadOnlyEnabled(t *testing.T) {
	tt := []struct {
		desc                 string
		opts                 Options
		wantOptions          *config.AlertManagerConfig
		wantOverridesOptions map[string]config.LogstoreOverrides
	}{
		{
			desc: "static mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Static,
					},
				},
				Timeouts: testTimeoutConfig(),
			},
			wantOptions:          nil,
			wantOverridesOptions: nil,
		},
		{
			desc: "dynamic mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.Dynamic,
					},
				},
				Timeouts: testTimeoutConfig(),
			},
			wantOptions:          nil,
			wantOverridesOptions: nil,
		},
		{
			desc: "openshift-logging mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
					Rules: &logstorev1.RulesSpec{
						Enabled: true,
					},
				},
				Timeouts: testTimeoutConfig(),
				Ruler: Ruler{
					Spec: &logstorev1.RulerConfigSpec{
						AlertManagerSpec: &logstorev1.AlertManagerSpec{
							EnableV2: false,
							DiscoverySpec: &logstorev1.AlertManagerDiscoverySpec{
								EnableSRV:       false,
								RefreshInterval: "2m",
							},
							Endpoints: []string{"http://my-alertmanager"},
						},
					},
				},
				OpenShiftOptions: openshift.Options{
					BuildOpts: openshift.BuildOptions{
						AlertManagerEnabled:             false,
						UserWorkloadAlertManagerEnabled: true,
					},
				},
			},
			wantOptions: &config.AlertManagerConfig{
				EnableV2:        false,
				EnableDiscovery: false,
				RefreshInterval: "2m",
				Hosts:           "http://my-alertmanager",
			},
			wantOverridesOptions: map[string]config.LogstoreOverrides{
				"application": {
					Ruler: config.RulerOverrides{
						AlertManager: &config.AlertManagerConfig{
							Hosts:           "https://_web._tcp.alertmanager-operated.openshift-user-workload-monitoring.svc",
							EnableV2:        true,
							EnableDiscovery: true,
							RefreshInterval: "1m",
							Notifier: &config.NotifierConfig{
								TLS: config.TLSConfig{
									ServerName: pointer.String("alertmanager-user-workload.openshift-user-workload-monitoring.svc.cluster.local"),
									CAPath:     pointer.String("/var/run/ca/alertmanager/service-ca.crt"),
								},
								HeaderAuth: config.HeaderAuth{
									Type:            pointer.String("Bearer"),
									CredentialsFile: pointer.String("/var/run/secrets/kubernetes.io/serviceaccount/token"),
								},
							},
						},
					},
				},
			},
		},
		{
			desc: "openshift-network mode",
			opts: Options{
				Stack: logstorev1.LogstoreStackSpec{
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftNetwork,
					},
					Rules: &logstorev1.RulesSpec{
						Enabled: true,
					},
				},

				Timeouts: testTimeoutConfig(),
				Ruler: Ruler{
					Spec: &logstorev1.RulerConfigSpec{
						AlertManagerSpec: &logstorev1.AlertManagerSpec{
							EnableV2: false,
							DiscoverySpec: &logstorev1.AlertManagerDiscoverySpec{
								EnableSRV:       false,
								RefreshInterval: "2m",
							},
							Endpoints: []string{"http://my-alertmanager"},
						},
					},
				},
				OpenShiftOptions: openshift.Options{
					BuildOpts: openshift.BuildOptions{
						AlertManagerEnabled: false,
					},
				},
			},
			wantOptions: &config.AlertManagerConfig{
				EnableV2:        false,
				EnableDiscovery: false,
				RefreshInterval: "2m",
				Hosts:           "http://my-alertmanager",
			},
			wantOverridesOptions: nil,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			cfg := ConfigOptions(tc.opts)
			err := ConfigureOptionsForMode(&cfg, tc.opts)
			require.Nil(t, err)
			require.EqualValues(t, tc.wantOverridesOptions, cfg.Overrides)
			require.EqualValues(t, tc.wantOptions, cfg.Ruler.AlertManager)
		})
	}
}

func TestConfigOptions_Replication(t *testing.T) {
	tt := []struct {
		desc        string
		spec        logstorev1.LogstoreStackSpec
		wantOptions logstorev1.ReplicationSpec
	}{
		{
			desc: "nominal case",
			spec: logstorev1.LogstoreStackSpec{
				Replication: &logstorev1.ReplicationSpec{
					Factor: 2,
					Zones: []logstorev1.ZoneSpec{
						{
							TopologyKey: "us-east-1a",
							MaxSkew:     1,
						},
						{
							TopologyKey: "us-east-1b",
							MaxSkew:     2,
						},
					},
				},
			},
			wantOptions: logstorev1.ReplicationSpec{
				Factor: 2,
				Zones: []logstorev1.ZoneSpec{
					{
						TopologyKey: "us-east-1a",
						MaxSkew:     1,
					},
					{
						TopologyKey: "us-east-1b",
						MaxSkew:     2,
					},
				},
			},
		},
		{
			desc: "using deprecated ReplicationFactor",
			spec: logstorev1.LogstoreStackSpec{
				ReplicationFactor: 3,
			},
			wantOptions: logstorev1.ReplicationSpec{
				Factor: 3,
			},
		},
		{
			desc: "using deprecated ReplicationFactor with ReplicationSpec",
			spec: logstorev1.LogstoreStackSpec{
				ReplicationFactor: 2,
				Replication: &logstorev1.ReplicationSpec{
					Factor: 4,
					Zones: []logstorev1.ZoneSpec{
						{
							TopologyKey: "us-east-1a",
							MaxSkew:     1,
						},
						{
							TopologyKey: "us-east-1b",
							MaxSkew:     2,
						},
					},
				},
			},
			wantOptions: logstorev1.ReplicationSpec{
				Factor: 4,
				Zones: []logstorev1.ZoneSpec{
					{
						TopologyKey: "us-east-1a",
						MaxSkew:     1,
					},
					{
						TopologyKey: "us-east-1b",
						MaxSkew:     2,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			inOpt := Options{
				Stack:    tc.spec,
				Timeouts: testTimeoutConfig(),
			}
			options := ConfigOptions(inOpt)
			require.Equal(t, tc.wantOptions, *options.Stack.Replication)
		})
	}
}

func TestConfigOptions_ServerOptions(t *testing.T) {
	opt := Options{
		Stack:    logstorev1.LogstoreStackSpec{},
		Timeouts: testTimeoutConfig(),
	}
	got := ConfigOptions(opt)

	want := config.HTTPTimeoutConfig{
		IdleTimeout:  time.Second,
		ReadTimeout:  time.Minute,
		WriteTimeout: 10 * time.Minute,
	}

	require.Equal(t, want, got.HTTPTimeouts)
}
