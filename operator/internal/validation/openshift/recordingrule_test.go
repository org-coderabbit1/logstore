package openshift

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
)

func TestRecordingRuleValidator(t *testing.T) {
	tt := []struct {
		desc       string
		spec       *logstorev1.RecordingRule
		wantErrors field.ErrorList
	}{
		{
			desc: "success",
			spec: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rule",
					Namespace: "example",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="example", level="error"}[5m])) by (job) > 0.1`,
								},
							},
						},
					},
				},
			},
			wantErrors: nil,
		},
		{
			desc: "success audit tenant",
			spec: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rule",
					Namespace: "openshift-logging",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "audit",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Expr: `sum(rate({level="error"}[5m])) by (job) > 0.1`,
								},
							},
						},
					},
				},
			},
			wantErrors: nil,
		},
		{
			desc: "success infrastructure tenant",
			spec: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rule",
					Namespace: "openshift-logging",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "infrastructure",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Expr: `sum(rate({level="error"}[5m])) by (job) > 0.1`,
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
			spec: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rule",
					Namespace: "openshift-example",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="openshift-example", level="error"}[5m])) by (job) > 0.1`,
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
					Detail:   `RecordingRule does not use correct tenant ["infrastructure"]`,
				},
			},
		},
		{
			desc: "custom tenant topology enabled",
			spec: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rule",
					Namespace: "openshift-example",
					Annotations: map[string]string{
						logstorev1.AnnotationDisableTenantValidation: "true",
					},
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "foobar",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="openshift-example", level="error"}[5m])) by (job) > 0.1`,
								},
							},
						},
					},
				},
			},
		},
		{
			desc: "wrong tenant topology",
			spec: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rule",
					Namespace: "openshift-example",
					Annotations: map[string]string{
						logstorev1.AnnotationDisableTenantValidation: "not-valid",
					},
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "foobar",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="openshift-example", level="error"}[5m])) by (job) > 0.1`,
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
			spec: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rule",
					Namespace: "example",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Expr: "invalid",
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
			spec: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rule",
					Namespace: "example",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Expr: `{kubernetes_namespace_name="example", level="error"}`,
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
			spec: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rule",
					Namespace: "example",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Expr: `sum(rate({level="error"}[5m])) by (job) > 0.1`,
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
			desc: "matcher does not match RecordingRule namespace",
			spec: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rule",
					Namespace: "example",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="other-ns", level="error"}[5m])) by (job) > 0.1`,
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
			desc: "query contains both namespace labels",
			spec: &logstorev1.RecordingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "recording-rule",
					Namespace: "example",
				},
				Spec: logstorev1.RecordingRuleSpec{
					TenantID: "application",
					Groups: []*logstorev1.RecordingRuleGroup{
						{
							Rules: []*logstorev1.RecordingRuleGroupSpec{
								{
									Expr: `sum(rate({kubernetes_namespace_name="example", k8s_namespace_name="example", level="error"}[5m])) by (job) > 0.1`,
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
			errors := RecordingRuleValidator(ctx, tc.spec)
			require.Equal(t, tc.wantErrors, errors)
		})
	}
}
