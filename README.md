# FIR Go Client
[![Go Report Card](https://goreportcard.com/badge/github.com/byt3smith/fir-go)](https://goreportcard.com/report/github.com/byt3smith/fir-go)
[![Build Status](https://travis-ci.org/byt3smith/fir-go.svg?branch=master)](https://travis-ci.org/byt3smith/fir-go)
[![GoDoc](https://godoc.org/github.com/byt3smith/fir-go?status.svg)](https://godoc.org/github.com/byt3smith/fir-go)

FirGo is a Go client library for accessing the Fast Incident Response API.

### Installation
`go get github.com/byt3smith/fir-go`

### Usage
```
import "github.com/byt3smith/fir-go"
```

##### Examples live in the `cmd/` directory

To execute an example, you can:
- use `go run` and execute `main.go`
- run the command as it's directory name (e.g. `fir-get`) if you have $GOBIN set in your path

*Note* The examples rely on environment variables `FIR_APIKEY` and `FIR_BASE_URL`
