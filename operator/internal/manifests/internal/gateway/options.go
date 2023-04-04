package gateway

import (
	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"
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

// Secret for clientID, clientSecret and issuerCAPath for tenant's authentication.
type Secret struct {
	TenantName   string
	ClientID     string
	ClientSecret string
	IssuerCAPath string
}
