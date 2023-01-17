# QuickFIX/Go

[![Build Status](https://github.com/quickfixgo/quickfix/workflows/CI/badge.svg)](https://github.com/quickfixgo/quickfix/actions) [![GoDoc](https://godoc.org/github.com/quickfixgo/quickfix?status.png)](https://godoc.org/github.com/quickfixgo/quickfix) [![Go Report Card](https://goreportcard.com/badge/github.com/quickfixgo/quickfix)](https://goreportcard.com/report/github.com/quickfixgo/quickfix)

Open Source [FIX Protocol](http://www.fixprotocol.org/) library implemented in Go

## About
<p>QuickFIX/Go is a <a href="https://www.fixtrading.org/">FIX Protocol Community</a> implementation for the <a href="https://golang.org">Go programming language</a>.</p> 

<ul>
  <li>100% free and open source with a liberal <a href="https://github.com/quickfixgo/quickfix/blob/master/LICENSE.txt">license</a></li>
  <li>Supports FIX versions 4.0 - 5.0SP2</li>
  <li>Runs on any hardware and operating system supported by Go (1.18+ required)</li>
  <li>Spec driven run-time message validation</li>
  <li>Spec driven code generation of type-safe FIX messages, fields, and repeating groups</li>
  <li>Support for protocol customizations</li>
  <li>Session state storage options: SQL, MongoDB, On-disk, or In-memory</li>
  <li>Logging options: File, Screen</li>
  <li>Failover and High Availability</li>
  <li>Daily and weekly scheduling of session connections</li>
  <li>Integrated support for SSL communicaitons</li>
  <li>Automated unit and acceptance tests</li>
  <li><a href="https://www.connamara.com/">Commercial Support available</a></li>
</ul>

<br>
<img width="208" alt="Sponsored by Connamara" src="https://user-images.githubusercontent.com/3065126/212457799-abd6408a-972d-4168-9feb-b80ce1f1ec83.png">

## Installation

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/quickfixgo/quickfix"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `quickfix` package:

```sh
go get -u github.com/quickfixgo/quickfix
```

## Getting Started

* [QuickFIX User Manual](http://quickfixgo.org/docs)
* [Go API Documentation](https://godoc.org/github.com/quickfixgo/quickfix)
* See [examples](https://github.com/quickfixgo/examples) for some simple examples of using QuickFIX/Go.

## FIX Messaging Model
To send and receive messages, your application will need a few additional packages.

QuickFIX/Go maintains separate packages for tags, fields, enums, messages, and message components auto-generated from the FIX 4.0 - FIX5.0SP2 XML specifications-

* [Tag](https://github.com/quickfixgo/tag)
* [Field](https://github.com/quickfixgo/field)
* [Enum](https://github.com/quickfixgo/enum)
* [FIX 4.0](https://github.com/quickfixgo/fix40)
* [FIX 4.1](https://github.com/quickfixgo/fix41)
* [FIX 4.2](https://github.com/quickfixgo/fix42)
* [FIX 4.3](https://github.com/quickfixgo/fix43)
* [FIX 4.4](https://github.com/quickfixgo/fix44)
* [FIX 5.0](https://github.com/quickfixgo/fix50)
* [FIX 5.0 SP1](https://github.com/quickfixgo/fix50sp1)
* [FIX 5.0 SP2](https://github.com/quickfixgo/fix50sp2)
* [FIXT 1.1](https://github.com/quickfixgo/fixt11)

For most FIX applications, these generated resources are sufficient. Custom FIX applications may generate source specific to the FIX spec of that application using the `generate-fix` tool included with QuickFIX/Go.

Following installation, `generate-fix` is installed to `$GOPATH/bin/generate-fix`. Run `$GOPATH/bin/generate-fix --help` for usage instructions.

## General Support
<h3>Github Discussions</h3>

<p>Our <a href="https://github.com/quickfixgo/quickfix/discussions/categories/q-a">Github Discussions Board</a> is free, public, and easily searchable. It’s the preferred method of user support from the QuickFIX/Go team.

<p>Please provide as much detail as you can when asking a question, and include relevant configurations and code snippets.</p>

<h3>FIX Protocol</h3>

<p>More information about the FIX protocol can be found at the <a href="http://fixtradingcommunity.org">FIX Protocol website</a>.

<h3>Bugs and Issues</h3>

<p>Bugs and issues can be submitted by anyone through our GitHub repository issues list.</p>

<p><strong>Note:</strong> Please do not submit questions or help requests to the issues list. It is for bugs and issues. If you need help, please use the Discussions board as described above and you’ll be able to send your question to the entire community.</p>

<p><a href="https://github.com/quickfixgo/quickfix/issues">GitHub Issues</a></p>

<p>Please provide sample code, logs, and a description of the problem when the issue is submitted.</p>

<p>We will try to address new issues as quickly as possible, and we welcome contributions for bug fixes and new features!</p>

## Commercial Support
<p><a href="https://connamara.com">Connamara Systems</a> offers commercial support for developers who are integrating any of the QuickFIX implementations (Go, C++, Java, .NET). The support is offered in 10-hour bundles and grants developers access, via telephone or email, to the team that created QuickFIX/Go, QuickFIX/n, and are maintainers of QuickFIX.</p>

<p>In addition to offering QuickFIX support, Connamara delivers Made-To-Measure Trading Solutions by bridging the gap between buy and build. By using internally developed trading platform components, Connamara delivers the best of off-the-shelf ISV solutions and custom application development. Coupled with Connamara’s unique licensing model, trading firms can get the best of both build and buy.</p>


## Contributing

If you wish to work on QuickFIX/Go itself, you will need [Docker](https://docs.docker.com/get-docker/) and [VSCode](https://code.visualstudio.com/download) on your machine.

* Clone the repo and open it with VSCode with Docker running
* This repo comes with vscode devcontainer configs in `./.devcontainer/`
* Click the pop-up to re-open the project in the Dev Container
* This opens the project in a docker container pre-configured with everything you need

### Build and Test

The default make target runs [go vet](https://godoc.org/golang.org/x/tools/cmd/vet) and unit tests.

```sh
make
```

If this exits with exit status 0, then everything is working!

### Generated Code

Generated code from the FIX40-FIX50SP2 specs are available as separate repos under the [QuickFIX/Go organization](https://github.com/quickfixgo).  The source specifications for this generated code is located in `spec/`.  Generated code can be identified by the `.generated.go` suffix.  Any changes to generated code must be captured by changes to source in `cmd/generate-fix`.  After making changes to the code generator source, run the following to re-generate the source

```sh
make generate
```

If you are making changes to the generated code, please create Pull Requests for these changes for the affected repos.

### Acceptance Tests

QuickFIX/Go has a comprehensive acceptance test suite covering the FIX protocol.  These are the same tests used across all QuickFIX implementations.

QuickFIX/Go acceptance tests depend on ruby in path, if you are using the dev container, it is already installed

To run acceptance tests,

```sh
# generate code locally
make generate

# build acceptance test rig
make build-test-srv

# run acceptance tests
make accept
```

## Licensing

This software is available under the QuickFIX Software License. Please see the [LICENSE.txt](https://github.com/quickfixgo/quickfix/blob/main/LICENSE.txt) for the terms specified by the QuickFIX Software License.
