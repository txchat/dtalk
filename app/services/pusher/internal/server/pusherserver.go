// Code generated by goctl. DO NOT EDIT.
// Source: pusher.proto

package server

import (
	"context"

	"github.com/txchat/dtalk/app/services/pusher/internal/logic"
	"github.com/txchat/dtalk/app/services/pusher/internal/svc"
	"github.com/txchat/dtalk/app/services/pusher/pusher"
)

type PusherServer struct {
	svcCtx *svc.ServiceContext
	pusher.UnimplementedPusherServer
}

func NewPusherServer(svcCtx *svc.ServiceContext) *PusherServer {
	return &PusherServer{
		svcCtx: svcCtx,
	}
}

func (s *PusherServer) PushGroup(ctx context.Context, in *pusher.PushGroupReq) (*pusher.PushGroupResp, error) {
	l := logic.NewPushGroupLogic(ctx, s.svcCtx)
	return l.PushGroup(in)
}

func (s *PusherServer) PushList(ctx context.Context, in *pusher.PushListReq) (*pusher.PushListResp, error) {
	l := logic.NewPushListLogic(ctx, s.svcCtx)
	return l.PushList(in)
}
