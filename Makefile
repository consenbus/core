# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

test:
	go test  -v -bench=. ./...

install-dep:
	go get -u github.com/golang/dep/cmd/dep

dep-init:
	dep init

deps:
	dep ensure
