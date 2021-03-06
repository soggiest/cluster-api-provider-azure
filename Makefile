# Copyright 2018 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# If you update this file, please follow
# https://suva.sh/posts/well-documented-makefiles

.DEFAULT_GOAL:=help

# A release should define this with quay.io/k8s
# TODO: Uncomment once production and staging images are stored in GCR
#REGISTRY ?= gcr.io/$(shell gcloud config get-value project)
REGISTRY ?= quay.io/k8s-staging

# A release should define this with IfNotPresent
PULL_POLICY ?= Always

# A release does not need to define this
MANAGER_IMAGE_NAME ?= cluster-api-azure-controller

# A release should define this with the next version after 0.1.0
MANAGER_IMAGE_TAG ?= dev

MANAGER_IMAGE ?= $(REGISTRY)/$(MANAGER_IMAGE_NAME):$(MANAGER_IMAGE_TAG)

CLUSTER_NAME ?= test1

## Image URL to use all building/pushing image targets
BAZEL_ARGS ?=

BAZEL_BUILD_ARGS := --define=REGISTRY=$(REGISTRY)\
 --define=PULL_POLICY=$(PULL_POLICY)\
 --define=MANAGER_IMAGE_NAME=$(MANAGER_IMAGE_NAME)\
 --define=MANAGER_IMAGE_TAG=$(MANAGER_IMAGE_TAG)\
$(BAZEL_ARGS)

# Bazel variables
BAZEL_VERSION := $(shell command -v bazel 2> /dev/null)

# Determine the OS
HOSTOS := $(shell go env GOHOSTOS)
HOSTARCH := $(shell go env GOARCH)
BINARYPATHPATTERN :=${HOSTOS}_${HOSTARCH}_*

ifndef BAZEL_VERSION
    $(error "Bazel is not available. \
		Installation instructions can be found at \
		https://docs.bazel.build/versions/master/install.html")
endif

.PHONY: all
all: check-install test binaries

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


## --------------------------------------
## Testing
## --------------------------------------

.PHONY: test
test: generate verify ## Run tests
	bazel test --nosandbox_debug //pkg/... //cmd/... $(BAZEL_ARGS)

# TODO: Enable integration tests
#.PHONY: integration
#integration: generate verify ## Run integraion tests
#	bazel test --define='gotags=integration' --test_output all //test/integration/...

# TODO: Enable e2e tests
#JANITOR_ENABLED ?= 0
#.PHONY: e2e
#e2e: generate verify ## Run e2e tests
#	JANITOR_ENABLED=$(JANITOR_ENABLED) ./hack/e2e.sh

# TODO: Enable e2e tests
#.PHONY: e2e-janitor
#e2e-janitor:
#	./hack/e2e-azure-janitor.sh

## --------------------------------------
## Docker
## --------------------------------------

.PHONY: docker-build
docker-build: generate ## Build the production docker image
	docker build . -t $(MANAGER_IMAGE)

# TODO: Move this to docker-build target once we figure out multi-stage builds and using a thinner image
.PHONY: docker-build-new
docker-build-new: generate ## Build the production docker image
	bazel run //cmd/manager:manager-image $(BAZEL_BUILD_ARGS)

.PHONY: docker-push
docker-push: ## Push production docker image
	docker push $(MANAGER_IMAGE)

# TODO: Move this to docker-push target once we figure out multi-stage builds and using a thinner image
.PHONY: docker-push-new
docker-push-new: ## Push production docker image
	bazel run //cmd/manager:manager-push $(BAZEL_BUILD_ARGS)

## --------------------------------------
## Cleanup / Verification
## --------------------------------------

.PHONY: clean
clean: ## Remove all generated files
	rm -rf cmd/clusterctl/examples/azure/out/
	rm -f kubeconfig
	rm -f minikube.kubeconfig
	rm -f bazel-*
	rm -rf out/
	rm -f cmd/clusterctl/examples/azure/provider-components-base.yaml

.PHONY: check-install
check-install: ## Checks that you've installed this repository correctly
	@./scripts/check-install.sh

.PHONY: verify
verify: ## Runs verification scripts to ensure correct execution
	./hack/verify_boilerplate.py

## --------------------------------------
## Manifests
## --------------------------------------

.PHONY: manifests
manifests: cmd/clusterctl/examples/azure/provider-components-base.yaml
	./cmd/clusterctl/examples/azure/generate-yaml.sh

.PHONY: cmd/clusterctl/examples/azure/provider-components-base.yaml
cmd/clusterctl/examples/azure/provider-components-base.yaml:
	bazel build //cmd/clusterctl/examples/azure:provider-components-base $(BAZEL_BUILD_ARGS)
	install bazel-genfiles/cmd/clusterctl/examples/azure/provider-components-base.yaml cmd/clusterctl/examples/azure

## --------------------------------------
## Generate
## --------------------------------------

.PHONY: vendor
vendor: ## Runs go mod to ensure proper vendoring.
	./hack/update-vendor.sh
	$(MAKE) gazelle

.PHONY: gazelle
gazelle: ## Run Bazel Gazelle
	(which bazel && ./hack/update-bazel.sh) || true

.PHONY: generate
generate: vendor ## Generate mocks, CRDs and runs `go generate` through Bazel
	GOPATH=$(shell go env GOPATH) bazel run //:generate $(BAZEL_ARGS)
	bazel build $(BAZEL_ARGS) //pkg/cloud/azure/mocks:mocks \
		//pkg/cloud/azure/services/disks/mock_disks:mocks \
		//pkg/cloud/azure/services/availabilityzones/mock_availabilityzones:mocks \
		//pkg/cloud/azure/services/groups/mock_groups:mocks \
		//pkg/cloud/azure/services/internalloadbalancers/mock_internalloadbalancers:mocks \
		//pkg/cloud/azure/services/networkinterfaces/mock_networkinterfaces:mocks \
		//pkg/cloud/azure/services/publicips/mock_publicips:mocks \
		//pkg/cloud/azure/services/publicloadbalancers/mock_publicloadbalancers:mocks \
		//pkg/cloud/azure/services/routetables/mock_routetables:mocks \
		//pkg/cloud/azure/services/securitygroups/mock_securitygroups:mocks \
		//pkg/cloud/azure/services/subnets/mock_subnets:mocks \
		//pkg/cloud/azure/services/virtualmachineextensions/mock_virtualmachineextensions:mocks \
		//pkg/cloud/azure/services/virtualmachines/mock_virtualmachines:mocks \
		//pkg/cloud/azure/services/virtualnetworks/mock_virtualnetworks:mocks
	./hack/copy-bazel-mocks.sh
	$(MAKE) generate-deepcopy
	$(MAKE) generate-crds

.PHONY: generate-deepcopy
generate-deepcopy:
	cd pkg/apis && go run ../../vendor/k8s.io/code-generator/cmd/deepcopy-gen/main.go -O zz_generated.deepcopy -i ./... -h ../../hack/boilerplate/boilerplate.go.txt

.PHONY: generate-crds
generate-crds:
	bazel build //config
	cp -R bazel-genfiles/config/crds/* config/crds/
	cp -R bazel-genfiles/config/rbac/* config/rbac/

## --------------------------------------
## Linting
## --------------------------------------

.PHONY: lint
lint: ## Lint codebase
	bazel run //:lint $(BAZEL_ARGS)

lint-full: ## Run slower linters to detect possible issues
	bazel run //:lint-full $(BAZEL_ARGS)

## --------------------------------------
## Binaries
## --------------------------------------

# TODO: Add clusterazureadm target once it exists
.PHONY: binaries
binaries: generate manager clusterctl ## Builds and installs all binaries

.PHONY: manager
manager: ## Build manager binary.
	bazel build //cmd/manager $(BAZEL_ARGS)
	install bazel-bin/cmd/manager/${BINARYPATHPATTERN}/manager $(shell go env GOPATH)/bin/azure-manager

.PHONY: clusterctl
clusterctl: ## Build clusterctl binary.
	bazel build --workspace_status_command=./hack/print-workspace-status.sh //cmd/clusterctl $(BAZEL_ARGS)
	install bazel-bin/cmd/clusterctl/${BINARYPATHPATTERN}/clusterctl $(shell go env GOPATH)/bin/clusterctl

# TODO: Uncomment clusterazureadm once it exists
#.PHONY: clusterazureadm
#clusterazureadm: ## Build clusterazureadm binary.
#	bazel build --workspace_status_command=./hack/print-workspace-status.sh //cmd/clusterazureadm $(BAZEL_ARGS)
#	install bazel-bin/cmd/clusterazureadm/${BINARYPATHPATTERN}/clusterazureadm $(shell go env GOPATH)/bin/clusterazureadm

## --------------------------------------
## Release
## --------------------------------------

# TODO: Uncomment clusterazureadm once it exists
.PHONY: release-artifacts
release-artifacts: ## Build release artifacts
	bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/clusterctl #//cmd/clusterazureadm
	bazel build --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64 //cmd/clusterctl #//cmd/clusterazureadm
	bazel build //cmd/clusterctl/examples/azure $(BAZEL_BUILD_ARGS)
	mkdir -p out
#	install bazel-bin/cmd/clusterazureadm/darwin_amd64_pure_stripped/clusterazureadm out/clusterazureadm-darwin-amd64
#	install bazel-bin/cmd/clusterazureadm/linux_amd64_pure_stripped/clusterazureadm out/clusterazureadm-linux-amd64
	install bazel-bin/cmd/clusterctl/darwin_amd64_pure_stripped/clusterctl out/clusterctl-darwin-amd64
	install bazel-bin/cmd/clusterctl/linux_amd64_pure_stripped/clusterctl out/clusterctl-linux-amd64
	install bazel-bin/cmd/clusterctl/examples/azure/azure.tar out/cluster-api-provider-azure-examples.tar

## --------------------------------------
## Define local development targets here
## --------------------------------------

.PHONY: binaries-dev
binaries-dev: ## Builds and installs all development binaries using go get
	go get -v ./...

.PHONY: create-cluster
create-cluster: binaries-dev ## Create a development Kubernetes cluster on Azure using examples
	clusterctl create cluster -v 4 \
	--provider azure \
	--bootstrap-type kind \
	-m ./cmd/clusterctl/examples/azure/out/machines.yaml \
	-c ./cmd/clusterctl/examples/azure/out/cluster.yaml \
	-p ./cmd/clusterctl/examples/azure/out/provider-components.yaml \
	-a ./cmd/clusterctl/examples/azure/out/addons.yaml

.PHONY: delete-cluster
delete-cluster: binaries-dev ## Deletes the development Kubernetes Cluster
	clusterctl delete cluster -v 4 \
	--bootstrap-type kind \
	--kubeconfig ./kubeconfig \
	-p ./cmd/clusterctl/examples/azure/out/provider-components.yaml

kind-reset: ## Destroys the "clusterapi" kind cluster.
	kind delete cluster --name=clusterapi || true

.PHONY: reset-bazel
reset-bazel: ## Deep cleaning for bazel
	bazel clean --expunge
