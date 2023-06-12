package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txchat/dtalk/pkg/util"
)

func TestStorageRepository_AddRecordFocus(t *testing.T) {
	err := repo.AddRecordFocus("testA", "testa2b", util.TimeNowUnixMilli())
	assert.Nil(t, err)
}

func TestStorageRepository_GetRecordFocusNumber(t *testing.T) {
	num, err := repo.GetRecordFocusNumber("testa2b")
	assert.Nil(t, err)
	assert.Equal(t, int32(1), num)
}
