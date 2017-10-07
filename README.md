QuickFIX/Go
===========

[![GoDoc](https://godoc.org/github.com/quickfixgo/quickfix?status.png)](https://godoc.org/github.com/quickfixgo/quickfix) [![Build Status](https://travis-ci.org/quickfixgo/quickfix.svg?branch=master)](https://travis-ci.org/quickfixgo/quickfix) [![Go Report Card](https://goreportcard.com/badge/github.com/quickfixgo/quickfix)](https://goreportcard.com/report/github.com/quickfixgo/quickfix)

- Website: http://www.quickfixgo.org
- Mailing list: [Google Groups](https://groups.google.com/forum/#!forum/quickfixgo)

Open Source [FIX Protocol](http://www.fixprotocol.org/) library implemented in Go

Getting Started and Documentation
---------------------------------

* [User Manual](http://quickfixgo.org/docs)
* [API Documentation](https://godoc.org/github.com/quickfixgo/quickfix)

### Installation

To install QuickFIX/Go, use `go get`:

```sh
$ go get github.com/quickfixgo/quickfix
```

### Staying up to date

To update QuickFIX/Go to the latest version, use `go get -u github.com/quickfixgo/quickfix`.

### Example Apps

See [examples](https://github.com/quickfixgo/examples) for some simple examples of using QuickFIX/Go.

### FIX Message Generation

QuickFIX/Go includes separate packages for tags, fields, enums, messages, and message components generated from the FIX 4.0 - FIX5.0SP2 specs. See:

* [github.com/quickfixgo/tag](https://github.com/quickfixgo/tag)
* [github.com/quickfixgo/field](https://github.com/quickfixgo/field)
* [github.com/quickfixgo/enum](https://github.com/quickfixgo/enum)
* [github.com/quickfixgo/fix40](https://github.com/quickfixgo/fix40)
* [github.com/quickfixgo/fix41](https://github.com/quickfixgo/fix41)
* [github.com/quickfixgo/fix42](https://github.com/quickfixgo/fix42)
* [github.com/quickfixgo/fix43](https://github.com/quickfixgo/fix43)
* [github.com/quickfixgo/fix44](https://github.com/quickfixgo/fix44)
* [github.com/quickfixgo/fix50](https://github.com/quickfixgo/fix50)
* [github.com/quickfixgo/fix50sp1](https://github.com/quickfixgo/fix50sp1)
* [github.com/quickfixgo/fix50sp2](https://github.com/quickfixgo/fix50sp2)
* [github.com/quickfixgo/fixt11](https://github.com/quickfixgo/fixt11)

For most FIX applications, these generated resources are sufficient. Custom FIX applications may generate source specific to the FIX spec of that application using the `generate-fix` tool included with QuickFIX/Go.

Following installation, `generate-fix` is installed to `$GOPATH/bin/generate-fix`. Run `$GOPATH/bin/generate-fix --help` for usage instructions.

Developing QuickFIX/Go
----------------------

If you wish to work on QuickFIX/Go itself, you will first need [Go](http://www.golang.org) installed on your machine (version 1.6+ is *required*).

For local dev first make sure Go is properly installed, including setting up a [GOPATH](http://golang.org/doc/code.html#GOPATH).

Next, using [Git](https://git-scm.com/), clone this repository into `$GOPATH/src/github.com/quickfixgo/quickfix`. 

### Installing Dependencies

QuickFIX/Go uses [dep](https://github.com/golang/dep) to manage the vendored dependencies. Install dep with `go get`:

```sh
$ go get -u github.com/golang/dep/cmd/dep
```

Run `dep ensure` to install the correct versioned dependencies into `vendor/`, which Go 1.6+ automatically recognizes and loads.

```sh
$ $GOPATH/bin/dep ensure
```

**Note:** No vendored dependencies are included in the QuickFIX/Go source.

### Build and Test

The default make target runs [go vet](https://godoc.org/golang.org/x/tools/cmd/vet) and unit tests.

```sh
$ make
```

If this exits with exit status 0, then everything is working!

### Generated Code

Generated code from the FIX40-FIX50SP2 specs are available as separate repos under the [QuickFIX/Go organization](https://github.com/quickfixgo).  The source specifications for this generated code is located in `spec/`.  Generated code can be identified by the `.generated.go` suffix.  Any changes to generated code must be captured by changes to source in `cmd/generate-fix`.  After making changes to the code generator source, run the following to re-generate the source

```sh
$ make generate-dist
```

If you are making changes to the generated code, please create Pull Requests for these changes for the affected repos.

### Acceptance Tests

QuickFIX/Go has a comprehensive acceptance test suite covering the FIX protocol.  These are the same tests used across all QuickFIX implementations.

QuickFIX/Go acceptance tests depend on ruby in path.

To run acceptance tests,

        # generate code locally
        make generate

		# build acceptance test rig
		make build_accept

		# run acceptance tests
		make accept

### Dependencies

If you are developing QuickFIX/Go, there are a few tasks you might need to perform related to dependencies.

#### Adding a dependency

If you are adding a dependency, you will need to update the dep manifest in the same Pull Request as the code that depends on it. You should do this in a separate commit from your code, as this makes PR review easier and Git history simpler to read in the future.

To add a dependency:

1. Add the dependency using `dep`:
```bash
$ dep ensure -add github.com/foo/bar
```
2. Review the changes in git and commit them.

#### Updating a dependency

To update a dependency to the latest version allowed by constraints in `Gopkg.toml`:

1. Run:
```bash
$ dep ensure -update github.com/foo/bar
```
2. Review the changes in git and commit them.

To change the allowed version/branch/revision of a dependency:

1. Manually edit `Gopkg.toml`
2. Run:
```bash
$ dep ensure
```
3. Review the changes in git and commit them.

Licensing
---------

This software is available under the QuickFIX Software License. Please see the [LICENSE.txt](https://github.com/quickfixgo/quickfix/blob/master/LICENSE.txt) for the terms specified by the QuickFIX Software License.
