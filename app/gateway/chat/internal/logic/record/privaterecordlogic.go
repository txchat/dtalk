package record

import (
	"context"

	"github.com/txchat/dtalk/internal/recordutil"

	"github.com/txchat/dtalk/app/services/storage/storageclient"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/imparse/proto/common"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrivateRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewPrivateRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrivateRecordLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &PrivateRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *PrivateRecordLogic) PrivateRecord(req *types.PrivateRecordReq) (resp *types.PrivateRecordResp, err error) {
	operator := l.custom.UID
	getRecordResp, err := l.svcCtx.StorageRPC.GetRecordsAfterMid(l.ctx, &storageclient.GetRecordsAfterMidReq{
		Tp:     common.Channel_ToUser,
		Mid:    util.MustToInt64(req.Mid),
		From:   operator,
		Target: req.TargetId,
		Count:  req.RecordCount,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.PrivateRecordResp{
		RecordCount: len(getRecordResp.GetRecords()),
		Records:     toChatRecord(getRecordResp.GetRecords()),
	}
	return
}

// toChatRecord []MsgContent -> []Record , content json -> proto
func toChatRecord(records []*storageclient.GetRecordReply) []*types.Record {
	rlt := make([]*types.Record, len(records))
	for i, msg := range records {
		rlt[i] = &types.Record{
			Mid:        msg.Mid,
			Seq:        msg.Seq,
			FromId:     msg.SenderId,
			TargetId:   msg.ReceiverId,
			MsgType:    int32(msg.MsgType),
			Content:    recordutil.CommonMsgJSONDataToProtobuf(msg.MsgType, []byte(msg.Content)),
			CreateTime: msg.CreateTime,
		}
	}
	return rlt
}
