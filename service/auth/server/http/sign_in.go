package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/auth/model"
)

// SignIn
// @Summary 注册一个新的 AppId
// @Author wyt@33.cn
// @Tags backend
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Param data formData model.SignInRequest true "body"
// @Success 200 {object} model.GeneralResponse{data=model.SignInResponse}
// @Router	/auth/sign-in [post]
func SignIn(c *gin.Context) {
	request := model.SignInRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("ShouldBindHeader"+err.Error()))
		return
	}
	request.ConfigFile, err = c.FormFile("configFile")
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("FormFile"+err.Error()))
		return
	}

	ret, err := svc.SignIn(&request)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, err)
}
