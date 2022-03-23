版本号`major.minor.patch`具体规则如下：
- major：主版本号，如有重大版本重构则该字段递增，通常各主版本间接口不兼容。
- minor：次版本号，各次版本号间接口保持兼容，如有接口新增或优化则该字段递增。
- patch：补丁号，如有功能改善或缺陷修复则该字段递增。

## version 0.6.4

**Feature**
- 修改imparse本地依赖为远程仓库 2021_12_07_17_44_27


## version 0.6.3

**配置文件更新**

所有 `[xxxRPCClient]` 增加 `RegAddrs = "127.0.0.1:2379"` 字段

**Feature**
- 重构 pusher @v0.6.0-pre 2021.10.14
- 更新 etcdv3.5.0 v0.6.2
- 修复新用户连接时没有群聊也强行加入群聊的 bug v0.6.3


## version 0.5

**Feature**
- 新增 3 中通知类型 @v0.5.2
- 转账交易消息格式支持 @v0.5.3 2021.8.16


## example x.x.x @yy.mm.dd

**Feature**

**Bug Fixes**

**Improvement**

**Breaking Change**
