#!/bin/bash

ROOTDIR=$(cd $(dirname $0);pwd)

cd $ROOTDIR

for dir in $(find ./servers -type d -depth 1)
do
    echo "Start build $dir"
    cd $dir
    # cp ../*.go ./
    # go mod init
    go get -u
    go mod tidy
    CGO_ENABLED=0 go build -o gowebbenchmark -ldflags='-s -w' .
    cd -
done
