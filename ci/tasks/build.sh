#!/bin/bash

set -e -x

export GOPATH=$PWD
export PATH=$GOPATH/bin:$PATH
cd ./src/github.com/dnem/ci-web
go get -u github.com/Masterminds/glide
glide install

go test -v $(glide nv)
go build -o $GOPATH/binary/ci-web
