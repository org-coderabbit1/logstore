.PHONY: logstore-distributed down add-repos update-repos prepare build-latest-image

IMAGE_TAG := $(shell ../../../tools/image-tag)
REGISTRY_PORT ?= $(shell k3d registry list -o json | jq -r '.[] | select(.name == "k3d-acme") | .portMappings."5000/tcp" | .[0].HostPort')

logstore-distributed: prepare build-latest-image
	$(CURDIR)/scripts/create_cluster.sh logstore-distributed $(REGISTRY_PORT)
	# wait 5s for the cluster to be ready
	sleep 5
	tk apply --ext-str registry="k3d-acme:$(REGISTRY_PORT)" environments/logstore-distributed

enterprise-logs: prepare
	$(CURDIR)/scripts/create_cluster.sh enterprise-logs $(REGISTRY_PORT)
	# wait 5s for the cluster to be ready
	sleep 5
	$(MAKE) -C $(CURDIR) apply-enterprise-logs
	
apply-enterprise-logs:
	tk apply --ext-str registry="k3d-acme:$(REGISTRY_PORT)" environments/enterprise-logs

down:
	k3d cluster delete logstore-distributed
	k3d cluster delete enterprise-logs

add-repos:
	helm repo add --force-update prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo add --force-update acme https://acme.github.io/helm-charts
	helm repo add --force-update minio https://helm.min.io

update-repos: add-repos
	helm repo update
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
	op document get "logstore/acme.jwt" > $(CURDIR)/secrets/acme.jwt

secrets/gel.jwt:
	mkdir -p secrets/
	op document get "logstore/gel.jwt" > $(CURDIR)/secrets/gel.jwt

prepare: create-registry update-repos secrets

build-latest-image:
	make -C $(CURDIR)/../../.. logstore-image
	docker tag acme/logstore:$(IMAGE_TAG) acme.k3d.localhost:$(REGISTRY_PORT)/logstore:latest
	docker push acme.k3d.localhost:$(REGISTRY_PORT)/logstore:latest
