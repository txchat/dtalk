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

// GetGroupInfoHandler
// @Summary 查询群信息
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetGroupInfoReq false "body"
// @Success 200 {object} types.GeneralResp{data=types.GetGroupInfoResp}
// @Router	/group/app/group-info [post]
func GetGroupInfoHandler(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &types.GetGroupInfoReq{}
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
			req.Id = util.ToInt64(req.IdStr)
		}

		l := logic.NewGroupLogic(c, ctx)
		res, err := l.GetPriGroupInfo(req)

		c.Set(api.ReqResult, res)
		c.Set(api.ReqError, err)
	}
}
