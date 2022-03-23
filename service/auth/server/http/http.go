package http

import (
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/txchat/dtalk/pkg/api"
	_ "github.com/txchat/dtalk/service/auth/docs"
	"github.com/txchat/dtalk/service/auth/service"
	"net/http"
)

var (
	svc *service.Service
	log = log15.New("module", "auth/http")
)

func Init(s *service.Service) *http.Server {
	addr := s.Config().Server.Addr
	engine := Default()
	InitService(s)
	SetupEngine(engine)

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

// SetupEngine
// @title auth服务接口
// @version 0.0.1
// @host 127.0.0.1:18103
func SetupEngine(e *gin.Engine) *gin.Engine {

	auth := e.Group("/auth", api.RespMiddleWare())
	{
		auth.POST("/sign-in", SignIn)
		auth.POST("/auth", Auth)
	}

	// swagger 文档接口
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e
}
