package group

import (
	"net/http"

	"github.com/txchat/dtalk/app/gateway/chat/internal/logic/group"
	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetAdminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetAdminReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.Error(w, r, xerror.ErrInvalidParams)
			return
		}

		l := group.NewSetAdminLogic(r.Context(), svcCtx)
		resp, err := l.SetAdmin(&req)
		if err != nil {
			xhttp.Error(w, r, err)
		} else {
			xhttp.OkJSON(w, r, resp)
		}
	}
}
