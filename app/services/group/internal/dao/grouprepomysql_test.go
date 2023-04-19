package dao

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/txchat/dtalk/app/services/group/internal/model"
	xmysql "github.com/txchat/dtalk/pkg/mysql"
	"github.com/zeromicro/go-zero/core/service"
)

var (
	mysqlRootPassword string
	repo              *GroupRepositoryMysql
)

func TestMain(m *testing.M) {
	mysqlRootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	repo = NewGroupRepositoryMysql(service.TestMode, xmysql.Config{
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		User:   "root",
		Passwd: mysqlRootPassword,
		DBName: "dtalk",
	})
	os.Exit(m.Run())
}

var (
	now = time.Now()
	g1  = model.GroupInfo{
		GroupId:         1,
		GroupMarkId:     "g1",
		GroupName:       "my test group1",
		GroupAvatar:     "https://xxxxxx.jpg",
		GroupMemberNum:  1,
		GroupMaximum:    10000,
		GroupIntroduce:  "this is the group example",
		GroupStatus:     0,
		GroupOwnerId:    "m1",
		GroupCreateTime: now.UnixMilli(),
		GroupUpdateTime: now.UnixMilli(),
		GroupJoinType:   0,
		GroupMuteType:   0,
		GroupFriendType: 0,
		GroupAESKey:     "",
		GroupPubName:    "pub my test group1",
		GroupType:       0,
	}
	g2 = model.GroupInfo{
		GroupId:         2,
		GroupMarkId:     "g2",
		GroupName:       "my test group2",
		GroupAvatar:     "https://xxxxxx.jpg",
		GroupMemberNum:  1,
		GroupMaximum:    10000,
		GroupIntroduce:  "this is the group example",
		GroupStatus:     1,
		GroupOwnerId:    "m1",
		GroupCreateTime: now.UnixMilli(),
		GroupUpdateTime: now.UnixMilli(),
		GroupJoinType:   1,
		GroupMuteType:   1,
		GroupFriendType: 1,
		GroupAESKey:     "haha",
		GroupPubName:    "pub my test group1",
		GroupType:       1,
	}
	g1m1 = model.GroupMember{
		GroupId:               1,
		GroupMemberId:         "m1",
		GroupMemberName:       "group1 test member1",
		GroupMemberType:       2,
		GroupMemberJoinTime:   now.UnixMilli(),
		GroupMemberUpdateTime: now.UnixMilli(),
		GroupMemberMute:       model.GroupMemberMute{},
	}
	g2m1 = model.GroupMember{
		GroupId:               2,
		GroupMemberId:         "m1",
		GroupMemberName:       "group2 test member1",
		GroupMemberType:       2,
		GroupMemberJoinTime:   now.UnixMilli(),
		GroupMemberUpdateTime: now.UnixMilli(),
		GroupMemberMute:       model.GroupMemberMute{},
	}
)

func TestGroupRepositoryMysql_InsertGroupAndGetGroupInfo(t *testing.T) {
	// init group 1
	tx, err := repo.NewTx()
	if err != nil {
		assert.Nil(t, tx.Rollback())
	}
	assert.Nil(t, err)
	_, _, err = repo.InsertGroupInfo(tx, &g1)
	if err != nil {
		assert.Nil(t, tx.Rollback())
	}
	assert.Nil(t, err)
	_, _, err = repo.InsertGroupMembers(tx, []*model.GroupMember{&g1m1}, now.UnixMilli())
	if err != nil {
		assert.Nil(t, tx.Rollback())
	}
	assert.Nil(t, tx.Commit())

	// init group 2
	tx, err = repo.NewTx()
	if err != nil {
		assert.Nil(t, tx.Rollback())
	}
	assert.Nil(t, err)
	_, _, err = repo.InsertGroupInfo(tx, &g2)
	if err != nil {
		assert.Nil(t, tx.Rollback())
	}
	assert.Nil(t, err)
	_, _, err = repo.InsertGroupMembers(tx, []*model.GroupMember{&g2m1}, now.UnixMilli())
	if err != nil {
		assert.Nil(t, tx.Rollback())
	}
	assert.Nil(t, tx.Commit())

	g1Info, err := repo.GetGroupById(1)
	assert.Nil(t, err)
	assert.EqualValues(t, &g1, g1Info)

	g2Info, err := repo.GetGroupByMarkId("g2")
	assert.Nil(t, err)
	assert.EqualValues(t, &g2, g2Info)

	g1m1Info, err := repo.GetMemberById(1, "m1")
	assert.Nil(t, err)
	assert.EqualValues(t, &g1m1, g1m1Info)

	joinedGroups, err := repo.JoinedGroups("m1")
	assert.Nil(t, err)
	assert.EqualValues(t, []int64{1, 2}, joinedGroups)
}
