package logic

import (
	"context"
	"strconv"

	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/imparse/proto/auth"
	"github.com/txchat/imparse/proto/common"
	"github.com/zeromicro/go-zero/core/logx"
)

type StorePrivateMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStorePrivateMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StorePrivateMsgLogic {
	return &StorePrivateMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StorePrivateMsgLogic) appendMsg(isSelfRead bool, p *common.Common) error {
	tx, err := l.svcCtx.Repo.NewTx()
	if err != nil {
		return err
	}
	_, _, err = l.svcCtx.Repo.AppendMsgContent(tx, &model.MsgContent{
		Mid:        strconv.FormatInt(p.Mid, 10),
		Seq:        p.Seq,
		SenderId:   p.From,
		ReceiverId: p.Target,
		MsgType:    uint32(p.MsgType),
		Content:    string(model.ParseCommonMsg(p)),
		CreateTime: p.Datetime,
		Source:     string(model.ParseSource(p)),
		Reference:  string(model.ParseReference(p)),
	})
	if err != nil {
		tx.RollBack()
		return err
	}
	var selfRead uint8 = model.Received
	if !isSelfRead {
		selfRead = model.UnReceive
	}
	_, _, err = l.svcCtx.Repo.AppendMsgRelation(tx, &model.MsgRelation{
		Mid:        strconv.FormatInt(p.Mid, 10),
		OwnerUid:   p.From,
		OtherUid:   p.Target,
		Type:       model.Send,
		State:      selfRead,
		CreateTime: p.Datetime,
	})
	if err != nil {
		tx.RollBack()
		return err
	}
	_, _, err = l.svcCtx.Repo.AppendMsgRelation(tx, &model.MsgRelation{
		Mid:        strconv.FormatInt(p.Mid, 10),
		OwnerUid:   p.Target,
		OtherUid:   p.From,
		Type:       model.Rev,
		State:      model.UnReceive,
		CreateTime: p.Datetime,
	})
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

func (l *StorePrivateMsgLogic) pushMem(p *common.Common) error {
	fromVer, err := l.svcCtx.Repo.IncMsgVersion(p.GetFrom())
	if err != nil {
		l.Error("PushMem IncMsgVersion failed", "uid", p.GetFrom())
	}
	toVer, err := l.svcCtx.Repo.IncMsgVersion(p.Target)
	if err != nil {
		l.Error("PushMem IncMsgVersion failed", "uid", p.GetTarget())
	}
	//store sender
	if fromVer != 0 {
		err = l.svcCtx.Repo.AddRecordCache(p.From, fromVer, &model.MsgCache{
			Mid:        strconv.FormatInt(p.Mid, 10),
			Seq:        p.Seq,
			SenderId:   p.From,
			ReceiverId: p.Target,
			MsgType:    uint32(p.MsgType),
			Content:    string(model.ParseCommonMsg(p)),
			CreateTime: p.Datetime,
			Source:     string(model.ParseSource(p)),
			Reference:  string(model.ParseReference(p)),
			Prev:       0,
			Version:    fromVer,
		})
		if err != nil {
			l.Error("PushMem AddRecordCache failed", "uid", p.From, "version", fromVer)
		}
	}
	//store receiver
	if toVer != 0 {
		err = l.svcCtx.Repo.AddRecordCache(p.Target, toVer, &model.MsgCache{
			Mid:        strconv.FormatInt(p.Mid, 10),
			Seq:        p.Seq,
			SenderId:   p.From,
			ReceiverId: p.Target,
			MsgType:    uint32(p.MsgType),
			Content:    string(model.ParseCommonMsg(p)),
			CreateTime: p.Datetime,
			Source:     string(model.ParseSource(p)),
			Reference:  string(model.ParseReference(p)),
			Prev:       0,
			Version:    toVer,
		})
		if err != nil {
			l.Error("PushMem AddRecordCache failed", "uid", p.From, "version", fromVer)
		}
	}
	return nil
}

func (l *StorePrivateMsgLogic) StoreMsg(deviceType auth.Device, pro *common.Common) error {
	//step 1.存数据库
	isSelfRead := false
	if deviceType == auth.Device_IOS || deviceType == auth.Device_Android {
		isSelfRead = true
	}
	err := l.appendMsg(isSelfRead, pro)
	if err != nil {
		l.Error("AppendMsg failed", "err", err)
		return model.ErrConsumeRedo
	}
	if l.svcCtx.Config.SyncCache {
		//step 2.存缓存队列
		err = l.pushMem(pro)
		if err != nil {
			l.Error("PushMem failed", "err", err)
		}
	}
	return nil
}
