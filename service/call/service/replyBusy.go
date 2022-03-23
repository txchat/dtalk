package service

import (
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/call/model"
)

func (s *Service) ReplyBusy(req *model.ReplyBusyRequest) (res *model.ReplyBusyResponse, err error) {
	defer func() {
		if err != nil {
			s.log.Error().Err(err).Interface("req", req).Str("personId", req.PersonId).Msg("ReplyBusy")
		} else {
			s.log.Info().Interface("req", req).Str("personId", req.PersonId).Msg("ReplyBusy")
		}
	}()

	// 从 redis 中获得 session
	session, err := s.dao.GetSession(req.TraceId)
	if err != nil {
		return nil, xerror.NewError(xerror.CodeInnerError).SetExtMessage(err.Error())
	}

	if session.GroupId == 0 {
		// 给 Caller 发 ActionStopCall with Busy
		if err := s.noticeStopCall(session.Caller, session.TraceId, model.Busy); err != nil {
			s.log.Error().Err(err).Msg("ReplyBusy() noticeStopCall()")
		}
	} else {
		// TODO 如果是群聊
		return nil, xerror.NewError(xerror.ParamsError)
	}

	return &model.ReplyBusyResponse{}, nil
}
