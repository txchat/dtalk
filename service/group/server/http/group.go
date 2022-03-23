package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	logic "github.com/txchat/dtalk/service/group/logic/http"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/types"
)

// CreateGroup
// @Summary 创建群
// @Author chy@33.cn
// @Tags group 群动作
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.CreateGroupRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.CreateGroupResponse}
// @Router	/app/create-group [post]
func CreateGroup(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.CreateGroupRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	req.Owner.MemberId = userId.(string)
	req.Owner.MemberName = ""
	req.Owner.MemberType = biz.GroupMemberTypeOwner

	req.Members = make([]types.GroupMember, len(req.MemberIds), len(req.MemberIds))
	for i, id := range req.MemberIds {
		req.Members[i].MemberId = id
		req.Members[i].MemberName = ""
		req.Members[i].MemberType = biz.GroupMemberTypeNormal
	}

	res, err := svc.CreateGroupSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)

}

// InviteGroupMembers
// @Summary 邀请新群员
// @Author chy@33.cn
// @Tags group 群动作
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.InviteGroupMembersRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.InviteGroupMembersResponse}
// @Router	/app/invite-group-members [post]
func InviteGroupMembers(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.InviteGroupMembersRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	req.Inviter.MemberId = userId.(string)
	req.NewMembers = make([]types.GroupMember, len(req.NewMemberIds), len(req.NewMemberIds))
	for i, id := range req.NewMemberIds {
		req.NewMembers[i].MemberId = id
		req.NewMembers[i].MemberName = ""
		req.NewMembers[i].MemberType = biz.GroupMemberTypeNormal
	}

	res, err := svc.InviteGroupMembersSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// JoinGroup
// @Summary 直接进群
// @Author chy@33.cn
// @Tags group 群动作
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.JoinGroupReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.JoinGroupResp}
// @Router	/app/join-group [post]
func JoinGroup(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.JoinGroupReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	req.PersonId = userId.(string)

	// 走邀请函数
	nreq := &types.InviteGroupMembersRequest{}
	nreq.Id = req.Id
	nreq.Inviter.MemberId = req.InviterId
	nreq.NewMembers = make([]types.GroupMember, 1, 1)
	nreq.NewMembers[0].MemberId = req.PersonId
	nreq.NewMembers[0].MemberName = ""
	nreq.NewMembers[0].MemberType = biz.GroupMemberTypeNormal

	res, err := svc.InviteGroupMembersSvc(c, nreq)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GroupExit
// @Summary 退群
// @Author chy@33.cn
// @Tags group 群动作
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GroupExitRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GroupExitResponse}
// @Router	/app/group-exit [post]
func GroupExit(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.GroupExitRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	req.PersonId = userId.(string)

	res, err := svc.GroupExitHttp(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GroupRemove
// @Summary 踢人
// @Author chy@33.cn
// @Tags group 群动作
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GroupRemoveRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GroupRemoveResponse}
// @Router	/app/group-remove [post]
func GroupRemove(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.GroupRemoveRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	req.PersonId = userId.(string)

	res, err := svc.GroupRemoveSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GroupDisband
// @Summary 解散群
// @Author chy@33.cn
// @Tags group 群动作
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GroupDisbandRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GroupDisbandResponse}
// @Router	/app/group-disband [post]
func GroupDisband(c *gin.Context) {
	req := &types.GroupDisbandRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.GroupDisbandHttp(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// ChangeOwner
// @Summary 转让群
// @Author chy@33.cn
// @Tags group 群动作
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.ChangeOwnerRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.ChangeOwnerResponse}
// @Router	/app/change-owner [post]
func ChangeOwner(c *gin.Context) {
	req := &types.ChangeOwnerRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	l := logic.NewChangeOwnerLogic(c, svc)
	res, err := l.ChangeOwner(req)

	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetGroupInfo
// @Summary 查询群信息
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetGroupInfoRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupInfoResponse}
// @Router	/app/group-info [post]
func GetGroupInfo(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.GetGroupInfoRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	req.PersonId = userId.(string)
	if req.DisPlayNum == 0 {
		req.DisPlayNum = biz.DisPlayNum
	}

	res, err := svc.GetGroupInfoHttp(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)

}

// GetGroupPubInfo
// @Summary 查询群公开信息
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetGroupPubInfoRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupPubInfoResponse}
// @Router	/app/group-pub-info [post]
func GetGroupPubInfo(c *gin.Context) {
	req := &types.GetGroupPubInfoRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.GetGroupPubInfoSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)

}

// GetGroupList
// @Summary 查询群列表
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetGroupListRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupListResponse}
// @Router	/app/group-list [post]
func GetGroupList(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.GetGroupListRequest{}
	req.PersonId = userId.(string)

	res, err := svc.GetGroupListSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)

}

// GetGroupMemberList
// @Summary 查询群成员列表
// @Author chy@33.cn
// @Tags group 群成员信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetGroupMemberListRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupMemberListResponse}
// @Router	/app/group-member-list [post]
func GetGroupMemberList(c *gin.Context) {
	req := &types.GetGroupMemberListRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.GetGroupMemberListSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetGroupMemberInfo
// @Summary 查询群成员信息
// @Author chy@33.cn
// @Tags group 群成员信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetGroupMemberInfoRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupMemberInfoResponse}
// @Router	/app/group-member-info [post]
func GetGroupMemberInfo(c *gin.Context) {
	req := &types.GetGroupMemberInfoRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.GetGroupMemberInfoSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// UpdateGroupName
// @Summary 更新群名称
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.UpdateGroupNameRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.UpdateGroupNameResponse}
// @Router	/app/name [post]
func UpdateGroupName(c *gin.Context) {
	req := &types.UpdateGroupNameRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.UpdateGroupNameSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// UpdateGroupAvatar
// @Summary 更新群头像
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.UpdateGroupAvatarRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.UpdateGroupAvatarResponse}
// @Router	/app/avatar [post]
func UpdateGroupAvatar(c *gin.Context) {
	req := &types.UpdateGroupAvatarRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.UpdateGroupAvatarSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// UpdateGroupJoinType
// @Summary 更新加群设置
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.UpdateGroupJoinTypeRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.UpdateGroupJoinTypeResponse}
// @Router	/app/joinType [post]
func UpdateGroupJoinType(c *gin.Context) {
	req := &types.UpdateGroupJoinTypeRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.UpdateGroupJoinTypeSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// UpdateGroupFriendType
// @Summary 更新群内加好友设置
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.UpdateGroupFriendTypeRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.UpdateGroupFriendTypeResponse}
// @Router	/app/friendType [post]
func UpdateGroupFriendType(c *gin.Context) {
	req := &types.UpdateGroupFriendTypeRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.UpdateGroupFriendTypeSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// UpdateGroupMemberName
// @Summary 更新群成员昵称
// @Author chy@33.cn
// @Tags group 群成员信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.UpdateGroupMemberNameRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.UpdateGroupMemberNameResponse}
// @Router	/app/member/name [post]
func UpdateGroupMemberName(c *gin.Context) {
	req := &types.UpdateGroupMemberNameRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.UpdateMemberNameSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// SetAdmin
// @Summary 设置管理员
// @Author chy@33.cn
// @Tags group 群成员信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.SetAdminRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.SetAdminResponse}
// @Router	/app/member/type [post]
func SetAdmin(c *gin.Context) {
	req := &types.SetAdminRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.SetAdminSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// UpdateGroupMuteType
// @Summary 更新群禁言设置
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.UpdateGroupMuteTypeRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.UpdateGroupMuteTypeResponse}
// @Router	/app/muteType [post]
func UpdateGroupMuteType(c *gin.Context) {
	req := &types.UpdateGroupMuteTypeRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.UpdateGroupMuteTypeSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// SetMembersMuteTime
// @Summary 更新群成员禁言时间
// @Author chy@33.cn
// @Tags group 禁言
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.UpdateGroupMemberMuteTimeRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.UpdateGroupMemberMuteTimeResponse}
// @Router	/app/member/muteTime [post]
func SetMembersMuteTime(c *gin.Context) {
	req := &types.UpdateGroupMemberMuteTimeRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.UpdateMembersMuteTimeSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetMuteList
// @Summary 查询群内被禁言成员名单
// @Author chy@33.cn
// @Tags group 禁言
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetMuteListRequest false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetMuteListResponse}
// @Router	/app/mute-list [post]
func GetMuteList(c *gin.Context) {
	req := &types.GetMuteListRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.GetMuteListSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetGroupInfoById
// @Summary 查询群信息
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param id path integer true "群id"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupInfoResponse}
// @Router	/app/group/{id} [get]
func GetGroupInfoById(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.GetGroupInfoRequest{}
	err := c.ShouldBindUri(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	req.PersonId = userId.(string)
	if req.DisPlayNum == 0 {
		req.DisPlayNum = biz.DisPlayNum
	}

	res, err := svc.GetGroupInfoHttp(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetGroups
// @Summary 查询群列表
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupListResponse}
// @Router	/app/groups [get]
func GetGroups(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.GetGroupListRequest{}
	req.PersonId = userId.(string)

	res, err := svc.GetGroupListSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)

}

// GetGroupMemberInfoByUri
// @Summary 查询群成员信息
// @Author chy@33.cn
// @Tags group 群成员信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param id path integer true "群id"
// @Param memberId path integer true "群成员id"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupMemberInfoResponse}
// @Router	/app/group/{id}/member/{memberId} [get]
func GetGroupMemberInfoByUri(c *gin.Context) {
	req := &types.GetGroupMemberInfoRequest{}
	err := c.ShouldBindUri(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.GetGroupMemberInfoSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetGroupMemberListByUri
// @Summary 查询群成员列表
// @Author chy@33.cn
// @Tags group 群成员信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param id path integer true "群id"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupMemberListResponse}
// @Router	/app/group/{id}/members [get]
func GetGroupMemberListByUri(c *gin.Context) {
	req := &types.GetGroupMemberListRequest{}
	err := c.ShouldBindUri(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	if req.Id == 0 && req.IdStr == "" {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError))
		return
	}

	if req.IdStr != "" {
		req.Id = util.ToInt64(req.IdStr)
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}

	req.PersonId = userId.(string)

	res, err := svc.GetGroupMemberListSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetGroupInfoByCondition
// @Summary 搜索群列表
// @Author chy@33.cn
// @Tags group 群信息
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body types.GetGroupInfoByConditionReq false "body"
// @Success 200 {object} types.GeneralResponse{data=types.GetGroupInfoByConditionResp}
// @Router	/app/group-search [post]
func GetGroupInfoByCondition(c *gin.Context) {
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req := &types.GetGroupInfoByConditionReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.GetGroupInfoByConditionSvc(c, req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}
