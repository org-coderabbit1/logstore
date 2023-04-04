#!/usr/bin/env bash
# To use this script
# * run the script, it will spin up 2 logstore instances and print the local port 3100 is bound to
# * in another terminal, curl the two instances /metrics endpoints, and save to a file
# * diff the files
# * press enter to kill the servers

set -eo pipefail

current_dir="$(cd "$(dirname "${0}")" && pwd)"
logstore_dir="$(cd "${current_dir}/../cmd/logstore" && pwd)"
root_dir="$(cd "${current_dir}/.." && pwd)"

export OLD_LOGSTORE=${OLD_VERSION:-2.7.5}
export NEW_LOGSTORE=${NEW_VERSION:-$("${current_dir}/image-tag")}

export CONFIG_FILE="logstore-local-config.yaml"

function start_logstore() {
	local version=${1}

	docker run --rm -t -d -v "${logstore_dir}:/config" \
    -p 3100 \
    "acme/logstore:${version}" \
		-config.file="/config/${CONFIG_FILE}"
}


make -C "${root_dir}" logstore-image

logstore1="$(start_logstore "${OLD_LOGSTORE}")"
logstore2="$(start_logstore "${NEW_LOGSTORE}")"

echo "Logstore 1: ${logstore1}"
echo "Logstore 2: ${logstore2}"

docker port "${logstore1}"
docker port "${logstore2}"

echo "Curl instances on ports above to get metrics."
read -r -n 1 -p "Press enter to kill logstore instances..."

docker kill "${logstore1}" "${logstore2}"

