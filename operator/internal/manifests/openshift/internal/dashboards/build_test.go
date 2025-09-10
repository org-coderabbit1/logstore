package dashboards

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	logstoreStackChunkDashboardFile     = "acme-dashboard-logstorestack-chunks.json"
	logstoreStackReadsDashboardFile     = "acme-dashboard-logstorestack-reads.json"
	logstoreStackWritesDashboardFile    = "acme-dashboard-logstorestack-writes.json"
	logstoreStackRetentionDashboardFile = "acme-dashboard-logstorestack-retention.json"
)

func TestContent(t *testing.T) {
	m, r := Content()
	require.Len(t, m, 4)
	require.Equal(t, dashboardMap, m)
	require.Equal(t, dashboardRules, r)

	require.Contains(t, m, logstoreStackChunkDashboardFile)
	require.Contains(t, m, logstoreStackReadsDashboardFile)
	require.Contains(t, m, logstoreStackWritesDashboardFile)
	require.Contains(t, m, logstoreStackRetentionDashboardFile)
	require.NotContains(t, m, logstoreStackDashboardRulesFile)
}
