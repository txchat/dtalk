版本号`major.minor.patch`具体规则如下：

- major：主版本号，如有重大版本重构则该字段递增，通常各主版本间接口不兼容。
- minor：次版本号，各次版本号间接口保持兼容，如有接口新增或优化则该字段递增。
- patch：补丁号，如有功能改善或缺陷修复则该字段递增。

## version 0.1.0 2022_01_20

**Feature**

- 增加 `group` 相关 api 接口

## version 0.0.8

**Feature**

- 更新gateway下服务修改signal和noticemsg名称 2021_12_07_17_52_18

## version 0.0.7

**配置文件新增**

所有 `[xxxRPCClient]` 增加 `RegAddrs = "127.0.0.1:2379"` 字段

**Feature**  
- 更新 etcdv3.5.0 v0.0.7 2021.11.03

## version 0.0

**Feature**
- 模块开启接口 @v0.0.2
- 撤回消息 @v0.0.4
- 改用 imparse 中的 proto @v0.0.5 2021.10.14
- 新增撤回消息时限配置文件 v0.0.6 2021.10.22


## example x.x.x @yy.mm.dd

**Feature**

**Bug Fixes**

**Improvement**

**Breaking Change**
