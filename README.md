QuickFIX/Go [![GoDoc](https://godoc.org/github.com/quickfixgo/quickfix?status.png)](https://godoc.org/github.com/quickfixgo/quickfix) [![Build Status](https://travis-ci.org/quickfixgo/quickfix.svg?branch=master)](https://travis-ci.org/quickfixgo/quickfix) [![Go Report Card](https://goreportcard.com/badge/github.com/quickfixgo/quickfix)](https://goreportcard.com/report/github.com/quickfixgo/quickfix)
===========

- Website: http://www.quickfixgo.org
- Mailing list: [Google Groups](https://groups.google.com/forum/#!forum/quickfixgo)

Open Source [FIX Protocol](http://www.fixprotocol.org/) library implemented in Go

FIX versions 4.0-5.0

Getting Started and Documentation
---------------------------------

All documentation is available on [GoDoc](https://godoc.org/github.com/quickfixgo/quickfix).

### Installation

To install QuickFIX/Go, use `go get`:

```
go get github.com/quickfixgo/quickfix
```

### Staying up to date

To update QuickFIX/Go to the latest version, use `go get -u github.com/quickfixgo/quickfix`.

### Example Apps

See [examples](https://github.com/quickfixgo/examples) for some simple examples of using QuickFIX/Go.

Developing QuickFIX/Go
----------------------

If you wish to work on QuickFIX/Go itself, you will first need [Go](http://www.golang.org) installed on your machine.

### Build and Test

The default make target runs [go vet](https://godoc.org/golang.org/x/tools/cmd/vet) and unit tests.

### Dependencies

QuickFIX/Go stores its dependencies under `vendor/`, which Go 1.6+ automatically recognize and load.  We use [govendor](https://github.com/kardianos/govendor) to manage the vendored dependencies.


If you are developing QuickFIX/Go, there are a few tasks you might need to perform.

#### Adding a dependency

If you are adding a dependency, you will need to vendor it in the same Pull Request as the code that depends on it. You should do this in a separate commit from your code, as this makes PR review easier and Git history simpler to read in the future.

To add a dependency:

Assuming your work is on a branch called `my-feature-branch`, the steps look like this:

1. Add the new package to your GOPATH:

    ```bash
    go get github.com/quickfixgo/my-project
    ```

2.  Add the new package to your vendor/ directory:

    ```bash
    govendor add github.com/quickfixgo/my-project
    ```

3. Review the changes in git and commit them.

#### Updating a dependency

To update a dependency:

1. Fetch the dependency:

    ```bash
    govendor fetch github.com/quickfixgo/my-project
    ```

2. Review the changes in git and commit them.

### Acceptance Tests

QuickFIX/Go has a comprehensive acceptance test suite covering the FIX protocol.  These are the same tests used across all QuickFIX implementations.

QuickFIX/Go acceptance tests depend on ruby in path.

To run acceptance tests,

		# build acceptance test rig
		make build_accept

		# run acceptance tests
		make accept

Licensing
---------

This software is available under the QuickFIX Software License. Please see the [LICENSE.txt](https://github.com/quickfixgo/quickfix/blob/master/LICENSE.txt) for the terms specified by the QuickFIX Software License.
