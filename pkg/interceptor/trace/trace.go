package trace

import (
	"context"
	"github.com/txchat/dtalk/pkg/api"
	"github.com/txchat/dtalk/pkg/api/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	// todo rename, oa那里也要一起改 X-Gateway-*
	OAOpeId   string = "X-OA-Ope-Id"
	OATraceId string = "X-OA-Trace-Id"
)

// NewTraceIdFromMD 从grpc metadata获取traceId
func NewTraceIdFromMD(md metadata.MD) (string, bool) {
	if md == nil {
		return "", false
	}

	var t string

	if traceId := md.Get(OATraceId); len(traceId) > 0 {
		t = traceId[0]
	} else {
		return "", false
	}

	return t, true
}

// NewAddressFromMD 从grpc metadata获取操作者 ID
func NewAddressFromMD(md metadata.MD) (string, bool) {
	if md == nil {
		return "", false
	}

	var t string

	if opeId := md.Get(OAOpeId); len(opeId) > 0 {
		t = opeId[0]
	} else {
		return "", false
	}

	return t, true
}

// ServerUnaryInterceptor 默认令牌服务端一元拦截器
func ServerUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return handler(wrapServerContext(ctx), req)
}

// wrapServerContext 包装服务端上下文
func wrapServerContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	token, ok := NewTraceIdFromMD(md)
	if !ok {
		// 如果没有就自己创一个 ID
		ctx = context.WithValue(ctx, trace.DtalkTraceId, trace.NewTraceIdWithContext(ctx))
	} else {
		ctx = context.WithValue(ctx, trace.DtalkTraceId, token)
	}

	address, ok := NewAddressFromMD(md)
	if !ok {
		return ctx
	}
	ctx = context.WithValue(ctx, api.Address, address)

	return ctx
}

// UnaryClientInterceptor 默认令牌客户端一元拦截器
func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return invoker(wrapClientContext(ctx), method, req, reply, cc, opts...)
}

// wrapClientContext 包装客户端上下文
func wrapClientContext(ctx context.Context) context.Context {
	t, ok := ctx.Value(trace.DtalkTraceId).(string)
	if ok {
		ctx = metadata.AppendToOutgoingContext(ctx, OATraceId, t)
	}

	token := api.NewAddrWithContext(ctx)
	if ok {
		ctx = metadata.AppendToOutgoingContext(ctx, OAOpeId, token)
	}

	return ctx
}
