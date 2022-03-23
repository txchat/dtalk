####  获取模块启用状态
URL: /app/modules/all

`post`

**请求参数：**

无

**返回参数：**

| **参数**    | **名字**     | **类型** | **说明**   |
| ----------- | ------------ | -------- | ---------- |
| name | 模块名称 | string   | 钱包模块：wallet，企业模块：oa，红包模块：redpacket |
| isEnabled | 是否启用 | bool |               |
| endPoints | 模块访问地址 | string | |

```json
{
    "result": 0,
    "message": "",
    "data": [
        {
            "name": "wallet",
            "isEnabled": true,
            "endPoints": [
                "https://172.16.101.126:8083"
            ]
        },
        {
            "name": "oa",
            "isEnabled": true,
            "endPoints": [
                "http://127.0.0.1"
            ]
        }
    ]
}
```

#### 消息撤回
URL: /app/record/revoke

`post`
**请求参数：**

| **参数**    | **名字**     | **类型** | **约束**   | **说明**   |
| ----------- | ------------ | -------- | ---------- |---------- |
| type | 类型 | int  | 必填 | 0->撤回私聊消息；1->撤回群聊消息 |
| logId | 消息id | int64 | 必填|               |

```json
{
    "type": 0,
    "logId": 0
}
```

**返回参数：**
无
```json
{
    "result": 0,
    "message": "",
    "data": {}
}
```

#### 关注消息
URL: /app/record/focus

`post` `auth`  
**请求参数：**

| **参数**    | **名字**     | **类型** | **约束**   | **说明**   |
| ----------- | ------------ | -------- | ---------- |---------- |
| type | 类型 | int  | 必填 | 0->关注私聊消息；1->关注群聊消息 |
| logId | 消息id | int64 | 必填|               |

```json
{
    "type": 0,
    "logId": 0
}
```

**返回参数：**
无
```json
{
    "result": 0,
    "message": "",
    "data": {}
}
```
