package connmq

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/pusher/internal/config"
	innerLogic "github.com/txchat/dtalk/app/services/pusher/internal/logic"
	"github.com/txchat/dtalk/app/services/pusher/internal/model"
	"github.com/txchat/dtalk/app/services/pusher/internal/svc"
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
	"github.com/txchat/dtalk/pkg/util"
	comet "github.com/txchat/im/api/comet/grpc"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	"github.com/txchat/imparse/proto/auth"
	"github.com/txchat/imparse/proto/signal"
	"github.com/zeromicro/go-zero/core/logx"
)

type Service struct {
	logx.Logger
	Config        config.Config
	svcCtx        *svc.ServiceContext
	batchConsumer *xkafka.BatchConsumer
}

func NewService(cfg config.Config, svcCtx *svc.ServiceContext) *Service {
	s := &Service{
		Logger: logx.WithContext(context.TODO()),
		Config: cfg,
		svcCtx: svcCtx,
	}
	//topic config
	cfg.ConnDealConsumerConfig.Topic = fmt.Sprintf("goim-%s-topic", cfg.AppID)
	cfg.ConnDealConsumerConfig.Group = fmt.Sprintf("goim-%s-group", cfg.AppID)
	//new batch consumer
	consumer := xkafka.NewConsumer(cfg.ConnDealConsumerConfig, nil)
	logx.Info("dial kafka broker success")
	bc := xkafka.NewBatchConsumer(cfg.ConnDealBatchConsumerConf, xkafka.WithHandle(s.handleFunc), consumer)
	s.batchConsumer = bc
	return s
}

func (s *Service) Serve() {
	s.batchConsumer.Start()
}

func (s *Service) Shutdown(ctx context.Context) {
	s.batchConsumer.GracefulStop(ctx)
}

func (s *Service) handleFunc(key string, data []byte) error {
	bizMsg := new(logic.BizMsg)
	if err := proto.Unmarshal(data, bizMsg); err != nil {
		s.Error("logic.BizMsg proto.Unmarshal error", "err", err)
		return err
	}
	if bizMsg.AppId != s.Config.AppID {
		return model.ErrAppID
	}
	switch bizMsg.GetOp() {
	case int32(comet.Op_Auth), int32(comet.Op_Disconnect), int32(comet.Op_ReceiveMsgReply), int32(comet.Op_SyncMsgReq):
		if err := s.DealConn(context.TODO(), bizMsg); err != nil {
			//TODO redo consume message
			return err
		}
	default:
		return model.ErrCustomNotSupport
	}
	return nil
}

func (s *Service) DealConn(ctx context.Context, m *logic.BizMsg) error {
	switch m.Op {
	case int32(comet.Op_Auth):
		s.Info("user login with key")
		//将用户设备信息存入缓存
		dev, err := parseDevice(m)
		if err != nil {
			s.Error("parseDevice failed", "err", err)
		} else {
			now := uint64(util.TimeNowUnixNano() / int64(time.Millisecond))
			_, err = s.svcCtx.DeviceRPC.AddDevice(ctx, &deviceclient.DeviceInfo{
				Uid:         m.FromId,
				ConnectId:   m.Key,
				DeviceUUid:  dev.GetUuid(),
				DeviceType:  int32(dev.Device),
				DeviceName:  dev.GetDeviceName(),
				Username:    dev.Username,
				DeviceToken: dev.DeviceToken,
				IsEnabled:   false,
				AddTime:     now,
			})
			if err != nil {
				s.Error("AddDeviceInfo failed", "err", err)
			}
			//发送登录通知
			err = s.svcCtx.SignalHub.EndpointLogin(ctx, m.GetFromId(), &signal.SignalEndpointLogin{
				Uuid:       dev.GetUuid(),
				Device:     dev.Device,
				DeviceName: dev.GetDeviceName(),
				Datetime:   now,
			})
			if err != nil {
				s.Error("UniCastSignalEndpointLogin failed", "err", err)
				return err
			}
		}
		if dev != nil && (dev.Device == auth.Device_Android || dev.Device == auth.Device_IOS) {
			err = s.svcCtx.StoragePublish.BatchPush(ctx, m.Key, m.GetFromId())
			if err != nil {
				s.Error("BatchPushPublish failed", "err", err)
			}
		}
		//连接群聊
		l := innerLogic.NewJoinGroupsLogic(ctx, s.svcCtx)
		err = l.JoinGroups(m.GetFromId(), m.GetKey())
		if err != nil {
			s.Error("JoinGroups failed", "err", err)
		}
	case int32(comet.Op_Disconnect):
		s.Info("user logout with key")
		err := s.svcCtx.Repo.ClearConnSeq(m.Key)
		if err != nil {
			s.Error("ClearConnSeq failed", "err", err)
		}
		_, err = s.svcCtx.DeviceRPC.EnableThreadPushDevice(ctx, &deviceclient.EnableThreadPushDeviceRequest{
			Uid:    m.GetFromId(),
			ConnId: m.Key,
		})
		if err != nil {
			s.Error("EnableDeviceInfo failed", "err", err)
		}
	case int32(comet.Op_ReceiveMsgReply):
		var p comet.Proto
		err := proto.Unmarshal(m.Msg, &p)
		if err != nil {
			s.Error("unmarshal proto error", "err", err)
			return err
		}
		item, err := s.svcCtx.RecordHelper.GetLogsIndex(m.Key, p.Ack)
		if err != nil {
			s.Error("GetConnSeqIndex failed", "err", err, "ack", p.GetAck())
			return err
		}
		if item == nil {
			s.Error("message not find", "key", m.Key, "ack", p.Ack)
			return nil
		}
		dev, err := s.svcCtx.DeviceRPC.GetDeviceByConnId(ctx, &deviceclient.GetDeviceByConnIdRequest{
			Uid:    m.GetFromId(),
			ConnID: m.GetKey(),
		})
		if err != nil || dev == nil {
			s.Error("GetDeviceByConnID failed",
				"err", err,
				"ack", p.GetAck(),
				"uid", m.GetFromId(),
				"connID", m.GetKey(),
			)
		}
		if auth.Device(dev.GetDeviceType()) == auth.Device_Android || auth.Device(dev.GetDeviceType()) == auth.Device_IOS {
			err = s.svcCtx.StoragePublish.MarkRead(ctx, m.Key, m.GetFromId(), imparse.FrameType(item.Type), item.Logs)
			if err != nil {
				s.Error("MarkReadPublish failed",
					"err", err,
					"ack", p.GetAck(),
				)
			}
			switch item.Type {
			case string(chat.PrivateFrameType):
				err := s.svcCtx.SignalHub.MessageReceived(ctx, item)
				if err != nil {
					s.Error("UniCastSignalReceived failed",
						"err", err,
						"ack", p.GetAck(),
					)
					return err
				}
			default:
			}
		}
	case int32(comet.Op_SyncMsgReq):
		//TODO disabled
		return nil
		var p comet.Proto
		var pro comet.SyncMsg
		err := proto.Unmarshal(m.Msg, &p)
		if err != nil {
			s.Error("unmarshal proto error", "err", err)
			return err
		}
		err = proto.Unmarshal(p.Body, &pro)
		if err != nil {
			s.Error("Unmarshal failed",
				"err", err,
				"option", comet.Op_SyncMsgReq,
			)
			break
		}
		err = s.svcCtx.StoragePublish.Sync(ctx, m.Key, m.FromId, pro.LogId)
		if err != nil {
			s.Error("SyncMsg failed",
				"err", err,
				"start", pro.LogId,
			)
			break
		}
	default:
		return model.ErrCustomNotSupport
	}
	return nil
}

func parseDevice(m *logic.BizMsg) (*auth.Login, error) {
	var p comet.Proto
	err := proto.Unmarshal(m.Msg, &p)
	if err != nil {
		return nil, err
	}
	var authMsg comet.AuthMsg
	err = proto.Unmarshal(p.Body, &authMsg)
	if err != nil {
		return nil, err
	}
	if len(authMsg.Ext) == 0 {
		return nil, errors.New("ext is nil")
	}
	var device auth.Login
	err = proto.Unmarshal(authMsg.Ext, &device)
	if err != nil {
		return nil, err
	}
	return &device, nil
}
