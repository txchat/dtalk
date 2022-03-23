package service

import "context"

// NoticeInviteMembers 新成员加入群发起通知
func (s *Service) NoticeInviteMembers(ctx context.Context, groupId int64, inviterId string, newMembers []string) {
	log := s.GetLogWithTrace(ctx)

	// 通知 logic
	if err := s.LogicNoticeJoin(ctx, groupId, newMembers); err != nil {
		log.Error().Err(err).Msg("inviteMemberNotice logic")
	}

	// 发送给 pusher
	if err := s.PusherSignalJoin(ctx, groupId, newMembers); err != nil {
		log.Error().Err(err).Msg("inviteMemberNotice pusher")
	}

	if err := s.NoticeMsgSignInGroup(ctx, groupId, inviterId, newMembers); err != nil {
		log.Error().Err(err).Msg("inviteMemberNotice alert")
	}
}

//
