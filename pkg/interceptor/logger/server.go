package logger

import (
	"context"
	"strings"

	api "github.com/txchat/dtalk/pkg/api/logger"

	"github.com/txchat/dtalk/pkg/logger"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Logger interface {
	Info(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
}

type ServerInterceptor struct {
	log    zerolog.Logger
	filter []string
}

func NewServerInterceptor(log zerolog.Logger, filter []string) *ServerInterceptor {
	return &ServerInterceptor{
		log:    log,
		filter: filter,
	}
}

func (s *ServerInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//log.Infof("%s req: %v", info.FullMethod, req)
	for _, method := range s.filter {
		if strings.HasPrefix(info.FullMethod, method) {
			return handler(ctx, req)
		}
	}

	logt := logger.NewLogWithCtx(ctx, s.log)

	logt.Info().
		Str("Path", info.FullMethod).
		Interface("|body", req).
		Msg("rpc req")

	m, err := handler(ctx, req)

	if err != nil {
		//log.Error("%s err: %v", info.FullMethod, err)
		code, msg := api.ParseErr(err)
		if code != 0 {
			logt.Error().
				Int("code", code).
				Str("msg", msg).
				Msg("rpc err")
		}
	}
	//log.Info("%s resp: %v", info.FullMethod, m)
	return m, err
}
