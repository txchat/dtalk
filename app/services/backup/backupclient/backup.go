// Code generated by goctl. DO NOT EDIT!
// Source: backup.proto

package backupclient

import (
	"context"

	"github.com/txchat/dtalk/app/services/backup/backup"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddressInfo          = backup.AddressInfo
	QueryBindReq         = backup.QueryBindReq
	QueryBindReqEmail    = backup.QueryBindReqEmail
	QueryBindReqPhone    = backup.QueryBindReqPhone
	QueryBindResp        = backup.QueryBindResp
	QueryRelatedReq      = backup.QueryRelatedReq
	QueryRelatedReqPhone = backup.QueryRelatedReqPhone
	QueryRelatedResp     = backup.QueryRelatedResp

	Backup interface {
		QueryBind(ctx context.Context, in *QueryBindReq, opts ...grpc.CallOption) (*QueryBindResp, error)
		QueryRelated(ctx context.Context, in *QueryRelatedReq, opts ...grpc.CallOption) (*QueryRelatedResp, error)
	}

	defaultBackup struct {
		cli zrpc.Client
	}
)

func NewBackup(cli zrpc.Client) Backup {
	return &defaultBackup{
		cli: cli,
	}
}

func (m *defaultBackup) QueryBind(ctx context.Context, in *QueryBindReq, opts ...grpc.CallOption) (*QueryBindResp, error) {
	client := backup.NewBackupClient(m.cli.Conn())
	return client.QueryBind(ctx, in, opts...)
}

func (m *defaultBackup) QueryRelated(ctx context.Context, in *QueryRelatedReq, opts ...grpc.CallOption) (*QueryRelatedResp, error) {
	client := backup.NewBackupClient(m.cli.Conn())
	return client.QueryRelated(ctx, in, opts...)
}
