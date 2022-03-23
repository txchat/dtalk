package service

import (
	"context"
	"strconv"

	logic "github.com/txchat/im/api/logic/grpc"
)

func (s *Service) JoinGroups(ctx context.Context, uid, key string) error {
	groups, err := s.dao.GetAllJoinedGroups(ctx, uid)
	if err != nil {
		return err
	}

	if len(groups) == 0 {
		return nil
	}

	var gids = make([]string, len(groups))
	for i, group := range groups {
		gids[i] = strconv.FormatInt(group, 10)
	}
	_, err = s.logicClient.JoinGroupsByKeys(ctx, &logic.GroupsKey{
		AppId: s.cfg.AppId,
		Keys:  []string{key},
		Gid:   gids,
	})
	return err
}
