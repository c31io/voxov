#!/usr/bin/env fish

cd (dirname (status -f))

protoc sms.proto --go_out=.         --go_opt=paths=source_relative \
                 --go-grpc_out=.    --go-grpc_opt=paths=source_relative \
