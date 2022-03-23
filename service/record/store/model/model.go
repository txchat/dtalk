package model

type MsgContent struct {
	Mid        string
	Seq        string
	SenderId   string
	ReceiverId string
	MsgType    uint32
	Content    string
	CreateTime uint64
	Source     string
	Reference  string
}

type MsgRelation struct {
	Mid        string
	OwnerUid   string
	OtherUid   string
	Type       uint8
	State      uint8
	CreateTime uint64
}

type MsgCache struct {
	Mid        string
	Seq        string
	SenderId   string
	ReceiverId string
	MsgType    uint32
	Content    string
	CreateTime uint64
	Source     string
	Reference  string
	Prev       uint64
	Version    uint64
}

type SignalContent struct {
	Id         string
	Uid        string
	Type       uint8
	State      uint8
	Content    string
	CreateTime uint64
	UpdateTime uint64
}
