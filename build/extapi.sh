#!/usr/bin/env bash

cd $(dirname $0)
cd ../api
protoc voxov.proto --go_out=.         --go_opt=paths=source_relative \
                   --go-grpc_out=.    --go-grpc_opt=paths=source_relative
