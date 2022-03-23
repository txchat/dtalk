package record

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/gateway/api/v1/internal/logic/record"
	"github.com/txchat/dtalk/gateway/api/v1/internal/model"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
)

// GetPriRecords
// @Summary 获得聊天记录
// @Author chy@33.cn
// @Tags record 消息模块
// @Accept       json
// @Produce      json
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body model.GetPriRecordsReq false "body"
// @Success 200 {object} model.GeneralResponse{data=model.GetPriRecordsResp}
// @Router	/app/pri-chat-record [post]
func GetPriRecords(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &model.GetPriRecordsReq{}
		if err := c.ShouldBind(req); err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}

		userId, ok := c.Get(api.Address)
		if !ok {
			c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
			return
		}
		req.FromId = userId.(string)

		if req.Mid == "" {
			req.Mid = "999999999999999999"
		}

		l := record.NewLogic(c.Request.Context(), ctx)
		data, err := l.GetPriRecord(req)
		c.Set(api.ReqResult, data)
		c.Set(api.ReqError, err)
	}
}

// SyncRecords
// @Summary 同步聊天记录
// @Author dld@33.cn
// @Tags record 消息模块
// @Accept       json
// @Produce      json
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body model.SyncRecordsReq false "body"
// @Success 200 {object} model.GeneralResponse{data=model.SyncRecordsResp}
// @Router	/app/sync-record [post]
func SyncRecords(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &model.SyncRecordsReq{}
		if err := c.ShouldBind(req); err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}

		userId, ok := c.Get(api.Address)
		if !ok {
			c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
			return
		}
		uid := userId.(string)

		l := record.NewLogic(c.Request.Context(), ctx)
		data, err := l.GetSyncRecord(uid, req)
		c.Set(api.ReqResult, data)
		c.Set(api.ReqError, err)
	}
}
