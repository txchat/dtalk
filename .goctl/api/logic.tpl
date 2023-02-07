package {{.pkgName}}

import (
	{{.imports}}

    //xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
    //custom *xhttp.Custom
}

func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
    //c, ok := xhttp.FromContext(ctx)
    //if !ok {
    //    c = &xhttp.Custom{}
    //}
	return &{{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
        //custom: c,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
	// todo: add your logic here and delete this line

	{{.returnString}}
}
