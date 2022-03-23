package logic

import (
	"context"

	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
)

type TestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) TestLogic {
	return TestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) GetHelloWorld() (*types.Hello, error) {
	return &types.Hello{Content: "hello world"}, nil
}
