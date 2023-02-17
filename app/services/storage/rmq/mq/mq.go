package rdmq

import (
	"bytes"
	"context"
	"fmt"

	"github.com/txchat/dtalk/app/services/storage/internal/exec"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/storage/internal/config"
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/internal/bizproto"
	record "github.com/txchat/dtalk/proto/record"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/proto/auth"
	xkafka "github.com/txchat/pkg/mq/kafka"
	"github.com/zeromicro/go-zero/core/logx"
)

type Service struct {
	logx.Logger
	Config        config.Config
	svcCtx        *svc.ServiceContext
	batchConsumer *xkafka.BatchConsumer
	storageExec   imparse.Storage
	parser        bizproto.StandardParse
}

func NewService(cfg config.Config, svcCtx *svc.ServiceContext) *Service {
	s := &Service{
		Logger:      logx.WithContext(context.TODO()),
		Config:      cfg,
		svcCtx:      svcCtx,
		storageExec: imparse.NewStandardStorage(exec.NewStorageExec(svcCtx)),
	}
	//topic config
	cfg.StoreDealConsumerConfig.Topic = fmt.Sprintf("biz-%s-store", cfg.AppID)
	cfg.StoreDealConsumerConfig.Group = fmt.Sprintf("biz-%s-store-storage", cfg.AppID)
	//new batch consumer
	consumer := xkafka.NewConsumer(cfg.StoreDealConsumerConfig, nil)
	logx.Info("dial kafka broker success")
	bc := xkafka.NewBatchConsumer(cfg.StoreDealBatchConsumerConf, xkafka.WithHandle(s.handleFunc), consumer)
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
	msg := new(record.StoreMsgMQ)
	if err := proto.Unmarshal(data, msg); err != nil {
		s.Error("logic.BizMsg proto.Unmarshal error", "err", err)
		return err
	}
	if msg.AppId != s.Config.AppID {
		return model.ErrAppID
	}
	if err := s.DealStore(context.TODO(), msg); err != nil {
		//TODO redo consume message
		return err
	}
	return nil
}

func (s *Service) DealStore(ctx context.Context, m *record.StoreMsgMQ) error {

	

	frame, err := s.parser.NewFrame(m.GetKey(), m.GetFromId(), bytes.NewReader(m.GetMsg()), bizproto.WithMid(m.GetMid()), bizproto.WithTarget(m.GetTarget()), bizproto.WithTransmissionMethod(imparse.Channel(m.GetType())))
	if err != nil {
		s.Error("NewFrame error", "key", m.Key, "from", m.FromId, "err", err)
		return err
	}
	//TODO 暂时处理，之后device信息统一通过answer服务传递
	devType := auth.Device_Android
	dev, err := s.svcCtx.DeviceRPC.GetDeviceByConnId(ctx, &deviceclient.GetDeviceByConnIdRequest{
		Uid:    m.GetFromId(),
		ConnID: m.GetKey(),
	})
	if dev != nil {
		devType = auth.Device(dev.GetDeviceType())
	}
	if err := s.storageExec.SaveMsg(ctx, devType, frame); err != nil {
		s.Error("Store error", "key", m.Key, "from", m.FromId, "err", err)
		return err
	}
	s.Slow("pass Store", "appId", m.AppId, "key", m.Key, "from", m.FromId)
	return nil
}
