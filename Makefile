.PHONY: build test check golint govendor run
.DEFAULT_GOAL: build

PACKAGES := ./poller ./types ./web ./model
BUILD_DIR := $(CURDIR)/bin/

run: test build
	heroku local

build:
	GOBIN=$(BUILD_DIR) go install ./cmd/...
	echo "Build Complete"

watch:
	reflex -r '\.go$$' make fmt build

test: fmt
	govendor test +local -test.timeout 60s -test.race

fmt:
# golint ${PACKAGES}
	go vet ${PACKAGES} ./utils/... ./cmd/...
	go fmt ${PACKAGES} ./utils/... ./cmd/...
	goimports -w ${PACKAGES} ./utils ./cmd

init:
	go get -u github.com/cespare/reflex
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/kardianos/govendor
	go get -u github.com/golang/lint/golint
	go get -u github.com/kardianos/govendor
