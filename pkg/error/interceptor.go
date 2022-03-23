package error

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// WrapErr 返回gRPC状态码包装后的业务错误
func WrapErr(err error) error {
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case interface{ GRPCStatus() *status.Status }:
		return e.GRPCStatus().Err()
	case *Error:
		return status.Error(codes.Code(-e.Code()), e.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}

// ErrInterceptor 业务错误服务端一元拦截器
func ErrInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	return resp, WrapErr(err)
}

// ErrClientInterceptor 业务错误客户端一元拦截器
func ErrClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	return UnwrapErr(err)
}

// UnwrapErr 返回gRPC状态码解包后的业务错误
func UnwrapErr(err error) error {
	if err == nil {
		return nil
	}

	s, _ := status.FromError(err)
	c := int(s.Code())

	if _, ok := errorMsg[-c]; ok {
		return NewError(-c)
	}

	return err
}
