package exec

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/txchat/dtalk/app/services/storage/internal/logic"
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
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
	case chat.PrivateFrameType:
		pv := frame.(*chat.PrivateFrame)
		if pv.GetTransmissionMethod() != imparse.UniCast {
			return model.ErrFrameType
		}
		l := logic.NewStorePrivateMsgLogic(context.TODO(), e.svcCtx)
		return l.StoreMsg(deviceType, pv.GetBase())
	case chat.GroupFrameType:
		pv := frame.(*chat.GroupFrame)
		if pv.GetTransmissionMethod() != imparse.GroupCast {
			return model.ErrFrameType
		}
		l := logic.NewStoreGroupMsgLogic(context.TODO(), e.svcCtx)
		return l.StoreMsg(deviceType, pv.GetBase())
	case chat.SignalFrameType:
		pv := frame.(*chat.SignalFrame)
		base := pv.GetBase()
		if !base.GetReliable() {
			return nil
		}
		data := model.ParseSignal(base)
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