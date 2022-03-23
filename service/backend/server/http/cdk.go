package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/backend/model/types"
)

//userId, ok := c.Get(api.Address)
//	if !ok {
//		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
//		return
//	}
// @Summary 创建 CdkType
// @Author chy@33.cn
// @Tags Cdk 后台
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.CreateGroupRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.CreateGroupResponse}
// @Router	/app/create-group [post]

// CreateCdkTypeHandler
// @Summary 创建 CdkType
// @Author chy@33.cn
// @Tags Cdk 后台
// @Param data body types.CreateCdkTypeReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.CreateCdkTypeResp}
// @Router	/backend/cdk/create-cdk-type [post]
func CreateCdkTypeHandler(c *gin.Context) {
	req := &types.CreateCdkTypeReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	res, err := svc.CdkService.CreateCdkTypeSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// CreateCdksHandler
// @Summary 创建 Cdks
// @Author chy@33.cn
// @Tags Cdk 后台
// @Param data body types.CreateCdksReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.CreateCdksResp}
// @Router	/backend/cdk/create-cdks [post]
func CreateCdksHandler(c *gin.Context) {
	req := &types.CreateCdksReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	res, err := svc.CdkService.CreateCdksSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetCdkTypesHandler
// @Summary 分页获得 cdkType
// @Author chy@33.cn
// @Tags Cdk 查询
// @Param data body types.GetCdkTypesReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetCdkTypesResp}
// @Router	/backend/cdk/get-cdk-types [post]
func GetCdkTypesHandler(c *gin.Context) {
	req := &types.GetCdkTypesReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	res, err := svc.CdkService.GetCdkTypesSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetCdksHandler
// @Summary 分页获得 cdks
// @Author chy@33.cn
// @Tags Cdk 查询
// @Param data body types.GetCdksReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetCdksResp}
// @Router	/backend/cdk/get-cdks [post]
func GetCdksHandler(c *gin.Context) {
	req := &types.GetCdksReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	res, err := svc.CdkService.GetCdksSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// DeleteCdksHandler
// @Summary 删除 cdks
// @Author chy@33.cn
// @Tags Cdk 后台
// @Param data body types.DeleteCdksReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.DeleteCdksResp}
// @Router	/backend/cdk/delete-cdks [post]
func DeleteCdksHandler(c *gin.Context) {
	req := &types.DeleteCdksReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	res, err := svc.CdkService.DeleteCdksSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// DeleteCdkTypesHandler
// @Summary 删除 cdkTypes
// @Author chy@33.cn
// @Tags Cdk 后台
// @Param data body types.DeleteCdkTypesReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.DeleteCdkTypesResp}
// @Router	/backend/cdk/delete-cdk-types [post]
func DeleteCdkTypesHandler(c *gin.Context) {
	req := &types.DeleteCdkTypesReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	res, err := svc.CdkService.DeleteCdkTypesSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// ExchangeCdksHandler
// @Summary 兑换 cdks
// @Author chy@33.cn
// @Tags Cdk 后台
// @Param data body types.ExchangeCdksReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.ExchangeCdksResp}
// @Router	/backend/cdk/exchange-cdks [post]
func ExchangeCdksHandler(c *gin.Context) {
	req := &types.ExchangeCdksReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	res, err := svc.CdkService.ExchangeCdksSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// UpdateCdkTypeHandler
// @Summary 更新 cdkType
// @Author chy@33.cn
// @Tags Cdk 后台
// @Param data body types.UpdateCdkTypeReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.UpdateCdkTypeResp}
// @Router	/backend/cdk/update-cdk-type [post]
func UpdateCdkTypeHandler(c *gin.Context) {
	req := &types.UpdateCdkTypeReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	res, err := svc.CdkService.UpdateCdkTypeSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetCdksByUserIdHandler
// @Summary 分页获得一个人拥有的 cdks
// @Author chy@33.cn
// @Tags Cdk App
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetCdksByUserIdReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetCdksByUserIdResp}
// @Router	/app/cdk/get-cdks-by-user-id [post]
func GetCdksByUserIdHandler(c *gin.Context) {
	req := &types.GetCdksByUserIdReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.CdkService.GetCdksByUserIdSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetCdkTypeByCoinNameHandler
// @Summary 查询一个票券对应的 cdkType
// @Author chy@33.cn
// @Tags Cdk App
// @Param data body types.GetCdkTypeByCoinNameReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetCdkTypeByCoinNameResp}
// @Router	/app/cdk/get-cdk-type-by-coin-name [post]
func GetCdkTypeByCoinNameHandler(c *gin.Context) {
	req := &types.GetCdkTypeByCoinNameReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	res, err := svc.CdkService.GetCdkTypeByCoinNameSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// CreateCdkOrderHandler
// @Summary 创建兑换订单
// @Author chy@33.cn
// @Tags Cdk App
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.CreateCdkOrderReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.CreateCdkOrderResp}
// @Router	/app/cdk/create-cdk-order [post]
func CreateCdkOrderHandler(c *gin.Context) {
	req := &types.CreateCdkOrderReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.CdkService.CreateCdkOrderSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// DealCdkOrderHandler
// @Summary 处理兑换订单
// @Author chy@33.cn
// @Tags Cdk App
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.DealCdkOrderReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.DealCdkOrderResp}
// @Router	/app/cdk/deal-cdk-order [post]
func DealCdkOrderHandler(c *gin.Context) {
	req := &types.DealCdkOrderReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.CdkService.DealCdkOrderSvc(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}
