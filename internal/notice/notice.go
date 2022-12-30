package notice

import (
	"context"

	"github.com/txchat/dtalk/internal/group"
)

type Notice interface {
	GroupAddNewMembers(ctx context.Context, gid int64, operator string, members []string) error
	GroupSignOut(ctx context.Context, gid int64, target string) error
	GroupKickOutMembers(ctx context.Context, gid int64, operator string, members []string) error
	GroupDeleted(ctx context.Context, gid int64, operator string) error
	UpdateGroupMuteType(ctx context.Context, gid int64, operator string, tp int32) error
	UpdateGroupName(ctx context.Context, gid int64, operator, name string) error
	UpdateGroupMemberRole(ctx context.Context, gid int64, operator, uid string, role group.RoleType) error
	UpdateMembersMuteTime(ctx context.Context, operator string, gid, muteTime int64, members []string) error
}
