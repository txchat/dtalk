package model

import (
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"
	"github.com/txchat/dtalk/app/services/storage/storageclient"
	"github.com/txchat/dtalk/internal/recordutil"
)

func ToChatRecord(records []*storageclient.Record) []*types.Record {
	rlt := make([]*types.Record, len(records))
	for i, msg := range records {
		rlt[i] = &types.Record{
			Mid:        msg.Mid,
			Cid:        msg.Cid,
			FromId:     msg.SenderId,
			TargetId:   msg.ReceiverId,
			MsgType:    msg.MsgType,
			Content:    recordutil.CommonMsgJSONDataToProtobuf(msg.MsgType, []byte(msg.Content)),
			CreateTime: msg.CreateTime,
		}
	}
	return rlt
}
