package logic

import (
	"context"

	"github.com/txchat/dtalk/internal/recordutil"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/app/services/storage/storage"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/imparse/proto/common"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSyncRecordsAfterMidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSyncRecordsAfterMidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSyncRecordsAfterMidLogic {
	return &GetSyncRecordsAfterMidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSyncRecordsAfterMidLogic) GetSyncRecordsAfterMid(in *storage.GetSyncRecordsAfterMidReq) (*storage.GetSyncRecordsAfterMidReply, error) {
	items, err := l.getSyncMsgJustBizLevel("", in.GetUid(), in.GetMid(), in.GetCount())
	if err != nil {
		return nil, err
	}
	return &storage.GetSyncRecordsAfterMidReply{
		Records: items,
	}, nil
}

func (l *GetSyncRecordsAfterMidLogic) getSyncMsgJustBizLevel(key, uid string, startMid, count int64) ([][]byte, error) {
	bizP := &common.Proto{
		EventType: common.Proto_common,
	}

	//私聊 DESC
	uMsg, err := l.svcCtx.Repo.UserMsgAfter(uid, startMid, count)
	if err != nil {
		return nil, err
	}

	//to ASC
	for i := 0; i < len(uMsg)/2; i++ {
		uMsg[i], uMsg[len(uMsg)-1-i] = uMsg[len(uMsg)-1-i], uMsg[i]
	}

	//群聊
	gMsg, err := l.svcCtx.Repo.GroupMsgAfter(uid, startMid, count)
	if err != nil {
		return nil, err
	}
	//to ASC
	for i := 0; i < len(gMsg)/2; i++ {
		gMsg[i], gMsg[len(gMsg)-1-i] = gMsg[len(gMsg)-1-i], gMsg[i]
	}

	var records = make([][]byte, len(uMsg)+len(gMsg))
	var num = 0
	for _, m := range uMsg {
		eveP := &common.Common{
			ChannelType: common.Channel_ToUser,
			Mid:         util.MustToInt64(m.Mid),
			Seq:         m.Seq,
			From:        m.SenderId,
			Target:      m.ReceiverId,
			MsgType:     common.MsgType(m.MsgType),
			Msg:         recordutil.CommonMsgJSONDataToProtobufData(m.MsgType, []byte(m.Content)),
			Source:      recordutil.SourceJSONUnmarshal([]byte(m.Source)),
			Reference:   recordutil.ReferenceJSONUnmarshal([]byte(m.Reference)),
			Datetime:    m.CreateTime,
		}
		bizP.Body, err = proto.Marshal(eveP)
		bytes, err := proto.Marshal(bizP)
		if err != nil {
			l.Error("Push msg Marshal failed", "err", err)
			continue
		}
		records[num] = bytes
		num++
	}

	for _, m := range gMsg {
		eveP := &common.Common{
			ChannelType: common.Channel_ToGroup,
			Mid:         util.MustToInt64(m.Mid),
			Seq:         m.Seq,
			From:        m.SenderId,
			Target:      m.ReceiverId,
			MsgType:     common.MsgType(m.MsgType),
			Msg:         recordutil.CommonMsgJSONDataToProtobufData(m.MsgType, []byte(m.Content)),
			Source:      recordutil.SourceJSONUnmarshal([]byte(m.Source)),
			Reference:   recordutil.ReferenceJSONUnmarshal([]byte(m.Reference)),
			Datetime:    m.CreateTime,
		}
		bizP.Body, err = proto.Marshal(eveP)
		bytes, err := proto.Marshal(bizP)
		if err != nil {
			l.Error("Push msg Marshal failed", "err", err)
			continue
		}
		records[num] = bytes
		num++
	}
	return records[0:num], nil
}
