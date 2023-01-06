package oss

import (
	"net/http"

	"github.com/txchat/dtalk/app/gateway/chat/internal/logic/oss"
	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadPartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadPartReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.Error(w, r, xerror.ErrInvalidParams)
			return
		}

		fh, err := xhttp.FromFile(r, "file")
		if err != nil {
			if err.Error() == "http: request body too large" {
				xhttp.Error(w, r, xerror.ErrOssFileTooBig)
			} else {
				xhttp.Error(w, r, err)
			}
			return
		}

		l := oss.NewUploadPartLogic(r.Context(), svcCtx)
		resp, err := l.UploadPart(&req, fh)
		if err != nil {
			xhttp.Error(w, r, err)
		} else {
			xhttp.OkJSON(w, r, resp)
		}
	}
}
