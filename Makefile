.PHONY: build test check golint govendor run
.DEFAULT_GOAL: build

PACKAGES := ./poller ./types ./web ./model

run: test build
	heroku local

build:
	go build
	echo "Build Complete"

watch:
	reflex -r '\.go$$' make fmt build

test: fmt
	govendor test +local -test.timeout 60s -test.race

fmt:
# golint ${PACKAGES}
	go vet ${PACKAGES} ./utils/... .
	go fmt ${PACKAGES} ./utils/... .
	goimports -w ${PACKAGES} ./utils *.go

init:
	go get -u github.com/cespare/reflex
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/kardianos/govendor
	go get -u github.com/golang/lint/golint
	go get -u github.com/kardianos/govendor
