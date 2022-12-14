#!/bin/sh

PREFIX="github.com/txchat/dtalk/proto"
protoc --proto_path="${GOPATH}/src/" --go_out=. --go_opt=module=$PREFIX \
    "${GOPATH}/src/${PREFIX}"/*.proto