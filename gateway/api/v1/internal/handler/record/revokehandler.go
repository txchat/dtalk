package record

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/gateway/api/v1/internal/logic/record"
	"github.com/txchat/dtalk/gateway/api/v1/internal/model"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
)

// RevokeHandler
// @Summary 撤回消息
// @Description
// @Author dld@33.cn
// @Tags record 消息模块
// @Accept       json
// @Produce      json
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body model.RevokeMsgReq true "body"
// @Success 200 {object} model.GeneralResponse{}
// @Router	/app/record/revoke [post]
func RevokeHandler(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &model.RevokeMsgReq{}
		if err := c.ShouldBind(req); err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}
		var err error
		operator := c.MustGet(api.Address).(string)
		l := record.NewLogic(c.Request.Context(), ctx)
		switch req.Type {
		case model.Private:
			err = l.RevokePersonal(operator, req.Mid)
		case model.Group:
			err = l.RevokeGroup(operator, req.Mid)
		default:
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("undefined type"))
			return
		}
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.CodeInnerError).SetExtMessage(err.Error()))
			return
		}
		c.Set(api.ReqResult, nil)
		c.Set(api.ReqError, err)
	}
}
