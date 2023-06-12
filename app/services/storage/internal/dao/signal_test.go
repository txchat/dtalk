package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txchat/dtalk/app/services/storage/internal/model"
)

var (
	signal1 = &model.SignalContent{
		Uid:        "test1",
		Seq:        1,
		Type:       0,
		Content:    "",
		CreateTime: 0,
		UpdateTime: 0,
	}
	signal2 = &model.SignalContent{
		Uid:        "test1",
		Seq:        2,
		Type:       0,
		Content:    "",
		CreateTime: 0,
		UpdateTime: 0,
	}
	signal3 = &model.SignalContent{
		Uid:        "test1",
		Seq:        3,
		Type:       0,
		Content:    "",
		CreateTime: 0,
		UpdateTime: 0,
	}
)

func TestStorageRepository_AppendSignal(t *testing.T) {
	_, _, err := repo.AppendSignal(signal1)
	assert.Nil(t, err)
}

func TestStorageRepository_BatchAppendSignal(t *testing.T) {
	_, _, err := repo.BatchAppendSignal([]*model.SignalContent{signal2, signal3})
	assert.Nil(t, err)
}
