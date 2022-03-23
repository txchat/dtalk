package grpc

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	pb "github.com/txchat/dtalk/service/record/pusher/api"
	"github.com/txchat/dtalk/service/record/pusher/service"
	"github.com/txchat/im-pkg/trace"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
)

func New(c *xgrpc.ServerConfig, svr *service.Service, logger zerolog.Logger) *xgrpc.Server {
	connectionTimeout := grpc.ConnectionTimeout(time.Duration(c.Timeout))
	ws := xgrpc.NewServer(
		c,
		connectionTimeout,
		grpc.ChainUnaryInterceptor(
			trace.OpentracingServerInterceptor,
			OpentracingServerLogXInterceptor(logger, trace.TraceIDLogKey),
		),
	)
	pb.RegisterPusherServer(ws.Server(), &server{pb.UnimplementedPusherServer{}, svr})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

func OpentracingServerLogXInterceptor(logger zerolog.Logger, fieldKey string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		span := opentracing.SpanFromContext(ctx)
		var traceID string
		if jgSpan, ok := span.Context().(jaeger.SpanContext); ok {
			traceID = jgSpan.TraceID().String()
		}
		logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str(fieldKey, traceID)
		})
		logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("method", info.FullMethod)
		})
		return handler(ctx, req)
	}
}

type server struct {
	pb.UnimplementedPusherServer
	svr *service.Service
}

func (s *server) PushClient(ctx context.Context, req *pb.PushReq) (*pb.PushReply, error) {
	err := s.svr.PushClient(ctx, req.Key, req.From, req.Mid, req.Target, req.Type, req.FrameType, req.Data)
	return &pb.PushReply{}, err
}
