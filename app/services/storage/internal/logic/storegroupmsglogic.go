package logic

import (
	"context"
	"strconv"

	"github.com/txchat/dtalk/internal/recordutil"

	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/imparse/proto/auth"
	"github.com/txchat/imparse/proto/common"
	"github.com/zeromicro/go-zero/core/logx"
)

type StoreGroupMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStoreGroupMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreGroupMsgLogic {
	return &StoreGroupMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StoreGroupMsgLogic) appendMsg(isSelfRead bool, p *common.Common) error {
	//获取所有群成员
	getGroupMembersLogic := NewGetGroupMembersLogic(l.ctx, l.svcCtx)
	members, err := getGroupMembersLogic.GetGroupMembersLogic(p.Target)
	if err != nil {
		return err
	}
	var selfRead uint8 = model.Received
	if !isSelfRead {
		selfRead = model.UnReceive
	}
	var msgRelate = make([]*model.MsgRelation, len(members))
	for i, member := range members {
		var item *model.MsgRelation
		if member == p.From {
			//发送者
			item = &model.MsgRelation{
				Mid:        strconv.FormatInt(p.Mid, 10),
				OwnerUid:   p.From,
				OtherUid:   p.Target,
				Type:       model.Send,
				State:      selfRead,
				CreateTime: p.Datetime,
			}
		} else {
			item = &model.MsgRelation{
				Mid:        strconv.FormatInt(p.Mid, 10),
				OwnerUid:   member,
				OtherUid:   p.Target,
				Type:       model.Rev,
				State:      model.UnReceive,
				CreateTime: p.Datetime,
			}
		}
		msgRelate[i] = item
	}
	tx, err := l.svcCtx.Repo.NewTx()
	if err != nil {
		return err
	}
	_, _, err = l.svcCtx.Repo.AppendGroupMsgContent(tx, &model.MsgContent{
		Mid:        strconv.FormatInt(p.Mid, 10),
		Seq:        p.Seq,
		SenderId:   p.From,
		ReceiverId: p.Target,
		MsgType:    uint32(p.MsgType),
		Content:    string(recordutil.CommonMsgProtobufDataToJSONData(p)),
		CreateTime: p.Datetime,
		Source:     string(recordutil.SourceJSONMarshal(p)),
		Reference:  string(recordutil.ReferenceJSONMarshal(p)),
	})
	if err != nil {
		tx.RollBack()
		return err
	}
	_, _, err = l.svcCtx.Repo.AppendGroupMsgRelation(tx, msgRelate)
	if err != nil {
		tx.RollBack()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.RollBack()
		return err
	}
	return nil
}

func (l *StoreGroupMsgLogic) StoreMsg(deviceType auth.Device, pro *common.Common) error {
	//step 1.存数据库
	isSelfRead := false
	if deviceType == auth.Device_IOS || deviceType == auth.Device_Android {
		isSelfRead = true
	}
	err := l.appendMsg(isSelfRead, pro)
	if err != nil {
		l.Error("AppendGroupMsg failed", "err", err)
		return model.ErrConsumeRedo
	}
	return nil
}
