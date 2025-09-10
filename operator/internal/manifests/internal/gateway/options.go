package gateway

import (
	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/manifests/openshift"
)

// Options is used to render the rbac.yaml and tenants.yaml file template
type Options struct {
	Stack logstorev1.LogstoreStackSpec

	Namespace        string
	Name             string
	StorageDirectory string

	OpenShiftOptions openshift.Options
	TenantSecrets    []*Secret
}

// Secret for tenant's authentication.
type Secret struct {
	TenantName string
	OIDC       *OIDC
	MTLS       *MTLS
}

// OIDC secret for tenant's authentication.
type OIDC struct {
	ClientID     string
	ClientSecret string
	IssuerCAPath string
}

// MTLS config for tenant's authentication.
type MTLS struct {
	CAPath string
}
