#!/usr/bin/env bash

cd $(dirname $0)
cd ..
mkdir -p bin && cd $_

go build ../../service/gate/gate.go
