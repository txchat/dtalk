// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	call "github.com/txchat/dtalk/app/gateway/chat/internal/handler/call"
	record "github.com/txchat/dtalk/app/gateway/chat/internal/handler/record"
	user "github.com/txchat/dtalk/app/gateway/chat/internal/handler/user"
	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AppAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/app/start-call",
					Handler: call.StartCallHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/app/reply-busy",
					Handler: call.ReplyBusyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/app/check-call",
					Handler: call.CheckCallHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/app/handle-call",
					Handler: call.HandleCallHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AppAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/app/record/revoke",
					Handler: record.RevokeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/app/record/focus",
					Handler: record.FocusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/app/record/sync-record",
					Handler: record.SyncHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/app/record/pri-chat-record",
					Handler: record.PrivateRecordHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/record/push",
					Handler: record.PushHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/record/push2",
					Handler: record.Push2Handler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AppAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/app/user/login",
					Handler: user.LoginHandler(serverCtx),
				},
			}...,
		),
	)
}
