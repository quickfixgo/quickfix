all:
	go build
	cd test; go build -o echo_server

fmt:
	go fmt ./...

unit:
	go test ./...

accept:
	cd test; ./runat.sh

check: unit accept
