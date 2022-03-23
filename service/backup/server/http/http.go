package http

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/pkg/api"
	"github.com/txchat/dtalk/service/backup/service"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/txchat/dtalk/service/backup/docs"
)

var (
	svc *service.Service
	log = log15.New("module", "backup/http")
)

func Init(s *service.Service) *http.Server {
	addr := s.Config().Server.Addr
	env := s.Config().Env
	gin.SetMode(env)
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

// @title backup
// @version 1.0
//@host 127.0.0.1:18004
func SetupEngine(e *gin.Engine) *gin.Engine {
	root := e.Group("/", api.RespMiddleWare())
	//获取服务器列表
	root.Any("/phone-query", QueryPhone)
	root.Any("/email-query", QueryEmail)

	//@deprecated
	root.POST("/phone-send", SendPhoneCode)
	root.POST("/email-send", SendEmailCode)
	root.POST("/phone-retrieve", PhoneRetrieve)
	root.POST("/email-retrieve", EmailRetrieve)
	//@deprecated end

	root.POST("/v2/phone-send", SendPhoneCodeV2)
	root.POST("/v2/email-send", SendEmailCodeV2)
	root.POST("/v2/phone-retrieve", PhoneRetrieve)
	root.POST("/v2/email-retrieve", EmailRetrieve)
	root.POST("/v2/phone-export", PhoneExport)
	root.POST("/v2/email-export", EmailExport)
	root.Use(api.AuthMiddleWare())
	{
		//@deprecated
		root.POST("/phone-binding", PhoneBinding)
		root.POST("/email-binding", EmailBinding)
		//@deprecated end

		root.POST("/v2/phone-binding", PhoneBindingV2)
		root.POST("/v2/email-binding", EmailBindingV2)
		root.POST("/phone-relate", PhoneRelate)
		root.POST("/address-retrieve", AddressRetrieve)
		root.POST("/edit-mnemonic", EditMnemonic)
		root.POST("/get-address", GetAddress)
	}

	move := root.Group("transform")
	move.Use(api.AuthMiddleWare())
	{
		move.POST("/addressEnrolment", AddressEnrolment)
	}

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e
}
