#!/usr/bin/env bash

# This can be run like:
# ./tools/tsdb/tsdb-map/diff.sh /tmp/logstore-scratch/logstore-ops-daily.r main $(git rev-parse --abbrev-ref HEAD) <rounds>

boltdb_base=$1
branch_a=$2
branch_b=$3
COUNT="${4:-8}"
echo running "${COUNT}" rounds

echo building from "${branch_a}"
git checkout "${branch_a}"
go run tools/tsdb/tsdb-map/main.go  -source "${boltdb_base}" -dest /tmp/logstore-tsdb-a
echo benchmarking "${branch_a}"
LOGSTORE_TSDB_PATH=/tmp/logstore-tsdb-a go test example.com/acme/logstore/tools/tsdb/tsdb-map -count="${COUNT}" -bench=BenchmarkQuery -run '^$' -benchmem > /tmp/logstore-tsdb-bench-a

echo building from "${branch_b}"
git checkout "${branch_b}"
go run tools/tsdb/tsdb-map/main.go  -source "${boltdb_base}" -dest /tmp/logstore-tsdb-b
echo benchmarking "${branch_b}"
LOGSTORE_TSDB_PATH=/tmp/logstore-tsdb-b go test example.com/acme/logstore/tools/tsdb/tsdb-map -count="${COUNT}" -bench=BenchmarkQuery -run '^$' -benchmem > /tmp/logstore-tsdb-bench-b


echo benchmarks:
echo
benchstat /tmp/logstore-tsdb-bench-a /tmp/logstore-tsdb-bench-b

echo
echo sizing:
echo

ls -lh /tmp/logstore-tsdb-a
ls -lh /tmp/logstore-tsdb-b
