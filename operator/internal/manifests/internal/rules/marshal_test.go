package rules_test

import (
	"fmt"
	"testing"

	logstorev1beta1 "example.com/acme/logstore/operator/apis/logstore/v1beta1"
	"example.com/acme/logstore/operator/internal/manifests/internal/rules"
	"github.com/stretchr/testify/require"
)

func TestMarshalAlertingRule(t *testing.T) {
	expCfg := `
groups:
  - name: an-alert
    interval: 1m
    limit: 2
    rules:
      - expr: |-
          sum(rate({app="foo", env="production"} |= "error" [5m])) by (job)
            /
          sum(rate({app="foo", env="production"}[5m])) by (job)
            > 0.05
        alert: HighPercentageErrors
        for: 10m
        annotations:
          playbook: http://link/to/playbook
          summary: High Percentage Latency
        labels:
          environment: production
          severity: page
      - expr: |-
          sum(rate({app="foo", env="production"} |= "error" [5m])) by (job)
            /
          sum(rate({app="foo", env="production"}[5m])) by (job)
            > 0.05
        alert: LowPercentageErrors
        for: 10m
        annotations:
          playbook: http://link/to/playbook
          summary: Low Percentage Latency
        labels:
          environment: production
          severity: low
`

	a := logstorev1beta1.AlertingRule{
		Spec: logstorev1beta1.AlertingRuleSpec{
			Groups: []*logstorev1beta1.AlertingRuleGroup{
				{
					Name:     "an-alert",
					Interval: logstorev1beta1.PrometheusDuration("1m"),
					Limit:    2,
					Rules: []*logstorev1beta1.AlertingRuleGroupSpec{
						{
							Alert: "HighPercentageErrors",
							Expr: `sum(rate({app="foo", env="production"} |= "error" [5m])) by (job)
  /
sum(rate({app="foo", env="production"}[5m])) by (job)
  > 0.05`,
							For: logstorev1beta1.PrometheusDuration("10m"),
							Labels: map[string]string{
								"severity":    "page",
								"environment": "production",
							},
							Annotations: map[string]string{
								"summary":  "High Percentage Latency",
								"playbook": "http://link/to/playbook",
							},
						},
						{
							Alert: "LowPercentageErrors",
							Expr: `sum(rate({app="foo", env="production"} |= "error" [5m])) by (job)
  /
sum(rate({app="foo", env="production"}[5m])) by (job)
  > 0.05`,
							For: logstorev1beta1.PrometheusDuration("10m"),
							Labels: map[string]string{
								"severity":    "low",
								"environment": "production",
							},
							Annotations: map[string]string{
								"summary":  "Low Percentage Latency",
								"playbook": "http://link/to/playbook",
							},
						},
					},
				},
			},
		},
	}

	cfg, err := rules.MarshalAlertingRule(a)
	require.NoError(t, err)
	require.YAMLEq(t, expCfg, cfg)
}

func TestMarshalRecordingRule(t *testing.T) {
	expCfg := `
groups:
  - name: a-recording
    interval: 2d
    limit: 1
    rules:
      - expr: |-
          sum(
            rate({container="nginx"}[1m])
          )
        record: nginx:requests:rate1m
      - expr: |-
          sum(
            rate({container="banana"}[5m])
          )
        record: banana:requests:rate5m
`

	r := logstorev1beta1.RecordingRule{
		Spec: logstorev1beta1.RecordingRuleSpec{
			Groups: []*logstorev1beta1.RecordingRuleGroup{
				{
					Name:     "a-recording",
					Interval: logstorev1beta1.PrometheusDuration("2d"),
					Limit:    1,
					Rules: []*logstorev1beta1.RecordingRuleGroupSpec{
						{
							Expr: `sum(
  rate({container="nginx"}[1m])
)`,
							Record: "nginx:requests:rate1m",
						},
						{
							Expr: `sum(
  rate({container="banana"}[5m])
)`,
							Record: "banana:requests:rate5m",
						},
					},
				},
			},
		},
	}

	cfg, err := rules.MarshalRecordingRule(r)
	fmt.Print(cfg)
	require.NoError(t, err)
	require.YAMLEq(t, expCfg, cfg)
}
