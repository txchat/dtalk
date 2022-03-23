package db

import (
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/pkg/util"
)

type GroupInfo struct {
	GroupId     int64  `json:"groupId" form:"groupId"`
	GroupMarkId string `json:"groupMarkId" form:"groupMarkId"`
	GroupName   string `json:"groupName" form:"groupName"`
	GroupAvatar string `json:"groupAvatar" form:"groupAvatar"`
	// 群人数
	GroupMemberNum int32 `json:"groupMemberNum" form:"groupMemberNum"`
	// 群人数上限
	GroupMaximum   int32  `json:"groupMaximum" form:"groupMaximum"`
	GroupIntroduce string `json:"groupIntroduce" form:"groupIntroduce"`
	// 群状态，0=正常 1=封禁 2=解散
	GroupStatus     int32  `json:"groupStatus" form:"groupStatus"`
	GroupOwnerId    string `json:"groupOwnerId" form:"groupOwnerId"`
	GroupCreateTime int64  `json:"groupCreateTime" form:"groupCreateTime"`
	GroupUpdateTime int64  `json:"groupUpdateTime" form:"groupUpdateTime"`
	// 加群方式，0=无需审批（默认），1=禁止加群，群主和管理员邀请加群, 2=普通人邀请需要审批,群主和管理员直接加群
	GroupJoinType int32 `json:"groupJoinType" form:"groupJoinType"`
	// 禁言， 0=全员可发言， 1=全员禁言(除群主和管理员)
	GroupMuteType int32 `json:"groupMuteType" form:"groupMuteType"`
	// 加好友限制， 0=群内可加好友，1=群内禁止加好友
	GroupFriendType int32
	//
	GroupAESKey string
	//
	GroupPubName string
	// 群类型 (0: 普通群, 1: 全员群, 2: 部门群)
	GroupType int32
}

func ConvertGroupInfos(maps []map[string]string) []*GroupInfo {
	dtos := make([]*GroupInfo, 0, len(maps))
	for _, res := range maps {
		dtos = append(dtos, ConvertGroupInfo(res))
	}

	return dtos
}

func ConvertGroupInfo(res map[string]string) *GroupInfo {
	return &GroupInfo{
		GroupId:         util.ToInt64(res["group_id"]),
		GroupMarkId:     res["group_mark_id"],
		GroupName:       res["group_name"],
		GroupAvatar:     res["group_avatar"],
		GroupMemberNum:  util.ToInt32(res["group_member_num"]),
		GroupMaximum:    util.ToInt32(res["group_maximum"]),
		GroupIntroduce:  res["group_Introduce"],
		GroupStatus:     util.ToInt32(res["group_status"]),
		GroupOwnerId:    res["group_owner_id"],
		GroupCreateTime: util.ToInt64(res["group_create_time"]),
		GroupUpdateTime: util.ToInt64(res["group_update_time"]),
		GroupJoinType:   util.ToInt32(res["group_join_type"]),
		GroupMuteType:   util.ToInt32(res["group_mute_type"]),
		GroupFriendType: util.ToInt32(res["group_friend_type"]),
		GroupAESKey:     util.ToString(res["group_aes_key"]),
		GroupPubName:    util.ToString(res["group_pub_name"]),
		GroupType:       util.ToInt32(res["group_type"]),
	}
}

func (g *GroupInfo) ToBiz() *biz.GroupInfo {
	return &biz.GroupInfo{
		GroupId:         g.GroupId,
		GroupMarkId:     g.GroupMarkId,
		GroupName:       g.GroupName,
		GroupAvatar:     g.GroupAvatar,
		GroupMemberNum:  g.GroupMemberNum,
		GroupMaximum:    g.GroupMaximum,
		GroupIntroduce:  g.GroupIntroduce,
		GroupStatus:     g.GroupStatus,
		GroupOwnerId:    g.GroupOwnerId,
		GroupCreateTime: g.GroupCreateTime,
		GroupUpdateTime: g.GroupUpdateTime,
		GroupJoinType:   g.GroupJoinType,
		GroupMuteType:   g.GroupMuteType,
		GroupFriendType: g.GroupFriendType,
		MuteNum:         0,
		AdminNum:        0,
		AESKey:          g.GroupAESKey,
		GroupPubName:    g.GroupPubName,
		GroupType:       g.GroupType,
	}
}
