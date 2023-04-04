package client

import (
	"net/url"
	"testing"
	"time"

	corestoreflag "example.com/acme/kit/flagext"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/require"

	"example.com/acme/logstore/clients/pkg/promtail/api"

	"example.com/acme/logstore/pkg/logproto"
	util_log "example.com/acme/logstore/pkg/util/log"
)

func TestNewLogger(t *testing.T) {
	_, err := NewLogger(nilMetrics, util_log.Logger, []Config{}...)
	require.Error(t, err)

	l, err := NewLogger(nilMetrics, util_log.Logger, []Config{{URL: corestoreflag.URLValue{URL: &url.URL{Host: "string"}}}}...)
	require.NoError(t, err)
	l.Chan() <- api.Entry{Labels: model.LabelSet{"foo": "bar"}, Entry: logproto.Entry{Timestamp: time.Now(), Line: "entry"}}
	l.Stop()
}
