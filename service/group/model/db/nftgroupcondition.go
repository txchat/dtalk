package db

import (
	"github.com/txchat/dtalk/pkg/util"
)

type NFTGroupCondition struct {
	GroupId int64
	NFTType int32
	NFTId   string
	NFTName string
}

func ConvertNFTGroupCondition(res map[string]string) *NFTGroupCondition {
	return &NFTGroupCondition{
		GroupId: util.ToInt64(res["group_id"]),
		NFTType: util.ToInt32(res["nft_type"]),
		NFTId:   res["nft_id"],
		NFTName: res["nft_name"],
	}
}
