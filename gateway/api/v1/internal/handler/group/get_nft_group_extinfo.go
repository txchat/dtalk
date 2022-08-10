package group

import (
	"github.com/gin-gonic/gin"
	logic "github.com/txchat/dtalk/gateway/api/v1/internal/logic/group"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
)

// GetNFTGroupExtInfoHandler
// @Summary 查询NFT群拓展信息
// @Author dld@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetNFTGroupExtInfoReq false "body"
// @Success 200 {object} types.GeneralResp{data=types.GetNFTGroupExtInfoResp}
// @Router	/group/app/nft-group-ext-info [post]
func GetNFTGroupExtInfoHandler(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &types.GetNFTGroupExtInfoReq{}
		err := c.ShouldBind(req)
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}

		l := logic.NewGroupLogic(c, ctx)
		res, err := l.GetNFTGroupExtInfo(req)

		c.Set(api.ReqResult, res)
		c.Set(api.ReqError, err)
	}
}
