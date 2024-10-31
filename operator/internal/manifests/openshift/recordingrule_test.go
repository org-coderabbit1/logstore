package openshift

import (
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
)

func TestRecordingRuleTenantLabels(t *testing.T) {
	tt := []struct {
		rule *logstorev1.RecordingRule
		want *logstorev1.RecordingRule
	}{
		{
			rule: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: tenantApplication,
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Record: "record",
								},
							},
						},
					},
				},
			},
			want: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: tenantApplication,
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Record: "record",
									Labels: map[string]string{
										opaDefaultLabelMatcher:    "test-ns",
										ocpMonitoringGroupByLabel: "test-ns",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			rule: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: tenantInfrastructure,
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Record: "record",
								},
							},
						},
					},
				},
			},
			want: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: tenantInfrastructure,
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Record: "record",
									Labels: map[string]string{
										ocpMonitoringGroupByLabel: "test-ns",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			rule: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: tenantAudit,
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Record: "record",
								},
							},
						},
					},
				},
			},
			want: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: tenantAudit,
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Record: "record",
									Labels: map[string]string{
										ocpMonitoringGroupByLabel: "test-ns",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			rule: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: tenantNetwork,
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Record: "record",
								},
							},
						},
					},
				},
			},
			want: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: tenantNetwork,
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Record: "record",
									Labels: map[string]string{
										ocpMonitoringGroupByLabel: "test-ns",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			rule: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "unknown",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Record: "record",
								},
							},
						},
					},
				},
			},
			want: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "unknown",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Name: "test-group",
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Record: "record",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.rule.Spec.TenantID, func(t *testing.T) {
			t.Parallel()
			RecordingRuleTenantLabels(tc.rule)

			require.Equal(t, tc.want, tc.rule)
		})
	}
}
