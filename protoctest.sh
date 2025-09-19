#!/bin/bash

protoc --proto_path=./idl \
              --go_out=paths=source_relative:./ \
              --proto_path=./third_party \
              --openapi_out=fq_schema_naming=true,default_response=false:. \
              --go-aggenall_out=paths=source_relative:./api \
              idl/api/**/*.proto 2>&1 | tee xxxx.log
              #idl/api/**/*.proto 2>&1 | tee xxxx.log
#idl/api/hzw/hzw.proto 2>&1 | tee xxxx.log
