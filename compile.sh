#!/bin/sh

set -e

mkdir -p bin/

go build -o bin/minjsinstall minjsinstall.go

exit 0
