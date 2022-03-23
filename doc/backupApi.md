# 找回服务
### 手机绑定查询
URL: /backup/phone-query

`get`|`post`

**请求参数：**

| 参数名 | 必选 | 类型   | 说明   |
| ------ | ---- | ------ | ------ |
| area   | true | string | 区号   |
| phone  | true | string | 手机号 |

**返回参数：**

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
        "exists":false
    }
}
```

### 邮箱绑定查询
URL: /backup/email-query

`get`|`post`

**请求参数：**

| 参数名 | 必选 | 类型   | 说明 |
| ------ | ---- | ------ | ---- |
| email  | true | string | 邮箱 |

**返回参数：**

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
        "exists":false
    }
}
```

### 发送手机验证码
URL: /backup/v2/phone-send

`post`

**请求参数：**

| 参数名 | 必选  | 类型   | 说明   |
| ------ | ----- | ------ | ------ |
| area   | false | string | 区号   |
| phone  | true  | string | 手机号 |
| codeType | true  | string | 验证码类型：quick->快速登录；bind->绑定；export->导出通讯录 |

**返回参数：**

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```

### 发送邮箱验证码

URL: /backup/v2/email-send

`post`

**请求参数：**

| 参数名 | 必选 | 类型   | 说明 |
| ------ | ---- | ------ | ---- |
| email  | true | string | 邮箱 |
| codeType | true  | string | 验证码类型：quick->快速登录；bind->绑定；export->导出通讯录 |

**返回参数：**

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```

### 绑定手机

URL: /backup/v2/phone-binding

`post`|`auth`

type: bind

**请求参数：**

| 参数名   | 必选  | 类型   | 说明       |
| -------- | ----- | ------ | ---------- |
| area     | false | string | 区号       |
| phone    | true  | string | 手机号     |
| code     | true  | string | 验证码     |
| mnemonic | true  | string | 加密助记词 |

**返回参数：**

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```

### 绑定邮箱

URL: /backup/v2/email-binding

`post`|`auth`

type: bind

**请求参数：**

| 参数名   | 必选 | 类型   | 说明       |
| -------- | ---- | ------ | ---------- |
| email    | true | string | 手机号     |
| code     | true | string | 验证码     |
| mnemonic | true | string | 加密助记词 |

**返回参数：**

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```

### 手机找回私钥

URL: /backup/v2/phone-retrieve

`post`

type: quick

**请求参数：**

| 参数名 | 必选 | 类型   | 说明   |
| ------ | ---- | ------ | ------ |
| area   | false | string | 区号   |
| phone  | true | string | 手机号 |
| code   | true | string | 验证码 |

**返回参数：**

| 参数名      | 必选 | 类型   | 说明       |
| ----------- | ---- | ------ | ---------- |
| address     | true | string | 用户地址   |
| area        | true | string | 区号       |
| phone       | true | string | 手机号     |
| email       | true | string | 邮箱       |
| mnemonic    | true | string | 加密助记词 |
| update_time | true | string | 更新时间   |

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
        "address": "0xffffffff",
        "area":"",
        "phone":"",
        "email":"",
        "mnemonic":"",
        "update_time":""
    }
}
```
### 邮箱找回私钥

URL: /backup/v2/email-retrieve

`post`

type: quick

**请求参数：**

| 参数名 | 必选 | 类型   | 说明   |
| ------ | ---- | ------ | ------ |
| email  | true | string | 邮箱   |
| code   | true | string | 验证码 |

**返回参数：**

| 参数名      | 必选 | 类型   | 说明       |
| ----------- | ---- | ------ | ---------- |
| address     | true | string | 用户地址   |
| area        | true | string | 区号       |
| phone       | true | string | 手机号     |
| email       | true | string | 邮箱       |
| mnemonic    | true | string | 加密助记词 |
| update_time | true | string | 更新时间   |

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
        "address": "0xffffffff",
        "area":"",
        "phone":"",
        "email":"",
        "mnemonic":"",
        "update_time":""
    }
}
```

### 手机导出通讯录验证

URL: /backup/v2/phone-export

`post`

type: export

**请求参数：**

| 参数名 | 必选 | 类型   | 说明   |
| ------ | ---- | ------ | ------ |
| area   | false | string | 区号   |
| phone  | true | string | 手机号 |
| code   | true | string | 验证码 |
| address   | true | string | 地址 |

**返回参数：**

ERR:-4010 账号与绑定手机不一致
```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```

### 邮箱导出通讯录验证

URL: /backup/v2/email-export

`post`

type: export

**请求参数：**

| 参数名 | 必选 | 类型   | 说明   |
| ------ | ---- | ------ | ------ |
| email  | true | string | 邮箱   |
| code   | true | string | 验证码 |
| address   | true | string | 地址 |

**返回参数：**

ERR:-4011 账号与绑定邮箱不一致

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```

### 修改助记词（修改密聊密码）

URL: /backup/edit-mnemonic

`post`|`auth`

**请求参数：**

| 参数名   | 必选 | 类型   | 说明       |
| -------- | ---- | ------ | ---------- |
| mnemonic | true | string | 加密助记词 |

**返回参数：**

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```


### 发送手机验证码
URL: /backup/phone-send（弃用）

`post`

**请求参数：**

| 参数名 | 必选  | 类型   | 说明   |
| ------ | ----- | ------ | ------ |
| area   | false | string | 区号   |
| phone  | true  | string | 手机号 |

**返回参数：**

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```

### 发送邮箱验证码

URL: /backup/email-send（弃用）

`post`

**请求参数：**

| 参数名 | 必选 | 类型   | 说明 |
| ------ | ---- | ------ | ---- |
| email  | true | string | 邮箱 |

**返回参数：**

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```

### 绑定手机

URL: /backup/phone-binding（弃用）

`post`|`auth`

**请求参数：**

| 参数名   | 必选  | 类型   | 说明       |
| -------- | ----- | ------ | ---------- |
| area     | false | string | 区号       |
| phone    | true  | string | 手机号     |
| code     | true  | string | 验证码     |
| mnemonic | true  | string | 加密助记词 |

**返回参数：**

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```

### 绑定邮箱

URL: /backup/email-binding（弃用）

`post`|`auth`

**请求参数：**

| 参数名   | 必选 | 类型   | 说明       |
| -------- | ---- | ------ | ---------- |
| email    | true | string | 手机号     |
| code     | true | string | 验证码     |
| mnemonic | true | string | 加密助记词 |

**返回参数：**

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```

### 手机找回私钥

URL: /backup/phone-retrieve（弃用）

`post`

**请求参数：**

| 参数名 | 必选 | 类型   | 说明   |
| ------ | ---- | ------ | ------ |
| area   | true | string | 区号   |
| phone  | true | string | 手机号 |
| code   | true | string | 验证码 |

**返回参数：**

| 参数名      | 必选 | 类型   | 说明       |
| ----------- | ---- | ------ | ---------- |
| address     | true | string | 用户地址   |
| area        | true | string | 区号       |
| phone       | true | string | 手机号     |
| email       | true | string | 邮箱       |
| mnemonic    | true | string | 加密助记词 |
| update_time | true | string | 更新时间   |

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
        "address": "0xffffffff",
        "area":"",
        "phone":"",
        "email":"",
        "mnemonic":"",
        "update_time":""
    }
}
```
### 邮箱找回私钥

URL: /backup/email-retrieve（弃用）

`post`

**请求参数：**

| 参数名 | 必选 | 类型   | 说明   |
| ------ | ---- | ------ | ------ |
| email  | true | string | 邮箱   |
| code   | true | string | 验证码 |

**返回参数：**

| 参数名      | 必选 | 类型   | 说明       |
| ----------- | ---- | ------ | ---------- |
| address     | true | string | 用户地址   |
| area        | true | string | 区号       |
| phone       | true | string | 手机号     |
| email       | true | string | 邮箱       |
| mnemonic    | true | string | 加密助记词 |
| update_time | true | string | 更新时间   |

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
        "address": "0xffffffff",
        "area":"",
        "phone":"",
        "email":"",
        "mnemonic":"",
        "update_time":""
    }
}
```

### bty和btc地址映射（第一阶段）

URL: /backup/transform/addressEnrolment

`post`

**请求参数：**

| 参数名 | 必选 | 类型   | 说明   |
| ------ | ---- | ------ | ------ |
| btcAddress  | true | string |   |
| ethAddress  | true | string |   |

**返回参数：**

| 参数名      | 必选 | 类型   | 说明       |
| ----------- | ---- | ------ | ---------- |

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```