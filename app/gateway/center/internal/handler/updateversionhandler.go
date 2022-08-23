package handler

import (
	"net/http"

	xcontext "github.com/gorilla/context"
	"github.com/txchat/dtalk/app/gateway/center/internal/logic"
	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"
	api "github.com/txchat/dtalk/pkg/newapi"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateVersionReq
		if err := httpx.Parse(r, &req); err != nil {
			xcontext.Set(r, api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}
		username, ok := api.GetStringOk(r, api.BackendJWTUsername)
		if !ok {
			xcontext.Set(r, api.ReqError, xerror.NewError(xerror.SignatureInvalid))
			return
		}
		req.OpeUser = username
		l := logic.NewUpdateVersionLogic(r.Context(), svcCtx)
		resp, err := l.UpdateVersion(&req)

		xcontext.Set(r, api.ReqResult, resp)
		xcontext.Set(r, api.ReqError, err)
	}
}
