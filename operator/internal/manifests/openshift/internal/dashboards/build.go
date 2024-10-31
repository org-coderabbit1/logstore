package dashboards

import (
	"embed"
	"encoding/json"
	"io/fs"

	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

const (
	staticDir                   = "static"
	logstoreStackDashboardRulesFile = "acme-dashboard-logstorestack-rules.json"
)

var (
	//go:embed static/*.json
	logstoreStackDashboards embed.FS
	dashboardMap        map[string][]byte
	dashboardRules      *monitoringv1.PrometheusRuleSpec
)

func init() {
	var err error
	subDir, err := fs.Sub(logstoreStackDashboards, staticDir)
	if err != nil {
		panic(err)
	}

	jsonFiles, err := fs.Glob(subDir, "*.json")
	if err != nil {
		panic(err)
	}

	dashboardMap = map[string][]byte{}
	for _, file := range jsonFiles {
		var content []byte
		content, err = fs.ReadFile(subDir, file)
		if err != nil {
			panic(err)
		}

		switch file {
		case logstoreStackDashboardRulesFile:
			err := json.Unmarshal(content, &dashboardRules)
			if err != nil {
				panic(err)
			}

		default:
			dashboardMap[file] = content
		}
	}
}

// Content returns the byte slices one for each LogstoreStack dashboard
// and a separate byte slice for the recording rules from the embedded filesystem.
func Content() (map[string][]byte, *monitoringv1.PrometheusRuleSpec) {
	return dashboardMap, dashboardRules
}
