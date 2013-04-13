all:
	go build
	cd test; go build -o echo_server

fmt:
	go fmt ./...

GEN = go run gen/generate-fix/main.go
FIXVERS = FIX40 FIX41 FIX42 FIX43 FIX44 FIX50 FIX50SP1 FIX50SP2 FIXT11

generate:
	$(foreach vers, $(FIXVERS), $(GEN) spec/$(vers).xml;)

unit:
	go test ./...

accept:
	cd test; ./runat.sh

check: unit accept
