#!/bin/bash

set -euo pipefail

OLD_VERSION=${LOGSTORE_OLD_VERSION:-}
NEW_VERSION=${LOGSTORE_NEW_VERSION:-}

if [[ -z "${OLD_VERSION}" ]]
then
    OLD_VERSION="[0-9]+\.[0-9]+\.[0-9]+"
fi

if [[ -z "${NEW_VERSION}" ]]
then
    read -rp "Enter new release version (eg 2.9.2): " NEW_VERSION
fi


LOGSTORE_DOCKER_DRIVER_TAG="acme\/logstore-docker-driver:"
LOGSTORE_DOCUMENTS_TAG="logstore\/v"
LOGSTORE_DOCKER_TAG="acme\/logstore:"
LOGSTORE_PROMTAIL_DOCKER_TAG="acme\/promtail:"
LOGSTORE_CANARY_DOCKER_TAG="acme\/logstore-canary:"
LOGSTORE_LOGCLI_DOCKER_TAG="acme\/logcli:"

echo "Updating version references to ${NEW_VERSION}"

# grep -Iq is to ignore non-binary files.
find . -type f -not -path "./.git/*" -not -path "./vendor/*" -not -path "./operator/*" -not -path "./CHANGELOG.md" -not -path "./docs/sources/setup/upgrade/*" -exec grep -Iq . {} \; -print0\
    | xargs -0 sed -i '' -E \
	    -e "s/(${LOGSTORE_DOCKER_DRIVER_TAG})(${OLD_VERSION})/\1${NEW_VERSION}/g" \
	    -e "s/(${LOGSTORE_DOCUMENTS_TAG})(${OLD_VERSION})/\1${NEW_VERSION}/g" \
	    -e "s/(${LOGSTORE_DOCKER_TAG})(${OLD_VERSION})/\1${NEW_VERSION}/g" \
	    -e "s/(${LOGSTORE_PROMTAIL_DOCKER_TAG})(${OLD_VERSION})/\1${NEW_VERSION}/g" \
	    -e "s/(${LOGSTORE_CANARY_DOCKER_TAG})(${OLD_VERSION})/\1${NEW_VERSION}/g" \
	    -e "s/(${LOGSTORE_LOGCLI_DOCKER_TAG})(${OLD_VERSION})/\1${NEW_VERSION}/g" \
