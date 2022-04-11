package handler

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	zlog "github.com/rs/zerolog/log"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/txchat/dtalk/gateway/api/v1/docs"
	"github.com/txchat/dtalk/gateway/api/v1/internal/handler/account"
	"github.com/txchat/dtalk/gateway/api/v1/internal/handler/group"
	modules "github.com/txchat/dtalk/gateway/api/v1/internal/handler/modules"
	"github.com/txchat/dtalk/gateway/api/v1/internal/handler/record"
	test "github.com/txchat/dtalk/gateway/api/v1/internal/handler/test"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/pkg/api"
	"github.com/txchat/dtalk/pkg/api/logger"
	"github.com/txchat/dtalk/pkg/api/trace"
	xlog "github.com/txchat/dtalk/pkg/logger"
	otrace "github.com/txchat/im-pkg/trace"
)

var (
	serverCtx *svc.ServiceContext
	log       = zlog.Logger
)

func Init(ctx *svc.ServiceContext) *http.Server {
	serverCtx = ctx
	addr := serverCtx.Config().Server.Addr
	engine := Default()
	SetupEngine(engine)
	SetupGroupRoutes(engine)
	SetupResourceRoutes(engine)
	pprof.Register(engine)

	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("engineInner.Start() err")
			panic(err)
		}
	}()
	return srv
}

// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *gin.Engine {
	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(api.Chat33GinLogFormatter))
	router.Use(gin.Recovery())
	return router
}

// @title           即时通讯系统后端接口
// @version         1.0
// @description
// @termsOfService
// @contact.name
// @contact.url
// @contact.email
// @schemes   https
// @host      localhost:8080
// @BasePath  /api/v1
func SetupEngine(e *gin.Engine) *gin.Engine {
	root := e.Group("/", otrace.OpenTracingServerMiddleWare(), api.RespMiddleWare())
	root.POST("/test", test.GetHelloWord(serverCtx))

	userRoute := root.Group("/user")
	//获取服务器列表
	userRoute.Use(api.AuthMiddleWare())
	{
		userRoute.POST("/login", account.AddressLogin(serverCtx))
	}

	app := root.Group("/app")
	{
		modulesRoute := app.Group("/modules")
		//获取模块启用状态
		modulesRoute.POST("/all", modules.GetModulesHandler(serverCtx))

		recordRoute := app.Group("/record")
		{
			recordRoute.POST("/revoke", api.AuthMiddleWare(), record.RevokeHandler(serverCtx))
			recordRoute.POST("/focus", api.AuthMiddleWare(), record.FocusHandler(serverCtx))
			recordRoute.POST("/sync-record", api.AuthMiddleWare(), record.SyncRecords(serverCtx))
		}
	}

	recordRoute := root.Group("/record")
	recordRoute.Use(api.AuthMiddleWare())
	{
		recordRoute.POST("/push", record.PushToUid(serverCtx))
		recordRoute.POST("/push2", record.PushToUid2(serverCtx))
	}

	store := root.Group("/store/app", api.RespMiddleWare())
	store.Use(api.AuthMiddleWare())
	{
		store.POST("/pri-chat-record", record.GetPriRecords(serverCtx))
	}

	return e
}

func SetupGroupRoutes(e *gin.Engine) *gin.Engine {
	logMiddleware := logger.NewMiddleware(xlog.New(serverCtx.Config().Env, "group"))

	root := e.Group("/group/app", api.RespMiddleWare())
	root.Use(api.AuthMiddleWare(), trace.TraceMiddleware(), logMiddleware.Handle())
	{
		root.POST("/create-group", group.CreateGroupHandler(serverCtx))
		root.POST("/invite-group-members", group.InviteGroupMembersHandler(serverCtx))
		root.POST("/group-exit", group.GroupExitHandler(serverCtx))
		root.POST("/group-disband", group.GroupDisbandHandler(serverCtx))
		root.POST("/group-remove", group.GroupRemoveHandler(serverCtx))
		root.POST("/change-owner", group.ChangeOwnerHandler(serverCtx))
		root.POST("/join-group", group.JoinGroupHandler(serverCtx))

		root.POST("/name", group.UpdateGroupNameHandler(serverCtx))
		root.POST("/avatar", group.UpdateGroupAvatarHandler(serverCtx))
		root.POST("/joinType", group.UpdateGroupJoinTypeHandler(serverCtx))
		root.POST("/friendType", group.UpdateGroupFriendTypeHandler(serverCtx))
		root.POST("/muteType", group.UpdateGroupMuteTypeHandler(serverCtx))
		root.POST("/member/name", group.UpdateGroupMemberNameHandler(serverCtx))
		root.POST("/member/type", group.SetAdminHandler(serverCtx))
		root.POST("/member/muteTime", group.UpdateGroupMemberMuteTimeHandler(serverCtx))

		root.POST("/group-info", group.GetGroupInfoHandler(serverCtx))
		root.POST("/group-pub-info", group.GetGroupPubInfoHandler(serverCtx))
		root.POST("/group-list", group.GetGroupListHandler(serverCtx))
		root.POST("/group-member-list", group.GetGroupMemberListHandler(serverCtx))
		root.POST("/group-member-info", group.GetGroupMemberInfoHandler(serverCtx))
		root.POST("/mute-list", group.GetMuteListHandler(serverCtx))

	}
	return e
}

func SetupResourceRoutes(e *gin.Engine) *gin.Engine {
	// swagger 文档接口
	if serverCtx.Config().Env == "debug" {
		// todo : 单独开一个 swagger 服务
		e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return e
}
