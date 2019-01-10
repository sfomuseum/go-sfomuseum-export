CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

rmdeps:
	if test -d src; then rm -rf src; fi 

self:   prep
	if test ! -d src; then rm -rf src; fi
	mkdir -p src/github.com/sfomuseum/go-sfomuseum-export
	cp export.go src/github.com/sfomuseum/go-sfomuseum-export/
	cp -r properties src/github.com/sfomuseum/go-sfomuseum-export/
	cp -r vendor/* src/

deps:   rmdeps
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-export"

vendor-deps: deps
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt export.go
	go fmt properties/*.go

bin:	self
	@GOPATH=$(GOPATH) go build -o bin/sfom-export-feature cmd/sfom-export-feature.go
