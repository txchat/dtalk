package group

import (
	"github.com/gin-gonic/gin"
	logic "github.com/txchat/dtalk/gateway/api/v1/internal/logic/group"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
)

// CreateNFTGroupHandler
// @Summary 创建藏品群
// @Author dld@33.cn
// @Tags group 群动作
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.CreateNFTGroupReq false "body"
// @Success 200 {object} types.GeneralResp{data=types.CreateNFTGroupResp}
// @Router	/group/app/create-nft-group [post]
func CreateNFTGroupHandler(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &types.CreateNFTGroupReq{}
		err := c.ShouldBind(req)
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}

		l := logic.NewGroupLogic(c, ctx)
		res, err := l.CreateNFTGroup(req)

		c.Set(api.ReqResult, res)
		c.Set(api.ReqError, err)
	}
}
