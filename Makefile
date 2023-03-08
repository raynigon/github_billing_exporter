GO    		 	:= GO111MODULE=on go
GOOPTS       ?=
GOHOSTOS     ?= $(shell $(GO) env GOHOSTOS)
GOHOSTARCH   ?= $(shell $(GO) env GOHOSTARCH)
FIRST_GOPATH 	:= $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))

ifeq (arm, $(GOHOSTARCH))
	GOHOSTARM ?= $(shell GOARM= $(GO) env GOARM)
	GO_BUILD_PLATFORM ?= $(GOHOSTOS)-$(GOHOSTARCH)v$(GOHOSTARM)
else
	GO_BUILD_PLATFORM ?= $(GOHOSTOS)-$(GOHOSTARCH)
endif

PROMU 		 	:= $(FIRST_GOPATH)/bin/promu
PROMU_VERSION 	?= 0.14.0
PROMU_URL 		:= https://github.com/prometheus/promu/releases/download/v$(PROMU_VERSION)/promu-$(PROMU_VERSION).$(GO_BUILD_PLATFORM).tar.gz

pkgs   			=  ./...

PREFIX                  ?= $(shell pwd)
BIN_DIR                 ?= $(shell pwd)
DOCKER_IMAGE_NAME       ?= github_billing_exporter
DOCKER_IMAGE_TAG        ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))

all: format build test

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

test: format vet

build: promu
	@echo ">> building binaries"
	@$(PROMU) build --prefix $(PREFIX)

tarball: promu
	@echo ">> building release tarball"
	@$(PROMU) tarball --prefix $(PREFIX) $(BIN_DIR)

docker:
	@echo ">> building docker image"
	@docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .

promu: $(PROMU)

$(PROMU):
	$(eval PROMU_TMP := $(shell mktemp -d))
	curl -s -L $(PROMU_URL) | tar -xvzf - -C $(PROMU_TMP)
	mkdir -p $(FIRST_GOPATH)/bin
	cp $(PROMU_TMP)/promu-$(PROMU_VERSION).$(GO_BUILD_PLATFORM)/promu $(FIRST_GOPATH)/bin/promu
	rm -r $(PROMU_TMP)

.PHONY: all style format build test vet tarball docker promu
