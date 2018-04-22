REPOSITORY = github.com/sgykfjsm/gu

## Initialize the development environment
init:
	go get -u github.com/golang/dep/cmd/dep
	dep init

## Resolve the dependencies
deps:
	dep ensure
	@if [ "$(ARGS)" = "update" ]; then \
		dep ensure --update; \
	fi

## Format files as Golang program
fmt:
	gofmt -s -w .
## After generating binary, start `go test`
test: install
	go test -race -v .

## Install this library into the
install:
	go install -v

################################################################################
# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=20
## Show help
help:
	@echo ''
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

