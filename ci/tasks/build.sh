#!/bin/bash

set -e -x

cd ci-web
go get -u github.com/Masterminds/glide
glide install
go test -v $(glide novendor)
