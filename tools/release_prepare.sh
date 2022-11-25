#!/bin/bash

# sed-wrap runs the appropriate sed command based on the
# underlying value of $OSTYPE
sed-wrap() {
  if [[ "${OSTYPE}" == "linux"* ]]; then
    # Linux
    sed -i "$1" "$2"
  else
    # macOS, BSD
    sed -i '' "$1" "$2"
  fi
}

echo
echo "Last 5 tags:"
git tag --sort=-taggerdate | head -n 5
echo

read -rp "Enter release version: " VERSION

if [[ ${VERSION} =~ ^v[0-9]+\.[0-9]+\.[0-9]+.*$ ]]; then
    echo "New Version: ${VERSION}"
else
    echo "Version must be in the format v0.1.0"
    exit 1
fi

LOGSTORE_CURRENT=$(sed -n -e 's/^version: //p' production/helm/logstore/Chart.yaml)
LOGSTORE_SUGGESTED=$(tools/increment_version.sh -m "${LOGSTORE_CURRENT}")
PROMTAIL_CURRENT=$(sed -n -e 's/^version: //p' production/helm/promtail/Chart.yaml)
PROMTAIL_SUGGESTED=$(tools/increment_version.sh -m "${PROMTAIL_CURRENT}")
LOGSTORE_STACK_CURRENT=$(sed -n -e 's/^version: //p' production/helm/logstore-stack/Chart.yaml)
LOGSTORE_STACK_SUGGESTED=$(tools/increment_version.sh -m "${LOGSTORE_STACK_CURRENT}")
echo
echo "Current Logstore helm chart version: ${LOGSTORE_CURRENT}"
read -rp "Enter new Logstore helm chart version [${LOGSTORE_SUGGESTED}]: " LOGSTORE_VERSION
LOGSTORE_VERSION=${LOGSTORE_VERSION:-${LOGSTORE_SUGGESTED}}
echo
echo "Current Promtail helm chart version: ${PROMTAIL_CURRENT}"
read -rp "Enter new Promtail helm chart version [${PROMTAIL_SUGGESTED}]: " PROMTAIL_VERSION
PROMTAIL_VERSION=${PROMTAIL_VERSION:-${PROMTAIL_SUGGESTED}}
echo
echo "Current Logstore-Stack helm chart version: ${LOGSTORE_STACK_CURRENT}"
read -rp "Enter new Logstore-Stack helm chart version [${LOGSTORE_STACK_SUGGESTED}]: " LOGSTORE_STACK_VERSION
LOGSTORE_STACK_VERSION=${LOGSTORE_STACK_VERSION:-${LOGSTORE_STACK_SUGGESTED}}
echo

echo "Creating Release"
echo "Release Version:       ${VERSION}"
echo "Logstore Helm Chart:       ${LOGSTORE_VERSION}"
echo "Promtail Helm Chart:   ${PROMTAIL_VERSION}"
echo "Logstore-Stack Helm Chart: ${LOGSTORE_STACK_VERSION}"
echo
read -rp "Is this correct? [y]: " CONTINUE
CONTINUE=${CONTINUE:-y}
echo

if [[ "${CONTINUE}" != "y" ]]; then
 exit 1
fi

echo "Updating helm and ksonnet image versions"
sed-wrap "s/.*promtail:.*/    promtail: 'acme\/promtail:${VERSION}',/" production/ksonnet/promtail/config.libsonnet
sed-wrap "s/.*logstore_canary:.*/    logstore_canary: 'acme\/logstore-canary:${VERSION}',/" production/ksonnet/logstore-canary/config.libsonnet
sed-wrap "s/.*logstore:.*/    logstore: 'acme\/logstore:${VERSION}',/" production/ksonnet/logstore/images.libsonnet
sed-wrap "s/.*tag:.*/  tag: ${VERSION}/" production/helm/logstore/values.yaml
sed-wrap "s/.*tag:.*/  tag: ${VERSION}/" production/helm/promtail/values.yaml

echo "Updating helm charts"
sed-wrap "s/^version:.*/version: ${LOGSTORE_VERSION}/" production/helm/logstore/Chart.yaml
sed-wrap "s/^version:.*/version: ${PROMTAIL_VERSION}/" production/helm/promtail/Chart.yaml
sed-wrap "s/^version:.*/version: ${LOGSTORE_STACK_VERSION}/" production/helm/logstore-stack/Chart.yaml

sed-wrap "s/^appVersion:.*/appVersion: ${VERSION}/" production/helm/logstore/Chart.yaml
sed-wrap "s/^appVersion:.*/appVersion: ${VERSION}/" production/helm/promtail/Chart.yaml
sed-wrap "s/^appVersion:.*/appVersion: ${VERSION}/" production/helm/logstore-stack/Chart.yaml

echo
echo "######################################################################################################"
echo
echo "Version numbers updated, create a new branch, commit and push"
echo
echo "######################################################################################################"

