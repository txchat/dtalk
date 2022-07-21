package service

import (
	"context"

	"github.com/txchat/dtalk/service/vip/dao"
	"github.com/txchat/dtalk/service/vip/model"
)

type Service struct {
	vipRepo dao.VIPRepository
}

func New(vipRepository dao.VIPRepository) *Service {
	s := &Service{
		vipRepo: vipRepository,
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

// GetVIP 查询VIP成员
func (s *Service) GetVIP(ctx context.Context, uid string) (*model.VIPEntity, error) {
	return s.vipRepo.GetVIP()
}

// AddVIP 添加VIP成员
func (s *Service) AddVIP(ctx context.Context, uid string) error {
	return s.vipRepo.AddVIP(&model.VIPEntity{UID: uid})
}

// GetScopeVIP 查询指定范围VIP成员
func (s *Service) GetScopeVIP(ctx context.Context, start, limit int32) ([]*model.VIPEntity, error) {
	return s.vipRepo.GetScopeVIP(start, limit)
}

// GetVIPCount 查询所有VIP数量
func (s *Service) GetVIPCount(ctx context.Context) (int32, error) {
	return s.vipRepo.GetVIPCount()
}
