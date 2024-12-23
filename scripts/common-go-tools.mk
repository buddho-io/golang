# Copyright 2024 BuddhoIO
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

MAKEFILE_PATH := $(abspath Makefile)
MAKEFILE_DIR  := $(dir $(MAKEFILE_PATH))
ROOT          := $(shell git rev-parse --show-toplevel)

GO_VERSION       := $(shell go version | grep -Eo '[0-9]\.[0-9]+\.?[0-9]+')
GO_VERSION_MAJOR := $(shell echo $(GO_VERSION) | cut -d. -f1)
GO_VERSION_MINOR := $(shell echo $(GO_VERSION) | cut -d. -f2)
ifeq (,$(wildcard go.mod))
MIN_GO_VERSION := $(shell cat go.work | grep -m1 go | sed '/^go /s/.* //g')
else
MIN_GO_VERSION := $(shell cat go.mod | grep -m1 go | sed '/^go /s/.* //g')
endif
MIN_GO_VERSION_MAJOR := $(shell echo $(MIN_GO_VERSION) | cut -d. -f1)
MIN_GO_VERSION_MINOR := $(shell echo $(MIN_GO_VERSION) | cut -d. -f2)

.PHONY: validate-go-version
validate-go-version: ## Validates the installed version of go against projects minimum requirement.
	@if [ $(GO_VERSION_MAJOR) -lt $(MIN_GO_VERSION_MAJOR) ]; then \
		echo "ERROR: Go version $(GO_VERSION) is too old. Minimum required version is $(MIN_GO_VERSION)."; \
		exit 1 ;\
	elif [ $(GO_VERSION_MAJOR) -eq $(MIN_GO_VERSION_MAJOR) ]; then \
		if [ $(GO_VERSION_MINOR) -lt $(MIN_GO_VERSION_MINOR) ]; then \
			echo "ERROR: Go version $(GO_VERSION) is too old. Minimum required version is $(MIN_GO_VERSION)." ;\
			exit 1 ;\
		fi \
	fi


LINT_OPTS ?= --timeout 5m
GOLANGCI_LINT_CONFIG ?= $(ROOT)/.golangci.yml
.PHONY: lint
lint: $(GOLANGCI_LINT_CONFIG) ## Lint checks for all Go code
	@printf "$(ansi_format_dark)" $(PROJECT) $@
	@go run $(GOLANGCI_LINT) run $(LINT_OPTS) --build-tags "$(TEST_TAGS)" --config $(GOLANGCI_LINT_CONFIG)
	@printf "$(ansi_format_bright)" $(PROJECT) "$@ ok"

.PHONY: format
format: go.mod ## Format all Go code
	@printf "$(ansi_format_dark)" $(PROJECT) $@
	@go run $(SWEYES) header fix -c $(ROOT)/.licenserc.yaml
	@# -local ensures consistent ordering of our module in imports
	@go run $(GOSIMPORTS) -local github.com/llinder/golang -w .
	@gofmt -w .
	@printf "$(ansi_format_bright)" $(PROJECT) "$@ ok"

check: go.mod tidy diff
	@printf "$(ansi_format_dark)" $(PROJECT) $@
	@$(MAKE) validate-go-version
	@printf "$(ansi_format_bright)" $(PROJECT) "$@ ok"

.PHONY: tidy
tidy: ## Run go mod tidy and import formatting
	@printf "$(ansi_format_dark)" $(PROJECT) $@
	@go mod tidy
	@make format lint
	@printf "$(ansi_format_bright)" $(PROJECT) "$@ ok"
	@$(MAKE) format lint
