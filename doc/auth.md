# 鉴权服务

#### 应用注册

URL: /auth/sign-in

`post`

**请求参数：**

| **参数**       | **名字**     | **类型** | **约束** | **说明** |
| -------------- | ------------ | -------- | -------- | -------- |
| Content-Type | body中的数据类型 | string | true | 值为multipart/form-data |
| configFile | 配置文件 | file | true | 包含第三方应用鉴权服务接口的相关信息 |
| appId | 应用ID | text | true   |    应用ID，位于请求头    |


**返回参数：**

| **参数**    | **名字**     | **类型** | **说明**   |
| ----------- | ------------ | -------- | ---------- |
| appId | appId | string | appId |
| key | key | string | 授权第三方应用使用本鉴权服务 |
| createTime | 创建时间 | int | app注册时间 |
| updateTime | 更新时间 | int | 信息更新时间 |

```json
{
  "result": 0,
  "message": "",
  "data": {
    "appId": "dtalk",
    "key": "key",
    "createTime": 1625552582748,
    "updateTime": 1625552582748
  }
}
```



#### 鉴权

**请求参数：**

| **参数**   | **名字**       | **类型** | **约束** | **说明**                                                     |
| ---------- | -------------- | -------- | -------- | ------------------------------------------------------------ |
| appId      | 应用的唯一标识 | string   | true     | 应用的唯一标识                                               |
| token      | token          | string   | true     | 第三方鉴权所要的token                                        |
| digest     | 消息摘要       | string   | true     | digest = SHA256 ( appId + token + createTime + key )（key为注册时所返回的key） |
| createTime | 请求创建时间   | int      | true     | 请求创建时的时间戳                                                 |

**返回参数：**

| **参数** | **名字** | **类型** | **说明** |
| -------- | -------- | -------- | -------- |
| uid      | 用户ID   | string   | 用户ID   |
