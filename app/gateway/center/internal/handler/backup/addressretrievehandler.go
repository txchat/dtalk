package backup

import (
	"net/http"

	"github.com/txchat/dtalk/app/gateway/center/internal/logic/backup"
	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddressRetrieveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddressRetrieveReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.Error(w, r, xerror.ErrInvalidParams)
			return
		}

		l := backup.NewAddressRetrieveLogic(r.Context(), svcCtx)
		resp, err := l.AddressRetrieve(&req)
		if err != nil {
			xhttp.Error(w, r, err)
		} else {
			xhttp.OkJSON(w, r, resp)
		}
	}
}