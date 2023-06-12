# 使用goctl生成代码框架

## 生成API网关
1. 生成空的API网关工程目录
```shell
goctl api new --home ../../.goctl center
```

2. 编辑.api文件定义网关接口
3. 从.api文件生成相应代码
```shell
goctl api go --home ../../../.goctl -api center.api -dir .
```

## 生成RPC服务
1. 创建空的rpc工程目录
```shell
goctl rpc new group --home ../../.goctl
```

2. 编辑.proto文件定义rpc接口
3. 从.proto文件生成相应代码
```shell
# 普通rpc
goctl rpc protoc group.proto --go_out=. --go-grpc_out=. --zrpc_out=. --home ../../../.goctl

# 引用了imparse的proto
goctl rpc protoc "$GOPATH"/src/github.com/txchat/dtalk/app/services/answer/answer.proto --proto_path="$GOPATH"/src/ --go_out=. --go-grpc_out=. --zrpc_out=.

goctl rpc protoc transfer.proto --proto_path=.:"$GOPATH"/src --go_out=. --go-grpc_out=. --zrpc_out=.
```