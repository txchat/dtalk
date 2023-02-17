package model

type PrivateMsgContent struct {
	Mid        string
	Cid        string
	SenderId   string
	ReceiverId string
	MsgType    uint32
	Content    string
	CreateTime uint64
	Source     string
	Reference  string
}

type PrivateMsgRelation struct {
	Mid        string
	OwnerUid   string
	OtherUid   string
	Type       uint8
	CreateTime uint64
}

type GroupMsgContent struct {
	Mid        string
	Cid        string
	SenderId   string
	ReceiverId string
	GroupId    string
	MsgType    uint32
	Content    string
	CreateTime uint64
	Source     string
	Reference  string
}

type SignalContent struct {
	Uid        string
	Seq        int64
	Type       uint8
	Content    string
	CreateTime uint64
	UpdateTime uint64
}
