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

var ltt = []struct {
	desc string
	spec logstorev1.LogstoreStack
	err  *apierrors.StatusError
}{
	{
		desc: "valid spec - no status",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Storage: logstorev1.ObjectStorageSpec{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       logstorev1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-13",
						},
					},
				},
			},
		},
	},
	{
		desc: "valid spec - with status",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Storage: logstorev1.ObjectStorageSpec{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       logstorev1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-13",
						},
					},
				},
			},
			Status: logstorev1.LogstoreStackStatus{
				Storage: logstorev1.LogstoreStackStorageStatus{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       logstorev1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-13",
						},
					},
				},
			},
		},
	},
	{
		desc: "not unique schema effective dates",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Storage: logstorev1.ObjectStorageSpec{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       logstorev1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-11",
						},
					},
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec").Child("storage").Child("schemas").Index(1).Child("effectiveDate"),
					"2020-10-11",
					logstorev1.ErrEffectiveDatesNotUnique.Error(),
				),
			},
		),
	},
	{
		desc: "schema effective dates bad format",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Storage: logstorev1.ObjectStorageSpec{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "2020/10/11",
						},
					},
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec").Child("storage").Child("schemas").Index(0).Child("effectiveDate"),
					"2020/10/11",
					logstorev1.ErrParseEffectiveDates.Error(),
				),
			},
		),
	},
	{
		desc: "missing valid starting date",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Storage: logstorev1.ObjectStorageSpec{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "9000-10-10",
						},
					},
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec").Child("storage").Child("schemas"),
					[]logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "9000-10-10",
						},
					},
					logstorev1.ErrMissingValidStartDate.Error(),
				),
			},
		),
	},
	{
		desc: "retroactively adding schema",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Storage: logstorev1.ObjectStorageSpec{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       logstorev1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-14",
						},
					},
				},
			},
			Status: logstorev1.LogstoreStackStatus{
				Storage: logstorev1.LogstoreStackStorageStatus{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
					},
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec").Child("storage").Child("schemas"),
					logstorev1.ObjectStorageSchema{
						Version:       logstorev1.ObjectStorageSchemaV12,
						EffectiveDate: "2020-10-14",
					},
					logstorev1.ErrSchemaRetroactivelyAdded.Error(),
				),
			},
		),
	},
	{
		desc: "retroactively removing schema",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Storage: logstorev1.ObjectStorageSpec{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
					},
				},
			},
			Status: logstorev1.LogstoreStackStatus{
				Storage: logstorev1.LogstoreStackStorageStatus{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       logstorev1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-14",
						},
					},
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec").Child("storage").Child("schemas"),
					[]logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
					},
					logstorev1.ErrSchemaRetroactivelyRemoved.Error(),
				),
			},
		),
	},
	{
		desc: "retroactively changing schema",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Storage: logstorev1.ObjectStorageSpec{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-11",
						},
					},
				},
			},
			Status: logstorev1.LogstoreStackStatus{
				Storage: logstorev1.LogstoreStackStorageStatus{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
					},
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec").Child("storage").Child("schemas"),
					logstorev1.ObjectStorageSchema{
						Version:       logstorev1.ObjectStorageSchemaV12,
						EffectiveDate: "2020-10-11",
					},
					logstorev1.ErrSchemaRetroactivelyChanged.Error(),
				),
			},
		),
	},
	{
		desc: "valid replication zones",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Storage: logstorev1.ObjectStorageSpec{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-11",
						},
					},
				},
				Replication: &logstorev1.ReplicationSpec{
					Zones: []logstorev1.ZoneSpec{
						{
							TopologyKey: "zone",
						},
					},
					Factor: 1,
				},
			},
		},
	},
	{
		desc: "using both replication and replicationFactor",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Storage: logstorev1.ObjectStorageSpec{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-11",
						},
					},
				},
				ReplicationFactor: 2,
				Replication: &logstorev1.ReplicationSpec{
					Zones: []logstorev1.ZoneSpec{
						{
							TopologyKey: "zone",
						},
						{
							TopologyKey: "region",
						},
						{
							TopologyKey: "planet",
						},
					},
					Factor: 1,
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec", "replicationFactor"),
					2,
					logstorev1.ErrReplicationSpecConflict.Error(),
				),
			},
		),
	},
	{
		desc: "using default InstanceAddrType and enableIPv6",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				HashRing: &logstorev1.HashRingSpec{
					Type: logstorev1.HashRingMemberList,
					MemberList: &logstorev1.MemberListSpec{
						EnableIPv6:       true,
						InstanceAddrType: logstorev1.InstanceAddrDefault,
					},
				},
				Storage: logstorev1.ObjectStorageSpec{
					Schemas: []logstorev1.ObjectStorageSchema{
						{
							Version:       logstorev1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-11",
						},
					},
				},
				Replication: &logstorev1.ReplicationSpec{
					Zones: []logstorev1.ZoneSpec{
						{
							TopologyKey: "zone",
						},
					},
					Factor: 1,
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec", "hashRing", "memberlist", "instanceAddrType"),
					logstorev1.InstanceAddrDefault,
					logstorev1.ErrIPv6InstanceAddrTypeNotAllowed.Error(),
				),
			},
		),
	},
	{
		desc: "logstorestack with custom OTLP configuration with a global stream label",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Limits: &logstorev1.LimitsSpec{
					Global: &logstorev1.LimitsTemplateSpec{
						OTLP: &logstorev1.OTLPSpec{
							StreamLabels: &logstorev1.OTLPStreamLabelSpec{
								ResourceAttributes: []logstorev1.OTLPAttributeReference{
									{
										Name: "global.stream.label",
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
					Mode: logstorev1.Static,
				},
			},
		},
		err: nil,
	},
	{
		desc: "logstorestack with custom OTLP configuration with a global stream label and a tenant with no stream label",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Limits: &logstorev1.LimitsSpec{
					Global: &logstorev1.LimitsTemplateSpec{
						OTLP: &logstorev1.OTLPSpec{
							StreamLabels: &logstorev1.OTLPStreamLabelSpec{
								ResourceAttributes: []logstorev1.OTLPAttributeReference{
									{
										Name: "global.stream.label",
									},
								},
							},
						},
					},
					Tenants: map[string]logstorev1.PerTenantLimitsTemplateSpec{
						"test-tenant": {
							OTLP: &logstorev1.OTLPSpec{},
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
					Mode: logstorev1.Static,
					Authentication: []logstorev1.AuthenticationSpec{
						{
							TenantName: "test-tenant",
						},
					},
				},
			},
		},
		err: nil,
	},
	{
		desc: "logstorestack with custom OTLP configuration with no global stream label and a tenant with a stream label",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Limits: &logstorev1.LimitsSpec{
					Tenants: map[string]logstorev1.PerTenantLimitsTemplateSpec{
						"test-tenant": {
							OTLP: &logstorev1.OTLPSpec{
								StreamLabels: &logstorev1.OTLPStreamLabelSpec{
									ResourceAttributes: []logstorev1.OTLPAttributeReference{
										{
											Name: "tenant.stream.label",
										},
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
					Mode: logstorev1.Static,
					Authentication: []logstorev1.AuthenticationSpec{
						{
							TenantName: "test-tenant",
						},
					},
				},
			},
		},
		err: nil,
	},
	{
		desc: "logstorestack with custom OTLP configuration missing a global stream label",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Limits: &logstorev1.LimitsSpec{
					Global: &logstorev1.LimitsTemplateSpec{
						OTLP: &logstorev1.OTLPSpec{},
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
					Mode: logstorev1.Static,
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec", "limits", "global", "otlp", "streamLabels", "resourceAttributes"),
					nil,
					logstorev1.ErrOTLPGlobalNoStreamLabel.Error(),
				),
			},
		),
	},
	{
		desc: "logstorestack with custom OTLP configuration missing a tenant",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Limits: &logstorev1.LimitsSpec{
					Tenants: map[string]logstorev1.PerTenantLimitsTemplateSpec{
						"test-tenant": {
							OTLP: &logstorev1.OTLPSpec{
								StreamLabels: &logstorev1.OTLPStreamLabelSpec{
									ResourceAttributes: []logstorev1.OTLPAttributeReference{
										{
											Name: "tenant.stream.label",
										},
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
					Mode: logstorev1.Static,
					Authentication: []logstorev1.AuthenticationSpec{
						{
							TenantName: "test-tenant",
						},
						{
							TenantName: "second-tenant",
						},
					},
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec", "limits", "tenants", "second-tenant", "otlp"),
					nil,
					logstorev1.ErrOTLPTenantMissing.Error(),
				),
			},
		),
	},
	{
		desc: "logstorestack with custom OTLP configuration with a tenant without stream label",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Limits: &logstorev1.LimitsSpec{
					Tenants: map[string]logstorev1.PerTenantLimitsTemplateSpec{
						"test-tenant": {
							OTLP: &logstorev1.OTLPSpec{},
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
					Mode: logstorev1.Static,
					Authentication: []logstorev1.AuthenticationSpec{
						{
							TenantName: "test-tenant",
						},
					},
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec", "limits", "tenants", "test-tenant", "otlp", "streamLabels", "resourceAttributes"),
					nil,
					logstorev1.ErrOTLPTenantNoStreamLabel.Error(),
				),
			},
		),
	},
	{
		desc: "logstorestack with custom OTLP configuration listing an attribute as both stream-label and drop",
		spec: logstorev1.LogstoreStack{
			Spec: logstorev1.LogstoreStackSpec{
				Limits: &logstorev1.LimitsSpec{
					Global: &logstorev1.LimitsTemplateSpec{
						OTLP: &logstorev1.OTLPSpec{
							StreamLabels: &logstorev1.OTLPStreamLabelSpec{
								ResourceAttributes: []logstorev1.OTLPAttributeReference{
									{
										Name: "global.stream.label",
									},
								},
							},
							Drop: &logstorev1.OTLPMetadataSpec{
								ResourceAttributes: []logstorev1.OTLPAttributeReference{
									{
										Name: "global.stream.label",
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
					Mode: logstorev1.Static,
				},
			},
		},
		err: apierrors.NewInvalid(
			schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
			"testing-stack",
			field.ErrorList{
				field.Invalid(
					field.NewPath("spec", "limits", "global", "otlp", "drop", "resourceAttributes").Index(0),
					"global.stream.label",
					logstorev1.ErrOTLPInvalidDrop.Error(),
				),
			},
		),
	},
}

func TestLogstoreStackValidationWebhook_ValidateCreate(t *testing.T) {
	for _, tc := range ltt {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			l := &logstorev1.LogstoreStack{
				ObjectMeta: metav1.ObjectMeta{
					Name: "testing-stack",
				},
				Spec: tc.spec.Spec,
			}
			ctx := context.Background()

			v := &validation.LogstoreStackValidator{}
			_, err := v.ValidateCreate(ctx, l)
			if err != nil {
				require.Equal(t, tc.err, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestLogstoreStackValidationWebhook_ValidateUpdate(t *testing.T) {
	for _, tc := range ltt {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			l := &logstorev1.LogstoreStack{
				ObjectMeta: metav1.ObjectMeta{
					Name: "testing-stack",
				},
				Spec: tc.spec.Spec,
			}
			ctx := context.Background()

			v := &validation.LogstoreStackValidator{}
			_, err := v.ValidateUpdate(ctx, &logstorev1.LogstoreStack{}, l)
			if err != nil {
				require.Equal(t, tc.err, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
