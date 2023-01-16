### 1. "撤回消息"

1. route definition

- Url: /app/record/revoke
- Method: POST
- Request: `RevokeReq`
- Response: `RevokeResp`

2. request definition



```golang
type RevokeReq struct {
	Type int `json:"type,optional,options=0|1"` //撤回类型: 0-&gt;私聊, 1-&gt;群聊
	Mid int64 `json:"logId"`
}
```


3. response definition



```golang
type RevokeResp struct {
}
```

### 2. "关注消息"

1. route definition

- Url: /app/record/focus
- Method: POST
- Request: `FocusReq`
- Response: `FocusResp`

2. request definition



```golang
type FocusReq struct {
	Type int `json:"type,optional,options=0|1"` //关注类型: 0-&gt;私聊, 1-&gt;群聊
	Mid int64 `json:"logId"`
}
```


3. response definition



```golang
type FocusResp struct {
}
```

### 3. "同步聊天记录"

1. route definition

- Url: /app/record/sync-record
- Method: POST
- Request: `SyncReq`
- Response: `SyncResp`

2. request definition



```golang
type SyncReq struct {
	MaxCount int64 `json:"count,range=[1:1000]"` // 消息数量
	StartMid int64 `json:"start,optional"` // 消息 ID
}
```


3. response definition



```golang
type SyncResp struct {
	RecordCount int `json:"record_count"` // 聊天记录数量
	Records []string `json:"records"` // 聊天记录 base64 encoding
}
```

### 4. "获取私聊消息"

1. route definition

- Url: /app/record/pri-chat-record
- Method: POST
- Request: `PrivateRecordReq`
- Response: `PrivateRecordResp`

2. request definition



```golang
type PrivateRecordReq struct {
	FromId string `json:"-"`
	TargetId string `json:"targetId"`
	RecordCount int64 `json:"count,range=[1:100]"`
	Mid string `json:"logId"`
}
```


3. response definition



```golang
type PrivateRecordResp struct {
	RecordCount int `json:"record_count"` // 聊天记录数量
	Records []*Record `json:"records"` // 聊天记录
}
```

### 5. "发送消息"

1. route definition

- Url: /record/push
- Method: POST
- Request: `PushReq`
- Response: `PushResp`

2. request definition



```golang
type PushReq struct {
	File string `form:"file"`
}
```


3. response definition



```golang
type PushResp struct {
	Mid int64 `json:"logId"`
	Datetime uint64 `json:"datetime"`
}
```

### 6. "发送消息"

1. route definition

- Url: /record/push2
- Method: POST
- Request: `PushReq`
- Response: `PushResp`

2. request definition



```golang
type PushReq struct {
	File string `form:"file"`
}
```


3. response definition



```golang
type PushResp struct {
	Mid int64 `json:"logId"`
	Datetime uint64 `json:"datetime"`
}
```

