package service

import (
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/call/model"
)

func (s *Service) HandleCall(req *model.HandleCallRequest) (res *model.HandleCallResponse, err error) {
	defer func() {
		if err != nil {
			s.log.Error().Err(err).Interface("req", req).Str("personId", req.PersonId).Msg("HandleCall")
		} else {
			s.log.Info().Interface("req", req).Str("personId", req.PersonId).Msg("HandleCall")
		}
	}()

	// 从 redis 中获得 session
	session, err := s.dao.GetSession(req.TraceId)
	if err != nil {
		return nil, xerror.NewError(xerror.CodeInnerError).SetExtMessage(err.Error())
	}

	// A 或 B 拒绝
	if req.Answer == false {
		if err := s.refuseCall(req.PersonId, session); err != nil {
			return &model.HandleCallResponse{}, xerror.NewError(xerror.CodeInnerError).SetExtMessage(err.Error())
		}
		return &model.HandleCallResponse{}, nil
	}

	// B 接受
	res, err = s.acceptCall(session)
	if err != nil {
		return &model.HandleCallResponse{}, xerror.NewError(xerror.CodeInnerError).SetExtMessage(err.Error())
	}
	return res, nil
}

// refuseCall 拒绝通话流程
func (s *Service) refuseCall(personId string, session *model.Session) error {
	// TODO 暂不支持群会议
	if session.GroupId != 0 {
		return model.ErrFeaturesUnSupported
	}

	if session.Status == model.READY {
		if personId == session.Caller {
			// 发起方主动取消
			// 给 对方 发 ActionStopCall with Cancel
			if err := s.noticeStopCall(session.Invitees[0], session.TraceId, model.Cancel); err != nil {
				s.log.Error().Err(err).Msg("refuseCall() noticeStopCall() 1")
			}
		} else {
			// 对方拒绝
			// 给 Caller 发 ActionStopCall with Reject
			if err := s.noticeStopCall(session.Caller, session.TraceId, model.Reject); err != nil {
				s.log.Error().Err(err).Msg("refuseCall() noticeStopCall() 2")
			}
		}
	} else {
		// 双方挂断
		// 给对方通知 ActionStopCall with Hangup
		if personId == session.Caller {
			if err := s.noticeStopCall(session.Invitees[0], session.TraceId, model.Hangup); err != nil {
				s.log.Error().Err(err).Msg("refuseCall() noticeStopCall() 1")
			}
		} else {
			if err := s.noticeStopCall(session.Caller, session.TraceId, model.Hangup); err != nil {
				s.log.Error().Err(err).Msg("refuseCall() noticeStopCall() 2")
			}
		}
	}

	session.Status = model.FINISH
	err := s.dao.SaveSession(session)
	if err != nil {
		return err
	}

	return nil
}

// acceptCall 接受通话流程, 只有被接收方可进行该流程
func (s *Service) acceptCall(session *model.Session) (*model.HandleCallResponse, error) {
	// TODO 暂不支持群聊
	if session.GroupId != 0 {
		return nil, model.ErrSessionNotExist
	}

	// TODO 判断是否在被接收方组内

	if session.Status != model.READY {
		return nil, model.ErrSessionNotExist
	}

	// 修改 session 状态并保存到 redis 中
	session.Status = model.INPROGRESS
	err := s.dao.SaveSession(session)
	if err != nil {
		return nil, err
	}

	// 获得 AppId
	sdkAppId := s.tlsSig.GetAppId()

	// 生成接收方的 userSig 和 privateMapKey
	userSigForInvitees, err := s.tlsSig.GetUserSig(session.Invitees[0])
	if err != nil {
		return nil, err
	}
	privateMapKeyForInvitees, err := s.tlsSig.GenPrivateMapKey(session.Invitees[0], session.RoomId, 255)
	if err != nil {
		return nil, err
	}

	// 生成发起方的 userSig 和 privateMapKey
	userSigForCaller, err := s.tlsSig.GetUserSig(session.Caller)
	if err != nil {
		return nil, err
	}
	privateMapKeyForCaller, err := s.tlsSig.GenPrivateMapKey(session.Caller, session.RoomId, 255)
	if err != nil {
		return nil, err
	}

	// 给Caller发 ActionAcceptCall
	if err = s.noticeAcceptCall(session.Caller, session.TraceId, session.RoomId, userSigForCaller, privateMapKeyForCaller, sdkAppId); err != nil {
		s.log.Error().Err(err).Msg("acceptCall() noticeAcceptCall()")
	}

	res := &model.HandleCallResponse{
		RoomId:        session.RoomId,
		UserSig:       userSigForInvitees,
		PrivateMapKey: privateMapKeyForInvitees,
		SDKAppId:      sdkAppId,
	}
	return res, nil
}
