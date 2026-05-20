#!/bin/sh

CFG=$1
PORT=$2
STORETYPE=$3
TESTS=$4

./echo_server $CFG $STORETYPE &
pid=$!

ruby -I. Runner.rb 127.0.0.1 $PORT $TESTS
result=$?
kill -kill $pid

exit $result
