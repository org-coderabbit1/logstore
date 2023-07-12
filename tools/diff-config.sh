#!/usr/bin/env bash

current_dir="$(cd "$(dirname "${0}")" && pwd)"
logstore_dir="$(cd "${current_dir}/../cmd/logstore" && pwd)"

export OLD_LOGSTORE=${OLD_VERSION:-2.7.5}
export NEW_LOGSTORE=${NEW_VERSION:-$("${current_dir}/image-tag")}

export CONFIG_FILE="logstore-local-config.yaml"

function get_config() {
	local version=${1}

	docker run --rm -t -v "${logstore_dir}:/config" "acme/logstore:${version}" \
		-config.file="/config/${CONFIG_FILE}" \
		-print-config-stderr 2>&1
}

function parse_config() {
	sed '/Starting Logstore/q' | tr -d '\r'
}

tmp_dir="$(mktemp -d)"
old_config="${tmp_dir}/config-${OLD_LOGSTORE}.yml"
new_config="${tmp_dir}/config-${NEW_LOGSTORE}.yml"

echo "Saving configs to ${tmp_dir}"
echo "Old config: ${old_config}"
echo "New config: ${new_config}"

get_config "${OLD_LOGSTORE}" | parse_config > "${old_config}"
get_config "${NEW_LOGSTORE}" | parse_config > "${new_config}"

diff --color=always \
  --side-by-side \
  "${old_config}" \
  "${new_config}"
