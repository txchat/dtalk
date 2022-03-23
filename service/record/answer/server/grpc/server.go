package grpc

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	pb "github.com/txchat/dtalk/service/record/answer/api"
	"github.com/txchat/dtalk/service/record/answer/service"
	"github.com/txchat/im-pkg/trace"
	"github.com/txchat/imparse"
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
	pb.RegisterAnswerServer(ws.Server(), &server{pb.UnimplementedAnswerServer{}, svr})
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
	pb.UnimplementedAnswerServer
	svr *service.Service
}

func (s *server) PushCommonMsg(ctx context.Context, req *pb.PushCommonMsgReq) (*pb.PushCommonMsgReply, error) {
	mid, createTime, err := s.svr.Push(ctx, req.GetKey(), req.GetFrom(), req.GetBody())
	if err != nil {
		return &pb.PushCommonMsgReply{}, err
	}
	return &pb.PushCommonMsgReply{
		Mid:  mid,
		Time: createTime,
	}, nil
}

func (s *server) PushNoticeMsg(ctx context.Context, req *pb.PushNoticeMsgReq) (*pb.PushNoticeMsgReply, error) {
	data, err := noticeMsgData(req.ChannelType, req.From, req.Target, req.Seq, req.Data)
	if err != nil {
		return &pb.PushNoticeMsgReply{}, err
	}
	mid, err := s.svr.InnerPush(ctx, "", req.From, req.Target, imparse.Undefined, data)
	if err != nil {
		return &pb.PushNoticeMsgReply{}, err
	}
	return &pb.PushNoticeMsgReply{
		Mid: mid,
	}, nil
}

func (s *server) UniCastSignal(ctx context.Context, req *pb.UniCastSignalReq) (*pb.UniCastSignalReply, error) {
	data, err := signalBody(req.GetTarget(), req.GetType(), req.GetBody())
	if err != nil {
		return &pb.UniCastSignalReply{}, err
	}
	mid, err := s.svr.InnerPush(ctx, "", "", req.GetTarget(), imparse.UniCast, data)
	if err != nil {
		return &pb.UniCastSignalReply{}, err
	}
	return &pb.UniCastSignalReply{
		Mid: mid,
	}, nil
}

func (s *server) GroupCastSignal(ctx context.Context, req *pb.GroupCastSignalReq) (*pb.GroupCastSignalReply, error) {
	data, err := signalBody(req.GetTarget(), req.GetType(), req.GetBody())
	if err != nil {
		return &pb.GroupCastSignalReply{}, err
	}
	mid, err := s.svr.InnerPush(ctx, "", "", req.GetTarget(), imparse.GroupCast, data)
	if err != nil {
		return &pb.GroupCastSignalReply{}, err
	}
	return &pb.GroupCastSignalReply{
		Mid: mid,
	}, nil
}
