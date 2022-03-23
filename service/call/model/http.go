package model

type GeneralResponse struct {
	Result  int         `json:"result"`
	Message int         `json:"message"`
	Data    interface{} `json:"data"`
}

type StartCallRequest struct {
	PersonId string   `json:"-"`
	GroupId  string   `json:"groupId"`
	Invitees []string `json:"invitees" binding:"required"`
	RTCType  int32    `json:"RTCType" binding:"oneof=1 2"`
}

type StartCallResponse struct {
	TraceId    int64    `json:"traceId"`
	TraceIdStr string   `json:"traceIdStr"`
	RTCType    int32    `json:"RTCType"`
	Invitees   []string `json:"invitees"`
	Caller     string   `json:"caller"`
	CreateTime int64    `json:"createTime"`
	Timeout    int32    `json:"timeout"`
	Deadline   int64    `json:"deadline"`
	// 0表示私聊, 其他表示群聊
	GroupId string `json:"groupId"`
}

type ReplyBusyRequest struct {
	PersonId string `json:"-"`
	TraceId  int64  `json:"traceId"`
	// 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
	TraceIdStr string `json:"traceIdStr"`
}

type ReplyBusyResponse struct {
}

type CheckCallRequest struct {
	PersonId string `json:"-"`
	TraceId  int64  `json:"traceId"`
	// 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
	TraceIdStr string `json:"traceIdStr"`
}

type CheckCallResponse struct {
	TraceId    int64    `json:"traceId"`
	TraceIdStr string   `json:"traceIdStr"`
	RTCType    int32    `json:"RTCType"`
	Invitees   []string `json:"invitees"`
	Caller     string   `json:"caller"`
	CreateTime int64    `json:"createTime"`
	Timeout    int32    `json:"timeout"`
	Deadline   int64    `json:"deadline"`
	GroupId    string   `json:"groupId"`
}

type HandleCallRequest struct {
	PersonId string `json:"-"`
	Answer   bool   `json:"answer"`
	TraceId  int64  `json:"traceId"`
	// 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
	TraceIdStr string `json:"traceIdStr"`
}

type HandleCallResponse struct {
	RoomId        int32  `json:"roomId"`
	UserSig       string `json:"userSig"`
	PrivateMapKey string `json:"privateMapKey"`
	SDKAppId      int32  `json:"sdkAppId"`
}
