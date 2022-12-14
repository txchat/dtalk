package record

import (
	"net/http"

	"github.com/txchat/dtalk/app/gateway/chat/internal/logic/record"
	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PrivateRecordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PrivateRecordReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.Error(w, r, xerror.ErrInvalidParams)
			return
		}

		l := record.NewPrivateRecordLogic(r.Context(), svcCtx)
		resp, err := l.PrivateRecord(&req)
		if err != nil {
			xhttp.Error(w, r, err)
		} else {
			xhttp.OkJSON(w, r, resp)
		}
	}
}