package service

import (
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/backup/model"
)

func (s *Service) AddressEnrolment(btyAddr, btcAddr string) error {
	record, err := s.dao.QueryAddressEnrolment(btyAddr)
	if err != nil {
		return xerror.NewError(xerror.ExecFailed)
	}
	if record == nil {
		err := s.dao.CreateAddressEnrolment(&model.AddrMove{
			BtyAddr: btyAddr,
			BtcAddr: btcAddr,
			State:   0,
		})
		if err != nil {
			return xerror.NewError(xerror.ExecFailed)
		}
	}
	if record.BtcAddr != btcAddr {
		return xerror.NewError(xerror.ExecFailed).SetExtMessage("new address is not consistent")
	}
	return nil
}
