package status_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s/k8sfakes"
	"example.com/acme/logstore/operator/internal/status"
)

func TestSetStorageSchemaStatus_WhenGetLogstoreStackReturnsError_ReturnError(t *testing.T) {
	k := &k8sfakes.FakeClient{}

	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		return apierrors.NewBadRequest("something wasn't found")
	}

	err := status.SetStorageSchemaStatus(context.TODO(), k, r, []logstorev1.ObjectStorageSchema{})
	require.Error(t, err)
}

func TestSetStorageSchemaStatus_WhenGetLogstoreStackReturnsNotFound_DoNothing(t *testing.T) {
	k := &k8sfakes.FakeClient{}

	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	err := status.SetStorageSchemaStatus(context.TODO(), k, r, []logstorev1.ObjectStorageSchema{})
	require.NoError(t, err)
}

func TestSetStorageSchemaStatus_WhenStorageStatusExists_OverwriteStorageStatus(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}

	k.StatusStub = func() client.StatusWriter { return sw }

	s := logstorev1.LogstoreStack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
		Status: logstorev1.LogstoreStackStatus{
			Storage: logstorev1.LogstoreStackStorageStatus{
				CredentialMode: logstorev1.CredentialModeStatic,
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
			},
		},
	}

	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	schemas := []logstorev1.ObjectStorageSchema{
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2020-10-11",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-10-11",
		},
	}

	expected := logstorev1.LogstoreStackStorageStatus{
		CredentialMode: logstorev1.CredentialModeStatic,
		Schemas: []logstorev1.ObjectStorageSchema{
			{
				Version:       logstorev1.ObjectStorageSchemaV11,
				EffectiveDate: "2020-10-11",
			},
			{
				Version:       logstorev1.ObjectStorageSchemaV12,
				EffectiveDate: "2021-10-11",
			},
		},
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, &s)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	sw.UpdateStub = func(_ context.Context, obj client.Object, _ ...client.SubResourceUpdateOption) error {
		stack := obj.(*logstorev1.LogstoreStack)
		require.Equal(t, expected, stack.Status.Storage)
		return nil
	}

	err := status.SetStorageSchemaStatus(context.TODO(), k, r, schemas)
	require.NoError(t, err)
	require.NotZero(t, k.StatusCallCount())
	require.NotZero(t, sw.UpdateCallCount())
}
