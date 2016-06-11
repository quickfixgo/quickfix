all: vet test

get:
	go get -t -d ./...

generate:
	go run cmd/generate/*.go spec/*.xml

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	go get github.com/golang/lint/golint
	golint .

test: get
	go test -v -cover . ./datadictionary

_build_all: get
	go build -v ./...

build_accept: get
	cd _test; go build -o echo_server

build: _build_all build_accept

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
