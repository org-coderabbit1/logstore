package validation_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/validation"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
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
}

func TestLogstoreStackValidationWebhook_ValidateCreate(t *testing.T) {
	for _, tc := range ltt {
		tc := tc
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
			err := v.ValidateCreate(ctx, l)
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
		tc := tc
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
			err := v.ValidateUpdate(ctx, &logstorev1.LogstoreStack{}, l)
			if err != nil {
				require.Equal(t, tc.err, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
