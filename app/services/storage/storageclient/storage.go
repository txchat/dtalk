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
	AddRecordFocusReply         = storage.AddRecordFocusReply
	AddRecordFocusReq           = storage.AddRecordFocusReq
	DelRecordReply              = storage.DelRecordReply
	DelRecordReq                = storage.DelRecordReq
	GetRecordReply              = storage.GetRecordReply
	GetRecordReq                = storage.GetRecordReq
	GetRecordsAfterMidReply     = storage.GetRecordsAfterMidReply
	GetRecordsAfterMidReq       = storage.GetRecordsAfterMidReq
	GetSyncRecordsAfterMidReply = storage.GetSyncRecordsAfterMidReply
	GetSyncRecordsAfterMidReq   = storage.GetSyncRecordsAfterMidReq

	Storage interface {
		DelRecord(ctx context.Context, in *DelRecordReq, opts ...grpc.CallOption) (*DelRecordReply, error)
		GetRecord(ctx context.Context, in *GetRecordReq, opts ...grpc.CallOption) (*GetRecordReply, error)
		AddRecordFocus(ctx context.Context, in *AddRecordFocusReq, opts ...grpc.CallOption) (*AddRecordFocusReply, error)
		GetRecordsAfterMid(ctx context.Context, in *GetRecordsAfterMidReq, opts ...grpc.CallOption) (*GetRecordsAfterMidReply, error)
		GetSyncRecordsAfterMid(ctx context.Context, in *GetSyncRecordsAfterMidReq, opts ...grpc.CallOption) (*GetSyncRecordsAfterMidReply, error)
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

func (m *defaultStorage) GetRecordsAfterMid(ctx context.Context, in *GetRecordsAfterMidReq, opts ...grpc.CallOption) (*GetRecordsAfterMidReply, error) {
	client := storage.NewStorageClient(m.cli.Conn())
	return client.GetRecordsAfterMid(ctx, in, opts...)
}

func (m *defaultStorage) GetSyncRecordsAfterMid(ctx context.Context, in *GetSyncRecordsAfterMidReq, opts ...grpc.CallOption) (*GetSyncRecordsAfterMidReply, error) {
	client := storage.NewStorageClient(m.cli.Conn())
	return client.GetSyncRecordsAfterMid(ctx, in, opts...)
}
