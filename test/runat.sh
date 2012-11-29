#!/bin/sh

PORT=5001

./echo_server &

ruby -I. Runner.rb 127.0.0.1 $PORT definitions/server/fix4*/1[a-e]_*.def definitions/server/fix5*/1[a-f]_*.def
result=$?
killall echo_server

exit $result
