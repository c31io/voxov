#!/usr/bin/env bash

cd $(dirname $0)
cd ../pkg/api

auth () {
    echo -n "Building auth... "
    cd auth
    protoc auth.proto --go_out=.         --go_opt=paths=source_relative \
                      --go-grpc_out=.    --go-grpc_opt=paths=source_relative
    cd ..
    echo "Done"
}

if [ $# -eq 0 ]; then
    echo "No arguments provided, building ALL!"
    auth
    echo "All internal gRPC modules generated."
    exit 0
fi
