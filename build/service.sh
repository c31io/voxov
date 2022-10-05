#!/usr/bin/env bash

cd $(dirname $0)
mkdir -p bin && cd $_

# Build by arguments
for i in "${@:1}"; do
    go build ../../service/$i/$i.go
done

# Build all without arguments
if [ $# -eq 0 ]; then
    for i in $(ls -d ../../service/*/); do
        go build $i$(basename $i).go;
    done
fi
