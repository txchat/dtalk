package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/imparse/proto/signal"
	"github.com/zeromicro/go-zero/core/logx"
)

type StoreSignalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStoreSignalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreSignalLogic {
	return &StoreSignalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StoreSignalLogic) AppendGroupCastSignal(mid, target string, signalType signal.SignalType, content []byte, createTime uint64) error {
	//获取所有群成员
	getGroupMembersLogic := NewGetGroupMembersLogic(l.ctx, l.svcCtx)
	members, err := getGroupMembersLogic.GetGroupMembersLogic(target)
	if err != nil {
		l.Error("AllGroupMembers failed", "err", err)
		return err
	}
	var signalItems = make([]*model.SignalContent, len(members))
	for i, member := range members {
		signalItems[i] = &model.SignalContent{
			Id:         mid,
			Uid:        member,
			Type:       uint8(signalType),
			State:      uint8(model.UnReceive),
			Content:    string(content),
			CreateTime: createTime,
			UpdateTime: createTime,
		}
	}
	_, _, err = l.svcCtx.Repo.BatchAppendSignalContent(signalItems)
	if err != nil {
		l.Error("BatchAppendSignalContent failed", "err", err)
		return err
	}
	return nil
}

func (l *StoreSignalLogic) AppendUniCastSignal(mid, target string, signalType signal.SignalType, content []byte, createTime uint64) error {
	msg := &model.SignalContent{
		Id:         mid,
		Uid:        target,
		Type:       uint8(signalType),
		State:      uint8(model.UnReceive),
		Content:    string(content),
		CreateTime: createTime,
		UpdateTime: createTime,
	}
	_, _, err := l.svcCtx.Repo.AppendSignalContent(msg)
	if err != nil {
		l.Error("AppendSignalContent failed", "msg", msg, "err", err)
		return err
	}
	return nil
}
