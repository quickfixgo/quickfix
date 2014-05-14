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
	go test -v . ./datadictionary/ ./message/ ./fix/

build_accept:
		cd test; go build -o echo_server

fix40:
	cd test; ./runat.sh $@.cfg 5001 "definitions/server/$@/*.def"
fix41:
	cd test; ./runat.sh $@.cfg 5002 "definitions/server/$@/*.def"
fix42:
	cd test; ./runat.sh $@.cfg 5003 "definitions/server/$@/*.def"
fix43:
	cd test; ./runat.sh $@.cfg 5004 "definitions/server/$@/*.def"
fix44:
	cd test; ./runat.sh $@.cfg 5005 "definitions/server/$@/*.def"
fix50:
	cd test; ./runat.sh $@.cfg 5006 "definitions/server/$@/*.def"
fix50sp1:
	cd test; ./runat.sh $@.cfg 5007 "definitions/server/$@/*.def"
fix50sp2:
	cd test; ./runat.sh $@.cfg 5008 "definitions/server/$@/*.def"

ACCEPT_SUITE=fix40 fix41 fix42 fix43 fix44 fix50 fix50sp1 fix50sp2 
accept: $(ACCEPT_SUITE)

# travis VMs have a difficult time if # of builds > 1
_travis_build_accept:
	cd test; go build -p 1 -o echo_server

travis_test: all _travis_build_accept accept

.PHONY: test $(ACCEPT_SUITE)
