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

func Push2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PushReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.Error(w, r, xerror.ErrInvalidParams)
			return
		}

		fh, err := xhttp.FromFile(r, "message")
		if err != nil {
			if err.Error() == "http: request body too large" {
				xhttp.Error(w, r, xerror.ErrOssFileTooBig)
			} else {
				xhttp.Error(w, r, err)
			}
			return
		}

		l := record.NewPush2Logic(r.Context(), svcCtx)
		resp, err := l.Push2(&req, fh)
		if err != nil {
			xhttp.Error(w, r, err)
		} else {
			xhttp.OkJSON(w, r, resp)
		}
	}
}