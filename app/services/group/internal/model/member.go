package model

import "github.com/txchat/dtalk/pkg/util"

const UnLimitedNumbers = -1

const (
	GroupMemberTypeOwner   = 2  // 群主
	GroupMemberTypeManager = 1  // 管理员
	GroupMemberTypeNormal  = 0  // 群员
	GroupMemberTypeOther   = 10 // 退群
)

type GroupMemberMute struct {
	GroupMemberMuteTime       int64 `json:"groupMemberMuteTime" form:"groupMemberMuteTime"`
	GroupMemberMuteUpdateTime int64 `json:"groupMemberMuteUpdateTime" form:"groupMemberMuteUpdateTime"`
}

type GroupMember struct {
	GroupId         int64  `json:"groupId" form:"groupId"`
	GroupMemberId   string `json:"groupMemberId" form:"groupMemberId"`
	GroupMemberName string `json:"groupMemberName" form:"groupMemberName"`
	// 用户角色，2=群主，1=管理员，0=群员，10=退群
	GroupMemberType       int32 `json:"groupMemberType" form:"groupMemberType"`
	GroupMemberJoinTime   int64 `json:"groupMemberJoinTime" form:"groupMemberJoinTime"`
	GroupMemberUpdateTime int64 `json:"groupMemberUpdateTime" form:"groupMemberUpdateTime"`
	GroupMemberMute
}

func ConvertGroupMember(res map[string]string) *GroupMember {
	if res["group_member_mute_time"] == "" {
		res["group_member_mute_time"] = "0"
	}
	return &GroupMember{
		GroupId:               util.MustToInt64(res["group_id"]),
		GroupMemberId:         res["group_member_id"],
		GroupMemberName:       res["group_member_name"],
		GroupMemberType:       util.MustToInt32(res["group_member_type"]),
		GroupMemberJoinTime:   util.MustToInt64(res["group_member_join_time"]),
		GroupMemberUpdateTime: util.MustToInt64(res["group_member_update_time"]),
		GroupMemberMute: GroupMemberMute{
			GroupMemberMuteTime:       util.MustToInt64(res["group_member_mute_time"]),
			GroupMemberMuteUpdateTime: util.MustToInt64(res["group_member_mute_update_time"]),
		},
	}
}
