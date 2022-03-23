版本号`major.minor.patch`具体规则如下：

- major：主版本号，如有重大版本重构则该字段递增，通常各主版本间接口不兼容。
- minor：次版本号，各次版本号间接口保持兼容，如有接口新增或优化则该字段递增。
- patch：补丁号，如有功能改善或缺陷修复则该字段递增。

## version 2.3.0 2022_01_20

**Feature**

- 增加一堆 rpc 方法
- 关闭 http 服务

## version 2.2.0 2022_01_06

**Feature**

- [x] 批量加群 rpc 接口 v2.2.1
- [x] 批量退群 rpc 接口 v2.2.2
- [x] 交换群主 rpc 接口 v2.2.2
- 修改创群 RPC 接口 v2.2.3

## version 2.1.0 2021_12_23

**数据库变动**

- dtalk_group_info 增加 group_type 字段 (0: 普通群, 1: 全员群, 2: 部门群)

## version 2.0.6

**配置文件新增**

```toml
[Redis]
network      = "tcp"
addr         = "127.0.0.1:6379"
auth         = ""
active       = 60000
idle         = 1024
dialTimeout  = "200ms"
readTimeout  = "500ms"
writeTimeout = "500ms"
idleTimeout  = "120s"
expire       = "30m"
```

**Feature**
- 实现退群 RPC 方法 2021_11_26_17_58_11
- 使用日志中间件记录请求, 日志设置 TraceId
- rpc log 增加 trace id 2021_12_06
- 所有方法增加 ctx
- 增加解散群 RPC 方法
- group 增加 cache 6 2021_12_07_15_37_50

## version 2.0.5 2021_11_05-

**数据库更新**
- group_name 修改长度
- dtalk_group_info 新增 group_aes_key 字段
- dtalk_group_info 新增 group_pub_name 字段

**Feature**
- 增加 aes 字段 v2.0.1 2021_11_05
- 增加 group_pub_name 字段 v2.0.2 2021_11_17
- 增加查询群信息 rpc 接口 v2.0.3 2021_11_18
- 增加 maintain 数据库名称 v2.0.4 2021_11_19
- 修复更新群名称重复发送通知的特性 v2.0.5 2021_11_23

## version 2.0 2021_11_03

**配置文件更新**  

group.toml  
所有 `[xxxRPCClient]` 增加 `RegAddrs = "127.0.0.1:2379"` 字段  

- 重构 group 代码 v2.0.0 2021_11_03


## version 1.6

**Feature**
- 增加加群申请接口 v1.6.0 2021.10.9
- 增加查询群成员 rpc 接口 v1.6.1 2021.10.10
- 改用 imparse 中的 proto @v1.6.2 2021.10.14

## version 1.5

**Feature**
- 增加加群 rpc 接口 v1.5.0 2021.9.24
- 增加搜索群列表接口 v1.5.4 2021.9.26

**Bug fix**
- 修复签名过期 bug v1.5.1 2021.9.26
- 生成不重复的 groupMarkId v1.5.2 2021.9.26
- 优化 groupMarkId 生成方法 v1.5.3 2021.9.26
- 修复批量邀请群成员全失败的逻辑 v1.5.5 2021.9.27 
- 修复创建群聊时群成员包含群主的产生的 bug v1.5.6 2021.9.27
- 修复成员主动退群后还能收到退群的通知 v1.5.7 2021.9.28

## version 1.4

**Feature**
- 增加查询群列表 rpc 接口 @v1.4.0 2021.9.2
- 增加直接加群的接口 @v1.4.1 2021.9.6
- 查询群公开信息接口返回用户在群里的信息 v1.4.2 2021.9.9
- 扫码加群走邀请加群通道 v1.4.3 2021.9.9

## version 1.3

**Feature**
- 增加创群 rpc 接口 @v1.3.0 2021.8.23
- 接口可以选择只传整形或者字符串类型 ID @v1.3.1 2021.8.26

## version 1.2

**Feature**
- 接口增加 string 类型 Id @v1.2.0 2021.8.20

## version 1.1 @2021.7.26

**Feature**

- 群人数上限和管理员数量上限可以通过配置文件配置 @v1.1.1 2021.7.7
- notice 从 answer 服务接入 @v1.1.2 2021.7.15
- change alert.from from groupId to personId @v1.1.3 2021.7.26
  
**Bug fix**
- 群主不能把自己设置为管理员 @v1.1.4 2021.8.17

## version 1.1.0 @2021.7.1

**Feature**

增加 Prometheus 的指标监控
- http_requests_total
- http_response_time_seconds

## version 1.0.0 @2021.6.30

group init, 实现第一批需求接口

## example x.x.x @yy.mm.dd

**Feature**

**Bug Fixes**

**Improvement**

**Breaking Change**
