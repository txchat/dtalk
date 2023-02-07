package backend

import (
	"net/http"

	"github.com/txchat/dtalk/app/gateway/center/internal/logic/backend"
	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateVersionReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.Error(w, r, xerror.NewCustomError(xerror.ErrInvalidParams, err))
			return
		}

		l := backend.NewCreateVersionLogic(r.Context(), svcCtx)
		resp, err := l.CreateVersion(&req)
		if err != nil {
			xhttp.Error(w, r, err)
		} else {
			xhttp.OkJSON(w, r, resp)
		}
	}
}
