package http

import (
	"net/http"

	"github.com/gin-contrib/pprof"

	"github.com/txchat/dtalk/pkg/api/logger"
	"github.com/txchat/dtalk/pkg/api/trace"

	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/txchat/dtalk/pkg/api"
	_ "github.com/txchat/dtalk/service/group/docs"
	"github.com/txchat/dtalk/service/group/service"
)

var (
	svc *service.Service
	log = log15.New("module", "group/http")
)

func Init(s *service.Service) *http.Server {
	addr := s.Config().HttpServer.Addr
	if s.Config().Env != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := Default()
	InitService(s)
	SetupEngine(engine)
	pprof.Register(engine)

	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("", "engineInner.Start() error(%v)", err)
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
	//router.Use(gin.LoggerWithFormatter(api.Chat33GinLogFormatter))
	router.Use(gin.Recovery())
	return router
}

func InitService(s *service.Service) {
	svc = s
}

// SetupEngine
// @title 群服务接口
// @version 1.0
// @host 127.0.0.1:18011
func SetupEngine(e *gin.Engine) *gin.Engine {
	group := e.Group("/app", api.RespMiddleWare())
	logMiddleware := logger.NewMiddleware(svc.GetLog())
	group.Use(api.AuthMiddleWare(), trace.TraceMiddleware(), logMiddleware.Handle())
	{
		group.POST("/create-group", CreateGroup)
		group.POST("/invite-group-members", InviteGroupMembers)
		group.POST("/group-exit", GroupExit)
		group.POST("/group-disband", GroupDisband)
		group.POST("/group-remove", GroupRemove)
		group.POST("/change-owner", ChangeOwner)
		group.POST("/join-group", JoinGroup)

		group.POST("/name", UpdateGroupName)
		group.POST("/avatar", UpdateGroupAvatar)
		group.POST("/joinType", UpdateGroupJoinType)
		group.POST("/friendType", UpdateGroupFriendType)
		group.POST("/muteType", UpdateGroupMuteType)
		group.POST("/member/name", UpdateGroupMemberName)
		group.POST("/member/type", SetAdmin)
		group.POST("/member/muteTime", SetMembersMuteTime)

		//group.GET("/group/:id", GetGroupInfoById)
		group.POST("/group-info", GetGroupInfo)
		group.POST("/group-pub-info", GetGroupPubInfo)
		//group.GET("/groups", GetGroups)
		group.POST("/group-list", GetGroupList)
		group.POST("/group-search", GetGroupInfoByCondition)
		//group.GET("/group/:id/members", GetGroupMemberListByUri)
		group.POST("/group-member-list", GetGroupMemberList)
		//group.GET("/group/:id/member/:memberId", GetGroupMemberInfoByUri)
		group.POST("/group-member-info", GetGroupMemberInfo)
		group.POST("/mute-list", GetMuteList)

		group.POST("/create-group-apply", CreateGroupApply)
		group.POST("/accept-group-apply", AcceptGroupApply)
		group.POST("/reject-group-apply", RejectGroupApply)
		group.POST("/get-group-apply", GetGroupApplyById)
		group.POST("/get-group-applys", GetGroupApplys)
	}

	// prometheus 接口
	//e.GET("/prometheus", midware.PrometheusHandler())
	// swagger 文档接口
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e
}
