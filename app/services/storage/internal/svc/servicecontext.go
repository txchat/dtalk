package svc

import (
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/api/proto/signal"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/app/services/pusher/pusherclient"
	"github.com/txchat/dtalk/app/services/storage/internal/config"
	"github.com/txchat/dtalk/app/services/storage/internal/dao"
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/internal/recordutil"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Repo   dao.Repository
	//need not init
	DeviceRPC deviceclient.Device
	PusherRPC pusherclient.Pusher
	GroupRPC  groupclient.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		Config: c,
		Repo:   dao.NewStorageRepository(c.RedisDB, c.MySQL),
		DeviceRPC: deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		PusherRPC: pusherclient.NewPusher(zrpc.MustNewClient(c.PusherRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		GroupRPC: groupclient.NewGroup(zrpc.MustNewClient(c.GroupRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
	}
	return s
}

func (s *ServiceContext) StorePrivateMessage(msg *message.Message) error {
	tx, err := s.Repo.NewTx()
	if err != nil {
		return err
	}
	_, _, err = s.Repo.AppendPrivateMsgContent(tx, &model.MsgContent{
		Mid:        msg.GetMid(),
		Cid:        msg.GetCid(),
		SenderId:   msg.GetFrom(),
		ReceiverId: msg.GetTarget(),
		MsgType:    int32(msg.GetMsgType()),
		Content:    string(recordutil.CommonMsgProtobufDataToJSONData(msg)),
		CreateTime: msg.GetDatetime(),
		Source:     string(recordutil.SourceJSONMarshal(msg)),
		Reference:  string(recordutil.ReferenceJSONMarshal(msg)),
	})
	if err != nil {
		tx.RollBack()
		return err
	}
	_, _, err = s.Repo.AppendPrivateMsgRelation(tx, &model.MsgRelation{
		Mid:        msg.GetMid(),
		OwnerUid:   msg.GetFrom(),
		OtherUid:   msg.GetTarget(),
		Type:       model.Send,
		CreateTime: msg.GetDatetime(),
	})
	if err != nil {
		tx.RollBack()
		return err
	}
	_, _, err = s.Repo.AppendPrivateMsgRelation(tx, &model.MsgRelation{
		Mid:        msg.GetMid(),
		OwnerUid:   msg.GetTarget(),
		OtherUid:   msg.GetFrom(),
		Type:       model.Rev,
		CreateTime: msg.GetDatetime(),
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

func (s *ServiceContext) StoreGroupMessage(members []string, msg *message.Message) error {
	var msgRelate = make([]*model.MsgRelation, len(members))
	for i, member := range members {
		sendOrRev := model.Rev
		if member == msg.GetFrom() {
			sendOrRev = model.Send
		}
		//发送者
		msgRelate[i] = &model.MsgRelation{
			Mid:        msg.GetMid(),
			OwnerUid:   member,
			OtherUid:   msg.GetTarget(),
			Type:       int8(sendOrRev),
			CreateTime: msg.GetDatetime(),
		}
	}
	tx, err := s.Repo.NewTx()
	if err != nil {
		return err
	}
	_, _, err = s.Repo.AppendGroupMsgContent(tx, &model.MsgContent{
		Mid:        msg.GetMid(),
		Cid:        msg.GetCid(),
		SenderId:   msg.GetFrom(),
		ReceiverId: msg.GetTarget(),
		MsgType:    int32(msg.GetMsgType()),
		Content:    string(recordutil.CommonMsgProtobufDataToJSONData(msg)),
		CreateTime: msg.GetDatetime(),
		Source:     string(recordutil.SourceJSONMarshal(msg)),
		Reference:  string(recordutil.ReferenceJSONMarshal(msg)),
	})
	if err != nil {
		tx.RollBack()
		return err
	}
	_, _, err = s.Repo.AppendGroupMsgRelation(tx, msgRelate)
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

func (s *ServiceContext) StoreSignal(target string, seq int64, sig *signal.Signal) error {
	now := util.TimeNowUnixMilli()
	m := &model.SignalContent{
		Uid:        target,
		Seq:        seq,
		Type:       int8(sig.GetType()),
		Content:    string(recordutil.SignalContentToJSONData(sig)),
		CreateTime: now,
		UpdateTime: now,
	}
	_, _, err := s.Repo.AppendSignalContent(m)
	return err
}
