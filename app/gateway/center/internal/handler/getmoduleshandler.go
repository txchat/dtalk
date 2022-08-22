package handler

import (
	"net/http"

	xcontext "github.com/gorilla/context"
	"github.com/txchat/dtalk/pkg/api"

	"github.com/txchat/dtalk/app/gateway/center/internal/logic"
	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetModulesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetModulesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetModulesLogic(r.Context(), svcCtx)
		resp, err := l.GetModules(&req)

		xcontext.Set(r, api.ReqResult, resp)
		xcontext.Set(r, api.ReqError, err)
	}
}
