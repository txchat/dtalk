### 1. "发起音视频通话"

1. route definition

- Url: /app/start-call
- Method: POST
- Request: `StartCallReq`
- Response: `StartCallResp`

2. request definition



```golang
type StartCallReq struct {
	GroupId string `json:"groupId"`
	Invitees []string `json:"invitees"`
	RTCType int32 `json:"RTCType,options=1|2"` //1-&gt;音频, 2-&gt;视频
}
```


3. response definition



```golang
type StartCallResp struct {
	TraceId int64 `json:"traceId"`
	TraceIdStr string `json:"traceIdStr"`
	RTCType int32 `json:"RTCType"` //1-&gt;音频, 2-&gt;视频
	Invitees []string `json:"invitees"`
	Caller string `json:"caller"`
	CreateTime int64 `json:"createTime"`
	Timeout int32 `json:"timeout"`
	Deadline int64 `json:"deadline"`
	GroupId string `json:"groupId"` // 0-&gt;私聊, ^0-&gt;群id
}
```

### 2. "通话响应-繁忙"

1. route definition

- Url: /app/reply-busy
- Method: POST
- Request: `ReplyBusyReq`
- Response: `ReplyBusyResp`

2. request definition



```golang
type ReplyBusyReq struct {
	TraceId int64 `json:"traceId"`
	TraceIdStr string `json:"traceIdStr,optional"` // 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
}
```


3. response definition



```golang
type ReplyBusyResp struct {
}
```

### 3. "检查通话"

1. route definition

- Url: /app/check-call
- Method: POST
- Request: `CheckCallReq`
- Response: `CheckCallResp`

2. request definition



```golang
type CheckCallReq struct {
	TraceId int64 `json:"traceId"`
	TraceIdStr string `json:"traceIdStr,optional"` // 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
}
```


3. response definition



```golang
type CheckCallResp struct {
	TraceId int64 `json:"traceId"`
	TraceIdStr string `json:"traceIdStr"`
	RTCType int32 `json:"RTCType"` //1-&gt;音频, 2-&gt;视频
	Invitees []string `json:"invitees"`
	Caller string `json:"caller"`
	CreateTime int64 `json:"createTime"`
	Timeout int32 `json:"timeout"`
	Deadline int64 `json:"deadline"`
	GroupId string `json:"groupId"` // 0-&gt;私聊, ^0-&gt;群id
}
```

### 4. "处理通话"

1. route definition

- Url: /app/handle-call
- Method: POST
- Request: `HandleCallReq`
- Response: `HandleCallResp`

2. request definition



```golang
type HandleCallReq struct {
	Answer bool `json:"answer"`
	TraceId int64 `json:"traceId"`
	TraceIdStr string `json:"traceIdStr,optional"` // 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
}
```


3. response definition



```golang
type HandleCallResp struct {
	RoomId int32 `json:"roomId"`
	UserSig string `json:"userSig"`
	PrivateMapKey string `json:"privateMapKey"`
	SDKAppId int32 `json:"sdkAppId"`
}
```

