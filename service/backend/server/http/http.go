package http

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/txchat/dtalk/pkg/api"
	"github.com/txchat/dtalk/service/backend/config"
	_ "github.com/txchat/dtalk/service/backend/docs"
	"github.com/txchat/dtalk/service/backend/midware"
	"github.com/txchat/dtalk/service/backend/service"
)

var (
	svc *service.Service
	log = log15.New("module", "backend/http")
)

func Init(s *service.Service) *http.Server {
	addr := s.Config().Server.Addr
	if s.Config().Env != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := Default()
	InitService(s)
	SetupEngine(engine)

	if s.Config().CdkMod {
		setCdkRoute(engine)
	}
	pprof.Register(engine)
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("engineInner.Start() error(%v)", err)
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

func InitService(s *service.Service) {
	svc = s
}

func SetupEngine(e *gin.Engine) *gin.Engine {
	{
		app := e.Group("/app", api.RespMiddleWare())
		version := app.Group("/version")
		//获取服务器列表
		version.Use(api.HeaderMiddleWare())
		{
			version.POST("/check", CheckAndUpdateVersion)
		}
	}
	{
		backend := e.Group("/backend", api.RespMiddleWare())
		token := backend.Group("/user")
		{
			token.POST("/login", GetToken)
		}
		version := backend.Group("/version")
		version.Use(midware.JWTAuthMiddleWare(config.Conf.Debug.Flag, config.Conf.Release.Key))
		{
			version.POST("/create", CreateVersion)
			version.PUT("/update", UpdateVersion)
			version.PUT("/change-status", ChangeVersionStatus)
			version.GET("/list", GetVersionList)
		}
	}

	// swagger 文档接口
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e
}

func setCdkRoute(e *gin.Engine) *gin.Engine {
	{
		app := e.Group("/app", api.RespMiddleWare())

		// cdk app
		appCkd := app.Group("/cdk")
		appCkd.Use(api.AuthMiddleWare())
		{
			appCkd.POST("get-cdks-by-user-id", GetCdksByUserIdHandler)
			appCkd.POST("get-cdk-type-by-coin-name", GetCdkTypeByCoinNameHandler)
			appCkd.POST("create-cdk-order", CreateCdkOrderHandler)
			appCkd.POST("deal-cdk-order", DealCdkOrderHandler)
		}
	}
	{
		backend := e.Group("/backend", api.RespMiddleWare())

		// cdk backend
		bakCdk := backend.Group("/cdk")
		bakCdk.POST("get-cdk-types", GetCdkTypesHandler)
		bakCdk.POST("get-cdks", GetCdksHandler)
		bakCdk.Use(midware.JWTAuthMiddleWare(config.Conf.Debug.Flag, config.Conf.Release.Key))
		{
			bakCdk.POST("create-cdk-type", CreateCdkTypeHandler)
			bakCdk.POST("create-cdks", CreateCdksHandler)
			bakCdk.POST("delete-cdks", DeleteCdksHandler)
			bakCdk.POST("delete-cdk-types", DeleteCdkTypesHandler)
			bakCdk.POST("update-cdk-type", UpdateCdkTypeHandler)
			bakCdk.POST("exchange-cdks", ExchangeCdksHandler)
		}
	}
	return e
}
