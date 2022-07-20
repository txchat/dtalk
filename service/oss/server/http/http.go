package http

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/pkg/api"
	"github.com/txchat/dtalk/service/oss/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/txchat/dtalk/service/oss/docs"
)

var (
	svc *service.Service
	log = log15.New("module", "oss/http")
)

func Init(s *service.Service) *http.Server {
	addr := s.Config().Server.Addr
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
// @title 云存储服务接口
// @version 1.0
// @host 127.0.0.1:18005
func SetupEngine(e *gin.Engine) *gin.Engine {
	root := e.Group("/", api.RespMiddleWare())
	//获取服务器列表
	root.Use(api.AuthMiddleWare())
	{
		root.POST("/get-token", GetOssToken)
		root.GET("/get-token", GetOssToken)
		root.POST("/get-huaweiyun-token", GetHuaweiyunOssToken)
		root.GET("/get-huaweiyun-token", GetHuaweiyunOssToken)
		root.POST("/upload", Upload)
		root.POST("/init-multipart-upload", InitMultipartUpload)
		root.POST("/upload-part", UploadPart)
		root.POST("/complete-multipart-upload", CompleteMultipartUpload)
		root.POST("/abort-multipart-upload", AbortMultipartUpload)
		root.POST("/get-host", GetHost)
	}

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e
}
