package dao

import (
	"database/sql/driver"

	"github.com/txchat/dtalk/app/services/group/internal/model"
)

type GroupRepository interface {
	NewTx() (driver.Tx, error)

	InsertGroupInfo(tx driver.Tx, group *model.GroupInfo) (int64, int64, error)
	InsertGroupMembers(tx driver.Tx, members []*model.GroupMember, updateTime int64) (int64, int64, error)
	GetGroupById(gid int64) (*model.GroupInfo, error)
	GetGroupByMarkId(markId string) (*model.GroupInfo, error)
	GetGroupMutedNumbers(gid int64, now int64) (int32, error)
	GetGroupManagerNumbers(gid int64) (int32, error)
	GetMemberById(gid int64, mid string) (*model.GroupMember, error)
	GetUnLimitedMembers(gid int64) ([]*model.GroupMember, error)
	GetLimitedMembers(gid, start, num int64) ([]*model.GroupMember, error)
	GetMutedMembers(gid, time int64) ([]*model.GroupMember, error)
	JoinedGroups(uid string) ([]int64, error)
	UpdateGroupStatus(tx driver.Tx, group *model.GroupInfo) (int64, int64, error)
	UpdateGroupMemberRole(tx driver.Tx, member *model.GroupMember) (int64, int64, error)
	UpdateGroupMembersMuteTime(tx driver.Tx, members []*model.GroupMember) (int64, int64, error)
	UpdateGroupOwner(tx driver.Tx, group *model.GroupInfo) (int64, int64, error)
	UpdateGroupAvatar(group *model.GroupInfo) (int64, int64, error)
	UpdateGroupFriendlyType(group *model.GroupInfo) (int64, int64, error)
	UpdateGroupJoinType(group *model.GroupInfo) (int64, int64, error)
	UpdateGroupMuteType(group *model.GroupInfo) (int64, int64, error)
	UpdateGroupName(group *model.GroupInfo) (int64, int64, error)
	UpdateGroupMembersNumber(tx driver.Tx, group *model.GroupInfo) (int64, int64, error)
	UpdateGroupMemberName(member *model.GroupMember) (int64, int64, error)
}
