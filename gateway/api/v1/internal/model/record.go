package model

type RevokeMsgReq struct {
	Type int   `json:"type" enums:"0,1"`
	Mid  int64 `json:"logId" binding:"required"`
}

type FocusMsgReq struct {
	Type  int   `json:"type" enums:"0,1"`
	LogId int64 `json:"logId" binding:"required"`
}

type Record struct {
	// log id
	Mid string `json:"logId"`
	// msg id (uuid)
	Seq string `json:"msgId"`
	// 发送者 id
	FromId string `json:"fromId"`
	// 接收者 id
	TargetId string `json:"targetId"`
	// 消息类型
	MsgType int32 `json:"msgType"`
	// 消息内容
	Content interface{} `json:"content"`
	// 消息发送时间
	CreateTime uint64 `json:"createTime"`
}

type GetPriRecordsReq struct {
	// 发送者 ID
	FromId string `json:"-"`

	// 发送者 ID
	//FromId    	string `json:"fromId" binding:"required"`

	// 接受者 ID
	TargetId string `json:"targetId" binding:"required"`
	// 消息数量
	RecordCount int64 `json:"count" binding:"required,min=1,max=100"`
	// 消息 ID
	Mid string `json:"logId"`
}

type GetPriRecordsResp struct {
	// 聊天记录数量
	RecordCount int `json:"record_count"`
	// 聊天记录
	Records []*Record `json:"records"`
}

type SyncRecordsReq struct {
	// 消息数量
	MaxCount int64 `json:"count" binding:"required,min=1,max=1000"`
	// 消息 ID
	StartMid int64 `json:"start"`
}

type SyncRecordsResp struct {
	// 聊天记录数量
	RecordCount int `json:"record_count"`
	// 聊天记录 base64 encoding
	Records []string `json:"records"`
}
