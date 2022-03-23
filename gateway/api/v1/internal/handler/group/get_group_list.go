package group

import (
	"github.com/gin-gonic/gin"
	logic "github.com/txchat/dtalk/gateway/api/v1/internal/logic/group"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	"github.com/txchat/dtalk/pkg/api"
)

// GetGroupListHandler
// @Summary 查询群列表
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetGroupListReq false "body"
// @Success 200 {object} types.GeneralResp{data=types.GetGroupListResp}
// @Router	/app/group-list [post]
func GetGroupListHandler(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &types.GetGroupListReq{}

		l := logic.NewGroupLogic(c, ctx)
		res, err := l.GetGroupList(req)

		c.Set(api.ReqResult, res)
		c.Set(api.ReqError, err)
	}
}
