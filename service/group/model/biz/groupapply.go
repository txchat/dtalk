package biz

import (
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model/types"
)

const (
	GroupApplyWait   = 0
	GroupApplyAccept = 1
	GroupApplyReject = 2
	GroupApplyIgnore = 10
)

type GroupApplyBiz struct {
	// 审批 ID
	ApplyId int64
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

// ToTypes 将业务模型转换为 api 展示信息
func (g *GroupApplyBiz) ToTypes() *types.GroupApplyInfo {
	return &types.GroupApplyInfo{
		ApplyId:      util.MustToString(g.ApplyId),
		GroupId:      util.MustToString(g.GroupId),
		InviterId:    g.InviterId,
		MemberId:     g.MemberId,
		ApplyNote:    g.ApplyNote,
		OperatorId:   g.OperatorId,
		ApplyStatus:  g.ApplyStatus,
		RejectReason: g.RejectReason,
		CreateTime:   g.CreateTime,
		UpdateTime:   g.UpdateTime,
	}
}

// IsWait 判断该审批是否被处理过
func (g *GroupApplyBiz) IsWait() error {
	if g.ApplyStatus == GroupApplyWait {
		return nil
	}
	return xerror.NewError(xerror.GroupApplyUsed)
}
