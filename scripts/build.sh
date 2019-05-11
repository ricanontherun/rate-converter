#! /bin/bash
EXECUTABLE=rate-converter
CLI_MAIN=cmd/rate-converter/main.go

go build -o $EXECUTABLE $CLI_MAIN