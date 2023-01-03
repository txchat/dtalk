package group

import (
	xrand "github.com/txchat/dtalk/pkg/rand"
	"github.com/txchat/dtalk/pkg/util"
)

type GroupManager struct {
}

func (mg *GroupManager) ChangeOwner(operator, newOwner *Member, group *Group) error {
	// check operator permission
	if operator.id != group.GetOwner() {
		return ErrPermissionDenied
	}
	// exec change info
	err := mg.dbExec.ChangeGroupOwner(group.id, operator.id, newOwner.id)
	if err != nil {
		return err
	}

	// unset mute
	err = mg.gmm.UnMute(operator, newOwner)
	if err != nil {
		return err
	}

	// send signal
	err = mg.signalHub.ChangeMemberRoleType(group.id, newOwner.id, Owner)
	if err != nil {
		return err
	}
	err = mg.signalHub.ChangeMemberRoleType(group.id, operator.id, Normal)
	if err != nil {
		return err
	}
	// send notify
	err = mg.noticeHub.UpdateGroupOwner(group.id, operator.id, newOwner.id)
	if err != nil {
		return err
	}
	return nil
}

func (mg *GroupManager) CreateNewGroup(gType TypeOfGroup, groupId int64, name, markId, ownerId string, maxMembers int) *Group {
	joinPermission := AnybodyCanJoinGroup
	switch gType {
	case NormalGroup:
		joinPermission = AnybodyCanJoinGroup
	default:
		joinPermission = JustManagerCanInvite
	}
	g := &Group{
		id:                   groupId,
		name:                 name,
		avatar:               "",
		markId:               markId,
		owner:                ownerId,
		maxMembers:           maxMembers,
		createTime:           util.TimeNowUnixMilli(),
		joinPermission:       joinPermission,
		mutePermission:       NotLimited,
		friendshipPermission: AllowedGroupFriendship,
		aesKey:               xrand.NewAESKey256(),
		members:              make([]*Member, 0),
	}
	return g
}
