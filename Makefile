.ONESHELL:
.PHONY: clean build release run run-docker
default: run

APP_NAME=$(shell basename $(CURDIR))
HOST_OS=$(shell uname -s | tr '[:upper:]' '[:lower:]')
HOST_ARCH=$(shell uname -m | tr '[:upper:]' '[:lower:]')
DIST_DIR=$(CURDIR)/dist

clean:
	rm -rf $(DIST_DIR)

build: clean
	CONTAINER_NAME=$(APP_NAME) goreleaser build --snapshot --single-target --rm-dist

run: build
	$(DIST_DIR)/$(APP_NAME)_$(HOST_OS)_$(HOST_ARCH)/$(APP_NAME)

release:
	CONTAINER_NAME=$(APP_NAME) goreleaser release --snapshot --rm-dist
	docker images $(APP_NAME)

run-docker: release
	docker run -v ~/.kube/config:/root/.kube/config --rm $(APP_NAME):latest-$(HOST_ARCH)