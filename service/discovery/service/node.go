package service

import (
	"github.com/txchat/dtalk/service/discovery/model"
)

func (s *Service) StoreNodes(cNodes []*model.CNode, dNodes []*model.DNode) {
	for i := 0; i < len(cNodes); i++ {
		err := s.dao.SetCNode(cNodes[i].Name, cNodes[i])
		if err != nil {
			s.log.Error("init chat node failed", "err", err)
		}
	}
	for i := 0; i < len(dNodes); i++ {
		err := s.dao.SetDNode(dNodes[i].Name, dNodes[i])
		if err != nil {
			s.log.Error("init chain33 node failed", "err", err)
		}
	}
}

func (s *Service) CNodes() ([]*model.CNode, error) {
	return s.dao.GetCNodes()
}

func (s *Service) DNodes() ([]*model.DNode, error) {
	return s.dao.GetDNodes()
}
