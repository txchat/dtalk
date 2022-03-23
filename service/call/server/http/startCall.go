package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/call/model"
)

// startCall
// @Summary 开始通话
// @Author chy@33.cn
// @Tags call
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body model.StartCallRequest false "body"
// @Success 200 {object} model.GeneralResponse{data=model.StartCallResponse}
// @Router	/app/start-call [post]
func startCall(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req := &model.StartCallRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.StartCall(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}
