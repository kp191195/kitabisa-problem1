GOCMD := go
GOBUILD := ${GOCMD} build
BIN_NAME := problem1

PROJECT_ROOT := ${shell pwd}
PROJECT_DEBUG_OUTPUT := ${PROJECT_ROOT}/bin/debug
PROJECT_RELEASE_OUTPUT := ${PROJECT_ROOT}/bin/release

.PHONY: test
test:
	@$(GOCMD) test -v

.PHONY: compile
compile:
	@$(GOBUILD) -o $(PROJECT_DEBUG_OUTPUT)/$(BIN_NAME)

.PHONY: run
run: compile
	@${PROJECT_DEBUG_OUTPUT}/${BIN_NAME}
