package msgfactory

import (
	"context"
	"time"

	"github.com/txchat/dtalk/app/services/answer/internal/dao"
	"github.com/txchat/dtalk/app/services/answer/internal/model"
	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	"github.com/txchat/imparse"
)

type MsgCache struct {
	repo           dao.AnswerRepository
	idGenRPCClient generatorclient.Generator
}

func NewMsgCache(repo dao.AnswerRepository, idGenRPCClient generatorclient.Generator) *MsgCache {
	return &MsgCache{
		repo:           repo,
		idGenRPCClient: idGenRPCClient,
	}
}

func (mc *MsgCache) GetMsg(ctx context.Context, from, seq string) (*imparse.MsgIndex, error) {
	r, err := mc.repo.GetRecordSeqIndex(from, seq)
	if err != nil {
		return nil, err
	}
	if r == nil {
		return nil, nil
	}
	return &imparse.MsgIndex{
		Mid:        r.Mid,
		Seq:        r.Seq,
		SenderId:   r.SenderId,
		CreateTime: r.CreateTime,
	}, nil
}

func (mc *MsgCache) AddMsg(ctx context.Context, uid string, m *imparse.MsgIndex) error {
	return mc.repo.AddRecordSeqIndex(uid, &model.MsgIndex{
		Mid:        m.Mid,
		Seq:        m.Seq,
		SenderId:   m.SenderId,
		CreateTime: m.CreateTime,
	})
}

func (mc *MsgCache) GetMid(ctx context.Context) (id int64, err error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	reply, err := mc.idGenRPCClient.GetID(ctx, &generatorclient.GetIDReq{})
	return reply.GetId(), err
}
