// Code generated by goctl. DO NOT EDIT.
package types

type StartCallReq struct {
	GroupId  string   `json:"groupId"`
	Invitees []string `json:"invitees"`
	RTCType  int32    `json:"RTCType,options=1|2"`
}

type StartCallResp struct {
	TraceId    int64    `json:"traceId"`
	TraceIdStr string   `json:"traceIdStr"`
	RTCType    int32    `json:"RTCType"`
	Invitees   []string `json:"invitees"`
	Caller     string   `json:"caller"`
	CreateTime int64    `json:"createTime"`
	Timeout    int32    `json:"timeout"`
	Deadline   int64    `json:"deadline"`
	GroupId    string   `json:"groupId"` // 0表示私聊, 其他表示群聊
}

type ReplyBusyReq struct {
	TraceId    int64  `json:"traceId"`
	TraceIdStr string `json:"traceIdStr,optional"` // 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
}

type ReplyBusyResp struct {
}

type CheckCallReq struct {
	TraceId    int64  `json:"traceId"`
	TraceIdStr string `json:"traceIdStr,optional"` // 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
}

type CheckCallResp struct {
	TraceId    int64    `json:"traceId"`
	TraceIdStr string   `json:"traceIdStr"`
	RTCType    int32    `json:"RTCType"`
	Invitees   []string `json:"invitees"`
	Caller     string   `json:"caller"`
	CreateTime int64    `json:"createTime"`
	Timeout    int32    `json:"timeout"`
	Deadline   int64    `json:"deadline"`
	GroupId    string   `json:"groupId"` // 0表示私聊, 其他表示群聊
}

type HandleCallReq struct {
	Answer     bool   `json:"answer"`
	TraceId    int64  `json:"traceId"`
	TraceIdStr string `json:"traceIdStr,optional"` // 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
}

type HandleCallResp struct {
	RoomId        int32  `json:"roomId"`
	UserSig       string `json:"userSig"`
	PrivateMapKey string `json:"privateMapKey"`
	SDKAppId      int32  `json:"sdkAppId"`
}
