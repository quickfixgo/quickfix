#!/bin/sh

CFG=$1
PORT=$2
TESTS=$3

./echo_server $CFG &
pid=$!

ruby -I. Runner.rb 127.0.0.1 $PORT $TESTS
result=$?
kill -kill $pid

exit $result
