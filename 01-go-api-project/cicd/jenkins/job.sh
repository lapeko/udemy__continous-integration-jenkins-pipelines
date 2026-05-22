#!/bin/bash

set -e
set -x

gofmt -d -e .
go vet ./...
go test -v --race ./...