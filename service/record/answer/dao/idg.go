package dao

import (
	"context"
	"time"

	idgen "github.com/txchat/dtalk/service/generator/api"
)

// Receive receive a message.
func (d *Dao) GetMid(ctx context.Context) (id int64, err error) {
	var (
		req   idgen.Empty
		reply *idgen.GetIDReply
	)
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	reply, err = d.idGenRPCClient.GetID(ctx, &req)
	if err != nil {
		return
	}

	return reply.Id, nil
}
