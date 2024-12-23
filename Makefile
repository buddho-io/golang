# This script uses automatic variables (ex $<, $(@D)) and substitution references $(<:.signed=)
# Please see GNU make's documentation if unfamiliar: https://www.gnu.org/software/make/manual/html_node/

default: test

include scripts/common-general.mk

PROJECT := "root"

# Projects to run tests, static analysis and module release for
LIB_PROJECTS := ext

test/%:
	@$(MAKE) -C $* $(subst /$*,,$@)

.PHONY: test
test: $(LIB_PROJECTS:%=test/%) ## Run all tests

benchmark/%:
	@$(MAKE) -C $* $(subst /$*,,$@)

.PHONY: benchmark
benchmark: $(LIB_PROJECTS:%=benchmark/%) ## Run all benchmarks

lint/%:
	@$(MAKE) -C $* $(subst /$*,,$@)

.PHONY: lint
lint: $(LIB_PROJECTS:%=lint/%) ## Run lint checks on all sub-projects

format/%:
	@$(MAKE) -C $* $(subst /$*,,$@)

.PHONY: format
format: $(LIB_PROJECTS:%=format/%) ## Run format on all sub-projects

coverage/%:
	@$(MAKE) -C $* $(subst /$*,,$@)

.PHONY: coverage
coverage: $(LIB_PROJECTS:%=coverage/%) ## Run coverage on all sub-projects

check/%:
	@$(MAKE) -C $* $(subst /$*,,$@)

.PHONY: check
check: $(LIB_PROJECTS:%=check/%) diff  ## CI blocks merge until this passes. If this fails, run "make check" locally and commit the difference.

.PHONY: tidy
## Runs tools like `go mod tidy` in all modules.
tidy: $(LIB_PROJECTS:%=tidy/%)

tidy/%:
	@$(MAKE) -C $* $(subst /$*,,$@)
