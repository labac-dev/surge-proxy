#!/usr/bin/env bash

# Generate protobuf files from https://github.com/Maasea/sgmodule/tree/e5d66ffc39b71e499c6e9b24ef13d44598f2c86f/Script/Youtube/protobuf

rm -rf gen
mkdir -p gen

protoc \
    --go_out=gen \
    --go_opt=module=github.com/labac-dev/surge-proxy/internal/youtube/gen \
    --proto_path=. \
    protobuf/**/*.proto
