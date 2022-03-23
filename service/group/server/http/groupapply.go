package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model/types"
)

// CreateGroupApply
// @Summary 创建加群审批
// @Author chy@33.cn
// @Tags group 群审批
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.CreateGroupApplyReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.CreateGroupApplyResp}
// @Router	/app/create-group-apply [post]
func CreateGroupApply(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.CreateGroupApplyReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.CreateGroupApplySvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// AcceptGroupApply
// @Summary 接受加群审批
// @Author chy@33.cn
// @Tags group 群审批
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.AcceptGroupApplyReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.AcceptGroupApplyResp}
// @Router	/app/accept-group-apply [post]
func AcceptGroupApply(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.AcceptGroupApplyReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.AcceptGroupApplySvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)

}

// RejectGroupApply
// @Summary 拒绝加群审批
// @Author chy@33.cn
// @Tags group 群审批
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.RejectGroupApplyReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.RejectGroupApplyResp}
// @Router	/app/reject-group-apply [post]
func RejectGroupApply(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.RejectGroupApplyReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.RejectGroupApplySvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)

}

// GetGroupApplyById
// @Summary 查询加群审批
// @Author chy@33.cn
// @Tags group 群审批
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetGroupApplyByIdReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupApplysResp}
// @Router	/app/get-group-apply [post]
func GetGroupApplyById(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.GetGroupApplyByIdReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.GetGroupApplyByIdSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)

}

// GetGroupApplys
// @Summary 查询加群审批列表
// @Author chy@33.cn
// @Tags group 群审批
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetGroupApplysReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupApplysResp}
// @Router	/app/get-group-applys [post]
func GetGroupApplys(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.GetGroupApplysReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.GetGroupApplysSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)

}
