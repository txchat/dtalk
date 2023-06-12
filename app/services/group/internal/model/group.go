package model

import "github.com/txchat/dtalk/pkg/util"

const (
	GroupStatusServing = 0 // 正常
	GroupStatusBlocked = 1 // 封禁
	GroupStatusDisBand = 2 // 解散
)

const (
	AnybodyCanJoinGroup   = 0 // 无需审批（默认）
	JustManagerCanInvite  = 1 // 禁止加群，群主和管理员邀请加群
	NormalMemberCanInvite = 2 // 普通人邀请需要审批,群主和管理员直接加群
)

type GroupInfo struct {
	GroupId     int64  `json:"groupId" form:"groupId" gorm:"group_id;"`
	GroupMarkId string `json:"groupMarkId" form:"groupMarkId" gorm:"group_mark_id;"`
	GroupName   string `json:"groupName" form:"groupName" gorm:"group_name;"`
	GroupAvatar string `json:"groupAvatar" form:"groupAvatar" gorm:"group_avatar;"`
	// 群人数
	GroupMemberNum int32 `json:"groupMemberNum" form:"groupMemberNum" gorm:"group_member_num;"`
	// 群人数上限
	GroupMaximum   int32  `json:"groupMaximum" form:"groupMaximum" gorm:"group_maximum;"`
	GroupIntroduce string `json:"groupIntroduce" form:"groupIntroduce" gorm:"group_Introduce;"`
	// 群状态，0=正常 1=封禁 2=解散
	GroupStatus     int32  `json:"groupStatus" form:"groupStatus" gorm:"group_status;"`
	GroupOwnerId    string `json:"groupOwnerId" form:"groupOwnerId" gorm:"group_owner_id;"`
	GroupCreateTime int64  `json:"groupCreateTime" form:"groupCreateTime" gorm:"group_create_time;"`
	GroupUpdateTime int64  `json:"groupUpdateTime" form:"groupUpdateTime" gorm:"group_update_time;"`
	// 加群方式，0=无需审批（默认），1=禁止加群，群主和管理员邀请加群, 2=普通人邀请需要审批,群主和管理员直接加群
	GroupJoinType int32 `json:"groupJoinType" form:"groupJoinType" gorm:"group_join_type;"`
	// 禁言， 0=全员可发言， 1=全员禁言(除群主和管理员)
	GroupMuteType int32 `json:"groupMuteType" form:"groupMuteType" gorm:"group_mute_type;"`
	// 加好友限制， 0=群内可加好友，1=群内禁止加好友
	GroupFriendType int32 `gorm:"group_friend_type;"`
	//
	GroupAESKey string `gorm:"group_aes_key;"`
	//
	GroupPubName string `gorm:"group_pub_name;"`
	// 群类型 (0: 普通群, 1: 全员群, 2: 部门群)
	GroupType int32 `gorm:"group_type;"`
}

func ConvertGroupInfo(res map[string]string) *GroupInfo {
	return &GroupInfo{
		GroupId:         util.MustToInt64(res["group_id"]),
		GroupMarkId:     res["group_mark_id"],
		GroupName:       res["group_name"],
		GroupAvatar:     res["group_avatar"],
		GroupMemberNum:  util.MustToInt32(res["group_member_num"]),
		GroupMaximum:    util.MustToInt32(res["group_maximum"]),
		GroupIntroduce:  res["group_Introduce"],
		GroupStatus:     util.MustToInt32(res["group_status"]),
		GroupOwnerId:    res["group_owner_id"],
		GroupCreateTime: util.MustToInt64(res["group_create_time"]),
		GroupUpdateTime: util.MustToInt64(res["group_update_time"]),
		GroupJoinType:   util.MustToInt32(res["group_join_type"]),
		GroupMuteType:   util.MustToInt32(res["group_mute_type"]),
		GroupFriendType: util.MustToInt32(res["group_friend_type"]),
		GroupAESKey:     util.MustToString(res["group_aes_key"]),
		GroupPubName:    util.MustToString(res["group_pub_name"]),
		GroupType:       util.MustToInt32(res["group_type"]),
	}
}
