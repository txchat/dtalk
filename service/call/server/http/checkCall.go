package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/call/model"
)

// checkCall
// @Summary 检查通话状态
// @Author chy@33.cn
// @Tags call
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body model.CheckCallRequest false "body"
// @Success 200 {object} model.GeneralResponse{data=model.CheckCallResponse}
// @Router	/app/check-call [post]
func checkCall(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req := &model.CheckCallRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.TraceId == 0 && req.TraceIdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.TraceIdStr != "" {
		traceId, err := util.ToInt64(req.TraceIdStr)
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}
		req.TraceId = traceId
	}

	req.PersonId = userId.(string)

	res, err := svc.CheckCall(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}
