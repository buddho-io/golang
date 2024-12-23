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

ROOT := $(shell git rev-parse --show-toplevel)

TEST_OPTS ?= -race
TEST_PACKAGES ?= $(shell go list ./...)
.PHONY: test
test:  ## Runs the test suites
	@printf "$(ansi_format_dark)" $(PROJECT) $@
	@go test $(TEST_OPTS) -p 1 --timeout 90s -tags "$(TEST_TAGS)" $(TEST_PACKAGES)
	@printf "$(ansi_format_bright)" $(PROJECT) "$@ ok"

.PHONY: benchmark
benchmark:  ## Runs the benchmark tests
	@printf "$(ansi_format_dark)" $(PROJECT) $@
	@go test $(TEST_OPTS) -tags "$(TEST_TAGS)" -run=XXX -benchmem -bench=. -benchtime=10x $(TEST_PACKAGES)
	@printf "$(ansi_format_bright)" $(PROJECT) "$@ ok"

COVERAGE_PKG ?= "$(PKG)/..."
COVERAGE_PACKAGES ?= $(shell go list ./...)
.PHONY: coverage
coverage: ## Creates coverage report for all projects
	@printf "$(ansi_format_dark)" $(PROJECT) $@
	@mkdir -p build
	@go test $(COVERAGE_OPTS) \
		-timeout 60s \
		-coverprofile build/coverage.out \
		-covermode atomic \
		-coverpkg $(COVERAGE_PKG) \
		-tags $(TEST_TAGS) $(COVERAGE_PACKAGES)
	@go tool cover -html="build/coverage.out" -o "build/coverage.html"
	@printf "$(ansi_format_bright)" $(PROJECT) "$@ ok"

.PHONY: clean-coverage
clean-coverage:
	@printf "$(ansi_format_dark)" $(PROJECT) $@
	@rm -rf build/coverage.out build/coverage.html
	@printf "$(ansi_format_bright)" $(PROJECT) "$@ ok"
