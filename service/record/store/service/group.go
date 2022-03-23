package service

import (
	"context"
	"strconv"

	"github.com/txchat/dtalk/service/record/store/model"
	xproto "github.com/txchat/imparse/proto"
)

func (s *Service) AppendGroupMsg(p *xproto.Common) error {
	//获取所有群成员
	members, err := s.dao.AllGroupMembers(context.TODO(), p.Target)
	if err != nil {
		return err
	}
	var msgRelate = make([]*model.MsgRelation, len(members))
	for i, member := range members {
		var item *model.MsgRelation
		if member == p.From {
			//发送者
			item = &model.MsgRelation{
				Mid:        strconv.FormatInt(p.Mid, 10),
				OwnerUid:   p.From,
				OtherUid:   p.Target,
				Type:       model.Send,
				State:      model.Received,
				CreateTime: p.Datetime,
			}
		} else {
			item = &model.MsgRelation{
				Mid:        strconv.FormatInt(p.Mid, 10),
				OwnerUid:   member,
				OtherUid:   p.Target,
				Type:       model.Rev,
				State:      model.UnReceive,
				CreateTime: p.Datetime,
			}
		}
		msgRelate[i] = item
	}
	tx, err := s.dao.NewTx()
	if err != nil {
		return err
	}
	_, _, err = s.dao.AppendGroupMsgContent(tx, &model.MsgContent{
		Mid:        strconv.FormatInt(p.Mid, 10),
		Seq:        p.Seq,
		SenderId:   p.From,
		ReceiverId: p.Target,
		MsgType:    uint32(p.MsgType),
		Content:    string(model.ParseCommonMsg(p)),
		CreateTime: p.Datetime,
		Source:     string(model.ParseSource(p)),
		Reference:  string(model.ParseReference(p)),
	})
	if err != nil {
		tx.RollBack()
		return err
	}
	_, _, err = s.dao.AppendGroupMsgRelation(tx, msgRelate)
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

func (s *Service) StoreGroupMsg(pro *xproto.Common) error {
	//step 1.存数据库
	err := s.AppendGroupMsg(pro)
	if err != nil {
		s.log.Error().Err(err).Msg("AppendGroupMsg failed")
		return model.ErrConsumeRedo
	}
	return nil
}
