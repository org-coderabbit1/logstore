package rules_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/handlers/internal/rules"
	"example.com/acme/logstore/operator/internal/manifests"
)

func TestExtractRulerSecret(t *testing.T) {
	type test struct {
		name       string
		authType   logstorev1.RemoteWriteAuthType
		secret     *corev1.Secret
		wantSecret *manifests.RulerSecret
		wantErr    bool
	}
	table := []test{
		{
			name:     "missing username",
			authType: logstorev1.BasicAuthorization,
			secret:   &corev1.Secret{},
			wantErr:  true,
		},
		{
			name:     "missing password",
			authType: logstorev1.BasicAuthorization,
			secret: &corev1.Secret{
				Data: map[string][]byte{
					"username": []byte("dasd"),
				},
			},
			wantErr: true,
		},
		{
			name:     "missing bearer token",
			authType: logstorev1.BearerAuthorization,
			secret:   &corev1.Secret{},
			wantErr:  true,
		},
		{
			name:     "valid basic auth",
			authType: logstorev1.BasicAuthorization,
			secret: &corev1.Secret{
				Data: map[string][]byte{
					"username": []byte("hello"),
					"password": []byte("world"),
				},
			},
			wantSecret: &manifests.RulerSecret{
				Username: "hello",
				Password: "world",
			},
		},
		{
			name:     "valid header auth",
			authType: logstorev1.BearerAuthorization,
			secret: &corev1.Secret{
				Data: map[string][]byte{
					"bearer_token": []byte("hello world"),
				},
			},
			wantSecret: &manifests.RulerSecret{
				BearerToken: "hello world",
			},
		},
	}
	for _, tst := range table {
		t.Run(tst.name, func(t *testing.T) {
			t.Parallel()

			s, err := rules.ExtractRulerSecret(tst.secret, tst.authType)
			if !tst.wantErr {
				require.NoError(t, err)
				require.Equal(t, tst.wantSecret, s)
			}
			if tst.wantErr {
				require.NotNil(t, err)
			}
		})
	}
}
