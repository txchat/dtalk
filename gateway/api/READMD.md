# Gateway

业务网关层往往又被称作业务聚合层，该层由若干个业务服务构成，负责接收通用网关转发过来的流量。核心作用如下：

### 1. 路由管理

业务网关层的服务负责自己服务的路由表维护。

### 2. 参数校验

业务网关层的服务负责执行与客户端约定的参数校验，验证通过之后再组装成后端微服务需要的数据结构请求到后端。

### 3. 权限校验

各个业务网关层的服务通过底层的用户服务调用来实现权限校验，对于哪些路由需要权限校验，哪些路由不需要，完全由业务网关层的服务自行维护。

### 4. 接口聚合

业务网关层的服务可能需要调用多个后端的微服务来组合实现一个接口，根据自身需求来对下层返回的数据进行聚合和处理。

### 5. 协议转换

业务网关层的服务接收转发过来的HTTP请求，并转换为内部的一个或者多个GRPC微服务调用来实现接口逻辑。

### 6. 数据转换

业务网关层的服务的输入和输出数据结构必须是表示层需要的，因此它所负责的数据结构和后端GRPC微服务的数据结构会不一样。业务网关层的服务需要负责数据结构的转换和封装处理。

> gateway 里不应该有复杂业务逻辑
>

文档路径 : http://172.16.101.107:8888/swagger/index.html