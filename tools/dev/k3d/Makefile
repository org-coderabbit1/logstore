.PHONY: logstore enterprise-logs helm-cluster down add-repos update-repos prepare prepare-gel build-latest-image
.PHONY: helm-cluster helm-install-enterprise-logs helm-upgrade-enterprise-logs helm-uninstall-enterprise-logs
.PHONY: helm-install-logstore helm-upgrade-logstore helm-uninstall-logstore

IMAGE_TAG := $(shell ../../../tools/image-tag)
EXISTING_REGISTRY_PORT := $(shell k3d registry list -o json | jq -r '.[] | select(.name == "k3d-acme") | .portMappings."5000/tcp" | .[0].HostPort')
REGISTRY_PORT ?= $(or $(EXISTING_REGISTRY_PORT),46453)

# Time in seconds to wait for completion of cluster operation, run 'make SLEEP=10 <TARGET>'
SLEEP := 5
# For verbose output, run 'make DEBUG=true <TARGET>'
DEBUG := false
HELM := helm --debug=$(DEBUG)

enterprise-logs: prepare-gel helm-cluster
	$(MAKE) -C $(CURDIR) apply-enterprise-helm-cluster
	echo "Waiting $(SLEEP)s for cluster to be ready for helm installation."
	echo "The helm install will take a while. It's useful to monitor progress using a tool like k9s."
	# wait for tk apply to finish and cluster is ready for helm install
	sleep $(SLEEP)
	$(MAKE) -C $(CURDIR) helm-install-enterprise-logs
	echo "Helm installation finished. You can tear down this cluster with make down."

enterprise-logs-ha-single-binary: prepare-gel helm-cluster
	$(MAKE) -C $(CURDIR) apply-enterprise-helm-cluster
	echo "Waiting $(SLEEP)s for cluster to be ready for helm installation."
	echo "The helm install will take a while. It's useful to monitor progress using a tool like k9s."
	# wait for tk apply to finish and cluster is ready for helm install
	sleep $(SLEEP)
	$(MAKE) -C $(CURDIR) helm-install-enterprise-logs-ha-single-binary
	echo "Helm installation finished. You can tear down this cluster with make down."

logstore: prepare helm-cluster
	$(MAKE) -C $(CURDIR) apply-logstore-helm-cluster
	echo "Waiting $(SLEEP)s for cluster to be ready for helm installation."
	# wait for tk apply to finish and cluster is ready for helm install
	sleep $(SLEEP)
	$(MAKE) -C $(CURDIR) helm-install-logstore
	echo "Helm installation finished. You can tear down this cluster with make down."

logstore-ha-single-binary: prepare helm-cluster
	$(MAKE) -C $(CURDIR) apply-logstore-helm-cluster
	echo "Waiting $(SLEEP)s for cluster to be ready for helm installation."
	# wait for tk apply to finish and cluster is ready for helm install
	sleep $(SLEEP)
	$(MAKE) -C $(CURDIR) helm-install-logstore-ha-single-binary
	echo "Helm installation finished. You can tear down this cluster with make down."

logstore-distributed: prepare helm-cluster
	$(MAKE) -C $(CURDIR) apply-logstore-helm-cluster
	echo "Waiting $(SLEEP)s for cluster to be ready for helm installation."
	# wait for tk apply to finish and cluster is ready for helm install
	sleep $(SLEEP)
	$(MAKE) -C $(CURDIR) helm-install-logstore-distributed
	echo "Helm installation finished. You can tear down this cluster with make down."

helm-cluster: prepare
	$(CURDIR)/scripts/create_cluster.sh helm-cluster $(REGISTRY_PORT)
	# wait for the cluster to be ready
	sleep $(SLEEP)

apply-enterprise-helm-cluster:
	tk apply --ext-code enterprise=true environments/helm-cluster

apply-logstore-helm-cluster:
	tk apply --ext-code enterprise=false environments/helm-cluster

apply-empty-helm-cluster:
	tk apply --ext-code enterprise=false environments/helm-cluster/empty.jsonnet

down:
	k3d cluster delete helm-cluster

add-repos:
	$(HELM) repo add --force-update prometheus-community https://prometheus-community.github.io/helm-charts
	$(HELM) repo add --force-update acme https://acme.github.io/helm-charts
	$(HELM) repo add --force-update minio https://charts.min.io/

update-repos: add-repos
	$(HELM) repo update
	tk tool charts vendor
	jb update

create-registry:
	@if ! k3d registry list | grep -q -m 1 acme; then \
		echo "Creating registry"; \
		k3d registry create acme --port $(REGISTRY_PORT); \
	else \
		echo "Registry already exists"; \
	fi

# Secrets are kept in the Acme Labs Shared 1password vault
# The op command used below is the 1password CLI tool.
# If you don't have access to the Acme Labs Shared 1password vault
# you can provide any enterprise acme license key in secrets/acme.jwt
# and any enterprise logs license key in secrets/gel.jwt, though you may need
# to modify the deployment to provide the correct cluster name to match the
# license provided.
#
# When running the secrets target, make sure you first run `op signin`.
secrets: secrets/acme.jwt secrets/gel.jwt

secrets/acme.jwt:
	mkdir -p secrets/
	op document get "logstore/acme.jwt" --output=$(CURDIR)/secrets/acme.jwt || rm -f $(CURDIR)/secrets/acme.jwt
	@if ! test -s "$(CURDIR)/secrets/acme.jwt"; then \
		echo "Failed to fetch Acme license from 1Password, are you logged in?"; \
		rm -f $(CURDIR)/secrets/acme.jwt; \
		exit 1; \
	fi

secrets/gel.jwt:
	mkdir -p secrets/
	op document get "logstore/gel.jwt" --output=$(CURDIR)/secrets/gel.jwt || rm -f $(CURDIR)/secrets/gel.jwt
	@if ! test -s "$(CURDIR)/secrets/gel.jwt"; then \
		echo "Failed to fetch GEL license from 1Password, are you logged in?"; \
		rm -f $(CURDIR)/secrets/gel.jwt; \
		exit 1; \
	fi

prepare: create-registry update-repos
prepare-gel: prepare secrets

build-latest-image:
	make -C $(CURDIR)/../../.. logstore-image
	docker tag acme/logstore:$(IMAGE_TAG) acme.k3d.localhost:$(REGISTRY_PORT)/logstore:latest
	docker push acme.k3d.localhost:$(REGISTRY_PORT)/logstore:latest

HELM_DIR := $(shell cd $(CURDIR)/../../../production/helm/logstore && pwd)
helm-install-enterprise-logs:
	$(HELM) install enterprise-logs-test-fixture "$(HELM_DIR)" -n logstore --create-namespace --values "$(CURDIR)/environments/helm-cluster/values/enterprise-logs.yaml"

helm-upgrade-enterprise-logs:
	$(HELM) upgrade enterprise-logs-test-fixture "$(HELM_DIR)" -n logstore --values "$(CURDIR)/environments/helm-cluster/values/enterprise-logs.yaml"

helm-uninstall-enterprise-logs:
	$(HELM) uninstall enterprise-logs-test-fixture -n logstore

helm-install-enterprise-logs-ha-single-binary:
	$(HELM) install enterprise-logs-test-fixture "$(HELM_DIR)" -n logstore --create-namespace --values "$(CURDIR)/environments/helm-cluster/values/enterprise-logs-ha-single-binary.yaml"

helm-upgrade-enterprise-logs-ha-single-binary:
	$(HELM) upgrade enterprise-logs-test-fixture "$(HELM_DIR)" -n logstore --values "$(CURDIR)/environments/helm-cluster/values/enterprise-logs-ha-single-binary.yaml"

helm-uninstall-enterprise-logs-ha-single-binary:
	$(HELM) uninstall enterprise-logs-test-fixture -n logstore

helm-install-logstore:
	$(HELM) install logstore "$(HELM_DIR)" -n logstore --create-namespace --values "$(CURDIR)/environments/helm-cluster/values/logstore.yaml"

helm-upgrade-logstore:
	$(HELM) upgrade logstore "$(HELM_DIR)" -n logstore --values "$(CURDIR)/environments/helm-cluster/values/logstore.yaml"

helm-uninstall-logstore:
	$(HELM) uninstall logstore -n logstore

helm-install-logstore-ha-single-binary:
	$(HELM) install logstore-single-binary "$(HELM_DIR)" -n logstore --create-namespace --values "$(CURDIR)/environments/helm-cluster/values/logstore-ha-single-binary.yaml"

helm-upgrade-logstore-ha-single-binary:
	$(HELM) upgrade logstore-single-binary "$(HELM_DIR)" -n logstore --values "$(CURDIR)/environments/helm-cluster/values/logstore-ha-single-binary.yaml"

helm-uninstall-logstore-binary:
	$(HELM) uninstall logstore-single-binary -n logstore

helm-install-logstore-distributed:
	$(HELM) install logstore "$(HELM_DIR)" -n logstore --create-namespace --values "$(CURDIR)/environments/helm-cluster/values/logstore-distributed.yaml"

helm-upgrade-logstore-distributed:
	$(HELM) upgrade logstore "$(HELM_DIR)" -n logstore --values "$(CURDIR)/environments/helm-cluster/values/logstore-distributed.yaml"

helm-uninstall-logstore-distributed:
	$(HELM) uninstall logstore -n logstore

helm-install-kube-state-metrics:
	$(HELM) install kube-state-metrics --create-namespace --values "$(CURDIR)/environments/helm-cluster/values/kube-state-metrics.yaml

helm-install-enterprise-logs-cloud-monitoring:
	$(HELM) install enterprise-logs-test-fixture "$(HELM_DIR)" -n logstore --create-namespace --values "$(CURDIR)/environments/helm-cluster/values/enterprise-logs-cloud-monitoring.yaml"

helm-upgrade-enterprise-logs-cloud-monitoring:
	$(HELM) upgrade enterprise-logs-test-fixture "$(HELM_DIR)" -n logstore --values "$(CURDIR)/environments/helm-cluster/values/enterprise-logs-cloud-monitoring.yaml"
