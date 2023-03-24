package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/pkg/util"
)

var (
	now           = util.TimeNowUnixMilli()
	privateRecord = &model.MsgContent{
		Mid:        "testa2b",
		Cid:        "testa2bcid",
		SenderId:   "testA",
		ReceiverId: "testB",
		MsgType:    0,
		Content:    "hello",
		CreateTime: now,
		Source:     "",
		Reference:  "",
	}
	privateRelationA = &model.MsgRelation{
		Mid:        "testa2b",
		OwnerUid:   "testA",
		OtherUid:   "testB",
		Type:       model.Send,
		CreateTime: now,
	}
	privateRelationB = &model.MsgRelation{
		Mid:        "testa2b",
		OwnerUid:   "testB",
		OtherUid:   "testA",
		Type:       model.Rev,
		CreateTime: now,
	}

	groupRecord = &model.MsgContent{
		Mid:        "testgroup1",
		Cid:        "testgroup1cid",
		SenderId:   "testA",
		ReceiverId: "gid1",
		MsgType:    0,
		Content:    "hello",
		CreateTime: now,
		Source:     "",
		Reference:  "",
	}
	groupRelation = []*model.MsgRelation{
		{
			Mid:        "testgroup1",
			OwnerUid:   "testA",
			OtherUid:   "gid1",
			Type:       model.Send,
			CreateTime: now,
		}, {
			Mid:        "testgroup1",
			OwnerUid:   "testB",
			OtherUid:   "gid1",
			Type:       model.Rev,
			CreateTime: now,
		}, {
			Mid:        "testgroup1",
			OwnerUid:   "testC",
			OtherUid:   "gid1",
			Type:       model.Rev,
			CreateTime: now,
		},
	}
)

func TestStorageRepository_AppendPrivateMsg(t *testing.T) {
	_, _, err := repo.AppendPrivateMsg(privateRecord, privateRelationA, privateRelationB)
	assert.Nil(t, err)
}

func TestStorageRepository_GetPrivateRecordByMid(t *testing.T) {
	record, err := repo.GetPrivateMsgByMid("testa2b")
	assert.Nil(t, err)
	assert.EqualValues(t, privateRecord, record)
}

func TestStorageRepository_AppendGroupMsg(t *testing.T) {
	_, _, err := repo.AppendGroupMsg(groupRecord, groupRelation)
	assert.Nil(t, err)
}

func TestStorageRepository_GetGroupRecordByMid(t *testing.T) {
	record, err := repo.GetGroupMsgByMid("testgroup1")
	assert.Nil(t, err)
	assert.EqualValues(t, groupRecord, record)
}

func TestStorageRepository_GetPrivateChatSessionMsg(t *testing.T) {
	records, err := repo.GetPrivateChatSessionMsg("testB", "testA", "", 10)
	assert.Nil(t, err)
	assert.EqualValues(t, []*model.MsgContent{
		privateRecord,
	}, records)
}

func TestStorageRepository_GetGroupChatSessionMsg(t *testing.T) {
	records, err := repo.GetGroupChatSessionMsg("testC", "gid1", "", 10)
	assert.Nil(t, err)
	assert.EqualValues(t, []*model.MsgContent{
		groupRecord,
	}, records)
}

func TestStorageRepository_DelPrivateMsg(t *testing.T) {
	_, _, err := repo.DelPrivateMsg("testa2b")
	assert.Nil(t, err)
}

func TestStorageRepository_DelGroupMsg(t *testing.T) {
	_, _, err := repo.DelGroupMsg("testgroup1")
	assert.Nil(t, err)
}
