package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
)

// 阶段一：登记各个地址关系
func AddressEnrolment(c *gin.Context) {
	var params struct {
		BtcAddress string `json:"btcAddress" binding:"required"`
		EthAddress string `json:"ethAddress" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	btyAddress := c.MustGet(api.Address).(string)
	err := svc.AddressEnrolment(btyAddress, params.BtcAddress)
	c.Set(api.ReqResult, nil)
	c.Set(api.ReqError, err)
}
