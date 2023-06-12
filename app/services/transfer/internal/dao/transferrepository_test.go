package dao

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txchat/dtalk/api/proto/chat"
	"github.com/txchat/dtalk/pkg/redis"
)

var (
	repo *TransferRepository
)

func TestMain(m *testing.M) {
	repo = NewTransferRepository(redis.Config{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
		Auth:    "",
		Active:  60000,
		Idle:    1024,
	})
	os.Exit(m.Run())
}

var (
	testUserACid   = "testUserACid"
	testUserAMid   = "testUserAMid"
	testUserAState = chat.SendMessageReply_UnSupportedMessageType
	testUserBCid   = "testUserBCid"
)

func TestTransferRepository_AddIndexCidMid(t *testing.T) {
	err := repo.AddIndexCidMid(context.Background(), testUserACid, testUserAMid, testUserAState)
	assert.Nil(t, err)
	mid, state, err := repo.GetMidByCid(context.Background(), testUserACid)
	assert.Nil(t, err)
	assert.Equal(t, testUserAMid, mid)
	assert.Equal(t, testUserAState, state)
	mid, state, err = repo.GetMidByCid(context.Background(), testUserBCid)
	assert.Nil(t, err)
	assert.Equal(t, "", mid)
	assert.Equal(t, chat.SendMessageReply_FailedType(0), state)
}
