package txchat

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/answer/answerclient"
	"github.com/txchat/dtalk/internal/group"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/imparse/proto/signal"
)

type SignalHub struct {
	answerClient answerclient.Answer
}

func NewSignalHub(conn answerclient.Answer) *SignalHub {
	return &SignalHub{answerClient: conn}
}

func (hub *SignalHub) GroupAddNewMembers(ctx context.Context, gid int64, members []string) error {
	action := &signal.SignalSignInGroup{
		Uid:   members,
		Group: gid,
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.GroupCastSignal(ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_SignInGroup,
		Target: util.MustToString(gid),
		Body:   body,
	})
	return err
}

func (hub *SignalHub) GroupRemoveMembers(ctx context.Context, gid int64, members []string) error {
	action := &signal.SignalSignOutGroup{
		Uid:   members,
		Group: gid,
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.GroupCastSignal(ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_SignOutGroup,
		Target: util.MustToString(gid),
		Body:   body,
	})
	return err
}

func (hub *SignalHub) GroupDeleted(ctx context.Context, gid int64) error {
	action := &signal.SignalDeleteGroup{
		Group: gid,
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.GroupCastSignal(ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_DeleteGroup,
		Target: util.MustToString(gid),
		Body:   body,
	})
	return err
}

func (hub *SignalHub) UpdateGroupJoinType(ctx context.Context, gid int64, tp int32) error {
	action := &signal.SignalUpdateGroupJoinType{
		Group: gid,
		Type:  signal.JoinType(tp),
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.GroupCastSignal(ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_UpdateGroupJoinType,
		Target: util.MustToString(gid),
		Body:   body,
	})
	return err
}

func (hub *SignalHub) UpdateGroupFriendlyType(ctx context.Context, gid int64, tp int32) error {
	action := &signal.SignalUpdateGroupFriendType{
		Group: gid,
		Type:  signal.FriendType(tp),
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.GroupCastSignal(ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_UpdateGroupFriendType,
		Target: util.MustToString(gid),
		Body:   body,
	})
	return err
}

func (hub *SignalHub) UpdateGroupMuteType(ctx context.Context, gid int64, tp int32) error {
	action := &signal.SignalUpdateGroupMuteType{
		Group: gid,
		Type:  signal.MuteType(tp),
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.GroupCastSignal(ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_UpdateGroupMuteType,
		Target: util.MustToString(gid),
		Body:   body,
	})
	return err
}

func (hub *SignalHub) UpdateGroupName(ctx context.Context, gid int64, name string) error {
	action := &signal.SignalUpdateGroupName{
		Group: gid,
		Name:  name,
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.GroupCastSignal(ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_UpdateGroupName,
		Target: util.MustToString(gid),
		Body:   body,
	})
	return err
}

func (hub *SignalHub) UpdateGroupAvatar(ctx context.Context, gid int64, avatar string) error {
	action := &signal.SignalUpdateGroupAvatar{
		Group:  gid,
		Avatar: avatar,
		Time:   util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.GroupCastSignal(ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_UpdateGroupAvatar,
		Target: util.MustToString(gid),
		Body:   body,
	})
	return err
}

func (hub *SignalHub) UpdateGroupMemberRole(ctx context.Context, gid int64, uid string, role group.RoleType) error {
	action := &signal.SignalUpdateGroupMemberType{
		Group: gid,
		Uid:   uid,
		Type:  signal.MemberType(role),
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.GroupCastSignal(ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_UpdateGroupMemberType,
		Target: util.MustToString(gid),
		Body:   body,
	})
	return err
}

func (hub *SignalHub) UpdateMembersMuteTime(ctx context.Context, gid, muteTime int64, members []string) error {
	action := &signal.SignalUpdateGroupMemberMuteTime{
		Group:    gid,
		Uid:      members,
		MuteTime: muteTime,
		Time:     util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.GroupCastSignal(ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_UpdateGroupMemberMuteTime,
		Target: util.MustToString(gid),
		Body:   body,
	})
	return err
}
