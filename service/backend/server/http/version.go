package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/backend/model"
)

//get all nodes
//func CheckVersion(c *gin.Context) {
//	ret := svc.CheckVersion(c.MustGet(api.DeviceType).(string))
//	c.Set(api.ReqResult, ret)
//	c.Set(api.ReqError, nil)
//}

func CreateVersion(c *gin.Context) {
	request := model.VersionCreateRequest{}
	err := c.ShouldBind(&request)
	request.OpeUser = c.GetString("userName")
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("ShouldBind"+err.Error()))
		return
	}

	ret, err := svc.CreateVersion(&request)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, err)
}

func UpdateVersion(c *gin.Context) {
	request := model.VersionUpdateRequest{}
	err := c.ShouldBind(&request)
	request.OpeUser = c.GetString("userName")
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("ShouldBind"+err.Error()))
		return
	}

	ret, err := svc.UpdateVersion(&request)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, err)
}

func ChangeVersionStatus(c *gin.Context) {
	request := model.VersionChangeStatusRequest{}
	err := c.ShouldBind(&request)
	request.OpeUser = c.GetString("userName")
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("ShouldBind"+err.Error()))
		return
	}

	err = svc.ChangeVersionStatus(&request)
	c.Set(api.ReqResult, nil)
	c.Set(api.ReqError, err)

}

func GetVersionList(c *gin.Context) {
	request := model.GetVersionListRequest{}
	request.Platform = c.DefaultQuery("platform", "%")
	request.DeviceType = c.DefaultQuery("deviceType", "%")
	request.Page = util.MustToInt64(c.DefaultQuery("page", "0"))
	ret, err := svc.GetVersionList(&request)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, err)
}

func CheckAndUpdateVersion(c *gin.Context) {
	request := model.VersionCheckAndUpdateRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("ShouldBind"+err.Error()))
		return
	}
	deviceType, ok := c.Get(api.DeviceType)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.DeviceTypeError))
		return
	}
	request.DeviceType = deviceType.(string)

	ret, err := svc.CheckAndUpdateVersion(&request)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, err)

}

func GetToken(c *gin.Context) {
	request := model.GetTokenRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("ShouldBind"+err.Error()))
		return
	}

	ret, err := svc.GetToken(&request)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, err)

}
