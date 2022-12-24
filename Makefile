
all: vet test

clean:
	rm -rf gen

generate: clean
	mkdir -p gen; cd gen; go run ../cmd/generate-fix/generate-fix.go -pkg-root=github.com/quickfixgo/quickfix/gen ../spec/*.xml

fmt:
	gofmt -l -w -s $(shell find . -type f -name '*.go')

vet:
	go vet `go list ./... | grep -v quickfix/gen`

test: 
	MONGODB_TEST_CXN=mongodb://db:27017 go test -v -cover . ./datadictionary ./internal

linters-install:
	@golangci-lint --version >/dev/null 2>&1 || { \
		echo "installing linting tools..."; \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.50.1; \
	}

lint: linters-install
	golangci-lint run

# ---------------------------------------------------------------
# Targets related to running acceptance tests -

build-test-srv:
	cd _test; go build -o echo_server ./test-server/
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
# ---------------------------------------------------------------

# ---------------------------------------------------------------
# These targets are specific to the Github CI Runner -

build-src:
	go build -v `go list ./...`

build: build-src build-test-srv

test-ci:
	go test -v -cover . ./datadictionary ./internal

generate-ci: clean
	mkdir -p gen; cd gen; go run ../cmd/generate-fix/generate-fix.go -pkg-root=github.com/quickfixgo/quickfix/gen ../spec/$(shell echo $(FIX_TEST) | tr  '[:lower:]' '[:upper:]').xml; 

# ---------------------------------------------------------------
