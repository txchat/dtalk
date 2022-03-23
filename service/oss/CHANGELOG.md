版本号`major.minor.patch`具体规则如下：

- major：主版本号，如有重大版本重构则该字段递增，通常各主版本间接口不兼容。
- minor：次版本号，各次版本号间接口保持兼容，如有接口新增或优化则该字段递增。
- patch：补丁号，如有功能改善或缺陷修复则该字段递增。

## version 1.5.1 2022_01_06

**Feature**

- 增加对 key 的验证
- 上传设置重试次数

## version 1.5

**Feature**

- minio 配置文件增加 publicUrl v1.5.0 2021_11_10

## version 1.4

**Feature**

- 上传成功接口同时返回 url 和 uri
- 增加 get-host 接口 @v1.4.0 2021.8.30

## version 1.3

**Feature**

- 支持 minio @v1.3.0 2021.7.28
- 简化接口参数, 限制单次上传最大5g, 分片单次上传最小5mb @v1.3.1 2021.8.2

**Bug fix**
- 分片上传取消最小文件大小限制 @v1.3.2 2021.8.17
- 修正取消分片上传接口文档错误 @v1.3.3 2021.8.23
- minio.getUrl 专门为 zbotc 只返回 uri, 而不是完整 url @v1.3.3.1 2021.8.26


## version 1.2.0 @2021.7.22

**Feature**

- 支持代理文件分段上传 @1.2.0 2021.7.22
- 更新文档 @v1.2.1
- 移动 aliyun 和 huaweiyun 至 pkg 文件夹中 @v1.2.2 2021.7.26

## version 1.1.2 @2021.7.15

**Feature**

- 通过 http 接口转发文件上传
- 支持阿里云普通上传 @2021.7.6
- 更新日志模块为 zerolog @2021.7.7
- 支持华为云普通上传 @2021.7.7
- 所有配置信息从配置文件中获取 @2021.7.8
- upload接口 OssType 字段可选填, 由配置文件中同APPID的最后一个Oss.OssType决定 @1.1.2 2021.7.15

## version 1.0.0 @2021.6.30

**Feature**

支持阿里云和华为云 oss 临时鉴权

## example x.x.x @yy.mm.dd

**Feature**

**Bug Fixes**

**Improvement**

**Breaking Change**
