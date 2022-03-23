package service

import (
	"github.com/txchat/dtalk/pkg/rand"
	"github.com/txchat/dtalk/service/group/model/db"
)

// MaintainGroupAESKey 给旧的群设置 aes key
func (s *Service) MaintainGroupAESKey() error {
	// 得到group列表
	groups, err := s.dao.GetAllGroupInfo()
	if err != nil {
		s.log.Error().Err(err).Msg("GetAllGroupInfo")
		return err
	}

	for _, group := range groups {
		if group.AESKey == "" {
			groupPo := &db.GroupInfo{
				GroupId:      group.GroupId,
				GroupName:    group.GroupName,
				GroupAESKey:  rand.NewAESKey256(),
				GroupPubName: group.GroupName,
			}
			_, _, err := s.dao.MaintainAESKeyAndPubName(groupPo)
			if err != nil {
				s.log.Error().Err(err).Interface("groupPo", groupPo).Msg("MaintainAESKeyAndPubName")
			}
		}
	}

	return nil
}
