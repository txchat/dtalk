package exec

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/txchat/dtalk/internal/recordutil"

	"github.com/txchat/dtalk/app/services/storage/internal/logic"
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/internal/bizproto"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/proto/auth"
)

type StorageExec struct {
	svcCtx *svc.ServiceContext
}

func NewStorageExec(svcCtx *svc.ServiceContext) *StorageExec {
	return &StorageExec{
		svcCtx: svcCtx,
	}
}

func (e *StorageExec) SaveMsg(deviceType auth.Device, frame imparse.Frame) error {
	switch frame.Type() {
	case bizproto.PrivateFrameType:
		pv := frame.(*bizproto.PrivateFrame)
		if pv.GetTransmissionMethod() != imparse.UniCast {
			return model.ErrFrameType
		}
		l := logic.NewStorePrivateMsgLogic(context.TODO(), e.svcCtx)
		return l.StoreMsg(deviceType, pv.GetBase())
	case bizproto.GroupFrameType:
		pv := frame.(*bizproto.GroupFrame)
		if pv.GetTransmissionMethod() != imparse.GroupCast {
			return model.ErrFrameType
		}
		l := logic.NewStoreGroupMsgLogic(context.TODO(), e.svcCtx)
		return l.StoreMsg(deviceType, pv.GetBase())
	case bizproto.SignalFrameType:
		pv := frame.(*bizproto.SignalFrame)
		base := pv.GetBase()
		if !base.GetReliable() {
			return nil
		}
		data := recordutil.SignalContentToJSONData(base)
		createTime := uint64(util.TimeNowUnixNano() / int64(time.Millisecond))
		l := logic.NewStoreSignalLogic(context.TODO(), e.svcCtx)
		switch pv.GetTransmissionMethod() {
		case imparse.UniCast:
			return l.AppendUniCastSignal(strconv.FormatInt(pv.GetMid(), 10), pv.GetTarget(), base.GetType(), data, createTime)
		case imparse.GroupCast:
			return l.AppendGroupCastSignal(strconv.FormatInt(pv.GetMid(), 10), pv.GetTarget(), base.GetType(), data, createTime)
		default:
			return fmt.Errorf("%v : frame type: %v, tarnsmission method : %v", model.ErrFrameType, frame.Type(), pv.GetTransmissionMethod())
		}
	}
	return fmt.Errorf("%v : frame type: %v, tarnsmission method : %v", model.ErrFrameType, frame.Type())
}
