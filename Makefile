.PHONY: build test check golint govendor run
.DEFAULT_GOAL: build

run: build
	forego run ./eve-dwh-golang

build: test
	go build

test: fmt
	govendor test +local -test.timeout 60s -test.race

fmt:
	golint *.go
	go vet *.go
	go fmt *.go
	goimports -w *.go

init:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/kardianos/govendor
	go get -u github.com/golang/lint/golint
	go get -u github.com/kardianos/govendor
