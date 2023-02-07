package group

import (
	xrand "github.com/txchat/dtalk/pkg/rand"
	"github.com/txchat/dtalk/pkg/util"
)

type Manager struct {
}

func NewGroupManager() *Manager {
	return &Manager{}
}

func (mg *Manager) CreateNewGroup(gType TypeOfGroup, groupId int64, name, markId, ownerId string, maxMembers int) *Group {
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
