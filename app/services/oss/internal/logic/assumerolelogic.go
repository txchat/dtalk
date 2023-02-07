package logic

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/oss/internal/svc"
	"github.com/txchat/dtalk/app/services/oss/oss"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssumeRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssumeRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssumeRoleLogic {
	return &AssumeRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssumeRoleLogic) AssumeRole(in *oss.AssumeRoleReq) (*oss.AssumeRoleResp, error) {
	engine, err := l.svcCtx.AppOssEngines.GetEngine(in.GetBase().GetAppId(), in.GetBase().GetOssType())
	if err != nil {
		return nil, xerror.ErrFeaturesUnSupported
	}

	assume, err := l.svcCtx.Repo.GetAssumeRole(in.GetBase().GetAppId(), in.GetBase().GetOssType(), engine.Config())
	if err == nil {
		return &oss.AssumeRoleResp{
			RequestId: assume.RequestId,
			Credentials: &oss.Credentials{
				AccessKeySecret: assume.Credentials.AccessKeySecret,
				Expiration:      assume.Credentials.Expiration,
				AccessKeyId:     assume.Credentials.AccessKeyId,
				SecurityToken:   assume.Credentials.SecurityToken,
			},
			AssumedRoleUser: &oss.AssumedRoleUser{
				AssumedRoleId: assume.AssumedRoleUser.AssumedRoleId,
				Arn:           assume.AssumedRoleUser.Arn,
			},
		}, nil
	}

	assume, err = engine.AssumeRole()
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.Repo.SaveAssumeRole(in.GetBase().GetAppId(), in.GetBase().GetOssType(), engine.Config(), assume)
	if err != nil {
		return nil, err
	}

	return &oss.AssumeRoleResp{
		RequestId: assume.RequestId,
		Credentials: &oss.Credentials{
			AccessKeySecret: assume.Credentials.AccessKeySecret,
			Expiration:      assume.Credentials.Expiration,
			AccessKeyId:     assume.Credentials.AccessKeyId,
			SecurityToken:   assume.Credentials.SecurityToken,
		},
		AssumedRoleUser: &oss.AssumedRoleUser{
			AssumedRoleId: assume.AssumedRoleUser.AssumedRoleId,
			Arn:           assume.AssumedRoleUser.Arn,
		},
	}, nil
}
