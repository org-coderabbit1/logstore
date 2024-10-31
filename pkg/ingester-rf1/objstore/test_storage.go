package objstore

import (
	"os"
	"testing"

	"github.com/prometheus/common/model"

	"example.com/acme/logstore/v3/pkg/storage"
	"example.com/acme/logstore/v3/pkg/storage/chunk/client/local"
	"example.com/acme/logstore/v3/pkg/storage/config"
)

var metrics *storage.ClientMetrics

func NewTestStorage(t testing.TB) (*Multi, error) {
	if metrics == nil {
		m := storage.NewClientMetrics()
		metrics = &m
	}
	dir := t.TempDir()
	t.Cleanup(func() {
		os.RemoveAll(dir)
		metrics.Unregister()
	})
	cfg := storage.Config{
		FSConfig: local.FSConfig{
			Directory: dir,
		},
	}
	return New([]config.PeriodConfig{
		{
			From:       config.DayTime{Time: model.Now()},
			ObjectType: "filesystem",
		},
	}, cfg, *metrics)
}
