package service

import (
	"strconv"

	"github.com/txchat/dtalk/service/record/store/model"
	xproto "github.com/txchat/imparse/proto"
)

func (s *Service) AppendMsg(p *xproto.Common) error {
	tx, err := s.dao.NewTx()
	if err != nil {
		return err
	}
	_, _, err = s.dao.AppendMsgContent(tx, &model.MsgContent{
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
	_, _, err = s.dao.AppendMsgRelation(tx, &model.MsgRelation{
		Mid:        strconv.FormatInt(p.Mid, 10),
		OwnerUid:   p.From,
		OtherUid:   p.Target,
		Type:       model.Send,
		State:      model.Received,
		CreateTime: p.Datetime,
	})
	if err != nil {
		tx.RollBack()
		return err
	}
	_, _, err = s.dao.AppendMsgRelation(tx, &model.MsgRelation{
		Mid:        strconv.FormatInt(p.Mid, 10),
		OwnerUid:   p.Target,
		OtherUid:   p.From,
		Type:       model.Rev,
		State:      model.UnReceive,
		CreateTime: p.Datetime,
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

func (s *Service) PushMem(p *xproto.Common) error {
	fromVer, err := s.dao.IncMsgVersion(p.From)
	if err != nil {
		s.log.Warn().Err(err).Str("uid", p.From).Msg("PushMem IncMsgVersion failed")
	}
	toVer, err := s.dao.IncMsgVersion(p.Target)
	if err != nil {
		s.log.Warn().Err(err).Str("uid", p.Target).Msg("PushMem IncMsgVersion failed")
	}
	//store sender
	if fromVer != 0 {
		err = s.dao.AddRecordCache(p.From, fromVer, &model.MsgCache{
			Mid:        strconv.FormatInt(p.Mid, 10),
			Seq:        p.Seq,
			SenderId:   p.From,
			ReceiverId: p.Target,
			MsgType:    uint32(p.MsgType),
			Content:    string(model.ParseCommonMsg(p)),
			CreateTime: p.Datetime,
			Source:     string(model.ParseSource(p)),
			Reference:  string(model.ParseReference(p)),
			Prev:       0,
			Version:    fromVer,
		})
		if err != nil {
			s.log.Warn().Err(err).Str("uid", p.From).Uint64("version", fromVer).Msg("PushMem AddRecordCache failed")
		}
	}
	//store receiver
	if toVer != 0 {
		err = s.dao.AddRecordCache(p.Target, toVer, &model.MsgCache{
			Mid:        strconv.FormatInt(p.Mid, 10),
			Seq:        p.Seq,
			SenderId:   p.From,
			ReceiverId: p.Target,
			MsgType:    uint32(p.MsgType),
			Content:    string(model.ParseCommonMsg(p)),
			CreateTime: p.Datetime,
			Source:     string(model.ParseSource(p)),
			Reference:  string(model.ParseReference(p)),
			Prev:       0,
			Version:    toVer,
		})
		if err != nil {
			s.log.Warn().Err(err).Str("uid", p.From).Uint64("version", toVer).Msg("PushMem AddRecordCache failed")
		}
	}
	return nil
}

func (s *Service) StoreMsg(pro *xproto.Common) error {
	//step 1.存数据库
	err := s.AppendMsg(pro)
	if err != nil {
		s.log.Error().Err(err).Msg("AppendMsg failed")
		return model.ErrConsumeRedo
	}
	if s.cfg.SyncCache {
		//step 2.存缓存队列
		err = s.PushMem(pro)
		if err != nil {
			s.log.Warn().Err(err).Msg("PushMem failed")
		}
	}
	return nil
}
