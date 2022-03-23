# 发现服务
#### 获取聊天服务列表
URL: /disc/nodes

`post`

**请求参数：**无

**返回参数：**

| 参数名  | 必填 | 类型     | 说明           |
| ------- | ---- | -------- | -------------- |
| servers | true | []object | 聊天服务器列表 |
| name    | true | string   | 名称           |
| address | true | string   | 地址           |
| nodes   | true | []object | 合约节点列表   |
| name    | true | string   | 名称           |
| address | true | string   | 地址           |

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
        "servers": [
            {
                "name": "默认服务器",
                "address":"127.0.0.1:8080"
            },{
                "name": "默认服务器2",
                "address":"chat.33.cn"
            }
        ],
        "nodes": [
            {
                "name": "合约节点1",
                "address":"127.0.0.1:8080"
            }
        ]
    }
}
```