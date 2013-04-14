all:
	go build
	cd test; go build -o echo_server

fmt:
	go fmt ./...

GEN_MESSAGES = go run gen/generate-messages/main.go
GEN_FIELDS = go run gen/generate-fields/main.go
FIXVERS = FIX40 FIX41 FIX42 FIX43 FIX44 FIX50 FIX50SP1 FIX50SP2 FIXT11

generate:
	$(GEN_FIELDS) $(foreach vers, $(FIXVERS), spec/$(vers).xml)
	$(foreach vers, $(FIXVERS), $(GEN_MESSAGES) spec/$(vers).xml;)

unit:
	go test ./...

accept:
	cd test; ./runat.sh

check: unit accept
