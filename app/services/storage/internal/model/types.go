package model

import "github.com/txchat/dtalk/app/services/storage/storage"

type MsgContent struct {
	Mid        string
	Cid        string
	SenderId   string
	ReceiverId string
	MsgType    int32
	Content    string
	CreateTime int64
	Source     string
	Reference  string
}

type MsgRelation struct {
	Mid        string
	OwnerUid   string
	OtherUid   string
	Type       int8
	CreateTime int64
}

type SignalContent struct {
	Uid        string
	Seq        int64
	Type       int8
	Content    string
	CreateTime int64
	UpdateTime int64
}

func ChatMsgRepoToRPC(record *MsgContent) *storage.Record {
	return &storage.Record{
		Mid:        record.Mid,
		Cid:        record.Cid,
		SenderId:   record.SenderId,
		ReceiverId: record.ReceiverId,
		MsgType:    record.MsgType,
		Content:    record.Content,
		CreateTime: record.CreateTime,
		Source:     record.Source,
		Reference:  record.Reference,
	}
}

func ChatMessagesRepoToRPC(records []*MsgContent) []*storage.Record {
	rlt := make([]*storage.Record, len(records))
	for i, record := range records {
		rlt[i] = ChatMsgRepoToRPC(record)
	}
	return rlt
}
