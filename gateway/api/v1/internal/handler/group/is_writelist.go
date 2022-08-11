package group

import (
	"github.com/gin-gonic/gin"
	logic "github.com/txchat/dtalk/gateway/api/v1/internal/logic/group"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/pkg/api"
)

// IsWhitelistHandler
// @Summary 查询是否是白名单成员
// @Author dld@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Success 200 {object} types.GeneralResp{data=types.IsWhitelistResp}
// @Router	/group/app/is-whitelist [post]
func IsWhitelistHandler(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := logic.NewGroupLogic(c, ctx)
		res, err := l.IsWhitelist()

		c.Set(api.ReqResult, res)
		c.Set(api.ReqError, err)
	}
}
