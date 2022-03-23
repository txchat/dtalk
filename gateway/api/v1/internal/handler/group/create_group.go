package group

import (
	"github.com/gin-gonic/gin"
	logic "github.com/txchat/dtalk/gateway/api/v1/internal/logic/group"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
)

// CreateGroupHandler
// @Summary 创建群
// @Author chy@33.cn
// @Tags group 群动作
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.CreateGroupReq false "body"
// @Success 200 {object} types.GeneralResp{data=types.CreateGroupResp}
// @Router	/group/app/create-group [post]
func CreateGroupHandler(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &types.CreateGroupReq{}
		err := c.ShouldBind(req)
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}

		l := logic.NewGroupLogic(c, ctx)
		res, err := l.CreateGroup(req)

		c.Set(api.ReqResult, res)
		c.Set(api.ReqError, err)
	}
}
