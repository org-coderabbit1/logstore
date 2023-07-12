package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	maxCharsPerLine = 20

	cases := []struct {
		name string
		in   string
		exp  string
	}{
		{
			name: "basic stream selector",
			in:   `{job="logstore", instance="localhost"}`,
			exp:  `{job="logstore", instance="localhost"}`,
		},
		{
			name: "pipeline_label_filter",
			in:   `{job="logstore", instance="localhost"}|logfmt|level="error" `,
			exp: `{job="logstore", instance="localhost"}
  | logfmt
  | level="error"`,
		},
		{
			name: "pipeline_line_filter",
			in:   `{job="logstore", instance="localhost"}|= "error" != "memcached" |= ip("192.168.0.1") |logfmt`,
			exp: `{job="logstore", instance="localhost"}
  |= "error"
  != "memcached"
  |= ip("192.168.0.1")
  | logfmt`,
		},
		{
			name: "pipeline_line_format",
			in:   `{job="logstore", instance="localhost"}|logfmt|line_format "{{.error}}"`,
			exp: `{job="logstore", instance="localhost"}
  | logfmt
  | line_format "{{.error}}"`,
		},
		{
			name: "pipeline_label_format",
			in:   `{job="logstore", instance="localhost"}|logfmt|label_format dst="{{.src}}"`,
			exp: `{job="logstore", instance="localhost"}
  | logfmt
  | label_format dst="{{.src}}"`,
		},
		{
			name: "aggregation",
			in:   `count_over_time({job="logstore", instance="localhost"}|logfmt[1m])`,
			exp: `count_over_time(
  {job="logstore", instance="localhost"}
    | logfmt [1m]
)`,
		},
		{
			name: "aggregation_with_offset",
			in:   `count_over_time({job="logstore", instance="localhost"}|= "error"[5m] offset 20m)`,
			exp: `count_over_time(
  {job="logstore", instance="localhost"}
    |= "error" [5m] offset 20m
)`,
		},
		{
			name: "unwrap",
			in:   `quantile_over_time(0.99,{container="ingress-nginx",service="hosted-acme"}| json| unwrap response_latency_seconds| __error__=""[1m]) by (cluster)`,
			exp: `quantile_over_time(
  0.99,
  {container="ingress-nginx", service="hosted-acme"}
    | json
    | unwrap response_latency_seconds
    | __error__="" [1m]
) by (cluster)`,
		},
		{
			name: "pipeline_aggregation_line_filter",
			in:   `count_over_time({job="logstore", instance="localhost"}|= "error" != "memcached" |= ip("192.168.0.1") |logfmt[1m])`,
			exp: `count_over_time(
  {job="logstore", instance="localhost"}
    |= "error"
    != "memcached"
    |= ip("192.168.0.1")
    | logfmt [1m]
)`,
		},
		{
			name: "jsonparserExpr",
			in:   `{job="logstore", namespace="logstore-prod", container="nginx-ingress"}| json first_server="servers[0]", ua="request.headers[\"User-Agent\"]" | level="error"`,
			exp: `{job="logstore", namespace="logstore-prod", container="nginx-ingress"}
  | json first_server="servers[0]",ua="request.headers[\"User-Agent\"]"
  | level="error"`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := ParseExpr(c.in)
			require.NoError(t, err)
			got := Prettify(expr)
			assert.Equal(t, c.exp, got)
		})
	}
}

func TestFormat_VectorAggregation(t *testing.T) {
	maxCharsPerLine = 20

	cases := []struct {
		name string
		in   string
		exp  string
	}{
		{
			name: "sum",
			in:   `sum(count_over_time({foo="bar",namespace="logstore",instance="localhost"}[5m])) by (container)`,
			exp: `sum by (container)(
  count_over_time(
    {foo="bar", namespace="logstore", instance="localhost"} [5m]
  )
)`,
		},
		{
			name: "topk",
			in:   `topk(5, count_over_time({foo="bar",namespace="logstore",instance="localhost"}[5m])) by (container)`,
			exp: `topk by (container)(
  5,
  count_over_time(
    {foo="bar", namespace="logstore", instance="localhost"} [5m]
  )
)`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := ParseExpr(c.in)
			require.NoError(t, err)
			got := Prettify(expr)
			assert.Equal(t, c.exp, got)
		})
	}
}

func TestFormat_LabelReplace(t *testing.T) {
	maxCharsPerLine = 20

	cases := []struct {
		name string
		in   string
		exp  string
	}{
		{
			name: "label_replace",
			in:   `label_replace(rate({job="api-server",service="a:c"}|= "err" [5m]), "foo", "$1", "service", "(.*):.*")`,
			exp: `label_replace(
  rate(
    {job="api-server", service="a:c"}
      |= "err" [5m]
  ),
  "foo",
  "$1",
  "service",
  "(.*):.*"
)`,
		},
		{
			name: "label_replace_nested",
			in:   `label_replace(label_replace(rate({job="api-server",service="a:c"}|= "err" [5m]), "foo", "$1", "service", "(.*):.*"), "foo", "$1", "service", "(.*):.*")`,
			exp: `label_replace(
  label_replace(
    rate(
      {job="api-server", service="a:c"}
        |= "err" [5m]
    ),
    "foo",
    "$1",
    "service",
    "(.*):.*"
  ),
  "foo",
  "$1",
  "service",
  "(.*):.*"
)`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := ParseExpr(c.in)
			require.NoError(t, err)
			got := Prettify(expr)
			assert.Equal(t, c.exp, got)
		})
	}
}

func TestFormat_BinOp(t *testing.T) {
	maxCharsPerLine = 20

	cases := []struct {
		name string
		in   string
		exp  string
	}{
		{
			name: "single binop",
			in:   `sum(rate({job="logstore", namespace="logstore-prod", instance="localhost"}[5m]))/sum(count_over_time({job="logstore", namespace="logstore-prod", instance="localhost"}[5m]))`,
			exp: `  sum(
    rate(
      {job="logstore", namespace="logstore-prod", instance="localhost"} [5m]
    )
  )
/
  sum(
    count_over_time(
      {job="logstore", namespace="logstore-prod", instance="localhost"} [5m]
    )
  )`,
		},
		{
			name: "multiple binops",
			in:   `sum(rate({job="logstore"}[5m])) + sum(rate({job="logstore-dev"}[5m])) / sum(rate({job="logstore-prod"}[5m]))`,
			exp: `  sum(
    rate(
      {job="logstore"} [5m]
    )
  )
+
    sum(
      rate(
        {job="logstore-dev"} [5m]
      )
    )
  /
    sum(
      rate(
        {job="logstore-prod"} [5m]
      )
    )`,
		},
		// NOTE: LogQL binary arithmetic ops have following precedences rules
		// 1. * / % - higher priority
		// 2. + -  - lower priority.
		// 3. Between same priority ops, whichever comes first takes precedence.
		// Following `_precedence*` tests makes sure LogQL formatter respects that.
		{
			name: "multiple binops check precedence",
			in:   `sum(rate({job="logstore"}[5m])) / sum(rate({job="logstore-dev"}[5m])) + sum(rate({job="logstore-prod"}[5m]))`,
			exp: `    sum(
      rate(
        {job="logstore"} [5m]
      )
    )
  /
    sum(
      rate(
        {job="logstore-dev"} [5m]
      )
    )
+
  sum(
    rate(
      {job="logstore-prod"} [5m]
    )
  )`,
		},
		{
			name: "multiple binops check precedence2",
			in:   `sum(rate({job="logstore"}[5m])) - sum(rate({job="logstore-stage"}[5m])) / sum(rate({job="logstore-dev"}[5m])) + sum(rate({job="logstore-prod"}[5m]))`,
			exp: `    sum(
      rate(
        {job="logstore"} [5m]
      )
    )
  -
      sum(
        rate(
          {job="logstore-stage"} [5m]
        )
      )
    /
      sum(
        rate(
          {job="logstore-dev"} [5m]
        )
      )
+
  sum(
    rate(
      {job="logstore-prod"} [5m]
    )
  )`,
		},
		{
			name: "multiple binops check precedence3",
			in:   `sum(rate({job="logstore"}[5m])) - sum(rate({job="logstore-stage"}[5m])) % sum(rate({job="logstore-dev"}[5m])) + sum(rate({job="logstore-prod"}[5m]))`,
			exp: `    sum(
      rate(
        {job="logstore"} [5m]
      )
    )
  -
      sum(
        rate(
          {job="logstore-stage"} [5m]
        )
      )
    %
      sum(
        rate(
          {job="logstore-dev"} [5m]
        )
      )
+
  sum(
    rate(
      {job="logstore-prod"} [5m]
    )
  )`,
		},
		{
			name: "multiple binops check precedence4",
			in:   `sum(rate({job="logstore"}[5m])) / sum(rate({job="logstore-stage"}[5m])) % sum(rate({job="logstore-dev"}[5m])) + sum(rate({job="logstore-prod"}[5m]))`,
			exp: `      sum(
        rate(
          {job="logstore"} [5m]
        )
      )
    /
      sum(
        rate(
          {job="logstore-stage"} [5m]
        )
      )
  %
    sum(
      rate(
        {job="logstore-dev"} [5m]
      )
    )
+
  sum(
    rate(
      {job="logstore-prod"} [5m]
    )
  )`,
		},
		{
			name: "multiple binops check precedence5",
			in:   `sum(rate({job="logstore"}[5m])) / sum(rate({job="logstore-stage"}[5m])) % sum(rate({job="logstore-dev"}[5m])) * sum(rate({job="logstore-prod"}[5m]))`,
			exp: `      sum(
        rate(
          {job="logstore"} [5m]
        )
      )
    /
      sum(
        rate(
          {job="logstore-stage"} [5m]
        )
      )
  %
    sum(
      rate(
        {job="logstore-dev"} [5m]
      )
    )
*
  sum(
    rate(
      {job="logstore-prod"} [5m]
    )
  )`,
		},
		{
			name: "binops with options", // options - on, ignoring, group_left, group_right
			in:   `sum(rate({job="logstore"}[5m])) * on(instance, job) group_left (node) sum(rate({job="logstore-prod"}[5m]))`,
			exp: `  sum(
    rate(
      {job="logstore"} [5m]
    )
  )
* on (instance, job) group_left (node)
  sum(
    rate(
      {job="logstore-prod"} [5m]
    )
  )`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := ParseExpr(c.in)
			require.NoError(t, err)
			got := Prettify(expr)
			assert.Equal(t, c.exp, got)
		})
	}
}
