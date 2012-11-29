all:
	go build
	cd test; go build -o echo_server

check:
	go test ./...

accept:
	cd test; ./runat.sh
