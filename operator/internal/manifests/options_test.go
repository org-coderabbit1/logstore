package manifests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"example.com/acme/logstore/operator/internal/manifests/internal/config"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
)

func TestNewTimeoutConfig_ReturnsDefaults_WhenLimitsSpecEmpty(t *testing.T) {
	s := logstorev1.LogstoreStack{}

	got, err := NewTimeoutConfig(s.Spec.Limits)
	require.NoError(t, err)
	require.Equal(t, defaultTimeoutConfig, got)
}

func TestNewTimeoutConfig_ReturnsCustomConfig_WhenLimitsSpecNotEmpty(t *testing.T) {
	s := logstorev1.LogstoreStack{
		Spec: logstorev1.LogstoreStackSpec{
			Limits: &logstorev1.LimitsSpec{
				Global: &logstorev1.LimitsTemplateSpec{
					QueryLimits: &logstorev1.QueryLimitSpec{
						QueryTimeout: "10m",
					},
				},
			},
		},
	}

	got, err := NewTimeoutConfig(s.Spec.Limits)
	require.NoError(t, err)

	want := TimeoutConfig{
		Logstore: config.HTTPTimeoutConfig{
			IdleTimeout:  30 * time.Second,
			ReadTimeout:  1 * time.Minute,
			WriteTimeout: 11 * time.Minute,
		},
		Gateway: GatewayTimeoutConfig{
			ReadTimeout:          1*time.Minute + gatewayReadDuration,
			WriteTimeout:         11*time.Minute + gatewayWriteDuration,
			UpstreamWriteTimeout: 11 * time.Minute,
		},
	}

	require.Equal(t, want, got)
}

func TestNewTimeoutConfig_ReturnsCustomConfig_WhenLimitsSpecNotEmpty_UseMaxTenantQueryTimeout(t *testing.T) {
	s := logstorev1.LogstoreStack{
		Spec: logstorev1.LogstoreStackSpec{
			Limits: &logstorev1.LimitsSpec{
				Global: &logstorev1.LimitsTemplateSpec{
					QueryLimits: &logstorev1.QueryLimitSpec{
						QueryTimeout: "10m",
					},
				},
				Tenants: map[string]logstorev1.LimitsTemplateSpec{
					"tenant-a": {
						QueryLimits: &logstorev1.QueryLimitSpec{
							QueryTimeout: "10m",
						},
					},
					"tenant-b": {
						QueryLimits: &logstorev1.QueryLimitSpec{
							QueryTimeout: "20m",
						},
					},
				},
			},
		},
	}

	got, err := NewTimeoutConfig(s.Spec.Limits)
	require.NoError(t, err)

	want := TimeoutConfig{
		Logstore: config.HTTPTimeoutConfig{
			IdleTimeout:  30 * time.Second,
			ReadTimeout:  2 * time.Minute,
			WriteTimeout: 21 * time.Minute,
		},
		Gateway: GatewayTimeoutConfig{
			ReadTimeout:          2*time.Minute + gatewayReadDuration,
			WriteTimeout:         21*time.Minute + gatewayWriteDuration,
			UpstreamWriteTimeout: 21 * time.Minute,
		},
	}

	require.Equal(t, want, got)
}

func TestNewTimeoutConfig_ReturnsCustomConfig_WhenTenantLimitsSpecOnly_ReturnsUseMaxTenantQueryTimeout(t *testing.T) {
	s := logstorev1.LogstoreStack{
		Spec: logstorev1.LogstoreStackSpec{
			Limits: &logstorev1.LimitsSpec{
				Tenants: map[string]logstorev1.LimitsTemplateSpec{
					"tenant-a": {
						QueryLimits: &logstorev1.QueryLimitSpec{
							QueryTimeout: "10m",
						},
					},
					"tenant-b": {
						QueryLimits: &logstorev1.QueryLimitSpec{
							QueryTimeout: "20m",
						},
					},
				},
			},
		},
	}

	got, err := NewTimeoutConfig(s.Spec.Limits)
	require.NoError(t, err)

	want := TimeoutConfig{
		Logstore: config.HTTPTimeoutConfig{
			IdleTimeout:  30 * time.Second,
			ReadTimeout:  2 * time.Minute,
			WriteTimeout: 21 * time.Minute,
		},
		Gateway: GatewayTimeoutConfig{
			ReadTimeout:          2*time.Minute + gatewayReadDuration,
			WriteTimeout:         21*time.Minute + gatewayWriteDuration,
			UpstreamWriteTimeout: 21 * time.Minute,
		},
	}

	require.Equal(t, want, got)
}

func TestNewTimeoutConfig_ReturnsDefaults_WhenGlobalQueryTimeoutParseError(t *testing.T) {
	s := logstorev1.LogstoreStack{
		Spec: logstorev1.LogstoreStackSpec{
			Limits: &logstorev1.LimitsSpec{
				Global: &logstorev1.LimitsTemplateSpec{
					QueryLimits: &logstorev1.QueryLimitSpec{
						QueryTimeout: "invalid",
					},
				},
			},
		},
	}

	_, err := NewTimeoutConfig(s.Spec.Limits)
	require.Error(t, err)
}

func TestNewTimeoutConfig_ReturnsDefaults_WhenTenantQueryTimeoutParseError(t *testing.T) {
	s := logstorev1.LogstoreStack{
		Spec: logstorev1.LogstoreStackSpec{
			Limits: &logstorev1.LimitsSpec{
				Global: &logstorev1.LimitsTemplateSpec{
					QueryLimits: &logstorev1.QueryLimitSpec{
						QueryTimeout: "10m",
					},
				},
				Tenants: map[string]logstorev1.LimitsTemplateSpec{
					"tenant-a": {
						QueryLimits: &logstorev1.QueryLimitSpec{
							QueryTimeout: "invalid",
						},
					},
					"tenant-b": {
						QueryLimits: &logstorev1.QueryLimitSpec{
							QueryTimeout: "20m",
						},
					},
				},
			},
		},
	}

	_, err := NewTimeoutConfig(s.Spec.Limits)
	require.Error(t, err)
}
