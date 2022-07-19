package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/service/device/config"
	"github.com/txchat/dtalk/service/device/dao"
	"github.com/txchat/dtalk/service/device/model"
)

type Service struct {
	log zerolog.Logger
	dao *dao.Dao
	cfg *config.Config
}

func New(c *config.Config) *Service {
	s := &Service{
		log: log.Logger,
		dao: dao.New(c),
		cfg: c,
	}
	return s
}

func (s *Service) Shutdown(ctx context.Context) {
	down := make(chan struct{})
	select {
	case <-ctx.Done():
		return
	case <-down:
		return
	}
}

func (s *Service) AddDevice(ctx context.Context, device *model.Device) error {
	return s.dao.AddDeviceInfo(device)
}

func (s *Service) EnableThreadPushDevice(ctx context.Context, uid, connId string) error {
	return s.dao.EnableDevice(uid, connId)
}

func (s *Service) GetUserAllDevices(ctx context.Context, uid string) ([]*model.Device, error) {
	return s.dao.GetAllDevices(uid)
}

func (s *Service) GetDeviceByConnId(ctx context.Context, uid, connID string) (*model.Device, error) {
	return s.dao.GetDevicesByConnID(uid, connID)
}
