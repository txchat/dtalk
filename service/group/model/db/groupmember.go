package db

import (
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model/biz"
)

type GroupMember struct {
	GroupId         int64  `json:"groupId" form:"groupId"`
	GroupMemberId   string `json:"groupMemberId" form:"groupMemberId"`
	GroupMemberName string `json:"groupMemberName" form:"groupMemberName"`
	// 用户角色，2=群主，1=管理员，0=群员，10=退群
	GroupMemberType       int32 `json:"groupMemberType" form:"groupMemberType"`
	GroupMemberJoinTime   int64 `json:"groupMemberJoinTime" form:"groupMemberJoinTime"`
	GroupMemberUpdateTime int64 `json:"groupMemberUpdateTime" form:"groupMemberUpdateTime"`
}

func ConvertGroupMember(res map[string]string) *GroupMember {
	return &GroupMember{
		GroupId:             util.ToInt64(res["group_id"]),
		GroupMemberId:       res["group_member_id"],
		GroupMemberName:     res["group_member_name"],
		GroupMemberType:     util.ToInt32(res["group_member_type"]),
		GroupMemberJoinTime: util.ToInt64(res["group_member_join_time"]),
		//GroupMemberUpdateTime: util.ToInt64(res["group_member_update_time"]),
	}
}

func (m *GroupMember) ToBiz() *biz.GroupMember {
	return &biz.GroupMember{
		GroupId:             m.GroupId,
		GroupMemberId:       m.GroupMemberId,
		GroupMemberName:     m.GroupMemberName,
		GroupMemberType:     m.GroupMemberType,
		GroupMemberMuteTime: 0,
		GroupMemberJoinTime: m.GroupMemberJoinTime,
	}
}

type GroupMemberMute struct {
	GroupId       int64
	GroupMemberId string
	// 该用户被禁言结束的时间 9223372036854775807=永久禁言
	GroupMemberMuteTime       int64
	GroupMemberMuteUpdateTime int64
}

func ConvertGroupMemberMute(res map[string]string) *GroupMemberMute {
	return &GroupMemberMute{
		GroupId:                   util.ToInt64(res["group_id"]),
		GroupMemberId:             res["group_member_id"],
		GroupMemberMuteTime:       util.ToInt64(res["group_member_mute_time"]),
		GroupMemberMuteUpdateTime: util.ToInt64(res["group_member_mute_update_time"]),
	}
}

type GroupMemberWithMute struct {
	GroupId             int64
	GroupMemberId       string
	GroupMemberName     string
	GroupMemberType     int32
	GroupMemberMuteTime int64
	GroupMemberJoinTime int64
}

func ConvertGroupMemberWithMute(res map[string]string) *GroupMemberWithMute {
	if res["group_member_mute_time"] == "" {
		res["group_member_mute_time"] = "0"
	}
	return &GroupMemberWithMute{
		GroupId:             util.ToInt64(res["group_id"]),
		GroupMemberId:       res["group_member_id"],
		GroupMemberMuteTime: util.ToInt64(res["group_member_mute_time"]),
		GroupMemberName:     res["group_member_name"],
		GroupMemberType:     util.ToInt32(res["group_member_type"]),
		GroupMemberJoinTime: util.ToInt64(res["group_member_join_time"]),
	}
}

func (m *GroupMemberWithMute) ToBiz() *biz.GroupMember {
	return &biz.GroupMember{
		GroupId:             m.GroupId,
		GroupMemberId:       m.GroupMemberId,
		GroupMemberName:     m.GroupMemberName,
		GroupMemberType:     m.GroupMemberType,
		GroupMemberMuteTime: m.GroupMemberMuteTime,
		GroupMemberJoinTime: m.GroupMemberJoinTime,
	}
}
