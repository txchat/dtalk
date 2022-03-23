版本号`major.minor.patch`具体规则如下：
- major：主版本号，如有重大版本重构则该字段递增，通常各主版本间接口不兼容。
- minor：次版本号，各次版本号间接口保持兼容，如有接口新增或优化则该字段递增。
- patch：补丁号，如有功能改善或缺陷修复则该字段递增。

## version 1.1.1 2021_12_23

**数据库**
- 新增 `dtalk_addr_move` 表

**Feature**
- 增加登记地址接口

## version 1.1.0 2021_12_09
**Feature**
- 增加 v2 接口
- backup服务手机邮箱验证分为三种模式bind、quick、export

## version 1.0.15 2021.10.22

**配置文件新增**
```toml
[[Whitelist]]
Account=""
Code=""
Enable=true
```

**Feature**
- 增加短信邮箱验证白名单 @v1.0.15 2021.10.22
- 更新 etcdv3.5.0 v1.0.18 2021.11.03

## version 1.0.9 @2021.8.3

**Feature**

- 支持通过手机或邮箱获得地址
- 优化验证码发送错误提示 @v1.0.9 2021.8.3
- 修复 grpc nil panic v1.0.12




## example x.x.x @yy.mm.dd

**Feature**

**Bug Fixes**

**Improvement**

**Breaking Change**
