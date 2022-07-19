package http

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/pkg/api"
	"github.com/txchat/dtalk/service/discovery/service"
)

var (
	svc *service.Service
	log = log15.New("module", "discovery/http")
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

func SetupEngine(e *gin.Engine) *gin.Engine {
	//TODO 这边鉴权还是调用base
	//inner := e.Group("/inner")
	//inner.GET("/userInfo", UserInfo, api.RespMiddleWare())
	root := e.Group("/", api.RespMiddleWare())
	//获取服务器列表
	root.Any("/nodes", Nodes)
	return e
}
