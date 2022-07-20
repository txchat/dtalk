package dao

import (
	"context"

	"github.com/txchat/imparse"

	pusher "github.com/txchat/dtalk/service/record/pusher/api"
)

func (d Dao) CheckOnline(ctx context.Context, key string) (bool, error) {
	//TODO check client online
	return true, nil
}

func (d Dao) PushClient(ctx context.Context, key, from, mid, target string, tp imparse.Channel, frameType imparse.FrameType, data []byte) error {
	_, err := d.pusherCli.PushClient(ctx, &pusher.PushReq{
		Key:       key,
		From:      from,
		Mid:       mid,
		Target:    target,
		Data:      data,
		Type:      int32(tp),
		FrameType: string(frameType),
	})
	return err
}
