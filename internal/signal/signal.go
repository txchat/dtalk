package signal

import (
	"context"

	"github.com/txchat/dtalk/internal/group"
	"github.com/txchat/dtalk/internal/recordhelper"
	"github.com/txchat/imparse/proto/signal"
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

	StartCall(ctx context.Context, target string, taskID int64) error
	AcceptCall(ctx context.Context, target string, action *signal.SignalAcceptCall) error
	StopCall(ctx context.Context, target string, action *signal.SignalStopCall) error

	MessageReceived(ctx context.Context, item *recordhelper.ConnSeqItem) error
	EndpointLogin(ctx context.Context, uid string, actionProto *signal.SignalEndpointLogin) error

	FocusPrivateMessage(ctx context.Context, users []string, action *signal.SignalFocusMessage) error
	FocusGroupMessage(ctx context.Context, gid int64, action *signal.SignalFocusMessage) error
	RevokePrivateMessage(ctx context.Context, users []string, action *signal.SignalRevoke) error
	RevokeGroupMessage(ctx context.Context, gid int64, action *signal.SignalRevoke) error
}
