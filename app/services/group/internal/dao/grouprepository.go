package dao

import (
	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/pkg/mysql"
)

type GroupRepository interface {
	NewTx() (*mysql.MysqlTx, error)

	InsertGroupInfo(tx *mysql.MysqlTx, group *model.GroupInfo) (int64, int64, error)
	InsertGroupMembers(tx *mysql.MysqlTx, members []*model.GroupMember, updateTime int64) (int64, int64, error)
	GetGroupById(gid int64) (*model.GroupInfo, error)
	GetGroupByMarkId(markId string) (*model.GroupInfo, error)
	GetGroupMutedNumbers(gid int64, now int64) (int32, error)
	GetGroupManagerNumbers(gid int64) (int32, error)
	GetMemberById(gid int64, mid string) (*model.GroupMember, error)
	GetLimitedMembers(gid, start, num int64) ([]*model.GroupMember, error)
	JoinedGroups(uid string) ([]int64, error)
	UpdateGroupStatus(tx *mysql.MysqlTx, group *model.GroupInfo) (int64, int64, error)
	UpdateGroupMemberRole(tx *mysql.MysqlTx, member *model.GroupMember) (int64, int64, error)
	UpdateGroupOwner(tx *mysql.MysqlTx, group *model.GroupInfo) (int64, int64, error)
	UpdateGroupAvatar(group *model.GroupInfo) (int64, int64, error)
	UpdateGroupFriendlyType(group *model.GroupInfo) (int64, int64, error)
	UpdateGroupJoinType(group *model.GroupInfo) (int64, int64, error)
	UpdateGroupMuteType(group *model.GroupInfo) (int64, int64, error)
	UpdateGroupName(group *model.GroupInfo) (int64, int64, error)
}
