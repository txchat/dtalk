package syncmq

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/storage/internal/config"
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
	record "github.com/txchat/dtalk/proto/record"
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
	cfg.SyncDealConsumerConfig.Topic = fmt.Sprintf("store-%s-topic", cfg.AppID)
	cfg.SyncDealConsumerConfig.Group = fmt.Sprintf("store-%s-storage-group", cfg.AppID)
	//new batch consumer
	consumer := xkafka.NewConsumer(cfg.SyncDealConsumerConfig, nil)
	logx.Info("dial kafka broker success")
	bc := xkafka.NewBatchConsumer(cfg.SyncDealBatchConsumerConf, xkafka.WithHandle(s.handleFunc), consumer)
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
	bizMsg := new(record.RecordDeal)
	if err := proto.Unmarshal(data, bizMsg); err != nil {
		s.Error("logic.BizMsg proto.Unmarshal error", "err", err)
		return err
	}
	if bizMsg.AppId != s.Config.AppID {
		return model.ErrAppID
	}
	err := s.DealSync(context.TODO(), bizMsg)
	if err == model.ErrConsumeRedo {
		//TODO redo consume message
	}
	return nil
}

func (s *Service) DealSync(ctx context.Context, m *record.RecordDeal) error {
	switch m.Opt {
	case record.Operation_BatchPush:
		s.Slow("user batch push", "uid", m.FromId, "key", m.Key)
		//检查用户ws是否保持连接
		if b, err := s.CheckOnline(context.TODO(), m.Key); !b {
			if err != nil {
				s.Error("CheckOnline failed", "uid", m.FromId, "key", m.Key, "err", err)
			}
			break
		}
		err := s.SendUnReadMsg(m.Key, m.FromId)
		if err != nil {
			s.Error("SendUnReadMsg failed", "uid", m.FromId, "key", m.Key, "err", err)
			break
		}
	case record.Operation_MarkRead:
		var mk record.Marked
		err := proto.Unmarshal(m.Msg, &mk)
		if err != nil {
			s.Error("unmarshal proto failed", "uid", m.FromId, "key", m.Key, "err", err)
			return err
		}
		for _, mid := range mk.GetMids() {
			s.Slow("msg received", "uid", m.FromId, "key", m.Key, "mid", mid)
			err = s.MarkReceived(mk.Type, m.GetFromId(), mid)
			if err != nil {
				s.Error("MarkMsgReceived failed", "uid", m.FromId, "key", m.Key, "mid", mid, "err", err)
			}
		}
	case record.Operation_SyncMsg:
		//disabled
		return nil
	default:
		return model.ErrCustomNotSupport
	}
	return nil
}
