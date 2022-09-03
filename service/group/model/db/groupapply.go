package db

import (
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model/biz"
)

type GroupApply struct {
	// 审批 ID
	Id int64
	// 群 ID
	GroupId int64
	// 邀请人 ID, 空表示是自己主动申请的
	InviterId string
	// 申请加入人 ID
	MemberId string
	// 申请备注
	ApplyNote string
	// 审批人 ID
	OperatorId string
	// 审批情况 0=待审批, 1=审批通过, 2=审批不通过, 10=审批忽略
	ApplyStatus int32
	// 拒绝原因
	RejectReason string
	// 创建时间 ms
	CreateTime int64
	// 修改时间 ms
	UpdateTime int64
}

func ConvertGroupApply(res map[string]string) *GroupApply {
	return &GroupApply{
		Id:           util.MustToInt64(res["id"]),
		GroupId:      util.MustToInt64(res["group_id"]),
		InviterId:    util.MustToString(res["inviter_id"]),
		MemberId:     util.MustToString(res["member_id"]),
		ApplyNote:    util.MustToString(res["apply_note"]),
		OperatorId:   util.MustToString(res["operator_id"]),
		ApplyStatus:  util.MustToInt32(res["apply_status"]),
		RejectReason: util.MustToString(res["reject_reason"]),
		CreateTime:   util.MustToInt64(res["create_time"]),
		UpdateTime:   util.MustToInt64(res["update_time"]),
	}
}

func (a *GroupApply) ToBiz() *biz.GroupApplyBiz {
	return &biz.GroupApplyBiz{
		ApplyId:      a.Id,
		GroupId:      a.GroupId,
		InviterId:    a.InviterId,
		MemberId:     a.MemberId,
		ApplyNote:    a.ApplyNote,
		OperatorId:   a.OperatorId,
		ApplyStatus:  a.ApplyStatus,
		RejectReason: a.RejectReason,
		CreateTime:   a.CreateTime,
		UpdateTime:   a.UpdateTime,
	}
}
