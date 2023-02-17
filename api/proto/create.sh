#!/bin/sh

#protoc --proto_path=$GOPATH/src:. \
#    --go_out=. --go_opt=paths=source_relative \
#    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
#    api.proto

#protoc -I=. -I=$GOPATH/src --go_out=plugins=grpc:. *.proto

#protoc --proto_path=.:"$GOPATH"/src --go_out=. --go_opt=paths=source_relative \
#    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
#    *.proto

PREFIX="github.com/txchat/dtalk/api/proto"
protoc --proto_path="${GOPATH}/src/" --go_out=. --go_opt=module=$PREFIX \
    "${GOPATH}/src/${PREFIX}"/*.proto
