package handlers_test

import (
	"context"
	"errors"
	"flag"
	"io"
	"os"
	"testing"

	configv1 "example.com/acme/logstore/operator/apis/config/v1"
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s/k8sfakes"
	"example.com/acme/logstore/operator/internal/handlers"
	"example.com/acme/logstore/operator/internal/status"

	"github.com/ViaQ/logerr/v2/log"
	"github.com/go-logr/logr"
	routev1 "github.com/openshift/api/route/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	logger logr.Logger

	scheme       = runtime.NewScheme()
	featureGates = configv1.FeatureGates{
		ServiceMonitors:            false,
		ServiceMonitorTLSEndpoints: false,
		BuiltInCertManagement: configv1.BuiltInCertManagement{
			Enabled:        true,
			CACertValidity: "10m",
			CACertRefresh:  "5m",
			CertValidity:   "2m",
			CertRefresh:    "1m",
		},
	}

	defaultSecret = corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "some-stack-secret",
			Namespace: "some-ns",
		},
		Data: map[string][]byte{
			"endpoint":          []byte("s3://your-endpoint"),
			"region":            []byte("a-region"),
			"bucketnames":       []byte("bucket1,bucket2"),
			"access_key_id":     []byte("a-secret-id"),
			"access_key_secret": []byte("a-secret-key"),
		},
	}

	defaultGatewaySecret = corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "some-stack-gateway-secret",
			Namespace: "some-ns",
		},
		Data: map[string][]byte{
			"clientID":     []byte("client-123"),
			"clientSecret": []byte("client-secret-xyz"),
			"issuerCAPath": []byte("/tmp/test/ca.pem"),
		},
	}

	rulesCM = corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind: "ConfigMap",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack-rules-0",
			Namespace: "some-ns",
		},
	}

	rulerSS = appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind: "StatefulSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack-ruler",
			Namespace: "some-ns",
		},
	}

	invalidSecret = corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "some-stack-secret",
			Namespace: "some-ns",
		},
		Data: map[string][]byte{},
	}

	invalidCAConfigMap = corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "some-stack-ca-configmap",
			Namespace: "some-ns",
		},
		Data: map[string]string{},
	}
)

func TestMain(m *testing.M) {
	testing.Init()
	flag.Parse()

	if testing.Verbose() {
		logger = log.NewLogger("testing", log.WithVerbosity(5))
	} else {
		logger = log.NewLogger("testing", log.WithOutput(io.Discard))
	}

	// Register the clientgo and CRD schemes
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(routev1.AddToScheme(scheme))
	utilruntime.Must(logstorev1.AddToScheme(scheme))

	os.Exit(m.Run())
}

func TestCreateOrUpdateLogstoreStack_WhenGetReturnsNotFound_DoesNotError(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	k.GetStub = func(ctx context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)
	require.NoError(t, err)

	// make sure create was NOT called because the Get failed
	require.Zero(t, k.CreateCallCount())
}

func TestCreateOrUpdateLogstoreStack_WhenGetReturnsAnErrorOtherThanNotFound_ReturnsTheError(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	badRequestErr := apierrors.NewBadRequest("you do not belong here")
	k.GetStub = func(ctx context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		return badRequestErr
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)

	require.Equal(t, badRequestErr, errors.Unwrap(err))

	// make sure create was NOT called because the Get failed
	require.Zero(t, k.CreateCallCount())
}

func TestCreateOrUpdateLogstoreStack_SetsNamespaceOnAllObjects(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	stack := logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
			Tenants: &logstorev1.TenantsSpec{
				Mode: "dynamic",
				Authentication: []logstorev1.AuthenticationSpec{
					{
						TenantName: "test",
						TenantID:   "1234",
						OIDC: &logstorev1.OIDCSpec{
							Secret: &logstorev1.TenantSecretSpec{
								Name: defaultGatewaySecret.Name,
							},
						},
					},
				},
				Authorization: &logstorev1.AuthorizationSpec{
					OPA: &logstorev1.OPASpec{
						URL: "some-url",
					},
				},
			},
		},
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, out client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(out, &stack)
			return nil
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(out, &defaultSecret)
			return nil
		}
		if defaultGatewaySecret.Name == name.Name {
			k.SetClientObject(out, &defaultGatewaySecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	k.CreateStub = func(_ context.Context, o client.Object, _ ...client.CreateOption) error {
		assert.Equal(t, r.Namespace, o.GetNamespace())
		return nil
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)
	require.NoError(t, err)

	// make sure create was called
	require.NotZero(t, k.CreateCallCount())
}

func TestCreateOrUpdateLogstoreStack_SetsOwnerRefOnAllObjects(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	stack := logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "someStack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
			Tenants: &logstorev1.TenantsSpec{
				Mode: "dynamic",
				Authentication: []logstorev1.AuthenticationSpec{
					{
						TenantName: "test",
						TenantID:   "1234",
						OIDC: &logstorev1.OIDCSpec{
							Secret: &logstorev1.TenantSecretSpec{
								Name: defaultGatewaySecret.Name,
							},
						},
					},
				},
				Authorization: &logstorev1.AuthorizationSpec{
					OPA: &logstorev1.OPASpec{
						URL: "some-url",
					},
				},
			},
		},
	}

	// Create looks up the CR first, so we need to return our fake stack
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, &stack)
			return nil
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(object, &defaultSecret)
			return nil
		}
		if defaultGatewaySecret.Name == name.Name {
			k.SetClientObject(object, &defaultGatewaySecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	expected := metav1.OwnerReference{
		APIVersion:         logstorev1.GroupVersion.String(),
		Kind:               stack.Kind,
		Name:               stack.Name,
		UID:                stack.UID,
		Controller:         pointer.BoolPtr(true),
		BlockOwnerDeletion: pointer.BoolPtr(true),
	}

	k.CreateStub = func(_ context.Context, o client.Object, _ ...client.CreateOption) error {
		// OwnerRefs are appended so we have to find ours in the list
		var ref metav1.OwnerReference
		var found bool
		for _, or := range o.GetOwnerReferences() {
			if or.UID == stack.UID {
				found = true
				ref = or
				break
			}
		}

		require.True(t, found, "expected to find a matching ownerRef, but did not")
		require.EqualValues(t, expected, ref)
		return nil
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)
	require.NoError(t, err)

	// make sure create was called
	require.NotZero(t, k.CreateCallCount())
}

func TestCreateOrUpdateLogstoreStack_WhenSetControllerRefInvalid_ContinueWithOtherObjects(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	stack := logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "someStack",
			// Set invalid namespace here, because
			// because cross-namespace controller
			// references are not allowed
			Namespace: "invalid-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
		},
	}

	// Create looks up the CR first, so we need to return our fake stack
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, &stack)
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(object, &defaultSecret)
		}
		return nil
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)

	// make sure error is returned to re-trigger reconciliation
	require.Error(t, err)
}

func TestCreateOrUpdateLogstoreStack_WhenGetReturnsNoError_UpdateObjects(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	stack := logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
		},
	}

	svc := corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack-gossip-ring",
			Namespace: "some-ns",
			Labels: map[string]string{
				"app.kubernetes.io/name":     "logstore",
				"app.kubernetes.io/provider": "openshift",
				"logstore.acme.com/name":      "my-stack",

				// Add custom label to fake semantic not equal
				"test": "test",
			},
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion:         "logstore.acme.com/v1",
					Kind:               "LogstoreStack",
					Name:               "my-stack",
					UID:                "b23f9a38-9672-499f-8c29-15ede74d3ece",
					Controller:         pointer.BoolPtr(true),
					BlockOwnerDeletion: pointer.BoolPtr(true),
				},
			},
		},
		Spec: corev1.ServiceSpec{
			ClusterIP: "None",
			Ports: []corev1.ServicePort{
				{
					Name:     "gossip",
					Port:     7946,
					Protocol: "TCP",
				},
			},
			Selector: map[string]string{
				"app.kubernetes.io/name":     "logstore",
				"app.kubernetes.io/provider": "openshift",
				"logstore.acme.com/name":      "my-stack",
			},
		},
	}

	// Create looks up the CR first, so we need to return our fake stack
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, &stack)
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(object, &defaultSecret)
		}
		if svc.Name == name.Name && svc.Namespace == name.Namespace {
			k.SetClientObject(object, &svc)
		}
		return nil
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)
	require.NoError(t, err)

	// make sure create not called
	require.Zero(t, k.CreateCallCount())

	// make sure update was called
	require.NotZero(t, k.UpdateCallCount())
}

func TestCreateOrUpdateLogstoreStack_WhenCreateReturnsError_ContinueWithOtherObjects(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	stack := logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "someStack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
		},
	}

	// GetStub looks up the CR first, so we need to return our fake stack
	// return NotFound for everything else to trigger create.
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, &stack)
			return nil
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(object, &defaultSecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something is not found")
	}

	// CreateStub returns an error for each resource to trigger reconciliation a new.
	k.CreateStub = func(_ context.Context, o client.Object, _ ...client.CreateOption) error {
		return apierrors.NewTooManyRequestsError("too many create requests")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)

	// make sure error is returned to re-trigger reconciliation
	require.Error(t, err)
}

func TestCreateOrUpdateLogstoreStack_WhenUpdateReturnsError_ContinueWithOtherObjects(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	stack := logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
		},
	}

	svc := corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack-gossip-ring",
			Namespace: "some-ns",
			Labels: map[string]string{
				"app.kubernetes.io/name":     "logstore",
				"app.kubernetes.io/provider": "openshift",
				"logstore.acme.com/name":      "my-stack",

				// Add custom label to fake semantic not equal
				"test": "test",
			},
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion:         "logstore.acme.com/v1",
					Kind:               "LogstoreStack",
					Name:               "someStack",
					UID:                "b23f9a38-9672-499f-8c29-15ede74d3ece",
					Controller:         pointer.BoolPtr(true),
					BlockOwnerDeletion: pointer.BoolPtr(true),
				},
			},
		},
		Spec: corev1.ServiceSpec{
			ClusterIP: "None",
			Ports: []corev1.ServicePort{
				{
					Name:     "gossip",
					Port:     7946,
					Protocol: "TCP",
				},
			},
			Selector: map[string]string{
				"app.kubernetes.io/name":     "logstore",
				"app.kubernetes.io/provider": "openshift",
				"logstore.acme.com/name":      "my-stack",
			},
		},
	}

	// GetStub looks up the CR first, so we need to return our fake stack
	// return NotFound for everything else to trigger create.
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, &stack)
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(object, &defaultSecret)
		}
		if svc.Name == name.Name && svc.Namespace == name.Namespace {
			k.SetClientObject(object, &svc)
		}
		return nil
	}

	// CreateStub returns an error for each resource to trigger reconciliation a new.
	k.UpdateStub = func(_ context.Context, o client.Object, _ ...client.UpdateOption) error {
		return apierrors.NewTooManyRequestsError("too many create requests")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)

	// make sure error is returned to re-trigger reconciliation
	require.Error(t, err)
}

func TestCreateOrUpdateLogstoreStack_WhenMissingSecret_SetDegraded(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	degradedErr := &status.DegradedError{
		Message: "Missing object storage secret",
		Reason:  logstorev1.ReasonMissingObjectStorageSecret,
		Requeue: false,
	}

	stack := &logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
		},
	}

	// GetStub looks up the CR first, so we need to return our fake stack
	// return NotFound for everything else to trigger create.
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, stack)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something is not found")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)

	// make sure error is returned
	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestCreateOrUpdateLogstoreStack_WhenInvalidSecret_SetDegraded(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	degradedErr := &status.DegradedError{
		Message: "Invalid object storage secret contents: missing secret field",
		Reason:  logstorev1.ReasonInvalidObjectStorageSecret,
		Requeue: false,
	}

	stack := &logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: invalidSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
		},
	}

	// GetStub looks up the CR first, so we need to return our fake stack
	// return NotFound for everything else to trigger create.
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, stack)
			return nil
		}
		if name.Name == invalidSecret.Name {
			k.SetClientObject(object, &invalidSecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something is not found")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)

	// make sure error is returned
	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestCreateOrUpdateLogstoreStack_WithInvalidStorageSchema_SetDegraded(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	degradedErr := &status.DegradedError{
		Message: "Invalid object storage schema contents: spec does not contain any schemas",
		Reason:  logstorev1.ReasonInvalidObjectStorageSchema,
		Requeue: false,
	}

	stack := &logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
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
						EffectiveDate: "2021-10-11",
					},
				},
			},
		},
	}

	// GetStub looks up the CR first, so we need to return our fake stack
	// return NotFound for everything else to trigger create.
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, stack)
			return nil
		}
		if name.Name == defaultSecret.Name {
			k.SetClientObject(object, &defaultSecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something is not found")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)

	// make sure error is returned
	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestCreateOrUpdateLogstoreStack_WhenMissingCAConfigMap_SetDegraded(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	degradedErr := &status.DegradedError{
		Message: "Missing object storage CA config map",
		Reason:  logstorev1.ReasonMissingObjectStorageCAConfigMap,
		Requeue: false,
	}

	stack := &logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
				TLS: &logstorev1.ObjectStorageTLSSpec{
					CA: "not-existing",
				},
			},
		},
	}

	// GetStub looks up the CR first, so we need to return our fake stack
	// return NotFound for everything else to trigger create.
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, stack)
			return nil
		}

		if name.Name == defaultSecret.Name {
			k.SetClientObject(object, &defaultSecret)
			return nil
		}

		return apierrors.NewNotFound(schema.GroupResource{}, "something is not found")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)

	// make sure error is returned
	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestCreateOrUpdateLogstoreStack_WhenInvalidCAConfigMap_SetDegraded(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	degradedErr := &status.DegradedError{
		Message: "Invalid object storage CA configmap contents: missing key or no contents",
		Reason:  logstorev1.ReasonInvalidObjectStorageCAConfigMap,
		Requeue: false,
	}

	stack := &logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
				TLS: &logstorev1.ObjectStorageTLSSpec{
					CA: invalidCAConfigMap.Name,
				},
			},
		},
	}

	// GetStub looks up the CR first, so we need to return our fake stack
	// return NotFound for everything else to trigger create.
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, stack)
			return nil
		}
		if name.Name == defaultSecret.Name {
			k.SetClientObject(object, &defaultSecret)
			return nil
		}

		if name.Name == invalidCAConfigMap.Name {
			k.SetClientObject(object, &invalidCAConfigMap)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something is not found")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)

	// make sure error is returned
	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestCreateOrUpdateLogstoreStack_WhenInvalidTenantsConfiguration_SetDegraded(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	degradedErr := &status.DegradedError{
		Message: "Invalid tenants configuration: mandatory configuration - missing OPA Url",
		Reason:  logstorev1.ReasonInvalidTenantsConfiguration,
		Requeue: false,
	}

	ff := configv1.FeatureGates{
		LogstoreStackGateway: true,
	}

	stack := &logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
			Tenants: &logstorev1.TenantsSpec{
				Mode: "dynamic",
				Authentication: []logstorev1.AuthenticationSpec{
					{
						TenantName: "test",
						TenantID:   "1234",
						OIDC: &logstorev1.OIDCSpec{
							Secret: &logstorev1.TenantSecretSpec{
								Name: defaultGatewaySecret.Name,
							},
						},
					},
				},
				Authorization: nil,
			},
		},
	}

	// GetStub looks up the CR first, so we need to return our fake stack
	// return NotFound for everything else to trigger create.
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, stack)
			return nil
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(object, &defaultSecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something is not found")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, ff)

	// make sure error is returned
	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestCreateOrUpdateLogstoreStack_WhenMissingGatewaySecret_SetDegraded(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	degradedErr := &status.DegradedError{
		Message: "Missing secrets for tenant test",
		Reason:  logstorev1.ReasonMissingGatewayTenantSecret,
		Requeue: true,
	}

	ff := configv1.FeatureGates{
		LogstoreStackGateway: true,
	}

	stack := &logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
			Tenants: &logstorev1.TenantsSpec{
				Mode: "dynamic",
				Authentication: []logstorev1.AuthenticationSpec{
					{
						TenantName: "test",
						TenantID:   "1234",
						OIDC: &logstorev1.OIDCSpec{
							Secret: &logstorev1.TenantSecretSpec{
								Name: defaultGatewaySecret.Name,
							},
						},
					},
				},
				Authorization: &logstorev1.AuthorizationSpec{
					OPA: &logstorev1.OPASpec{
						URL: "some-url",
					},
				},
			},
		},
	}

	// GetStub looks up the CR first, so we need to return our fake stack
	// return NotFound for everything else to trigger create.
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		o, ok := object.(*logstorev1.LogstoreStack)
		if r.Name == name.Name && r.Namespace == name.Namespace && ok {
			k.SetClientObject(o, stack)
			return nil
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(object, &defaultSecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something is not found")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, ff)

	// make sure error is returned to re-trigger reconciliation
	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestCreateOrUpdateLogstoreStack_WhenInvalidGatewaySecret_SetDegraded(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	degradedErr := &status.DegradedError{
		Message: "Invalid gateway tenant secret contents",
		Reason:  logstorev1.ReasonInvalidGatewayTenantSecret,
		Requeue: true,
	}

	ff := configv1.FeatureGates{
		LogstoreStackGateway: true,
	}

	stack := &logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
			Tenants: &logstorev1.TenantsSpec{
				Mode: "dynamic",
				Authentication: []logstorev1.AuthenticationSpec{
					{
						TenantName: "test",
						TenantID:   "1234",
						OIDC: &logstorev1.OIDCSpec{
							Secret: &logstorev1.TenantSecretSpec{
								Name: invalidSecret.Name,
							},
						},
					},
				},
				Authorization: &logstorev1.AuthorizationSpec{
					OPA: &logstorev1.OPASpec{
						URL: "some-url",
					},
				},
			},
		},
	}

	// GetStub looks up the CR first, so we need to return our fake stack
	// return NotFound for everything else to trigger create.
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		o, ok := object.(*logstorev1.LogstoreStack)
		if r.Name == name.Name && r.Namespace == name.Namespace && ok {
			k.SetClientObject(o, stack)
			return nil
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(object, &defaultSecret)
			return nil
		}
		if name.Name == invalidSecret.Name {
			k.SetClientObject(object, &invalidSecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something is not found")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, ff)

	// make sure error is returned to re-trigger reconciliation
	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestCreateOrUpdateLogstoreStack_MissingTenantsSpec_SetDegraded(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	degradedErr := &status.DegradedError{
		Message: "Invalid tenants configuration - TenantsSpec cannot be nil when gateway flag is enabled",
		Reason:  logstorev1.ReasonInvalidTenantsConfiguration,
		Requeue: false,
	}

	ff := configv1.FeatureGates{
		LogstoreStackGateway: true,
	}

	stack := &logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
			Tenants: nil,
		},
	}

	// GetStub looks up the CR first, so we need to return our fake stack
	// return NotFound for everything else to trigger create.
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		o, ok := object.(*logstorev1.LogstoreStack)
		if r.Name == name.Name && r.Namespace == name.Namespace && ok {
			k.SetClientObject(o, stack)
			return nil
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(object, &defaultSecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something is not found")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, ff)

	// make sure error is returned
	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestCreateOrUpdateLogstoreStack_WhenInvalidQueryTimeout_SetDegraded(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	degradedErr := &status.DegradedError{
		Message: `Error parsing query timeout: time: invalid duration "invalid"`,
		Reason:  logstorev1.ReasonQueryTimeoutInvalid,
		Requeue: false,
	}

	stack := &logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV12,
						EffectiveDate: "2023-05-22",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
			Tenants: &logstorev1.TenantsSpec{
				Mode: "openshift",
			},
			Limits: &logstorev1.LimitsSpec{
				Global: &logstorev1.LimitsTemplateSpec{
					QueryLimits: &logstorev1.QueryLimitSpec{
						QueryTimeout: "invalid",
					},
				},
			},
		},
	}

	// Create looks up the CR first, so we need to return our fake stack
	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(object, stack)
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(object, &defaultSecret)
		}
		return nil
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)

	// make sure error is returned
	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestCreateOrUpdateLogstoreStack_RemovesRulerResourcesWhenDisabled(t *testing.T) {
	sw := &k8sfakes.FakeStatusWriter{}
	k := &k8sfakes.FakeClient{}
	r := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "my-stack",
			Namespace: "some-ns",
		},
	}

	stack := logstorev1.LogstoreStack{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.LogstoreStackSpec{
			Size: logstorev1.SizeOneXExtraSmall,
			Storage: logstorev1.ObjectStorageSpec{
				Schemas: []logstorev1.ObjectStorageSchema{
					{
						Version:       logstorev1.ObjectStorageSchemaV11,
						EffectiveDate: "2020-10-11",
					},
				},
				Secret: logstorev1.ObjectStorageSecretSpec{
					Name: defaultSecret.Name,
					Type: logstorev1.ObjectStorageSecretS3,
				},
			},
			Rules: &logstorev1.RulesSpec{
				Enabled: true,
			},
			Tenants: &logstorev1.TenantsSpec{
				Mode: "dynamic",
				Authentication: []logstorev1.AuthenticationSpec{
					{
						TenantName: "test",
						TenantID:   "1234",
						OIDC: &logstorev1.OIDCSpec{
							Secret: &logstorev1.TenantSecretSpec{
								Name: defaultGatewaySecret.Name,
							},
						},
					},
				},
				Authorization: &logstorev1.AuthorizationSpec{
					OPA: &logstorev1.OPASpec{
						URL: "some-url",
					},
				},
			},
		},
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, out client.Object, _ ...client.GetOption) error {
		_, ok := out.(*logstorev1.RulerConfig)
		if ok {
			return apierrors.NewNotFound(schema.GroupResource{}, "no ruler config")
		}
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(out, &stack)
			return nil
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(out, &defaultSecret)
			return nil
		}
		if defaultGatewaySecret.Name == name.Name {
			k.SetClientObject(out, &defaultGatewaySecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	k.CreateStub = func(_ context.Context, o client.Object, _ ...client.CreateOption) error {
		assert.Equal(t, r.Namespace, o.GetNamespace())
		return nil
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	k.DeleteStub = func(_ context.Context, o client.Object, _ ...client.DeleteOption) error {
		assert.Equal(t, r.Namespace, o.GetNamespace())
		return nil
	}

	k.ListStub = func(_ context.Context, list client.ObjectList, options ...client.ListOption) error {
		switch list.(type) {
		case *corev1.ConfigMapList:
			k.SetClientObjectList(list, &corev1.ConfigMapList{
				Items: []corev1.ConfigMap{
					rulesCM,
				},
			})
		}
		return nil
	}

	err := handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)
	require.NoError(t, err)

	// make sure create was called
	require.NotZero(t, k.CreateCallCount())

	// make sure delete not called
	require.Zero(t, k.DeleteCallCount())

	// Disable the ruler
	stack.Spec.Rules.Enabled = false

	// Get should return ruler resources
	k.GetStub = func(_ context.Context, name types.NamespacedName, out client.Object, _ ...client.GetOption) error {
		_, ok := out.(*logstorev1.RulerConfig)
		if ok {
			return apierrors.NewNotFound(schema.GroupResource{}, "no ruler config")
		}
		if rulesCM.Name == name.Name {
			k.SetClientObject(out, &rulesCM)
			return nil
		}
		if rulerSS.Name == name.Name {
			k.SetClientObject(out, &rulerSS)
			return nil
		}
		if r.Name == name.Name && r.Namespace == name.Namespace {
			k.SetClientObject(out, &stack)
			return nil
		}
		if defaultSecret.Name == name.Name {
			k.SetClientObject(out, &defaultSecret)
			return nil
		}
		if defaultGatewaySecret.Name == name.Name {
			k.SetClientObject(out, &defaultGatewaySecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}
	err = handlers.CreateOrUpdateLogstoreStack(context.TODO(), logger, r, k, scheme, featureGates)
	require.NoError(t, err)

	// make sure delete was called twice (delete rules configmap and ruler statefulset)
	require.Equal(t, 2, k.DeleteCallCount())
}
