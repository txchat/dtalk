版本号`major.minor.patch`具体规则如下：
- major：主版本号，如有重大版本重构则该字段递增，通常各主版本间接口不兼容。
- minor：次版本号，各次版本号间接口保持兼容，如有接口新增或优化则该字段递增。
- patch：补丁号，如有功能改善或缺陷修复则该字段递增。

## version 0.1.5

**配置文件更新**
- 所有 `[xxxRPCClient]` 增加 `RegAddrs = "127.0.0.1:2379"` 字段  2021_11_12

## version 0.1

**Feature**
- 接口增加 string 类型 Id @v0.1.0 2021.8.20
- 接口可以选择只传整形或者字符串类型 ID @v0.1.2 2021.8.26
- 改用 imparse 中的 proto @v0.1.3 2021.10.14
- 更新 etcdv3.5.0 v0.1.4

**Bug Fix**
- string 类型转化为整形时增加error判断 @v0.1.1 2021.8.20

## version 0.0.8 @2021.7.19

**Feature**

- 从服务端获取 userSign, privateMapKey @0.0.6
- 从服务端获取 appid @0.0.7
- roomId支持简单分布式 @0.0.8

## version 0.0.5 @2021.7.15

**Feature**

- 通过 start-call, reply-busy, check-call, stop-call 四个接口实现私聊双方进入房间


## example x.x.x @yy.mm.dd

**Feature**

**Bug Fixes**

**Improvement**

**Breaking Change**
