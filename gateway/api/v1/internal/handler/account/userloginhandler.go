package account

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/gateway/api/v1/internal/logic/account"
	"github.com/txchat/dtalk/gateway/api/v1/internal/model"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	xproto "github.com/txchat/imparse/proto"
)

// AddressLogin
// @Summary 用户登录
// @Description 内部接口,comet层使用
// @Author dld@33.cn
// @Tags account 账户模块
// @Accept       json
// @Produce      json
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Success 200 {object} model.GeneralResponse{data=model.AddressLoginResp}
// @Router	/user/login [post]
func AddressLogin(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &model.AddressLoginReq{}
		if err := c.ShouldBind(req); err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}

		uuid := c.MustGet(api.Uuid).(string)
		deviceType := c.MustGet(api.DeviceType).(string)
		//deviceName := c.MustGet(api.DeviceName)
		log.Debug().Int32("connType", req.ConnType).Str("uuid", uuid).Str("deviceType", deviceType).Msg("AddressLogin params")

		if xproto.Login_ConnType(req.ConnType) == xproto.Login_Reconnect {
			l := account.NewLogic(c.Request.Context(), ctx)
			dev, err := l.GetConflictDevice(c.MustGet(api.Address).(string), deviceType, uuid)
			if err != nil {
				c.Set(api.ReqError, xerror.NewError(xerror.CodeInnerError).SetExtMessage(err.Error()))
				return
			}
			if dev != nil {
				c.Set(api.ReqError, xerror.NewError(xerror.ReconnectNotAllowed).SetChildErr("", xerror.ReconnectNotAllowed, &model.AddressLoginNotAllowedErr{
					UUid:       dev.GetDeviceUUid(),
					Device:     dev.GetDeviceType(),
					DeviceName: dev.GetDeviceName(),
					Datetime:   dev.GetAddTime(),
				}))
				return
			}
		}

		c.Set(api.ReqResult, &model.AddressLoginResp{
			Address: c.MustGet(api.Address).(string),
		})
		c.Set(api.ReqError, nil)
	}
}
