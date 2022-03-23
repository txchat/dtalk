# 找回服务
### 发送消息
URL: /record/push2

`post`  
**Request Headers**
- Content-Type: multipart/form-data

**Body:**`form-data`
  
| 参数名 | 必选 | 类型   | 说明   |
| ------ | ---- | ------ | ------ |
| name   | true | string | 区号   |
| filename  | true | string | 手机号 |

**请求示例**
```
请求头:
	Content-Type=multipart/form-data
请求参数:
	name="message"
	filename="message"
```

**返回参数：**
```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
    }
}
```