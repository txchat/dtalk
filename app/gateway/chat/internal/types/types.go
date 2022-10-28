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

type RevokeReq struct {
	Type int   `json:"type,optional" enums:"0,1"`
	Mid  int64 `json:"logId"`
}

type RevokeResp struct {
}

type FocusReq struct {
	Type int   `json:"type,optional" enums:"0,1"`
	Mid  int64 `json:"logId"`
}

type FocusResp struct {
}

type SyncReq struct {
	MaxCount int64 `json:"count,range=[1:1000]"`
	StartMid int64 `json:"start,optional"`
}

type SyncResp struct {
	RecordCount int      `json:"record_count"`
	Records     []string `json:"records"`
}

type PrivateRecordReq struct {
	FromId      string `json:"-"`
	TargetId    string `json:"targetId"`
	RecordCount int64  `json:"count,range=[1:100]"`
	Mid         string `json:"logId"`
}

type PrivateRecordResp struct {
	RecordCount int       `json:"record_count"`
	Records     []*Record `json:"records"`
}

type Record struct {
	Mid        string      `json:"logId"`
	Seq        string      `json:"msgId"`
	FromId     string      `json:"fromId"`
	TargetId   string      `json:"targetId"`
	MsgType    int32       `json:"msgType"`
	Content    interface{} `json:"content"`
	CreateTime uint64      `json:"createTime"`
}

type PushReq struct {
	File string `form:"file"`
}

type PushResp struct {
	Mid      int64  `json:"logId"`
	Datetime uint64 `json:"datetime"`
}

type LoginReq struct {
	ConnType int32 `json:"connType"`
}

type LoginResp struct {
	Address string `json:"address"`
}
