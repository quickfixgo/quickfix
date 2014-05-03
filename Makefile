.PHONY: test
all: vet test

get:
	go get github.com/golang/lint/golint
	go get gopkg.in/check.v1

GEN_MESSAGES = go run gen/generate-messages/main.go
GEN_FIELDS = go run gen/generate-fields/main.go
FIXVERS = FIX40 FIX41 FIX42 FIX43 FIX44 FIX50 FIX50SP1 FIX50SP2 FIXT11

generate:
	$(GEN_FIELDS) $(foreach vers, $(FIXVERS), spec/$(vers).xml)
	$(foreach vers, $(FIXVERS), $(GEN_MESSAGES) spec/$(vers).xml;)

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golint .

test:
	go test -v ./...

build_accept:
		cd test; go build -o echo_server

accept:
		cd test; ./runat.sh

# travis VMs have a difficult time if # of builds > 1
_travis_build_accept:
	cd test; go build -p 1 -o echo_server

travis_test: all _travis_build_accept accept
