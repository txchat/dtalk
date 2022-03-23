# 后台服务

####  检查更新
URL: /app/version/check

`post`

**请求参数：**

| **参数**       | **名字**     | **类型** | **约束** | **说明** |
| -------------- | ------------ | -------- | -------- | -------- |
| versionCode | 当前版本 | int   | true   |    版本code    |

**返回参数：**

| **参数**    | **名字**     | **类型** | **说明**   |
| ----------- | ------------ | -------- | ---------- |
| id | 最新的版本 | int   | 版本编号             |
| platform | 平台 | string | chat33pro               |
| status | 线上状态 | int | 0：历史；1：线上版本 |
| deviceType | 终端类型 | string | Android/IOS          |
| versionName | 版本名 | string | 3.6.8.10             |
| versionCode | 版本code | int | 36810                |
| url | 下载地址 | string |                      |
| force | 是否强制更新 | bool | false：非强制；true：强制 |
| description | 描述信息 | string array |            |
| opeUser | 操作者 | string | |
| updateTime | 更新时间 | int      |            |
| createTime | 创建时间       | int |            |
| size  | 包大小     | int          |             单位：byte              |
| md5  |  包的md5    | string          |                           |


```json
{
  "result": 0,
  "message": "",
  "data": {
      "id": 8,
      "platform": "Chat33Pro",
      "status": 1,
      "deviceType": "IOS",
      "versionName": "1.0.4",
      "versionCode": 10400,
      "url": "https://xxx",
      "force": true,
      "description": [
        "qqq",
        "ww"
      ],
      "opeUser": "root",
      "md5": "12345",
      "size": 123,
      "updateTime": 1621408163504,
      "createTime": 1621396088880
  }
}
```

#### 创建版本

URL: /backend/version/create

`post`

**请求参数：**

| 参数        | 名字         | 类型         | 约束 | 说明                      |
| ----------- | ------------ | ------------ | ---- | ------------------------- |
| platform    | 平台         | string       | true | chat33pro                    |
| description | 描述信息     | string array | true |                           |
| force       | 是否强制更新 | bool         | true | false：非强制；true：强制 |
| url         | 下载地址     | string       | true |                           |
| versionCode | 版本code     | int       | true | 36810                     |
| versionName | 版本名       | string       | true | 3.6.8.10                  |
| deviceType  | 终端类型     | string       | true | Android/IOS               |
| Authorization | 授权 | string | true | 用于传递token |
| size  | 包大小     | int          |             单位：byte              |
| md5  |  包的md5    | string          |                           |

**返回参数：**

| 参数        | 名字         | 类型         | 说明                      |
| ----------- | ------------ | ------------ | ------------------------- |
| version     | 版本信息     | object       | 创建的版本的全部信息      |
| id          | 版本编号     | int          |                           |
| platform    | 平台         | string       | chat33pro                    |
| status      | 线上状态     | int          | 0：历史；1：线上版本      |
| deviceType  | 终端类型     | string       | Android/IOS               |
| versionName | 版本名       | string       | 3.6.8.10                  |
| versionCode | 版本code     | int       | 36810                     |
| url         | 下载地址     | string       |                           |
| force       | 是否强制更新 | bool         | false：非强制；true：强制 |
| description | 描述信息     | string array |                           |
| opeUser     | 操作者       | string       |                           |
| updateTime  | 更新时间     | int          |                           |
| createTime  | 创建时间     | int          |                           |
| size  | 包大小     | int          |             单位：byte              |
| md5  |  包的md5    | string          |                           |

```json
{
    "result": 0,
    "message": "",
    "data": {
        "version": {
            "id": 3,
            "platform": "Chat33Pro",
            "status": 0,
            "deviceType": "Android",
            "versionName": "1.0.1",
            "versionCode": 10000,
            "url": "https://xxx",
            "force": true,
            "description": [
                "qqq",
                "ww"
            ],
            "opeUser": "root",
            "md5": "12345",
            "size": 123,
            "updateTime": 1621394358387,
            "createTime": 1621394358387
        }
    }
}
```



#### 更新版本

URL: /backend/version/update

`put`

**请求参数：**

| 参数        | 名字         | 类型         | 约束 | 说明                      |
| ----------- | ------------ | ------------ | ---- | ------------------------- |
| description | 描述信息     | string array | true |                           |
| force       | 是否强制更新 | bool         | true | false：非强制；true：强制 |
| url         | 下载地址     | string       | true |                           |
| versionCode | 版本code     | int       | true | 36810                     |
| versionName | 版本名       | string       | true | 3.6.8.10                  |
| id          | 版本编号     | int          | true |                           |
| Authorization | 授权 | string | true | 用于传递token                           |
| size  | 包大小     | int          |             单位：byte              |
| md5  |  包的md5    | string          |                           |

**返回参数：**

| 参数        | 名字         | 类型         | 说明                      |
| ----------- | ------------ | ------------ | ------------------------- |
| version     | 版本信息     | object       | 修改后的版本的全部信息    |
| id          | 版本编号     | int          |                           |
| platform    | 平台         | string       | chat33pro                    |
| status      | 线上状态     | int          | 0：历史；1：线上版本      |
| deviceType  | 终端类型     | string       | Android/IOS               |
| versionName | 版本名       | string       | 3.6.8.10                  |
| versionCode | 版本code     | int       | 36810                     |
| url         | 下载地址     | string       |                           |
| force       | 是否强制更新 | bool         | false：非强制；true：强制 |
| description | 描述信息     | string array |                           |
| opeUser     | 操作者       | string       |                           |
| updateTime  | 更新时间     | int          |                           |
| createTime  | 创建时间     | int          |                           |
| size  | 包大小     | int          |             单位：byte              |
| md5  |  包的md5    | string          |                           |

```json
{
  "result": 0,
  "message": "",
  "data": {
    "version": {
      "id": 3,
      "platform": "Chat33Pro",
      "status": 0,
      "deviceType": "Android",
      "versionName": "1.0.0",
      "versionCode": 10000,
      "url": "https://xxx",
      "force": false,
      "description": [
        "xx",
        "yy"
      ],
      "opeUser": "root",
      "md5": "234567",
      "size": 2345,
      "updateTime": 1621395843469,
      "createTime": 1621394358387
    }
  }
}
```

#### 修改版本线上状态

URL: /backend/version/change-status

`put`

**请求参数：**

| 参数   | 名字     | 类型   | 约束 | 说明                 |
| ------ | -------- | ------ | ---- | -------------------- |
| id     | 版本编号 | int    | true | 要修改的版本编号     |
| Authorization | 授权 | string | true | 用于传递token                      |

**返回参数：**

无

```json
{
  "result": 0,
  "message": "",
  "data": null
}
```

#### 获取全部版本信息

URL: /backend/version/list

`get`

**请求参数：**

| 参数       | 名字     | 类型   | 约束  | 说明         |
| ---------- | -------- | ------ | ----- | ------------ |
| page       | 页码     | int    | false  | 从0开始，不填默认是0      |
| platform   | 平台     | string | false | 要筛选的平台 |
| deviceType | 终端类型 | string | false | 要筛选的终端 |
| Authorization | 授权 | string | true | 用于传递token  |

**返回参数：**

| 参数          | 名字           | 类型         | 说明                       |
| ------------- | -------------- | ------------ | -------------------------- |
| totalElements | 所有的记录条数 | int          |                            |
| totalPages    | 总页数         | int          |                            |
| versionList   | 版本列表       | object array | 所查询的全部版本的全部信息 |
| id            | 版本编号       | int          |                            |
| platform      | 平台           | string       | chat33pro                     |
| status        | 线上状态       | int          | 0：历史；1：线上版本       |
| deviceType    | 终端类型       | string       | Android/IOS                |
| versionName   | 版本名         | string       | 3.6.8.10                   |
| versionCode   | 版本code       | int       | 36810                      |
| url           | 下载地址       | string       |                            |
| force         | 是否强制更新   | bool         | false：非强制；true：强制  |
| description   | 描述信息       | string array |                            |
| opeUser       | 操作者         | string       |                            |
| updateTime    | 更新时间       | int          |                            |
| createTime    | 创建时间       | int          |                            |
| size  | 包大小     | int          |             单位：byte              |
| md5  |  包的md5    | string          |                           |

```json
{
  "result": 0,
  "message": "",
  "data": {
    "totalElements": 5,
    "totalPages": 1,
    "versionList": [
      {
        "id": 12,
        "platform": "Chat33Pro",
        "status": 0,
        "deviceType": "IOS",
        "versionName": "1.0.0",
        "versionCode": 10000,
        "url": "https://xxx",
        "force": false,
        "description": [
          "qqq",
          "ww"
        ],
        "opeUser": "root",
        "md5": "12345",
        "size": 123,
        "updateTime": 1621396134321,
        "createTime": 1621396134321
      },
      {
        "id": 11,
        "platform": "Chat33Pro",
        "status": 1,
        "deviceType": "IOS",
        "versionName": "1.0.1",
        "versionCode": 10100,
        "url": "https://xxx",
        "force": true,
        "description": [
          "qqq",
          "ww"
        ],
        "opeUser": "root",
        "md5": "12345",
        "size": 123,
        "updateTime": 1621396320416,
        "createTime": 1621396123534
      },
      {
        "id": 10,
        "platform": "Chat33Pro",
        "status": 0,
        "deviceType": "IOS",
        "versionName": "1.0.2",
        "versionCode": 10200,
        "url": "https://xxx",
        "force": true,
        "description": [
          "qqq",
          "ww"
        ],
        "opeUser": "root",
        "md5": "12345",
        "size": 123,
        "updateTime": 1621396117028,
        "createTime": 1621396117028
      },
      {
        "id": 9,
        "platform": "Chat33Pro",
        "status": 0,
        "deviceType": "IOS",
        "versionName": "1.0.3",
        "versionCode": 10300,
        "url": "https://xxx",
        "force": false,
        "description": [
          "qqq",
          "ww"
        ],
        "opeUser": "root",
        "md5": "12345",
        "size": 123,
        "updateTime": 1621396108579,
        "createTime": 1621396108579
      },
      {
        "id": 8,
        "platform": "Chat33Pro",
        "status": 0,
        "deviceType": "IOS",
        "versionName": "1.0.4",
        "versionCode": 10400,
        "url": "https://xxx",
        "force": true,
        "description": [
          "qqq",
          "ww"
        ],
        "opeUser": "root",
        "md5": "12345",
        "size": 123,
        "updateTime": 1621396320416,
        "createTime": 1621396088880
      }
    ]
  }
}
```

#### 获取token

URL: /backend/user/login

`get`

**请求参数：**

| 参数     | 名字   | 类型   | 约束 | 说明         |
| -------- | ------ | ------ | ---- | ------------ |
| userName | 用户名 | string | true | 暂时为“root” |
| password | 口令   | string | true | 暂时为“root” |

**返回参数：**

| 参数  | 名字  | 类型   | 说明 |
| ----- | ----- | ------ | ---- |
| userInfo | 用户信息 | object |包含用户名和token      |
| userName | 用户名 | string |      |
| token | token | string |      |

```json
{
  "result": 0,
  "message": "",
  "data": {
    "userInfo": {
      "userName": "root",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InJvb3QiLCJleHAiOjE2MjA5ODcyMDEsImlzcyI6IkJvYiJ9.w_NoSezjjJLRJMjiU4jiMYozdYvL6NPwv2xuCMepws4"
    }
  }
}
```





