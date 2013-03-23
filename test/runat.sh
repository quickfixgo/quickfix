#!/bin/sh

PORT=5001

./echo_server &

ruby -I. Runner.rb 127.0.0.1 $PORT definitions/server/fix[4-5]*/1[a-z]_*.def definitions/server/fix[4-5]*/2[a-d]_*.def 
result=$?
killall echo_server

exit $result
