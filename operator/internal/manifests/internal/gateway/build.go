package gateway

import (
	"bytes"
	"embed"
	"io"
	"text/template"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"

	"github.com/ViaQ/logerr/v2/kverrors"
)

const (
	// LogstoreGatewayTenantFileName is the name of the tenant config file in the configmap
	LogstoreGatewayTenantFileName = "tenants.yaml"
	// LogstoreGatewayRbacFileName is the name of the rbac config file in the configmap
	LogstoreGatewayRbacFileName = "rbac.yaml"
	// LogstoreGatewayRegoFileName is the name of the logstorestack-gateway rego config file in the configmap
	LogstoreGatewayRegoFileName = "logstorestack-gateway.rego"
	// LogstoreGatewayMountDir is the path that is mounted from the configmap
	LogstoreGatewayMountDir = "/etc/logstorestack-gateway"
)

var (
	//go:embed gateway-rbac.yaml
	logstoreGatewayRbacYAMLTmplFile embed.FS

	//go:embed gateway-tenants.yaml
	logstoreGatewayTenantsYAMLTmplFile embed.FS

	//go:embed logstorestack-gateway.rego
	logstoreStackGatewayRegoTmplFile embed.FS

	logstoreGatewayRbacYAMLTmpl = template.Must(template.ParseFS(logstoreGatewayRbacYAMLTmplFile, "gateway-rbac.yaml"))

	logstoreGatewayTenantsYAMLTmpl = template.Must(template.ParseFS(logstoreGatewayTenantsYAMLTmplFile, "gateway-tenants.yaml"))

	logstoreStackGatewayRegoTmpl = template.Must(template.ParseFS(logstoreStackGatewayRegoTmplFile, "logstorestack-gateway.rego"))
)

// Build builds a logstore gateway configuration files
func Build(opts Options) (rbacCfg []byte, tenantsCfg []byte, regoCfg []byte, err error) {
	// Build logstore gateway rbac yaml
	w := bytes.NewBuffer(nil)
	err = logstoreGatewayRbacYAMLTmpl.Execute(w, opts)
	if err != nil {
		return nil, nil, nil, kverrors.Wrap(err, "failed to create logstore gateway rbac configuration")
	}
	rbacCfg, err = io.ReadAll(w)
	if err != nil {
		return nil, nil, nil, kverrors.Wrap(err, "failed to read configuration from buffer")
	}
	// Build logstore gateway tenants yaml
	w = bytes.NewBuffer(nil)
	err = logstoreGatewayTenantsYAMLTmpl.Execute(w, opts)
	if err != nil {
		return nil, nil, nil, kverrors.Wrap(err, "failed to create logstore gateway tenants configuration")
	}
	tenantsCfg, err = io.ReadAll(w)
	if err != nil {
		return nil, nil, nil, kverrors.Wrap(err, "failed to read configuration from buffer")
	}
	// Build logstore gateway observatorium rego for static mode
	if opts.Stack.Tenants.Mode == logstorev1.Static {
		w = bytes.NewBuffer(nil)
		err = logstoreStackGatewayRegoTmpl.Execute(w, opts)
		if err != nil {
			return nil, nil, nil, kverrors.Wrap(err, "failed to create logstorestack gateway rego configuration")
		}
		regoCfg, err = io.ReadAll(w)
		if err != nil {
			return nil, nil, nil, kverrors.Wrap(err, "failed to read configuration from buffer")
		}
		return rbacCfg, tenantsCfg, regoCfg, nil
	}
	return rbacCfg, tenantsCfg, nil, nil
}
