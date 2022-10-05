#!/usr/bin/env bash

cd $(dirname $0)
cd ../pkg/api

build () {
    protoc $1.proto --go_out=.         --go_opt=paths=source_relative \
                    --go-grpc_out=.    --go-grpc_opt=paths=source_relative
}

# Build by arguments
for i in "${@:1}"; do
    cd $i; build $(basename $i); cd ..
done

# Build all without arguments
if [ $# -eq 0 ]; then
    for i in $(ls -d */); do
        cd $i; build $(basename $i); cd ..
    done
fi
