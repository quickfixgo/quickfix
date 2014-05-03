QuickFIX/Go [![GoDoc](https://godoc.org/github.com/quickfixgo/quickfix?status.png)](https://godoc.org/github.com/quickfixgo/quickfix) [![Build Status](https://travis-ci.org/quickfixgo/quickfix.svg?branch=master)](https://travis-ci.org/quickfixgo/quickfix)
===========

Open Source [FIX Protocol](http://www.fixprotocol.org/) library implemented in Go

FIX versions 4.0-5.0

Example Apps
------------

See [examples](https://github.com/quickfixgo/examples) for some simple examples of using QuickFIX/Go.

Build and Test
--------------

QuickFIX/Go has build dependencies for testing. To fetch Go dependencies, run `make get`. Acceptance tests depend on ruby in path.

The default make target runs [go vet](http://godoc.org/code.google.com/p/go.tools/cmd/vet) and unit tests.

To run acceptance tests,

		# build acceptance test rig
		make build_accept

		# run acceptance tests
		make accept
