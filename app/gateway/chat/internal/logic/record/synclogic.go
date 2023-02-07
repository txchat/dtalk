package record

import (
	"context"
	"encoding/base64"

	"github.com/txchat/dtalk/app/services/storage/storageclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &SyncLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *SyncLogic) Sync(req *types.SyncReq) (resp *types.SyncResp, err error) {
	uid := l.custom.UID
	syncRecordResp, err := l.svcCtx.StorageRPC.GetSyncRecordsAfterMid(l.ctx, &storageclient.GetSyncRecordsAfterMidReq{
		Mid:   req.StartMid,
		Uid:   uid,
		Count: req.MaxCount,
	})
	if err != nil {
		return nil, xerror.ErrExec
	}

	records := make([]string, len(syncRecordResp.GetRecords()))
	for i, record := range syncRecordResp.GetRecords() {
		records[i] = base64.StdEncoding.EncodeToString(record)
	}

	resp = &types.SyncResp{
		RecordCount: len(syncRecordResp.GetRecords()),
		Records:     records,
	}
	return
}
