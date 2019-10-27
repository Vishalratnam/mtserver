# Simple Concurrent Web Server

Implementing a simple cli application which runs a web server built from scratch in Go.

## Installation

* Download the package using `go get github.com/pmk21/mtserver`.
* This installs the package in `GOPATH/bin`.

## Usage

Assuming `GOPATH/bin` is in your `PATH` environment variable. It can be run -

```terminal
$ mtserver --help
Implementation of a concurrent(multithreaded) and non-concurrent web server

Usage:
  mtserver [command]

Available Commands:
  help        Help about any command
  run         Start a webserver

Flags:
  -h, --help   help for mtserver

Use "mtserver [command] --help" for more information about a command.
```

## Commands

* `mtserver run ncs -p <port_num>` - Runs a basic non-concurrent server on given port number(default is 9000) which returns current time.
