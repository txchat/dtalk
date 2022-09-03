package group

import (
	"github.com/gin-gonic/gin"
	logic "github.com/txchat/dtalk/gateway/api/v1/internal/logic/group"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
)

// GetMuteListHandler
// @Summary 查询群内被禁言成员名单
// @Author chy@33.cn
// @Tags group 禁言
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetMuteListReq false "body"
// @Success 200 {object} types.GeneralResp{data=types.GetMuteListResp}
// @Router	/app/mute-list [post]
func GetMuteListHandler(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &types.GetMuteListReq{}
		err := c.ShouldBind(req)
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}

		if req.Id == 0 && req.IdStr == "" {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
			return
		}

		if req.IdStr != "" {
			req.Id = util.MustToInt64(req.IdStr)
		}

		l := logic.NewGroupLogic(c, ctx)
		res, err := l.GetMuteList(req)

		c.Set(api.ReqResult, res)
		c.Set(api.ReqError, err)
	}
}
