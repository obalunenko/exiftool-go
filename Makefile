NAME=exiftool-go

SHELL := env VERSION=$(VERSION) $(SHELL)
VERSION ?= $(shell git describe --tags $(git rev-list --tags --max-count=1))


# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


TARGET_MAX_CHAR_NUM=20


define colored
	@echo '${GREEN}$1${RESET}'
endef

## Show help
help:
	${call colored, help is running...}
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## vet project
vet:
	${call colored, vet is running...}
	./scripts/linting/run-vet.sh
.PHONY: vet

## Run full linting
lint-full:
	./scripts/linting/run-linters.sh
.PHONY: lint-full

## Run linting for build pipeline
lint-pipeline:
	./scripts/linting/run-linters-pipeline.sh
.PHONY: lint-pipeline

## Test the project (excluding integration tests).
test:
	${call colored, test is running...}
	./scripts/tests/run.sh
.PHONY: test

## Test the project including integration tests.
test-integration:
	${call colored, test-integration is running...}
	./scripts/tests/run-integration.sh
.PHONY: test-integration

## Test coverage report.
test-cover:
	${call colored, test-cover is running...}
	./scripts/tests/coverage.sh
.PHONY: test-cover

# Sync vendor
sync-vendor:
	${call colored, sync-vendor is running...}
	./scripts/sync-vendor.sh
.PHONY: sync-vendor


## Fix imports sorting.
imports:
	${call colored, fix-imports is running...}
	./scripts/style/fix-imports.sh
.PHONY: imports

## Format code with go fmt.
fmt:
	${call colored, fmt is running...}
	./scripts/style/fmt.sh
.PHONY: fmt

## Format code with golines.
fmt-golines:
	${call colored, fmt-golines is running...}
	./scripts/style/golines.sh
.PHONY: fmt-golines

## Format code and sort imports.
format-project: fmt-golines fmt imports
.PHONY: format-project

## Installs tools from vendor.
install-vendored-tools:
	./scripts/install/vendored-tools.sh
.PHONY: install-vendored-tools

## Installs all tools.
install-tools: install-vendored-tools install-exiftool
.PHONY: install-tools

## Install exiftool.
install-exiftool:
	./scripts/install/exiftool.sh
.PHONY: install-exiftool