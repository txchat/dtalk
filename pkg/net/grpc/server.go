package grpc

import (
	"context"
	"math"
	"net"
	xtime "time"

	"github.com/txchat/dtalk/pkg/time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var (
	_abortIndex int8 = math.MaxInt8 / 2
)

// ServerConfig is rpc server conf.
type ServerConfig struct {
	// Network is grpc listen network,default value is tcp
	Network string `dsn:"network"`
	// Addr is grpc listen addr,default value is 0.0.0.0:9000
	Addr string `dsn:"address"`
	// Timeout is context timeout for per rpc call.
	Timeout time.Duration `dsn:"query.timeout"`
	// IdleTimeout is a duration for the amount of time after which an idle connection would be closed by sending a GoAway.
	// Idleness duration is defined since the most recent time the number of outstanding RPCs became zero or the connection establishment.
	KeepAliveMaxConnectionIdle time.Duration `dsn:"query.idleTimeout"`
	// MaxLifeTime is a duration for the maximum amount of time a connection may exist before it will be closed by sending a GoAway.
	// A random jitter of +/-10% will be added to MaxConnectionAge to spread out connection storms.
	KeepAliveMaxConnectionAge time.Duration `dsn:"query.maxLife"`
	// ForceCloseWait is an additive period after MaxLifeTime after which the connection will be forcibly closed.
	KeepAliveMaxMaxConnectionAgeGrace time.Duration `dsn:"query.closeWait"`
	// KeepAliveInterval is after a duration of this time if the server doesn't see any activity it pings the client to see if the transport is still alive.
	KeepAliveTime time.Duration `dsn:"query.keepaliveInterval"`
	// KeepAliveTimeout  is After having pinged for keepalive check, the server waits for a duration of Timeout and if no activity is seen even after that
	// the connection is closed.
	KeepAliveTimeout time.Duration `dsn:"query.keepaliveTimeout"`
}

func NewServer(conf *ServerConfig, opt ...grpc.ServerOption) *Server {
	if conf == nil {
		//TODO 远程读取
		panic("no config")
	}

	s := new(Server)
	s.conf = conf

	keepParam := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     xtime.Duration(s.conf.KeepAliveMaxConnectionIdle),
		MaxConnectionAgeGrace: xtime.Duration(s.conf.KeepAliveMaxMaxConnectionAgeGrace),
		Time:                  xtime.Duration(s.conf.KeepAliveTime),
		Timeout:               xtime.Duration(s.conf.KeepAliveTimeout),
		MaxConnectionAge:      xtime.Duration(s.conf.KeepAliveMaxConnectionAge),
	})
	opt = append(opt, keepParam)
	s.server = grpc.NewServer(opt...)
	return s
}

type Server struct {
	conf   *ServerConfig
	server *grpc.Server

	handlers []grpc.UnaryServerInterceptor
}

// Server return the grpc server for registering service.
func (s *Server) Server() *grpc.Server {
	return s.server
}

// Use attachs a global inteceptor to the server.
// For example, this is the right place for a rate limiter or error management inteceptor.
func (s *Server) Use(handlers ...grpc.UnaryServerInterceptor) *Server {
	finalSize := len(s.handlers) + len(handlers)
	if finalSize >= int(_abortIndex) {
		panic("warden: server use too many handlers")
	}
	mergedHandlers := make([]grpc.UnaryServerInterceptor, finalSize)
	copy(mergedHandlers, s.handlers)
	copy(mergedHandlers[len(s.handlers):], handlers)
	s.handlers = mergedHandlers
	return s
}

// Start create a new goroutine run server with configured listen addr
// will panic if any error happend
// return server itself
func (s *Server) Start() (*Server, error) {
	lis, err := net.Listen(s.conf.Network, s.conf.Addr)
	if err != nil {
		return nil, err
	}
	reflection.Register(s.server)
	go func() {
		if err := s.server.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return s, nil
}

// Shutdown stops the server gracefully. It stops the server from
// accepting new connections and RPCs and blocks until all the pending RPCs are
// finished or the context deadline is reached.
func (s *Server) Shutdown(ctx context.Context) (err error) {
	ch := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(ch)
	}()
	select {
	case <-ctx.Done():
		s.server.Stop()
		err = ctx.Err()
	case <-ch:
	}
	return
}
