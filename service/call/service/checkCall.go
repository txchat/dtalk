package service

import (
	"time"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/call/model"
)

func (s *Service) CheckCall(req *model.CheckCallRequest) (res *model.CheckCallResponse, err error) {
	defer func() {
		if err != nil {
			s.log.Error().Err(err).Interface("req", req).Str("personId", req.PersonId).Msg("CheckCall")
		} else {
			s.log.Info().Interface("req", req).Str("personId", req.PersonId).Msg("CheckCall")
		}
	}()

	// 从 redis 中获得 session
	session, err := s.dao.GetSession(req.TraceId)
	if err != nil {
		return nil, xerror.NewError(xerror.CodeInnerError).SetExtMessage(err.Error())
	}

	// TODO 暂不支持群会议
	if session.GroupId != 0 {
		return nil, xerror.NewError(xerror.ParamsError)
	}

	// 判断 session 是否过期
	nowTime := time.Now().UnixNano() / 1e6
	if session.Deadline < nowTime {
		return nil, xerror.NewError(xerror.CodeInnerError).SetExtMessage("1")
	}

	// TODO 判断是否在被接收方组内
	//if session.Target != req.PersonId {
	//	return nil, xerror.NewError(xerror.CodeInnerError).SetExtMessage("2")
	//}

	// 判断 session 状态是否在准备中
	if session.Status != model.READY {
		return nil, xerror.NewError(xerror.CodeInnerError).SetExtMessage("3")
	}

	return &model.CheckCallResponse{
		TraceId:    session.TraceId,
		TraceIdStr: util.ToString(session.TraceId),
		RTCType:    session.RTCType,
		Invitees:   session.Invitees,
		Caller:     session.Caller,
		CreateTime: session.CreateTime,
		Deadline:   session.Deadline,
		GroupId:    util.ToString(session.GroupId),
		Timeout:    session.Timeout,
	}, nil
}
