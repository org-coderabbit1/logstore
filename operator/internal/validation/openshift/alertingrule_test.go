package openshift

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
)

func TestAlertingRuleValidator(t *testing.T) {
	var nilMap map[string]string

	tt := []struct {
		desc       string
		spec       *logstorev1.AlertingRule
		wantErrors field.ErrorList
	}{
		{
			desc: "success",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "example",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="example", level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: nil,
		},
		{
			desc: "allow audit in openshift-logging",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "openshift-logging",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "audit",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: nil,
		},
		{
			desc: "allow infrastructure in openshift-logging without namespace label",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "openshift-logging",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "infrastructure",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: nil,
		},
		{
			desc: "wrong tenant",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "openshift-example",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="openshift-example", level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "spec.tenantID",
					BadValue: "application",
					Detail:   `AlertingRule does not use correct tenant ["infrastructure"]`,
				},
			},
		},
		{
			desc: "custom tenant topology enabled",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "openshift-example",
					Annotations: map[string]string{
						logstorev1.AnnotationDisableTenantValidation: "true",
					},
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "foobar",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="openshift-example", level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			desc: "custom tenant topology disabled",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "openshift-example",
					Annotations: map[string]string{
						logstorev1.AnnotationDisableTenantValidation: "false",
					},
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "foobar",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="openshift-example", level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "spec.tenantID",
					BadValue: "foobar",
					Detail:   `AlertingRule does not use correct tenant ["infrastructure"]`,
				},
			},
		},
		{
			desc: "wrong tenant topology",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "openshift-example",
					Annotations: map[string]string{
						logstorev1.AnnotationDisableTenantValidation: "not-valid",
					},
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "foobar",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="openshift-example", level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    `metadata.annotations[logstore.acme.com/disable-tenant-validation]`,
					BadValue: "not-valid",
					Detail:   `strconv.ParseBool: parsing "not-valid": invalid syntax`,
				},
			},
		},
		{
			desc: "expression does not parse",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "example",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: "invalid",
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "spec.groups[0].rules[0].expr",
					BadValue: "invalid",
					Detail:   logstorev1.ErrParseLogQLExpression.Error(),
				},
			},
		},
		{
			desc: "expression does not produce samples",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "example",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `{kubernetes_namespace_name="example", level="error"}`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "spec.groups[0].rules[0].expr",
					BadValue: `{kubernetes_namespace_name="example", level="error"}`,
					Detail:   logstorev1.ErrParseLogQLNotSample.Error(),
				},
			},
		},
		{
			desc: "no namespace matcher",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "example",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "spec.groups[0].rules[0].expr",
					BadValue: `sum(rate({level="error"}[5m])) by (job) > 0.1`,
					Detail:   logstorev1.ErrRuleMustMatchNamespace.Error(),
				},
			},
		},
		{
			desc: "matcher does not match AlertingRule namespace",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "example",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="other-ns", level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "spec.groups[0].rules[0].expr",
					BadValue: `sum(rate({kubernetes_namespace_name="other-ns", level="error"}[5m])) by (job) > 0.1`,
					Detail:   logstorev1.ErrRuleMustMatchNamespace.Error(),
				},
			},
		},
		{
			desc: "missing severity label",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "example",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="example", level="error"}[5m])) by (job) > 0.1`,
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "spec.groups[0].rules[0].labels",
					BadValue: nilMap,
					Detail:   logstorev1.ErrSeverityLabelMissing.Error(),
				},
			},
		},
		{
			desc: "invalid severity label",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "example",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="example", level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "page",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:  field.ErrorTypeInvalid,
					Field: "spec.groups[0].rules[0].labels",
					BadValue: map[string]string{
						severityLabelName: "page",
					},
					Detail: logstorev1.ErrSeverityLabelInvalid.Error(),
				},
			},
		},
		{
			desc: "missing summary annotation",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "example",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="example", level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:  field.ErrorTypeInvalid,
					Field: "spec.groups[0].rules[0].annotations",
					BadValue: map[string]string{
						descriptionAnnotationName: "alert description",
					},
					Detail: logstorev1.ErrSummaryAnnotationMissing.Error(),
				},
			},
		},
		{
			desc: "missing description annotation",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "example",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="example", level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName: "alert summary",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:  field.ErrorTypeInvalid,
					Field: "spec.groups[0].rules[0].annotations",
					BadValue: map[string]string{
						summaryAnnotationName: "alert summary",
					},
					Detail: logstorev1.ErrDescriptionAnnotationMissing.Error(),
				},
			},
		},
		{
			desc: "AlertingRule expression using both namespace labels",
			spec: &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "alerting-rule",
					Namespace: "example",
				},
				Spec: logstorev1.AlertingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.AlertingRuleGroup{
						{
							Rules: []*logstorev1.AlertingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="example", k8s_namespace_name="example", level="error"}[5m])) by (job) > 0.1`,
									Labels: map[string]string{
										severityLabelName: "warning",
									},
									Annotations: map[string]string{
										summaryAnnotationName:     "alert summary",
										descriptionAnnotationName: "alert description",
									},
								},
							},
						},
					},
				},
			},
			wantErrors: []*field.Error{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "spec.groups[0].rules[0].expr",
					BadValue: `sum(rate({kubernetes_namespace_name="example", k8s_namespace_name="example", level="error"}[5m])) by (job) > 0.1`,
					Detail:   logstorev1.ErrRuleExclusiveNamespaceLabel.Error(),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			errors := AlertingRuleValidator(ctx, tc.spec)
			require.Equal(t, tc.wantErrors, errors)
		})
	}
}
