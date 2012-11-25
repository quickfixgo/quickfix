all:
	go build
	cd test; go build -o echo_server

check:
	go test ./...

accept:
	./test/echo_server &
	cd test; ruby -I. Runner.rb localhost 5001 definitions/server/fix40/1a_ValidLogonWithCorrectMsgSeqNum.def || true
	killall echo_server
