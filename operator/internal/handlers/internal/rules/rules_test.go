package rules

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/external/k8s/k8sfakes"
	"example.com/acme/logstore/operator/internal/status"
)

func TestBuildOptions_WhenMissingRemoteWriteSecret_SetDegraded(t *testing.T) {
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

	rulerCfg := &logstorev1.RulerConfig{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.RulerConfigSpec{
			RemoteWriteSpec: &logstorev1.RemoteWriteSpec{
				Enabled: true,
				ClientSpec: &logstorev1.RemoteWriteClientSpec{
					AuthorizationType:       logstorev1.BasicAuthorization,
					AuthorizationSecretName: "test",
				},
			},
		},
	}

	degradedErr := &status.DegradedError{
		Message: "Missing ruler remote write authorization secret",
		Reason:  logstorev1.ReasonMissingRulerSecret,
		Requeue: false,
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, out client.Object, _ ...client.GetOption) error {
		_, isRulerConfig := out.(*logstorev1.RulerConfig)
		if r.Name == name.Name && r.Namespace == name.Namespace && isRulerConfig {
			k.SetClientObject(out, rulerCfg)
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

	k.StatusStub = func() client.StatusWriter { return sw }

	_, _, _, _, err := BuildOptions(context.TODO(), logger, k, &stack)

	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestBuildOptions_WhenInvalidRemoteWriteSecret_SetDegraded(t *testing.T) {
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

	rulerCfg := &logstorev1.RulerConfig{
		TypeMeta: metav1.TypeMeta{
			Kind: "LogstoreStack",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "some-ns",
			UID:       "b23f9a38-9672-499f-8c29-15ede74d3ece",
		},
		Spec: logstorev1.RulerConfigSpec{
			RemoteWriteSpec: &logstorev1.RemoteWriteSpec{
				Enabled: true,
				ClientSpec: &logstorev1.RemoteWriteClientSpec{
					AuthorizationType:       logstorev1.BasicAuthorization,
					AuthorizationSecretName: "some-client-secret",
				},
			},
		},
	}

	invalidSecret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "some-client-secret",
			Namespace: "some-ns",
		},
		Data: map[string][]byte{},
	}

	degradedErr := &status.DegradedError{
		Message: "Invalid ruler remote write authorization secret contents",
		Reason:  logstorev1.ReasonInvalidRulerSecret,
		Requeue: false,
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, out client.Object, _ ...client.GetOption) error {
		_, isRulerConfig := out.(*logstorev1.RulerConfig)
		if r.Name == name.Name && r.Namespace == name.Namespace && isRulerConfig {
			k.SetClientObject(out, rulerCfg)
			return nil
		}

		_, isLogstoreStack := out.(*logstorev1.LogstoreStack)
		if r.Name == name.Name && r.Namespace == name.Namespace && isLogstoreStack {
			k.SetClientObject(out, &stack)
			return nil
		}
		if invalidSecret.Name == name.Name {
			k.SetClientObject(out, &invalidSecret)
			return nil
		}
		if defaultGatewaySecret.Name == name.Name {
			k.SetClientObject(out, &defaultGatewaySecret)
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	k.StatusStub = func() client.StatusWriter { return sw }

	_, _, _, _, err := BuildOptions(context.TODO(), logger, k, &stack)

	require.Error(t, err)
	require.Equal(t, degradedErr, err)
}

func TestList_AlertingRulesMatchSelector_WithDefaultStackNamespaceRules(t *testing.T) {
	const stackNs = "some-ns"

	k := &k8sfakes.FakeClient{}
	rs := &logstorev1.RulesSpec{
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"labelname": "labelvalue",
			},
		},
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if name.Name == stackNs {
			k.SetClientObject(object, &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: stackNs,
				},
			})
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	k.ListStub = func(_ context.Context, ol client.ObjectList, opt ...client.ListOption) error {
		switch ol.(type) {
		case *corev1.NamespaceList:
			k.SetClientObjectList(ol, &corev1.NamespaceList{})
			return nil
		case *logstorev1.RecordingRuleList:
			k.SetClientObjectList(ol, &logstorev1.RecordingRuleList{})
			return nil
		}

		l := opt[0].(*client.MatchingLabelsSelector)
		m := labels.Set(rs.Selector.MatchLabels)

		if l.Matches(m) {
			k.SetClientObjectList(ol, &logstorev1.AlertingRuleList{
				Items: []logstorev1.AlertingRule{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "rule-a",
							Namespace: stackNs,
							Labels: map[string]string{
								"labelname": "labelvalue",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "rule-b",
							Namespace: "other-ns",
							Labels: map[string]string{
								"labelname": "labelvalue",
							},
						},
					},
				},
			})
		}

		return nil
	}

	rules, _, err := listRules(context.TODO(), k, stackNs, rs)

	require.NoError(t, err)
	require.NotEmpty(t, rules)
	require.Len(t, rules, 1)
}

func TestList_AlertingRulesMatchSelector_FilteredByNamespaceSelector(t *testing.T) {
	const stackNs = "some-ns"

	k := &k8sfakes.FakeClient{}
	rs := &logstorev1.RulesSpec{
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"labelname": "labelvalue",
			},
		},
		NamespaceSelector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"group.acme.org/logs": "true",
			},
		},
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if name.Name == "some-ns" {
			k.SetClientObject(object, &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: stackNs,
				},
			})
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	k.ListStub = func(_ context.Context, ol client.ObjectList, opt ...client.ListOption) error {
		switch ol.(type) {
		case *logstorev1.RecordingRuleList:
			k.SetClientObjectList(ol, &logstorev1.RecordingRuleList{})
			return nil
		}

		l := opt[0].(*client.MatchingLabelsSelector)
		m := labels.Set(rs.Selector.MatchLabels)

		if l.Matches(m) {
			k.SetClientObjectList(ol, &logstorev1.AlertingRuleList{
				Items: []logstorev1.AlertingRule{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "rule-a",
							Namespace: "matching-ns",
							Labels: map[string]string{
								"labelname": "labelvalue",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "rule-b",
							Namespace: stackNs,
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "rule-c",
							Namespace: "not-matching-ns",
							Labels: map[string]string{
								"labelname": "labelvalue",
							},
						},
					},
				},
			})

			return nil
		}

		n := labels.Set(rs.NamespaceSelector.MatchLabels)
		if l.Matches(n) {
			k.SetClientObjectList(ol, &corev1.NamespaceList{
				Items: []corev1.Namespace{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "matching-ns",
							Labels: map[string]string{
								"group.acme.org/logs": "true",
							},
						},
					},
				},
			})

			return nil
		}

		k.SetClientObjectList(ol, &corev1.NamespaceList{})

		return nil
	}

	rules, _, err := listRules(context.TODO(), k, stackNs, rs)

	require.NoError(t, err)
	require.NotEmpty(t, rules)
	require.Len(t, rules, 2)
}

func TestList_RecordingRulesMatchSelector_WithDefaultStackNamespaceRules(t *testing.T) {
	const stackNs = "some-ns"

	k := &k8sfakes.FakeClient{}
	rs := &logstorev1.RulesSpec{
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"labelname": "labelvalue",
			},
		},
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if name.Name == stackNs {
			k.SetClientObject(object, &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: stackNs,
				},
			})
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	k.ListStub = func(_ context.Context, ol client.ObjectList, opt ...client.ListOption) error {
		switch ol.(type) {
		case *corev1.NamespaceList:
			k.SetClientObjectList(ol, &corev1.NamespaceList{})
			return nil
		case *logstorev1.AlertingRuleList:
			k.SetClientObjectList(ol, &logstorev1.AlertingRuleList{})
			return nil
		}

		l := opt[0].(*client.MatchingLabelsSelector)
		m := labels.Set(rs.Selector.MatchLabels)

		if l.Matches(m) {
			k.SetClientObjectList(ol, &logstorev1.RecordingRuleList{
				Items: []logstorev1.RecordingRule{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "rule-a",
							Namespace: stackNs,
							Labels: map[string]string{
								"labelname": "labelvalue",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "rule-b",
							Namespace: "other-ns",
							Labels: map[string]string{
								"labelname": "labelvalue",
							},
						},
					},
				},
			})
		}

		return nil
	}

	_, rules, err := listRules(context.TODO(), k, stackNs, rs)

	require.NoError(t, err)
	require.NotEmpty(t, rules)
	require.Len(t, rules, 1)
}

func TestList_RecordingRulesMatchSelector_FilteredByNamespaceSelector(t *testing.T) {
	const stackNs = "some-ns"

	k := &k8sfakes.FakeClient{}
	rs := &logstorev1.RulesSpec{
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"labelname": "labelvalue",
			},
		},
		NamespaceSelector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"group.acme.org/logs": "true",
			},
		},
	}

	k.GetStub = func(_ context.Context, name types.NamespacedName, object client.Object, _ ...client.GetOption) error {
		if name.Name == "some-ns" {
			k.SetClientObject(object, &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: stackNs,
				},
			})
			return nil
		}
		return apierrors.NewNotFound(schema.GroupResource{}, "something wasn't found")
	}

	k.ListStub = func(_ context.Context, ol client.ObjectList, opt ...client.ListOption) error {
		switch ol.(type) {
		case *logstorev1.AlertingRuleList:
			k.SetClientObjectList(ol, &logstorev1.AlertingRuleList{})
			return nil
		}

		l := opt[0].(*client.MatchingLabelsSelector)
		m := labels.Set(rs.Selector.MatchLabels)

		if l.Matches(m) {
			k.SetClientObjectList(ol, &logstorev1.RecordingRuleList{
				Items: []logstorev1.RecordingRule{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "rule-a",
							Namespace: "matching-ns",
							Labels: map[string]string{
								"labelname": "labelvalue",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "rule-b",
							Namespace: stackNs,
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "rule-c",
							Namespace: "not-matching-ns",
							Labels: map[string]string{
								"labelname": "labelvalue",
							},
						},
					},
				},
			})

			return nil
		}

		n := labels.Set(rs.NamespaceSelector.MatchLabels)
		if l.Matches(n) {
			k.SetClientObjectList(ol, &corev1.NamespaceList{
				Items: []corev1.Namespace{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "matching-ns",
							Labels: map[string]string{
								"group.acme.org/logs": "true",
							},
						},
					},
				},
			})
			return nil
		}

		k.SetClientObjectList(ol, &corev1.NamespaceList{})

		return nil
	}

	_, rules, err := listRules(context.TODO(), k, stackNs, rs)

	require.NoError(t, err)
	require.NotEmpty(t, rules)
	require.Len(t, rules, 2)
}
