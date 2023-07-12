package openshift

import (
	"testing"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAlertingRuleTenantLabels(t *testing.T) {
	tt := []struct {
		rule *logstorev1.AlertingRule
		want *logstorev1.AlertingRule
	}{
		{
			rule: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: tenantApplication,
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Alert: "alert",
								},
							},
						},
					},
				},
			},
			want: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: tenantApplication,
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Alert: "alert",
									Labels: map[string]string{
										opaDefaultLabelMatcher: "test-ns",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			rule: &logstorev1.AlertingRule{
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: tenantInfrastructure,
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Alert: "alert",
								},
							},
						},
					},
				},
			},
			want: &logstorev1.AlertingRule{
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: tenantInfrastructure,
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Alert: "alert",
								},
							},
						},
					},
				},
			},
		},
		{
			rule: &logstorev1.AlertingRule{
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: tenantAudit,
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Alert: "alert",
								},
							},
						},
					},
				},
			},
			want: &logstorev1.AlertingRule{
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: tenantAudit,
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Alert: "alert",
								},
							},
						},
					},
				},
			},
		},
		{
			rule: &logstorev1.AlertingRule{
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: tenantNetwork,
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Alert: "alert",
								},
							},
						},
					},
				},
			},
			want: &logstorev1.AlertingRule{
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: tenantNetwork,
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Alert: "alert",
								},
							},
						},
					},
				},
			},
		},
		{
			rule: &logstorev1.AlertingRule{
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "unknown",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Alert: "alert",
								},
							},
						},
					},
				},
			},
			want: &logstorev1.AlertingRule{
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "unknown",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Alert: "alert",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.rule.Spec.TenantID, func(t *testing.T) {
			t.Parallel()
			AlertingRuleTenantLabels(tc.rule)

			require.Equal(t, tc.want, tc.rule)
		})
	}
}
