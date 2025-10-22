
all: vet test

clean:
	rm -rf gen

generate: clean
	mkdir -p gen; cd gen; go run ../cmd/generate-fix/generate-fix.go -pkg-root=github.com/quickfixgo/quickfix/gen ../spec/*.xml

generate-udecimal: clean
	mkdir -p gen; cd gen; go run ../cmd/generate-fix/generate-fix.go -use-udecimal=true -pkg-root=github.com/quickfixgo/quickfix/gen ../spec/*.xml

fmt:
	gofmt -l -w -s $(shell find . -type f -name '*.go')

vet:
	go vet `go list ./... | grep -v quickfix/gen`

test: 
	MONGODB_TEST_CXN=mongodb://db:27017 go test -v -cover `go list ./... | grep -v quickfix/gen`

linters-install:
	@golangci-lint --version >/dev/null 2>&1 || { \
		echo "installing linting tools..."; \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.64.6; \
	}

lint: linters-install
	golangci-lint run

# An easy way to run the linter without going through the install process -
# docker run -t --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.64.6 golangci-lint run -v
# See https://golangci-lint.run/welcome/install/ for more details.

# ---------------------------------------------------------------
# Targets related to running acceptance tests -

ifdef STORE_TYPE
STORE := $(STORE_TYPE)
else
STORE := memory
endif

ifdef ACCEPTANCE_SET
TEST_SET := $(ACCEPTANCE_SET)
else
TEST_SET := server
endif

build-test-srv:
	cd _test; go build -v -o echo_server ./test-server/
fix40:
	cd _test; ./runat.sh cfg/$(TEST_SET)/$@.cfg 5001 $(STORE) "definitions/$(TEST_SET)/$@/*.def"
fix41:
	cd _test; ./runat.sh cfg/$(TEST_SET)/$@.cfg 5002 $(STORE) "definitions/$(TEST_SET)/$@/*.def"
fix42:
	cd _test; ./runat.sh cfg/$(TEST_SET)/$@.cfg 5003 $(STORE) "definitions/$(TEST_SET)/$@/*.def"
fix43:
	cd _test; ./runat.sh cfg/$(TEST_SET)/$@.cfg 5004 $(STORE) "definitions/$(TEST_SET)/$@/*.def"
fix44:
	cd _test; ./runat.sh cfg/$(TEST_SET)/$@.cfg 5005 $(STORE) "definitions/$(TEST_SET)/$@/*.def"
fix50:
	cd _test; ./runat.sh cfg/$(TEST_SET)/$@.cfg 5006 $(STORE) "definitions/$(TEST_SET)/$@/*.def"
fix50sp1:
	cd _test; ./runat.sh cfg/$(TEST_SET)/$@.cfg 5007 $(STORE) "definitions/$(TEST_SET)/$@/*.def"
fix50sp2:
	cd _test; ./runat.sh cfg/$(TEST_SET)/$@.cfg 5008 $(STORE) "definitions/$(TEST_SET)/$@/*.def"

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
	go test -v -cover `go list ./... | grep -v quickfix/gen`

generate-ci: clean
	mkdir -p gen; cd gen; go run ../cmd/generate-fix/generate-fix.go -pkg-root=github.com/quickfixgo/quickfix/gen ../spec/$(shell echo $(FIX_TEST) | tr  '[:lower:]' '[:upper:]').xml;

generate-ci-udecimal: clean
	mkdir -p gen; cd gen; go run ../cmd/generate-fix/generate-fix.go -use-udecimal=true -pkg-root=github.com/quickfixgo/quickfix/gen ../spec/$(shell echo $(FIX_TEST) | tr  '[:lower:]' '[:upper:]').xml;

# ---------------------------------------------------------------
