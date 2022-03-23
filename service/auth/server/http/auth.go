package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/auth/model"
)

func Auth(c *gin.Context) {
	request := model.AuthRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("ShouldBind"+err.Error()))
		return
	}

	ret, err := svc.Auth(&request)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, err)
}
