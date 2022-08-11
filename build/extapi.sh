#!/usr/bin/env bash

cd (dirname (status -f))
cd ../api
protoc voxov.proto --go_out=.         --go_opt=paths=source_relative \
                   --go-grpc_out=.    --go-grpc_opt=paths=source_relative
