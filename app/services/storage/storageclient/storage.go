// Code generated by goctl. DO NOT EDIT.
// Source: storage.proto

package storageclient

import (
	"context"

	"github.com/txchat/dtalk/app/services/storage/storage"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddRecordFocusReply    = storage.AddRecordFocusReply
	AddRecordFocusReq      = storage.AddRecordFocusReq
	DelRecordReply         = storage.DelRecordReply
	DelRecordReq           = storage.DelRecordReq
	GetChatSessionMsgReply = storage.GetChatSessionMsgReply
	GetChatSessionMsgReq   = storage.GetChatSessionMsgReq
	GetRecordReply         = storage.GetRecordReply
	GetRecordReq           = storage.GetRecordReq
	Record                 = storage.Record

	Storage interface {
		DelRecord(ctx context.Context, in *DelRecordReq, opts ...grpc.CallOption) (*DelRecordReply, error)
		GetRecord(ctx context.Context, in *GetRecordReq, opts ...grpc.CallOption) (*GetRecordReply, error)
		AddRecordFocus(ctx context.Context, in *AddRecordFocusReq, opts ...grpc.CallOption) (*AddRecordFocusReply, error)
		GetChatSessionMsg(ctx context.Context, in *GetChatSessionMsgReq, opts ...grpc.CallOption) (*GetChatSessionMsgReply, error)
	}

	defaultStorage struct {
		cli zrpc.Client
	}
)

func NewStorage(cli zrpc.Client) Storage {
	return &defaultStorage{
		cli: cli,
	}
}

func (m *defaultStorage) DelRecord(ctx context.Context, in *DelRecordReq, opts ...grpc.CallOption) (*DelRecordReply, error) {
	client := storage.NewStorageClient(m.cli.Conn())
	return client.DelRecord(ctx, in, opts...)
}

func (m *defaultStorage) GetRecord(ctx context.Context, in *GetRecordReq, opts ...grpc.CallOption) (*GetRecordReply, error) {
	client := storage.NewStorageClient(m.cli.Conn())
	return client.GetRecord(ctx, in, opts...)
}

func (m *defaultStorage) AddRecordFocus(ctx context.Context, in *AddRecordFocusReq, opts ...grpc.CallOption) (*AddRecordFocusReply, error) {
	client := storage.NewStorageClient(m.cli.Conn())
	return client.AddRecordFocus(ctx, in, opts...)
}

func (m *defaultStorage) GetChatSessionMsg(ctx context.Context, in *GetChatSessionMsgReq, opts ...grpc.CallOption) (*GetChatSessionMsgReply, error) {
	client := storage.NewStorageClient(m.cli.Conn())
	return client.GetChatSessionMsg(ctx, in, opts...)
}
