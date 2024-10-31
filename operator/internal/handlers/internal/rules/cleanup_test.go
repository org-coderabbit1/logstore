package rules

import (
	"context"
	"io"
	"testing"

	"github.com/ViaQ/logerr/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s/k8sfakes"
)

var (
	logger = log.NewLogger("testing", log.WithOutput(io.Discard))

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
)

func TestCleanup_RemovesRulerResourcesWhenDisabled(t *testing.T) {
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

		_, isLogstoreStack := out.(*logstorev1.LogstoreStack)
		if r.Name == name.Name && r.Namespace == name.Namespace && isLogstoreStack {
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

	err := Cleanup(context.TODO(), logger, k, &stack)
	require.NoError(t, err)

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

		_, isLogstoreStack := out.(*logstorev1.LogstoreStack)
		if r.Name == name.Name && r.Namespace == name.Namespace && isLogstoreStack {
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

	err = Cleanup(context.TODO(), logger, k, &stack)
	require.NoError(t, err)

	// make sure delete was called twice (delete rules configmap and ruler statefulset)
	require.Equal(t, 2, k.DeleteCallCount())
}
