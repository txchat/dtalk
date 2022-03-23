package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/gateway/api/v1/internal/model"
	"github.com/txchat/dtalk/pkg/api"

	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
)

// GetModulesHandler
// @Summary 获取模块启用状态
// @Description
// @Author dld@33.cn
// @Tags startup 初始化模块
// @Accept       json
// @Produce      json
// @Success 200 {object} model.GeneralResponse{data=[]model.GetModuleResp}
// @Router	/app/modules/all [post]
func GetModulesHandler(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		modules := ctx.Config().Modules
		var ret = make([]model.GetModuleResp, len(modules))
		for i, v := range modules {
			ret[i] = model.GetModuleResp{
				Name:      v.Name,
				IsEnabled: v.IsEnabled,
				EndPoints: v.EndPoints,
			}
		}
		c.Set(api.ReqResult, ret)
		c.Set(api.ReqError, nil)
	}
}
