package dao

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/service/group/model/db"
)

func TestDao_InsertNFTGroupInfoExt(t *testing.T) {
	d := Dao{
		log:         zerolog.Logger{},
		conn:        testConn,
		redis:       nil,
		redisExpire: 0,
	}
	tx, err := d.NewTx()
	if err != nil {
		t.Error(err)
		return
	}
	_, _, err = d.InsertNFTGroupInfoExt(tx, &db.NFTGroupInfoExt{
		GroupId:       123,
		ConditionType: 1,
	})
	if err != nil {
		t.Error(err)
		return
	}
	err = tx.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDao_InsertNFTGroupConditions(t *testing.T) {
	d := Dao{
		log:         zerolog.Logger{},
		conn:        testConn,
		redis:       nil,
		redisExpire: 0,
	}
	tx, err := d.NewTx()
	if err != nil {
		t.Error(err)
		return
	}
	_, _, err = d.InsertNFTGroupConditions(tx, []*db.NFTGroupCondition{
		{
			GroupId: 123,
			NFTType: 0,
			NFTId:   "12312",
			NFTName: "hh",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	err = tx.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}
