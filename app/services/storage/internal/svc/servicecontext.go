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
	Repo   dao.StorageRepository
	//need not init
	DeviceRPC deviceclient.Device
	PusherRPC pusherclient.Pusher
	GroupRPC  groupclient.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		Config: c,
		Repo:   dao.NewUniRepository(c.RedisDB, c.MySQL),
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
	_, _, err = s.Repo.AppendPrivateMsgContent(tx, &model.PrivateMsgContent{
		Mid:        msg.GetMid(),
		Cid:        msg.GetCid(),
		SenderId:   msg.GetFrom(),
		ReceiverId: msg.GetTarget(),
		MsgType:    uint32(msg.GetMsgType()),
		Content:    string(recordutil.CommonMsgProtobufDataToJSONData(msg)),
		CreateTime: msg.GetDatetime(),
		Source:     string(recordutil.SourceJSONMarshal(msg)),
		Reference:  string(recordutil.ReferenceJSONMarshal(msg)),
	})
	if err != nil {
		tx.RollBack()
		return err
	}
	_, _, err = s.Repo.AppendPrivateMsgRelation(tx, &model.PrivateMsgRelation{
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
	_, _, err = s.Repo.AppendPrivateMsgRelation(tx, &model.PrivateMsgRelation{
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

func (s *ServiceContext) StoreGroupMessage(member string, msg *message.Message) error {
	_, _, err := s.Repo.AppendGroupMsgContent(&model.GroupMsgContent{
		Mid:        msg.GetMid(),
		Cid:        msg.GetCid(),
		SenderId:   msg.GetFrom(),
		ReceiverId: member,
		GroupId:    msg.GetTarget(),
		MsgType:    uint32(msg.GetMsgType()),
		Content:    string(recordutil.CommonMsgProtobufDataToJSONData(msg)),
		CreateTime: msg.GetDatetime(),
		Source:     string(recordutil.SourceJSONMarshal(msg)),
		Reference:  string(recordutil.ReferenceJSONMarshal(msg)),
	})
	return err
}

func (s *ServiceContext) StoreSignal(target string, seq int64, sig *signal.Signal) error {
	now := util.TimeNowUnixMilli()
	m := &model.SignalContent{
		Uid:        target,
		Seq:        seq,
		Type:       uint8(sig.GetType()),
		Content:    string(recordutil.SignalContentToJSONData(sig)),
		CreateTime: uint64(now),
		UpdateTime: uint64(now),
	}
	_, _, err := s.Repo.AppendSignalContent(m)
	return err
}
