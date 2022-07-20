package http

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/pkg/api"
	"github.com/txchat/dtalk/pkg/logger"
	"github.com/txchat/dtalk/service/call/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/txchat/dtalk/service/call/docs"
)

const srvName = "call/http"

var (
	svc *service.Service
	log zerolog.Logger
)

func Init(s *service.Service) *http.Server {
	addr := s.Config().HttpServer.Addr
	engine := defaultEngine()
	initService(s)
	setupEngine(engine)
	log = logger.New(s.Config().Env, srvName)
	pprof.Register(engine)
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("engineInner.Start()")
			panic(err)
		}
	}()
	return srv
}

// defaultEngine returns an Engine instance with the Logger and Recovery middleware already attached.
func defaultEngine() *gin.Engine {
	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(api.Chat33GinLogFormatter))
	router.Use(gin.Recovery())
	return router
}

func initService(s *service.Service) {
	svc = s
}

// setupEngine
// @title 音视频信令服务接口
// @version 1.0
// @host 127.0.0.1:18013
func setupEngine(e *gin.Engine) *gin.Engine {
	app := e.Group("/app", api.RespMiddleWare(), api.AuthMiddleWare())
	{
		app.POST("/start-call", startCall)
		app.POST("/reply-busy", replyBusy)
		app.POST("/check-call", checkCall)
		app.POST("/handle-call", handleCall)
	}

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e
}
