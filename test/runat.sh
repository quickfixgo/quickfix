#!/bin/sh

PORT=5001

./echo_server &

ruby -I. Runner.rb 127.0.0.1 $PORT definitions/server/*/*.def 
result=$?
killall echo_server

exit $result
