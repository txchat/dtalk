package {{.PkgName}}

import (
	"net/http"

	xcontext "github.com/gorilla/context"
    xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/newapi"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
            xcontext.Set(r, api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})

		xcontext.Set(r, api.ReqResult, resp)
        xcontext.Set(r, api.ReqError, err)
	}
}
