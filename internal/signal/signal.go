package signal

import (
	"context"

	"github.com/txchat/dtalk/internal/group"
)

type Signal interface {
	GroupAddNewMembers(ctx context.Context, gid int64, members []string) error
	GroupRemoveMembers(ctx context.Context, gid int64, members []string) error
	GroupDeleted(ctx context.Context, gid int64) error
	UpdateGroupJoinType(ctx context.Context, gid int64, tp int32) error
	UpdateGroupFriendlyType(ctx context.Context, gid int64, tp int32) error
	UpdateGroupMuteType(ctx context.Context, gid int64, tp int32) error
	UpdateGroupName(ctx context.Context, gid int64, name string) error
	UpdateGroupAvatar(ctx context.Context, gid int64, avatar string) error
	UpdateGroupMemberRole(ctx context.Context, gid int64, uid string, role group.RoleType) error
	UpdateMembersMuteTime(ctx context.Context, gid, muteTime int64, members []string) error
}
