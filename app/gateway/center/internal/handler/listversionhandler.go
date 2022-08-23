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

func ListVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListVersionReq
		if err := httpx.Parse(r, &req); err != nil {
			xcontext.Set(r, api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}

		req.Platform = svcCtx.Config.Backend.Platform
		l := logic.NewListVersionLogic(r.Context(), svcCtx)
		resp, err := l.ListVersion(&req)

		xcontext.Set(r, api.ReqResult, resp)
		xcontext.Set(r, api.ReqError, err)
	}
}
