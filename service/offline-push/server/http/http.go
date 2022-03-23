package http

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Init() *http.Server {
	engine := gin.Default()
	pprof.Register(engine)
	srv := &http.Server{
		Addr:    ":0",
		Handler: engine,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	return srv
}
