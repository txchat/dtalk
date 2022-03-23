package service

import (
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/call/model"
	"time"
)

// StartCall 准备发起通话
func (s *Service) StartCall(req *model.StartCallRequest) (res *model.StartCallResponse, err error) {
	defer func() {
		if err != nil {
			s.log.Error().Err(err).Interface("req", req).Str("personId", req.PersonId).Msg("StartCall")
		} else {
			s.log.Info().Interface("req", req).Str("personId", req.PersonId).Msg("StartCall")
		}
	}()

	var groupId int64
	if req.GroupId != "" {
		groupId, err = util.ToInt64E(req.GroupId)
		if err != nil {
			return nil, err
		}
	}

	// 如果是私聊,Invitees只能有一人
	if groupId == 0 && len(req.Invitees) != 1 {
		return nil, xerror.NewError(xerror.ParamsError)
	}

	// 生成 session
	session, err := model.NewSession(
		req.RTCType, req.PersonId, req.Invitees,
		groupId, s.idGenRPCClient, s.room)
	if err != nil {
		return nil, xerror.NewError(xerror.CodeInnerError).SetExtMessage(err.Error())
	}

	// 保存在 redis 中
	if err = s.dao.SaveSession(session); err != nil {
		return nil, xerror.NewError(xerror.CodeInnerError).SetExtMessage(err.Error())
	}

	if session.GroupId == 0 {
		// 给 B 发 ActionStartCall
		if err = s.noticeStartCall(session.Invitees[0], session.TraceId); err != nil {
			s.log.Error().Err(err).Msg("StartCall() noticeStartCall()")
		}
	} else {
		// TODO 如果是群聊
		return nil, xerror.NewError(xerror.ParamsError)
	}

	// 给 Caller 返回 traceId
	res = &model.StartCallResponse{
		RTCType:    session.RTCType,
		TraceId:    session.TraceId,
		TraceIdStr: util.ToString(session.TraceId),
		Caller:     session.Caller,
		Invitees:   session.Invitees,
		CreateTime: session.CreateTime,
		Deadline:   session.Deadline,
		GroupId:    util.ToString(session.GroupId),
		Timeout:    session.Timeout,
	}
	return res, nil
}

// checkTimeout 弃用
func (s *Service) checkTimeout(traceId int64, readyTime int32) {
	durationTime := time.Second * time.Duration(readyTime)
	timeTickerChan := time.Tick(durationTime)
	<-timeTickerChan

	session, err := s.dao.GetSession(traceId)
	if err != nil {
		s.log.Error().Err(err).Msg("checkTimeout GetSession")
	}
	if session.Status != model.READY {
		return
	}
	if session.GroupId == 0 {
		// 给双方发 ActionStopCall with Timeout
		if err := s.noticeStopCall(session.Caller, session.TraceId, model.Timeout); err != nil {
			s.log.Error().Err(err).Msg("StartCall() checkTimeout() noticeStopCall()")
		}
		if err := s.noticeStopCall(session.Invitees[0], session.TraceId, model.Timeout); err != nil {
			s.log.Error().Err(err).Msg("StartCall() checkTimeout() noticeStopCall()")
		}

		session.Status = model.FINISH
		if err := s.dao.SaveSession(session); err != nil {
			s.log.Error().Err(err).Msg("checkTimeout SaveSession")
		}
	}

}
