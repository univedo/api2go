# Auto generated binary variables helper managed by https://github.com/bwplotka/bingo v0.9. DO NOT EDIT.
# All tools are designed to be build inside $GOBIN.
BINGO_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin
GO     ?= $(shell which go)

# Below generated variables ensure that every time a tool under each variable is invoked, the correct version
# will be used; reinstalling only if needed.
# For example for ginkgo variable:
#
# In your main Makefile (for non array binaries):
#
#include .bingo/Variables.mk # Assuming -dir was set to .bingo .
#
#command: $(GINKGO)
#	@echo "Running ginkgo"
#	@$(GINKGO) <flags/args..>
#
GINKGO := $(GOBIN)/ginkgo-v1.16.5
$(GINKGO): $(BINGO_DIR)/ginkgo.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/ginkgo-v1.16.5"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=ginkgo.mod -o=$(GOBIN)/ginkgo-v1.16.5 "github.com/onsi/ginkgo/ginkgo"

GOVER := $(GOBIN)/gover-v0.0.0-20171022184752-b58185e213c5
$(GOVER): $(BINGO_DIR)/gover.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/gover-v0.0.0-20171022184752-b58185e213c5"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=gover.mod -o=$(GOBIN)/gover-v0.0.0-20171022184752-b58185e213c5 "github.com/modocache/gover"

GOVERALLS := $(GOBIN)/goveralls-v0.0.12
$(GOVERALLS): $(BINGO_DIR)/goveralls.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/goveralls-v0.0.12"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=goveralls.mod -o=$(GOBIN)/goveralls-v0.0.12 "github.com/mattn/goveralls"

