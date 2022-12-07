GOBIN         = $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN         = $(shell go env GOPATH)/bin
endif
GOX           = $(GOBIN)/gox
GOIMPORTS     = $(GOBIN)/goimports
ARCH          = $(shell uname -p)

# ------------------------------------------------------------------------------
#  dependencies

# If go install is run from inside the project directory it will add the
# dependencies to the go.mod file. To avoid that we change to a directory
# without a go.mod file when downloading the following dependencies

$(GOX):
	(cd /; GO111MODULE=on go install github.com/mitchellh/gox@latest)

$(GOIMPORTS):
	(cd /; GO111MODULE=on go install golang.org/x/tools/cmd/goimports@latest)

# ------------------------------------------------------------------------------

all: vet test

clean:
	rm -rf gen

generate: clean
	mkdir -p gen; cd gen; go run ../cmd/generate-fix/generate-fix.go -pkg-root=github.com/quickfixgo/quickfix/gen ../spec/*.xml

generate-dist:
	cd ..; go run quickfix/cmd/generate-fix/generate-fix.go quickfix/spec/*.xml

test-style:
	GO111MODULE=on golangci-lint run

.PHONY: format
format: $(GOIMPORTS)
	GO111MODULE=on go list -f '{{.Dir}}' ./... | xargs $(GOIMPORTS) -w -local github.com/quickfixgo/quickfix

fmt:
	go fmt `go list ./... | grep -v quickfix/gen`

vet:
	go vet `go list ./... | grep -v quickfix/gen`

test: 
	MONGODB_TEST_CXN=mongodb://db:27017 go test -v -cover . ./datadictionary ./internal

build-src: 
	go build -v `go list ./...`

build-test-srv: 
	cd _test; go build -o echo_server ./test-server/

build: build-src build-test-srv

test-ci: 
	go test -v -cover . ./datadictionary ./internal

generate-ci: clean
	mkdir -p gen; cd gen; go run ../cmd/generate-fix/generate-fix.go -pkg-root=github.com/quickfixgo/quickfix/gen ../spec/$(shell echo $(FIX_TEST) | tr  '[:lower:]' '[:upper:]').xml; 

fix40:
	cd _test; ./runat.sh $@.cfg 5001 "definitions/server/$@/*.def"
fix41:
	cd _test; ./runat.sh $@.cfg 5002 "definitions/server/$@/*.def"
fix42:
	cd _test; ./runat.sh $@.cfg 5003 "definitions/server/$@/*.def"
fix43:
	cd _test; ./runat.sh $@.cfg 5004 "definitions/server/$@/*.def"
fix44:
	cd _test; ./runat.sh $@.cfg 5005 "definitions/server/$@/*.def"
fix50:
	cd _test; ./runat.sh $@.cfg 5006 "definitions/server/$@/*.def"
fix50sp1:
	cd _test; ./runat.sh $@.cfg 5007 "definitions/server/$@/*.def"
fix50sp2:
	cd _test; ./runat.sh $@.cfg 5008 "definitions/server/$@/*.def"

ACCEPT_SUITE=fix40 fix41 fix42 fix43 fix44 fix50 fix50sp1 fix50sp2 
accept: $(ACCEPT_SUITE)

.PHONY: test $(ACCEPT_SUITE)
