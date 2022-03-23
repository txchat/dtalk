package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/backup/model"
)

//get all nodes
func QueryPhone(c *gin.Context) {
	var params struct {
		Area  string `json:"area"`
		Phone string `json:"phone" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	b, err := svc.PhoneIsBound(params.Area, params.Phone)
	ret := make(map[string]interface{})
	ret["exists"] = b
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, err)
}

func QueryEmail(c *gin.Context) {
	var params struct {
		Email string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	b, err := svc.EmailIsBound(params.Email)
	ret := make(map[string]interface{})
	ret["exists"] = b
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, err)
}

//@deprecated
func SendPhoneCode(c *gin.Context) {
	var params struct {
		Phone string `json:"phone" binding:"required"`
		//图形验证码所需
		//Ticket     string `json:"ticket"`
		//BusinessId string `json:"businessId"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	err := svc.SendPhoneCode(model.Quick, params.Phone)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.VerifyCodeSendError).SetExtMessage(err.Error()))
		return
	}
	c.Set(api.ReqError, nil)
}

func SendPhoneCodeV2(c *gin.Context) {
	var params struct {
		CodeType string `json:"codeType" binding:"required"`
		Phone    string `json:"phone" binding:"required"`
		//图形验证码所需
		//Ticket     string `json:"ticket"`
		//BusinessId string `json:"businessId"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	err := svc.SendPhoneCode(params.CodeType, params.Phone)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.VerifyCodeSendError).SetExtMessage(err.Error()))
		return
	}
	c.Set(api.ReqError, nil)
}

//@deprecated
func SendEmailCode(c *gin.Context) {
	var params struct {
		Email string `json:"email" binding:"required"`
		//图形验证码所需
		//Ticket     string `json:"ticket"`
		//BusinessId string `json:"businessId"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	err := svc.SendEmailCode(model.Quick, params.Email)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.VerifyCodeSendError).SetExtMessage(err.Error()))
		return
	}
	c.Set(api.ReqError, nil)
}

func SendEmailCodeV2(c *gin.Context) {
	var params struct {
		CodeType string `json:"codeType" binding:"required"`
		Email    string `json:"email" binding:"required"`
		//图形验证码所需
		//Ticket     string `json:"ticket"`
		//BusinessId string `json:"businessId"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	err := svc.SendEmailCode(params.CodeType, params.Email)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.VerifyCodeSendError).SetExtMessage(err.Error()))
		return
	}
	c.Set(api.ReqError, nil)
}

//@deprecated
func PhoneBinding(c *gin.Context) {
	var params struct {
		Area     string `json:"area"`
		Phone    string `json:"phone" binding:"required"`
		Code     string `json:"code" binding:"required"`
		Mnemonic string `json:"mnemonic" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	addr := c.MustGet(api.Address).(string)
	err := svc.PhoneBinding(addr, params.Area, params.Phone, params.Code, params.Mnemonic)
	if err != nil {
		c.Set(api.ReqError, err)
		return
	}

	ret := make(map[string]interface{})
	ret["address"] = c.MustGet(api.Address)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, nil)
}

func PhoneBindingV2(c *gin.Context) {
	var params struct {
		Area     string `json:"area"`
		Phone    string `json:"phone" binding:"required"`
		Code     string `json:"code" binding:"required"`
		Mnemonic string `json:"mnemonic" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	addr := c.MustGet(api.Address).(string)
	err := svc.PhoneBindingV2(addr, params.Area, params.Phone, params.Code, params.Mnemonic)
	if err != nil {
		c.Set(api.ReqError, err)
		return
	}

	ret := make(map[string]interface{})
	ret["address"] = c.MustGet(api.Address)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, nil)
}

func PhoneExport(c *gin.Context) {
	var params struct {
		Area    string `json:"area"`
		Phone   string `json:"phone" binding:"required"`
		Code    string `json:"code" binding:"required"`
		Address string `json:"address" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	b, err := svc.PhoneExport(params.Address, params.Area, params.Phone, params.Code)
	if err != nil {
		c.Set(api.ReqError, err)
		return
	}

	if !b {
		err = xerror.NewError(xerror.ExportAddressPhoneInconsistent)
	}
	c.Set(api.ReqResult, nil)
	c.Set(api.ReqError, err)
}

func PhoneRelate(c *gin.Context) {
	var params struct {
		Area     string `json:"area"`
		Phone    string `json:"phone" binding:"required"`
		Mnemonic string `json:"mnemonic" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	addr := c.MustGet(api.Address).(string)
	err := svc.PhoneRelate(addr, params.Area, params.Phone, params.Mnemonic)
	if err != nil {
		c.Set(api.ReqError, err)
		return
	}

	ret := make(map[string]interface{})
	ret["address"] = c.MustGet(api.Address)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, nil)
}

func EmailBinding(c *gin.Context) {
	var params struct {
		Email    string `json:"email" binding:"required"`
		Code     string `json:"code" binding:"required"`
		Mnemonic string `json:"mnemonic" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	addr := c.MustGet(api.Address).(string)

	err := svc.EmailBinding(addr, params.Email, params.Code, params.Mnemonic)
	if err != nil {
		c.Set(api.ReqError, err)
		return
	}

	ret := make(map[string]interface{})
	ret["address"] = c.MustGet(api.Address)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, nil)
}

func EmailBindingV2(c *gin.Context) {
	var params struct {
		Email    string `json:"email" binding:"required"`
		Code     string `json:"code" binding:"required"`
		Mnemonic string `json:"mnemonic" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	addr := c.MustGet(api.Address).(string)

	err := svc.EmailBindingV2(addr, params.Email, params.Code, params.Mnemonic)
	if err != nil {
		c.Set(api.ReqError, err)
		return
	}

	ret := make(map[string]interface{})
	ret["address"] = c.MustGet(api.Address)
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, nil)
}

func EmailExport(c *gin.Context) {
	var params struct {
		Email   string `json:"email" binding:"required"`
		Code    string `json:"code" binding:"required"`
		Address string `json:"address" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	b, err := svc.EmailExport(params.Address, params.Email, params.Code)
	if err != nil {
		c.Set(api.ReqError, err)
		return
	}
	if !b {
		err = xerror.NewError(xerror.ExportAddressEmailInconsistent)
	}
	c.Set(api.ReqResult, nil)
	c.Set(api.ReqError, err)
}

func PhoneRetrieve(c *gin.Context) {
	var params struct {
		Area  string `json:"area"`
		Phone string `json:"phone" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	//检查手机号
	item, err := svc.PhoneRetrieve(params.Area, params.Phone, params.Code)
	if err != nil {
		c.Set(api.ReqError, err)
		return
	}
	c.Set(api.ReqResult, item)
	c.Set(api.ReqError, nil)
}

func EmailRetrieve(c *gin.Context) {
	var params struct {
		Email string `json:"email" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	//检查手机号
	item, err := svc.EmailRetrieve(params.Email, params.Code)
	if err != nil {
		c.Set(api.ReqError, err)
		return
	}
	c.Set(api.ReqResult, item)
	c.Set(api.ReqError, nil)
}

func AddressRetrieve(c *gin.Context) {
	addr := c.MustGet(api.Address).(string)
	//检查手机号
	item, err := svc.AddressRetrieve(addr)
	if err != nil {
		c.Set(api.ReqError, err)
		return
	}
	c.Set(api.ReqResult, item)
	c.Set(api.ReqError, nil)
}

func EditMnemonic(c *gin.Context) {
	var params struct {
		Mnemonic string `json:"mnemonic" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	addr := c.MustGet(api.Address).(string)

	err := svc.EditMnemonic(addr, params.Mnemonic)
	c.Set(api.ReqError, err)
}

// @Summary 通过手机或邮箱得到地址
// @Author chy@33.cn
// @Tags backup
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body model.GetAddressRequest false "body"
// @Success 200 {object} model.GeneralResponse{data=model.GetAddressResponse}
// @Router	/get-address [post]
func GetAddress(c *gin.Context) {
	req := &model.GetAddressRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	res, err := svc.GetAddress(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}
