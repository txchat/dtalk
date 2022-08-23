package logic

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if !l.svcCtx.UsersManager.IsMatch(req.UserName, req.Password) {
		return nil, xerror.NewError(xerror.ParamsError).SetExtMessage("用户名或密码错误")
	}
	// get token
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, req.UserName, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, err
	}
	resp = &types.LoginResp{UserInfo: types.UserInfo{
		UserName: req.UserName,
		Token:    token,
	}}
	return
}

func (l *LoginLogic) getJwtToken(secretKey, username string, iat, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	// custom payload
	claims["username"] = username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
