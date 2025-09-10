package openshift

import (
	"testing"

	"github.com/stretchr/testify/require"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
)

func TestValidateOTLPInvalidDrop(t *testing.T) {
	tt := []struct {
		desc string
		spec logstorev1.LogstoreStack
		err  *apierrors.StatusError
	}{
		{
			desc: "logstorestack using openshift-logging tenancy mode trying to remove a required stream label",
			spec: logstorev1.LogstoreStack{
				Spec: logstorev1.LogstoreStackSpec{
					Limits: &logstorev1.LimitsSpec{
						Global: &logstorev1.LimitsTemplateSpec{
							OTLP: &logstorev1.OTLPSpec{
								Drop: &logstorev1.OTLPMetadataSpec{
									ResourceAttributes: []logstorev1.OTLPAttributeReference{
										{
											Name: "kubernetes.namespace_name",
										},
									},
								},
							},
						},
					},
					Storage: logstorev1.ObjectStorageSpec{
						Schemas: []logstorev1.ObjectStorageSchema{
							{
								Version:       logstorev1.ObjectStorageSchemaV13,
								EffectiveDate: "2024-10-22",
							},
						},
					},
					Tenants: &logstorev1.TenantsSpec{
						Mode: logstorev1.OpenshiftLogging,
					},
				},
			},
			err: apierrors.NewInvalid(
				schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
				"testing-stack",
				field.ErrorList{
					field.Invalid(
						field.NewPath("spec", "limits", "global", "otlp", "drop", "resourceAttributes").Index(0).Child("name"),
						"kubernetes.namespace_name",
						logstorev1.ErrOTLPInvalidDrop.Error(),
					),
				},
			),
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			errList := ValidateOTLPInvalidDrop(&tc.spec.Spec)
			if tc.err != nil {
				testErr := apierrors.NewInvalid(
					schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
					"testing-stack",
					errList,
				)
				require.Equal(t, tc.err, testErr)
			} else {
				require.Nil(t, errList)
			}
		})
	}
}
