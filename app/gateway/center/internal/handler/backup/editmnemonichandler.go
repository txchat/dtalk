package backup

import (
	"net/http"

	xcontext "github.com/gorilla/context"
	xerror "github.com/txchat/dtalk/pkg/error"
	api "github.com/txchat/dtalk/pkg/newapi"

	"github.com/txchat/dtalk/app/gateway/center/internal/logic/backup"
	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditMnemonicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditMnemonicReq
		if err := httpx.Parse(r, &req); err != nil {
			xcontext.Set(r, api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}

		l := backup.NewEditMnemonicLogic(r.Context(), svcCtx)
		resp, err := l.EditMnemonic(&req)

		xcontext.Set(r, api.ReqResult, resp)
		xcontext.Set(r, api.ReqError, err)
	}
}
