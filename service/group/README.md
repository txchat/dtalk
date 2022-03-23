# group 群组服务

## TODO

- [ ] 删除service 中所有带 Svc 后缀的方法
- [ ] dao 层暴露的方法把 biz 对象作为方法参数, 而不是 db

## 目录结构

```
group
├── api
├── build
├── cmd
├── config
├── dao
├── docs
├── logic
│        └── http
├── model
│        ├── biz
│        ├── db
│        └── types
├── server
│        ├── grpc
│        └── http
└── service
```

### api

包含 grpc 的 `proto` , `*.pb.go`, `client` 文件  
client 提供 rpc 客户端的创建方法, 相当于手动给 grpc 客户端包了一层

### build

make build 指令生成的二进制文件保存目录

### cmd

项目启动目录  
包括资源的初始化, http 和 rpc 服务的启动

### config

配置文件描述目录

### dao

数据访问层   
注入到 service 中, 暴露数据操作方法, 不区分是操作数据库还是操作缓存(比如获得群信息, service 不关心是从数据库中拿还是从缓存中拿).  
如有需要,可以与业务逻辑服务分开,单独暴露 rpc 方法, 提供给多个业务服务,类似:
![](https://pic3.zhimg.com/80/v2-bf7456dbea800d3b7bc209ede853b7a6_720w.jpg)

### docs

http api swag 描述文件, 由 `make swag` 生成  
可以删去, 以后统一由 gateway 服务管理

### logic

使用`service`中的各种方法编排业务逻辑, 包含参数校验, 权限判断

### model

包含各种用到的结构体与 const 值

- `biz` service 中一直使用的结构体
- `db` 数据库字段映射(可以移动到 dao 层)
- `types` http服务的请求与返回(可以不要,因为 http 请求都移动到 gateway 了)

### server

http 服务和 grpc 服务的启动路径, 仅绑定参数, 不做任何处理  
服务中的各个方法与 logic 中的方法一一对应

### service

各种业务方法, 假定传入的参数都是正确,经过验证的, 不进行权限判断
