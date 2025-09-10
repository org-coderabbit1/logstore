package executor

import (
	"math"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/go-kit/log"
	"github.com/stretchr/testify/require"

	"example.com/acme/logstore/v3/pkg/dataobj"
	"example.com/acme/logstore/v3/pkg/dataobj/consumer/logsobj"
	"example.com/acme/logstore/v3/pkg/dataobj/sections/logs"
	"example.com/acme/logstore/v3/pkg/dataobj/sections/streams"
	"example.com/acme/logstore/v3/pkg/engine/internal/datatype"
	"example.com/acme/logstore/v3/pkg/engine/internal/types"
	"example.com/acme/logstore/v3/pkg/engine/planner/physical"
	"example.com/acme/logstore/v3/pkg/logproto"

	"example.com/acme/logstore/pkg/push"
)

var (
	labelMD    = datatype.ColumnMetadata(types.ColumnTypeLabel, datatype.Logstore.String)
	metadataMD = datatype.ColumnMetadata(types.ColumnTypeMetadata, datatype.Logstore.String)
)

func Test_dataobjScan(t *testing.T) {
	obj := buildDataobj(t, []logproto.Stream{
		{
			Labels: `{service="logstore", env="prod"}`,
			Entries: []logproto.Entry{
				{
					Timestamp:          time.Unix(5, 0),
					Line:               "hello world",
					StructuredMetadata: []push.LabelAdapter{{Name: "guid", Value: "aaaa-bbbb-cccc-dddd"}},
				},
				{
					Timestamp:          time.Unix(10, 0),
					Line:               "goodbye world",
					StructuredMetadata: []push.LabelAdapter{{Name: "guid", Value: "eeee-ffff-aaaa-bbbb"}},
				},
			},
		},
		{
			Labels: `{service="notlogstore", env="prod"}`,
			Entries: []logproto.Entry{
				{
					Timestamp:          time.Unix(2, 0),
					Line:               "hello world",
					StructuredMetadata: []push.LabelAdapter{{Name: "pod", Value: "notlogstore-pod-1"}},
				},
				{
					Timestamp:          time.Unix(3, 0),
					Line:               "goodbye world",
					StructuredMetadata: []push.LabelAdapter{{Name: "pod", Value: "notlogstore-pod-1"}},
				},
			},
		},
	})

	var (
		streamsSection *streams.Section
		logsSection    *logs.Section
	)

	for _, sec := range obj.Sections() {
		var err error

		switch {
		case streams.CheckSection(sec):
			streamsSection, err = streams.Open(t.Context(), sec)
			require.NoError(t, err, "failed to open streams section")

		case logs.CheckSection(sec):
			logsSection, err = logs.Open(t.Context(), sec)
			require.NoError(t, err, "failed to open logs section")
		}
	}

	t.Run("All columns", func(t *testing.T) {
		pipeline := newDataobjScanPipeline(dataobjScanOptions{
			StreamsSection: streamsSection,
			LogsSection:    logsSection,
			StreamIDs:      []int64{1, 2}, // All streams
			Projections:    nil,           // All columns

			BatchSize: 512,
		}, log.NewNopLogger())

		expectFields := []arrow.Field{
			{Name: "env", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},
			{Name: "service", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},
			{Name: "guid", Type: arrow.BinaryTypes.String, Metadata: metadataMD, Nullable: true},
			{Name: "pod", Type: arrow.BinaryTypes.String, Metadata: metadataMD, Nullable: true},
			{Name: "timestamp", Type: arrow.FixedWidthTypes.Timestamp_ns, Metadata: datatype.ColumnMetadataBuiltinTimestamp, Nullable: true},
			{Name: "message", Type: arrow.BinaryTypes.String, Metadata: datatype.ColumnMetadataBuiltinMessage, Nullable: true},
		}

		expectCSV := `prod,logstore,eeee-ffff-aaaa-bbbb,NULL,1970-01-01 00:00:10,goodbye world
prod,logstore,aaaa-bbbb-cccc-dddd,NULL,1970-01-01 00:00:05,hello world
prod,notlogstore,NULL,notlogstore-pod-1,1970-01-01 00:00:03,goodbye world
prod,notlogstore,NULL,notlogstore-pod-1,1970-01-01 00:00:02,hello world`

		expectRecord, err := CSVToArrow(expectFields, expectCSV)
		require.NoError(t, err)
		defer expectRecord.Release()

		AssertPipelinesEqual(t, pipeline, NewBufferedPipeline(expectRecord))
	})

	t.Run("Column subset", func(t *testing.T) {
		pipeline := newDataobjScanPipeline(dataobjScanOptions{
			StreamsSection: streamsSection,
			LogsSection:    logsSection,
			StreamIDs:      []int64{1, 2}, // All streams
			Projections: []physical.ColumnExpression{
				&physical.ColumnExpr{Ref: types.ColumnRef{Column: "env", Type: types.ColumnTypeLabel}},
				&physical.ColumnExpr{Ref: types.ColumnRef{Column: "timestamp", Type: types.ColumnTypeBuiltin}},
			},

			BatchSize: 512,
		}, log.NewNopLogger())

		expectFields := []arrow.Field{
			{Name: "env", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},
			{Name: "timestamp", Type: arrow.FixedWidthTypes.Timestamp_ns, Metadata: datatype.ColumnMetadataBuiltinTimestamp, Nullable: true},
		}

		expectCSV := `prod,1970-01-01 00:00:10
prod,1970-01-01 00:00:05
prod,1970-01-01 00:00:03
prod,1970-01-01 00:00:02`

		expectRecord, err := CSVToArrow(expectFields, expectCSV)
		require.NoError(t, err)
		defer expectRecord.Release()

		AssertPipelinesEqual(t, pipeline, NewBufferedPipeline(expectRecord))
	})

	t.Run("Streams subset", func(t *testing.T) {
		pipeline := newDataobjScanPipeline(dataobjScanOptions{
			StreamsSection: streamsSection,
			LogsSection:    logsSection,
			StreamIDs:      []int64{2}, // Only stream 2
			Projections:    nil,        // All columns

			BatchSize: 512,
		}, log.NewNopLogger())

		expectFields := []arrow.Field{
			{Name: "env", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},
			{Name: "service", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},
			{Name: "guid", Type: arrow.BinaryTypes.String, Metadata: metadataMD, Nullable: true},
			{Name: "pod", Type: arrow.BinaryTypes.String, Metadata: metadataMD, Nullable: true},
			{Name: "timestamp", Type: arrow.FixedWidthTypes.Timestamp_ns, Metadata: datatype.ColumnMetadataBuiltinTimestamp, Nullable: true},
			{Name: "message", Type: arrow.BinaryTypes.String, Metadata: datatype.ColumnMetadataBuiltinMessage, Nullable: true},
		}

		expectCSV := `prod,notlogstore,NULL,notlogstore-pod-1,1970-01-01 00:00:03,goodbye world
prod,notlogstore,NULL,notlogstore-pod-1,1970-01-01 00:00:02,hello world`

		expectRecord, err := CSVToArrow(expectFields, expectCSV)
		require.NoError(t, err)
		defer expectRecord.Release()

		AssertPipelinesEqual(t, pipeline, NewBufferedPipeline(expectRecord))
	})

	t.Run("Ambiguous column", func(t *testing.T) {
		// Here, we'll check for a column which only exists once in the dataobj but is
		// ambiguous from the perspective of the caller.
		pipeline := newDataobjScanPipeline(dataobjScanOptions{
			StreamsSection: streamsSection,
			LogsSection:    logsSection,
			StreamIDs:      []int64{1, 2}, // All streams
			Projections: []physical.ColumnExpression{
				&physical.ColumnExpr{Ref: types.ColumnRef{Column: "env", Type: types.ColumnTypeAmbiguous}},
			},
			BatchSize: 512,
		}, log.NewNopLogger())

		expectFields := []arrow.Field{
			{Name: "env", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},
		}

		expectCSV := `prod
prod
prod
prod`

		expectRecord, err := CSVToArrow(expectFields, expectCSV)
		require.NoError(t, err)
		defer expectRecord.Release()

		AssertPipelinesEqual(t, pipeline, NewBufferedPipeline(expectRecord))
	})
}

func Test_dataobjScan_DuplicateColumns(t *testing.T) {
	obj := buildDataobj(t, []logproto.Stream{
		// Case 1: A single row has a value for a label and metadata column with
		// the same name.
		{
			Labels: `{service="logstore", env="prod", pod="pod-1"}`,
			Entries: []logproto.Entry{
				{
					Timestamp:          time.Unix(1, 0),
					Line:               "message 1",
					StructuredMetadata: []push.LabelAdapter{{Name: "pod", Value: "override"}},
				},
			},
		},

		// Case 2: A label and metadata column share a name but have values in
		// different rows.
		{
			Labels: `{service="logstore", env="prod"}`,
			Entries: []logproto.Entry{{
				Timestamp:          time.Unix(2, 0),
				Line:               "message 2",
				StructuredMetadata: []push.LabelAdapter{{Name: "namespace", Value: "namespace-1"}},
			}},
		},
		{
			Labels: `{service="logstore", env="prod", namespace="namespace-2"}`,
			Entries: []logproto.Entry{{
				Timestamp:          time.Unix(3, 0),
				Line:               "message 3",
				StructuredMetadata: nil,
			}},
		},
	})

	var (
		streamsSection *streams.Section
		logsSection    *logs.Section
	)

	for _, sec := range obj.Sections() {
		var err error

		switch {
		case streams.CheckSection(sec):
			streamsSection, err = streams.Open(t.Context(), sec)
			require.NoError(t, err, "failed to open streams section")

		case logs.CheckSection(sec):
			logsSection, err = logs.Open(t.Context(), sec)
			require.NoError(t, err, "failed to open logs section")
		}
	}

	t.Run("All columns", func(t *testing.T) {
		pipeline := newDataobjScanPipeline(dataobjScanOptions{
			StreamsSection: streamsSection,
			LogsSection:    logsSection,
			StreamIDs:      []int64{1, 2, 3}, // All streams
			Projections:    nil,              // All columns
			BatchSize:      512,
		}, log.NewNopLogger())

		expectFields := []arrow.Field{
			{Name: "env", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},
			{Name: "namespace", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},
			{Name: "pod", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},
			{Name: "service", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},

			{Name: "namespace", Type: arrow.BinaryTypes.String, Metadata: metadataMD, Nullable: true},
			{Name: "pod", Type: arrow.BinaryTypes.String, Metadata: metadataMD, Nullable: true},

			{Name: "timestamp", Type: arrow.FixedWidthTypes.Timestamp_ns, Metadata: datatype.ColumnMetadataBuiltinTimestamp, Nullable: true},
			{Name: "message", Type: arrow.BinaryTypes.String, Metadata: datatype.ColumnMetadataBuiltinMessage, Nullable: true},
		}

		expectCSV := `prod,namespace-2,NULL,logstore,NULL,NULL,1970-01-01 00:00:03,message 3
prod,NULL,NULL,logstore,namespace-1,NULL,1970-01-01 00:00:02,message 2
prod,NULL,pod-1,logstore,NULL,override,1970-01-01 00:00:01,message 1`

		expectRecord, err := CSVToArrow(expectFields, expectCSV)
		require.NoError(t, err)
		defer expectRecord.Release()

		AssertPipelinesEqual(t, pipeline, NewBufferedPipeline(expectRecord))
	})

	t.Run("Ambiguous pod", func(t *testing.T) {
		pipeline := newDataobjScanPipeline(dataobjScanOptions{
			StreamsSection: streamsSection,
			LogsSection:    logsSection,
			StreamIDs:      []int64{1, 2, 3}, // All streams
			Projections: []physical.ColumnExpression{
				&physical.ColumnExpr{Ref: types.ColumnRef{Column: "pod", Type: types.ColumnTypeAmbiguous}},
			},
			BatchSize: 512,
		}, log.NewNopLogger())

		expectFields := []arrow.Field{
			{Name: "pod", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},
			{Name: "pod", Type: arrow.BinaryTypes.String, Metadata: metadataMD, Nullable: true},
		}

		expectCSV := `NULL,NULL
NULL,NULL
pod-1,override`

		expectRecord, err := CSVToArrow(expectFields, expectCSV)
		require.NoError(t, err)
		defer expectRecord.Release()

		AssertPipelinesEqual(t, pipeline, NewBufferedPipeline(expectRecord))
	})

	t.Run("Ambiguous namespace", func(t *testing.T) {
		pipeline := newDataobjScanPipeline(dataobjScanOptions{
			StreamsSection: streamsSection,
			LogsSection:    logsSection,
			StreamIDs:      []int64{1, 2, 3}, // All streams
			Projections: []physical.ColumnExpression{
				&physical.ColumnExpr{Ref: types.ColumnRef{Column: "namespace", Type: types.ColumnTypeAmbiguous}},
			},
			BatchSize: 512,
		}, log.NewNopLogger())

		expectFields := []arrow.Field{
			{Name: "namespace", Type: arrow.BinaryTypes.String, Metadata: labelMD, Nullable: true},
			{Name: "namespace", Type: arrow.BinaryTypes.String, Metadata: metadataMD, Nullable: true},
		}

		expectCSV := `namespace-2,NULL
NULL,namespace-1
NULL,NULL`

		expectRecord, err := CSVToArrow(expectFields, expectCSV)
		require.NoError(t, err)
		defer expectRecord.Release()

		AssertPipelinesEqual(t, pipeline, NewBufferedPipeline(expectRecord))
	})
}

func buildDataobj(t testing.TB, streams []logproto.Stream) *dataobj.Object {
	t.Helper()

	builder, err := logsobj.NewBuilder(logsobj.BuilderConfig{
		TargetPageSize:          8_000,
		TargetObjectSize:        math.MaxInt,
		TargetSectionSize:       32_000,
		BufferSize:              8_000,
		SectionStripeMergeLimit: 2,
	}, nil)
	require.NoError(t, err)

	for _, stream := range streams {
		require.NoError(t, builder.Append("tenant", stream))
	}

	obj, closer, err := builder.Flush()
	require.NoError(t, err)
	t.Cleanup(func() { closer.Close() })
	return obj
}
