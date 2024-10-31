package validation_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/validation"
)

var att = []struct {
	desc string
	spec logstorev1.AlertingRuleSpec
	err  *apierrors.StatusError
}{
	{
		desc: "valid spec",
		spec: logstorev1.AlertingRuleSpec{
			Groups: []*logstorev1.AlertingRuleGroup{
				{
					Name:     "first",
					Interval: logstorev1.PrometheusDuration("1m"),
					Limit:    10,
					Rules: []*logstorev1.AlertingRuleGroupSpec{
						{
							Alert: "first-alert",
							For:   logstorev1.PrometheusDuration("10m"),
							Expr:  `sum(rate({app="foo", env="production"} |= "error" [5m])) by (job)`,
							Annotations: map[string]string{
								"annot": "something",
							},
							Labels: map[string]string{
								"severity": "critical",
							},
						},
						{
							Alert: "second-alert",
							For:   logstorev1.PrometheusDuration("10m"),
							Expr:  `sum(rate({app="foo", env="stage"} |= "error" [5m])) by (job)`,
							Annotations: map[string]string{
								"env": "something",
							},
							Labels: map[string]string{
								"severity": "warning",
							},
						},
					},
				},
				{
					Name:     "second",
					Interval: logstorev1.PrometheusDuration("1m"),
					Limit:    10,
					Rules: []*logstorev1.AlertingRuleGroupSpec{
						{
							Alert: "third-alert",
							For:   logstorev1.PrometheusDuration("10m"),
							Expr:  `sum(rate({app="foo", env="production"} |= "error" [5m])) by (job)`,
							Annotations: map[string]string{
								"annot": "something",
							},
							Labels: map[string]string{
								"severity": "critical",
							},
						},
						{
							Alert: "fourth-alert",
							For:   logstorev1.PrometheusDuration("10m"),
							Expr:  `sum(rate({app="foo", env="stage"} |= "error" [5m])) by (job)`,
							Annotations: map[string]string{
								"env": "something",
							},
							Labels: map[string]string{
								"severity": "warning",
							},
						},
					},
				},
			},
		},
	},
	{
		desc: "not unique group names",
		spec: logstorev1.AlertingRuleSpec{
			Groups: []*logstorev1.AlertingRuleGroup{
				{
					Name:     "first",
					Interval: logstorev1.PrometheusDuration("1m"),
				},
				{
					Name:     "first",
					Interval: logstorev1.PrometheusDuration("1m"),
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "AlertingRule"},
			"testing-rule",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec").Child("groups").Index(1).Child("name"),
					"first",
					logstorev1.ErrGroupNamesNotUnique.Error(),
				),
			},
		),
	},
	{
		desc: "parse eval interval err",
		spec: logstorev1.AlertingRuleSpec{
			Groups: []*logstorev1.AlertingRuleGroup{
				{
					Name:     "first",
					Interval: logstorev1.PrometheusDuration("1mo"),
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "AlertingRule"},
			"testing-rule",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec").Child("groups").Index(0).Child("interval"),
					"1mo",
					logstorev1.ErrParseEvaluationInterval.Error(),
				),
			},
		),
	},
	{
		desc: "parse for interval err",
		spec: logstorev1.AlertingRuleSpec{
			Groups: []*logstorev1.AlertingRuleGroup{
				{
					Name:     "first",
					Interval: logstorev1.PrometheusDuration("1m"),
					Rules: []*logstorev1.AlertingRuleGroupSpec{
						{
							Alert: "an-alert",
							For:   logstorev1.PrometheusDuration("10years"),
							Expr:  `sum(rate({label="value"}[1m]))`,
						},
					},
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "AlertingRule"},
			"testing-rule",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec").Child("groups").Index(0).Child("rules").Index(0).Child("for"),
					"10years",
					logstorev1.ErrParseAlertForPeriod.Error(),
				),
			},
		),
	},
	{
		desc: "parse LogQL expression err",
		spec: logstorev1.AlertingRuleSpec{
			Groups: []*logstorev1.AlertingRuleGroup{
				{
					Name:     "first",
					Interval: logstorev1.PrometheusDuration("1m"),
					Rules: []*logstorev1.AlertingRuleGroupSpec{
						{
							Expr: "this is not a valid expression",
						},
					},
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "AlertingRule"},
			"testing-rule",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec").Child("groups").Index(0).Child("rules").Index(0).Child("expr"),
					"this is not a valid expression",
					logstorev1.ErrParseLogQLExpression.Error(),
				),
			},
		),
	},
	{
		desc: "LogQL not sample-expression",
		spec: logstorev1.AlertingRuleSpec{
			Groups: []*logstorev1.AlertingRuleGroup{
				{
					Name:     "first",
					Interval: logstorev1.PrometheusDuration("1m"),
					Rules: []*logstorev1.AlertingRuleGroupSpec{
						{
							Expr: `{message=~".+"}`,
						},
					},
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "AlertingRule"},
			"testing-rule",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec").Child("groups").Index(0).Child("rules").Index(0).Child("expr"),
					`{message=~".+"}`,
					logstorev1.ErrParseLogQLNotSample.Error(),
				),
			},
		),
	},
}

func TestAlertingRuleValidationWebhook_ValidateCreate(t *testing.T) {
	for _, tc := range att {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			l := &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name: "testing-rule",
				},
				Spec: tc.spec,
			}
			ctx := context.Background()

			v := &validation.AlertingRuleValidator{}
			_, err := v.ValidateCreate(ctx, l)
			if err != nil {
				require.Equal(t, tc.err, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAlertingRuleValidationWebhook_ValidateUpdate(t *testing.T) {
	for _, tc := range att {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			l := &logstorev1.AlertingRule{
				ObjectMeta: metav1.ObjectMeta{
					Name: "testing-rule",
				},
				Spec: tc.spec,
			}
			ctx := context.Background()

			v := &validation.AlertingRuleValidator{}
			_, err := v.ValidateUpdate(ctx, &logstorev1.AlertingRule{}, l)
			if err != nil {
				require.Equal(t, tc.err, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
