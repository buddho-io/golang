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

GOLANGCI_LINT ?= github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
GOSIMPORTS    ?= github.com/rinchsan/gosimports/cmd/gosimports@v0.3.8
SWEYES        ?= github.com/apache/skywalking-eyes/cmd/license-eye@v0.6.0

ifneq ($(strip $(RELEASE_TAG)),)
# Remove the suffix as we want N.N.N instead of vN.N.N or release-vN.N.N
VERSION ?= $(strip $(shell echo $(RELEASE_TAG) | sed 's/.*v//g'))
DEP_VERSION ?= "v$(VERSION)"
else
VERSION ?= $(shell git rev-parse HEAD)
DEP_VERSION ?= $(VERSION)
endif

# Set a default remote to push tags to
GIT_REMOTE ?= origin

GIT_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)

# Default is to not publish a release tag when releasing a library
# Set as 1 in module libraries to publish version tags.
PUBLISH_LIBRARY ?= 0

# Default build tags for tests
TEST_TAGS ?= ""

# Default build tags
BUILD_TAGS ?= ""

# ANSI escape codes. f_ means foreground, b_ background.
# See https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters
b_black            := $(shell printf "\33[40m")
f_white            := $(shell printf "\33[97m")
f_gray             := $(shell printf "\33[37m")
f_dark_gray        := $(shell printf "\33[90m")
f_bright_magenta   := $(shell printf "\33[95m")
b_bright_magenta   := $(shell printf "\33[105m")
ansi_reset         := $(shell printf "\33[0m")
ansi_ngac          := $(b_black)$(f_white)ngac$(ansi_reset)
ansi_format_dark   := $(f_gray)$(f_bright_magenta)%-10s$(ansi_reset) %s\n
ansi_format_bright := $(f_white)$(f_bright_magenta)%-10s$(ansi_reset) $(f_white)$(b_black)%s$(ansi_reset)\n
ansi_format_help   := $(f_gray)$(f_bright_magenta)%-22s$(ansi_reset) %s\n

# Useful variables to convert lists to comma-separated strings
comma := ,
empty :=
space := $(empty) $(empty)

.PHONY: version
version: ## Print the version
	@echo $(VERSION)

.PHONY: variables
variables:
	$(foreach v, $(.VARIABLES), $(info $(v) = $($(v))))

# This formats help statements in ANSI colors. To hide a target from help, don't comment it with a trailing '##'.
.PHONY: help
help: ## Describe how to use each target
	@printf "$(ansi_ngac)$(f_white)\n"
	@awk 'BEGIN {FS = ":.*?## "} /^[0-9a-zA-Z_-]+:.*?## / {printf "$(ansi_format_help)", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: diff
diff: ## CI blocks merge until this passes. If this fails, run "make check" locally and commit the difference.
	@printf "$(ansi_format_dark)" $(PROJECT) $@
	@if [ ! -z "`git status -s $(MAKEFILE_DIR) | grep -v '.proto-descriptor'`" ]; then \
		echo "The following differences will fail CI until committed:"; \
		git diff $(MAKEFILE_DIR); \
		exit 1; \
	fi
	@printf "$(ansi_format_bright)" $(PROJECT) "$@ ok"
