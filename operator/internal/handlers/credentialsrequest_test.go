package handlers

import (
	"context"
	"testing"

	cloudcredentialv1 "github.com/openshift/cloud-credential-operator/pkg/apis/cloudcredential/v1"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/config"
	"example.com/acme/logstore/operator/internal/external/k8s/k8sfakes"
)

func credentialsRequestFakeClient(cr *cloudcredentialv1.CredentialsRequest, logstorestack *logstorev1.LogstoreStack) *k8sfakes.FakeClient {
	k := &k8sfakes.FakeClient{}
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		switch object.(type) {
		case *cloudcredentialv1.CredentialsRequest:
			if cr == nil {
				return errors.NewNotFound(schema.GroupResource{}, name.Name)
			}
			k.SetClientObject(object, cr)
		case *logstorev1.LogstoreStack:
			if logstorestack == nil {
				return errors.NewNotFound(schema.GroupResource{}, name.Name)
			}
			k.SetClientObject(object, logstorestack)
		}
		return nil
	}

	return k
}

func TestCreateUpdateDeleteCredentialsRequest_CreateNewResource(t *testing.T) {
	wantServiceAccountNames := []string{
		"my-stack",
		"my-stack-ruler",
	}

	logstorestack := &logstorev1.LogstoreStack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "ns",
		},
	}

	k := credentialsRequestFakeClient(nil, logstorestack)
	req := ctrl.Request{
		NamespacedName: client.ObjectKey{Name: "my-stack", Namespace: "ns"},
	}

	tokenCCOAuth := &config.TokenCCOAuthConfig{
		AWS: &config.AWSEnvironment{
			RoleARN: "a-role-arn",
		},
	}

	err := CreateUpdateDeleteCredentialsRequest(context.Background(), logger, scheme, tokenCCOAuth, k, req)
	require.NoError(t, err)
	require.Equal(t, 1, k.CreateCallCount())

	_, obj, _ := k.CreateArgsForCall(0)
	credReq, ok := obj.(*cloudcredentialv1.CredentialsRequest)
	require.True(t, ok)

	require.Equal(t, wantServiceAccountNames, credReq.Spec.ServiceAccountNames)
}

func TestCreateUpdateDeleteCredentialsRequest_CreateNewResourceAzure(t *testing.T) {
	wantRegion := "test-region"

	logstorestack := &logstorev1.LogstoreStack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "ns",
		},
	}

	k := credentialsRequestFakeClient(nil, logstorestack)
	req := ctrl.Request{
		NamespacedName: client.ObjectKey{Name: "my-stack", Namespace: "ns"},
	}

	tokenCCOAuth := &config.TokenCCOAuthConfig{
		Azure: &config.AzureEnvironment{
			ClientID:       "test-client-id",
			SubscriptionID: "test-tenant-id",
			TenantID:       "test-subscription-id",
			Region:         "test-region",
		},
	}

	err := CreateUpdateDeleteCredentialsRequest(context.Background(), logger, scheme, tokenCCOAuth, k, req)
	require.NoError(t, err)

	require.Equal(t, 1, k.CreateCallCount())
	_, obj, _ := k.CreateArgsForCall(0)
	credReq, ok := obj.(*cloudcredentialv1.CredentialsRequest)
	require.True(t, ok)

	providerSpec := &cloudcredentialv1.AzureProviderSpec{}
	require.NoError(t, cloudcredentialv1.Codec.DecodeProviderSpec(credReq.Spec.ProviderSpec, providerSpec))

	require.Equal(t, wantRegion, providerSpec.AzureRegion)
}

func TestCreateUpdateDeleteCredentialsRequest_Update_WhenCredentialsRequestExist(t *testing.T) {
	req := ctrl.Request{
		NamespacedName: client.ObjectKey{Name: "my-stack", Namespace: "ns"},
	}

	tokenCCOAuth := &config.TokenCCOAuthConfig{
		AWS: &config.AWSEnvironment{
			RoleARN: "a-role-arn",
		},
	}

	cr := &cloudcredentialv1.CredentialsRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "ns",
		},
	}
	logstorestack := &logstorev1.LogstoreStack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "ns",
		},
	}

	k := credentialsRequestFakeClient(cr, logstorestack)

	err := CreateUpdateDeleteCredentialsRequest(context.Background(), logger, scheme, tokenCCOAuth, k, req)
	require.NoError(t, err)
	require.Equal(t, 2, k.GetCallCount())
	require.Equal(t, 0, k.CreateCallCount())
	require.Equal(t, 1, k.UpdateCallCount())
}

func TestCreateUpdateDeleteCredentialsRequest_DeleteExisting_WhenNotManagedMode(t *testing.T) {
	req := ctrl.Request{
		NamespacedName: client.ObjectKey{Name: "my-stack", Namespace: "ns"},
	}

	tokenCCOAuth := &config.TokenCCOAuthConfig{
		AWS: &config.AWSEnvironment{
			RoleARN: "a-role-arn",
		},
	}

	cr := &cloudcredentialv1.CredentialsRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "ns",
		},
	}
	logstorestack := &logstorev1.LogstoreStack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "ns",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Storage: logstorev1.ObjectStorageSpec{
				Secret: logstorev1.ObjectStorageSecretSpec{
					CredentialMode: logstorev1.CredentialModeStatic,
				},
			},
		},
	}

	k := credentialsRequestFakeClient(cr, logstorestack)

	err := CreateUpdateDeleteCredentialsRequest(context.Background(), logger, scheme, tokenCCOAuth, k, req)
	require.NoError(t, err)
	require.Equal(t, 2, k.GetCallCount())
	require.Equal(t, 0, k.CreateCallCount())
	require.Equal(t, 0, k.UpdateCallCount())
	require.Equal(t, 1, k.DeleteCallCount())
}

func TestCreateUpdateDeleteCredentialsRequest_DoNothing_WhenNotManagedMode(t *testing.T) {
	req := ctrl.Request{
		NamespacedName: client.ObjectKey{Name: "my-stack", Namespace: "ns"},
	}

	tokenCCOAuth := &config.TokenCCOAuthConfig{
		AWS: &config.AWSEnvironment{
			RoleARN: "a-role-arn",
		},
	}

	logstorestack := &logstorev1.LogstoreStack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "ns",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Storage: logstorev1.ObjectStorageSpec{
				Secret: logstorev1.ObjectStorageSecretSpec{
					CredentialMode: logstorev1.CredentialModeStatic,
				},
			},
		},
	}

	k := credentialsRequestFakeClient(nil, logstorestack)

	err := CreateUpdateDeleteCredentialsRequest(context.Background(), logger, scheme, tokenCCOAuth, k, req)
	require.NoError(t, err)
	require.Equal(t, 2, k.GetCallCount())
	require.Equal(t, 0, k.CreateCallCount())
	require.Equal(t, 0, k.UpdateCallCount())
	require.Equal(t, 0, k.DeleteCallCount())
}
