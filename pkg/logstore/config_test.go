package logstore

import (
	"flag"
	"testing"
	"time"

	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/require"

	"example.com/acme/logstore/pkg/ingester"
	"example.com/acme/logstore/pkg/storage/config"
)

func TestCrossComponentValidation(t *testing.T) {
	for _, tc := range []struct {
		desc string
		base *Config
		err  bool
	}{
		{
			desc: "correct shards",
			base: &Config{
				Ingester: ingester.Config{
					IndexShards: 32,
				},
				SchemaConfig: config.SchemaConfig{
					Configs: []config.PeriodConfig{
						{
							RowShards: 16,
							Schema:    "v11",
							From: config.DayTime{
								Time: model.Now(),
							},
						},
					},
				},
			},
			err: false,
		},
		{
			desc: "correct shards",
			base: &Config{
				Ingester: ingester.Config{
					IndexShards: 32,
				},
				SchemaConfig: config.SchemaConfig{
					Configs: []config.PeriodConfig{
						{
							RowShards: 16,
							Schema:    "v11",
							From: config.DayTime{
								Time: model.Now().Add(-48 * time.Hour),
							},
						},
						{
							RowShards: 17,
							Schema:    "v11",
							From: config.DayTime{
								Time: model.Now(),
							},
						},
					},
				},
			},
			err: true,
		},
	} {
		tc.base.RegisterFlags(flag.NewFlagSet(tc.desc, 0))
		err := tc.base.Validate()
		if tc.err {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
		}
	}
}
