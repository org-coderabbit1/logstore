package v1_test

import (
	"testing"

	v1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"github.com/stretchr/testify/require"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

var ltt = []struct {
	desc string
	spec v1.LogstoreStack
	err  *apierrors.StatusError
}{
	{
		desc: "valid spec - no status",
		spec: v1.LogstoreStack{
			Spec: v1.LogstoreStackSpec{
				Storage: v1.ObjectStorageSpec{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       v1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-13",
						},
					},
				},
			},
		},
	},
	{
		desc: "valid spec - with status",
		spec: v1.LogstoreStack{
			Spec: v1.LogstoreStackSpec{
				Storage: v1.ObjectStorageSpec{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       v1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-13",
						},
					},
				},
			},
			Status: v1.LogstoreStackStatus{
				Storage: v1.LogstoreStackStorageStatus{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       v1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-13",
						},
					},
				},
			},
		},
	},
	{
		desc: "not unique schema effective dates",
		spec: v1.LogstoreStack{
			Spec: v1.LogstoreStackSpec{
				Storage: v1.ObjectStorageSpec{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       v1.ObjectStorageSchemaV12,
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
					field.NewPath("Spec").Child("Storage").Child("Schemas").Index(1).Child("EffectiveDate"),
					"2020-10-11",
					v1.ErrEffectiveDatesNotUnique.Error(),
				),
			},
		),
	},
	{
		desc: "schema effective dates bad format",
		spec: v1.LogstoreStack{
			Spec: v1.LogstoreStackSpec{
				Storage: v1.ObjectStorageSpec{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
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
					field.NewPath("Spec").Child("Storage").Child("Schemas").Index(0).Child("EffectiveDate"),
					"2020/10/11",
					v1.ErrParseEffectiveDates.Error(),
				),
			},
		),
	},
	{
		desc: "missing valid starting date",
		spec: v1.LogstoreStack{
			Spec: v1.LogstoreStackSpec{
				Storage: v1.ObjectStorageSpec{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
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
					field.NewPath("Spec").Child("Storage").Child("Schemas"),
					[]v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
							EffectiveDate: "9000-10-10",
						},
					},
					v1.ErrMissingValidStartDate.Error(),
				),
			},
		),
	},
	{
		desc: "retroactively adding schema",
		spec: v1.LogstoreStack{
			Spec: v1.LogstoreStackSpec{
				Storage: v1.ObjectStorageSpec{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       v1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-14",
						},
					},
				},
			},
			Status: v1.LogstoreStackStatus{
				Storage: v1.LogstoreStackStorageStatus{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
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
					field.NewPath("Spec").Child("Storage").Child("Schemas"),
					v1.ObjectStorageSchema{
						Version:       v1.ObjectStorageSchemaV12,
						EffectiveDate: "2020-10-14",
					},
					v1.ErrSchemaRetroactivelyAdded.Error(),
				),
			},
		),
	},
	{
		desc: "retroactively removing schema",
		spec: v1.LogstoreStack{
			Spec: v1.LogstoreStackSpec{
				Storage: v1.ObjectStorageSpec{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
					},
				},
			},
			Status: v1.LogstoreStackStatus{
				Storage: v1.LogstoreStackStorageStatus{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
						{
							Version:       v1.ObjectStorageSchemaV12,
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
					field.NewPath("Spec").Child("Storage").Child("Schemas"),
					[]v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
							EffectiveDate: "2020-10-11",
						},
					},
					v1.ErrSchemaRetroactivelyRemoved.Error(),
				),
			},
		),
	},
	{
		desc: "retroactively changing schema",
		spec: v1.LogstoreStack{
			Spec: v1.LogstoreStackSpec{
				Storage: v1.ObjectStorageSpec{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV12,
							EffectiveDate: "2020-10-11",
						},
					},
				},
			},
			Status: v1.LogstoreStackStatus{
				Storage: v1.LogstoreStackStorageStatus{
					Schemas: []v1.ObjectStorageSchema{
						{
							Version:       v1.ObjectStorageSchemaV11,
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
					field.NewPath("Spec").Child("Storage").Child("Schemas"),
					v1.ObjectStorageSchema{
						Version:       v1.ObjectStorageSchemaV12,
						EffectiveDate: "2020-10-11",
					},
					v1.ErrSchemaRetroactivelyChanged.Error(),
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
			l := v1.LogstoreStack{
				ObjectMeta: metav1.ObjectMeta{
					Name: "testing-stack",
				},
				Spec: tc.spec.Spec,
			}

			err := l.ValidateCreate()
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
			l := v1.LogstoreStack{
				ObjectMeta: metav1.ObjectMeta{
					Name: "testing-stack",
				},
				Spec: tc.spec.Spec,
			}

			err := l.ValidateUpdate(&v1.LogstoreStack{})
			if err != nil {
				require.Equal(t, tc.err, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
