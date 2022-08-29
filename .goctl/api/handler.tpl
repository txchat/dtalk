package {{.PkgName}}

import (
	"net/http"

    xerror "github.com/txchat/dtalk/pkg/error"
    xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
            xhttp.Error(w, r, xerror.ErrInvalidParams)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        if err != nil {
            xhttp.Error(w, r, err)
        } else {
            {{if .HasResp}}xhttp.OkJSON(w, r, resp){{else}}httpx.Ok(w){{end}}
        }
	}
}
