// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	backend "github.com/txchat/dtalk/app/gateway/center/internal/handler/backend"
	backup "github.com/txchat/dtalk/app/gateway/center/internal/handler/backup"
	"github.com/txchat/dtalk/app/gateway/center/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/disc/nodes",
				Handler: GetNodesHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AppParseHeaderMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/app/modules/all",
					Handler: GetModulesHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/backend/user/login",
				Handler: backend.LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/backup/phone-query",
				Handler: backup.QueryPhoneHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/backup/email-query",
				Handler: backup.QueryEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/backup/phone-retrieve",
				Handler: backup.PhoneRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/backup/email-retrieve",
				Handler: backup.EmailRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/backup/phone-export",
				Handler: backup.PhoneExportHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/backup/email-export",
				Handler: backup.EmailExportHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AppAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/backup/phone-send",
					Handler: backup.SendPhoneCodeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/backup/email-send",
					Handler: backup.SendEmailCodeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/backup/phone-binding",
					Handler: backup.PhoneBindingHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/backup/email-binding",
					Handler: backup.EmailBindingHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/backup/phone-relate",
					Handler: backup.PhoneRelateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/backup/address-retrieve",
					Handler: backup.AddressRetrieveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/backup/edit-mnemonic",
					Handler: backup.EditMnemonicHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/backup/get-address",
					Handler: backup.GetAddressHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AppParseHeaderMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/app/version/check",
					Handler: VersionCheckHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/backend/version/create",
				Handler: backend.CreateVersionHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/backend/version/update",
				Handler: backend.UpdateVersionHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/backend/version/change-status",
				Handler: backend.ChangeVersionStateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/backend/version/list",
				Handler: backend.ListVersionHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
