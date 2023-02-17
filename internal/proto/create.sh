#!/bin/sh

PREFIX="github.com/txchat/dtalk/internal/proto"
protoc --proto_path="${GOPATH}/src/" --go_out=. --go_opt=module=$PREFIX \
    "${GOPATH}/src/${PREFIX}"/*.proto
