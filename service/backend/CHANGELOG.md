版本号`major.minor.patch`具体规则如下：
- major：主版本号，如有重大版本重构则该字段递增，通常各主版本间接口不兼容。
- minor：次版本号，各次版本号间接口保持兼容，如有接口新增或优化则该字段递增。
- patch：补丁号，如有功能改善或缺陷修复则该字段递增。

## version 0.2.3

**配置文件更新**
```toml
CdkMod = false # true 表示开启 cdk 模块, false 表示关闭
```
**Feature**
- cdk 模块设计开关 v0.2.3 2021_11_09

## version 0.2.0-0.2.2

**配置文件更新**
```toml
Env = "release"

[IdGenRPCClient]
RegAddrs = "127.0.0.1:2379"
Schema = "dtalk"
SrvName = "generator"
Dial = "1s"
Timeout = "1s"
```

**数据库更新**
新增两张表
- dtalk_cdk_info
- dtalk_cdk_list

**Feature**
- 新增 cdk 相关接口 v0.2.0 2021_10_22
- 新增接口参数约束, 每种优惠券兑换数量上限 v0.2.1 2021_10_31
- 异步处理订单, 后端查询交易是否成功 v0.2.2 2021_11_01

## version 0.1.0

- platform 从配置文件中获得 v0.1.0 2021_09_18

## version 0.0.10

init backend


## example x.x.x @yy.mm.dd

**Feature**

**Bug Fixes**

**Improvement**

**Breaking Change**
