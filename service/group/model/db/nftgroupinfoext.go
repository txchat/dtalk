package db

import (
	"github.com/txchat/dtalk/pkg/util"
)

type NFTGroupInfoExt struct {
	GroupId       int64
	ConditionType int32
}

func ConvertNFTGroupInfoExt(res map[string]string) *NFTGroupInfoExt {
	return &NFTGroupInfoExt{
		GroupId:       util.ToInt64(res["group_id"]),
		ConditionType: util.ToInt32(res["condition_type"]),
	}
}
