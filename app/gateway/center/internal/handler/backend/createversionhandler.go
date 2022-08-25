package backend

import (
	"net/http"

	xcontext "github.com/gorilla/context"
	xerror "github.com/txchat/dtalk/pkg/error"
	api "github.com/txchat/dtalk/pkg/newapi"

	"github.com/txchat/dtalk/app/gateway/center/internal/logic/backend"
	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateVersionReq
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
		req.Platform = svcCtx.Config.Backend.Platform
		l := backend.NewCreateVersionLogic(r.Context(), svcCtx)
		resp, err := l.CreateVersion(&req)

		xcontext.Set(r, api.ReqResult, resp)
		xcontext.Set(r, api.ReqError, err)
	}
}
