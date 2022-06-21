#!/usr/bin/env fish

cd (dirname (status -f))

if not test -d build
    mkdir build
end

go build -o build sms-dev/sms-dev.go
