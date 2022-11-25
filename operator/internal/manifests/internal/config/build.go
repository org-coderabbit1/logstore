package config

import (
	"bytes"
	"embed"
	"io"
	"text/template"

	"github.com/ViaQ/logerr/v2/kverrors"
)

const (
	// LogstoreConfigFileName is the name of the config file in the configmap
	LogstoreConfigFileName = "config.yaml"
	// LogstoreRuntimeConfigFileName is the name of the runtime config file in the configmap
	LogstoreRuntimeConfigFileName = "runtime-config.yaml"
	// LogstoreConfigMountDir is the path that is mounted from the configmap
	LogstoreConfigMountDir = "/etc/logstore/config"
)

var (
	//go:embed logstore-config.yaml
	logstoreConfigYAMLTmplFile embed.FS

	//go:embed logstore-runtime-config.yaml
	logstoreRuntimeConfigYAMLTmplFile embed.FS

	logstoreConfigYAMLTmpl = template.Must(template.ParseFS(logstoreConfigYAMLTmplFile, "logstore-config.yaml"))

	logstoreRuntimeConfigYAMLTmpl = template.Must(template.ParseFS(logstoreRuntimeConfigYAMLTmplFile, "logstore-runtime-config.yaml"))
)

// Build builds a logstore stack configuration files
func Build(opts Options) ([]byte, []byte, error) {
	// Build logstore config yaml
	w := bytes.NewBuffer(nil)
	err := logstoreConfigYAMLTmpl.Execute(w, opts)
	if err != nil {
		return nil, nil, kverrors.Wrap(err, "failed to create logstore configuration")
	}
	cfg, err := io.ReadAll(w)
	if err != nil {
		return nil, nil, kverrors.Wrap(err, "failed to read configuration from buffer")
	}
	// Build logstore runtime config yaml
	w = bytes.NewBuffer(nil)
	err = logstoreRuntimeConfigYAMLTmpl.Execute(w, opts)
	if err != nil {
		return nil, nil, kverrors.Wrap(err, "failed to create logstore runtime configuration")
	}
	rcfg, err := io.ReadAll(w)
	if err != nil {
		return nil, nil, kverrors.Wrap(err, "failed to read configuration from buffer")
	}
	return cfg, rcfg, nil
}
