package types

type GeneralResp struct {
	Result  int         `json:"result"`
	Message int         `json:"message"`
	Data    interface{} `json:"data"`
}

type GroupMember struct {
	// 用户 ID
	MemberId string `json:"memberId" form:"memberId"`
	// 用户群昵称
	MemberName string `json:"memberName" form:"memberName"`
	// 用户角色，2=群主，1=管理员，0=群员，10=退群
	MemberType int32 `json:"memberType" form:"memberType"`
	// 该用户被禁言结束的时间 9223372036854775807=永久禁言
	MemberMuteTime int64 `json:"memberMuteTime"`
}

type GroupInfo struct {
	// 群 ID
	Id    int64  `json:"id" form:"id"`
	IdStr string `json:"idStr"`
	// 群显示的 ID
	MarkId string `json:"markId" form:"markId"`
	// 群名称 加密的
	Name string `json:"name" form:"name"`
	// 公开的群名称 不加密的
	PublicName string `json:"publicName"`
	// 头像 url
	Avatar    string `json:"avatar" form:"avatar"`
	Introduce string `json:"introduce" form:"introduce"`
	// 群主 信息
	Owner *GroupMember `json:"owner" form:"owner"`
	// 本人在群内的信息
	Person *GroupMember `json:"person" form:"person"`
	// 群人数
	MemberNum int32 `json:"memberNum" form:"memberNum"`
	// 群人数上限
	Maximum int32 `json:"maximum" form:"maximum"`
	// 群状态，0=正常 1=封禁 2=解散
	Status int32 `json:"status" form:"status"`
	// 群创建时间
	CreateTime int64 `json:"createTime" form:"createTime"`
	// 加群方式，0=无需审批（默认），1=禁止加群，群主和管理员邀请加群, 2=普通人邀请需要审批,群主和管理员直接加群
	JoinType int32 `json:"joinType" form:"joinType"`
	// 禁言， 0=全员可发言， 1=全员禁言(除群主和管理员)
	MuteType int32 `json:"muteType" form:"muteType"`
	// 加好友限制， 0=群内可加好友，1=群内禁止加好友
	FriendType int32 `json:"friendType"`
	// 群内当前被禁言的人数
	MuteNum int32 `json:"muteNum"`
	// 群内管理员数量
	AdminNum int32 `json:"adminNum"`
	//
	AESKey string `json:"key"`
	// 群类型 (0: 普通群, 1: 全员群, 2: 部门群, 3: 藏品群)
	GroupType int32 `json:"groupType"`
}

type CreateGroupReq struct {
	Name      string   `json:"name" form:"name"`
	Avatar    string   `json:"avatar" form:"avatar"`
	Introduce string   `json:"introduce" form:"introduce"`
	MemberIds []string `json:"memberIds" form:"memberIds"`
}

type CreateGroupResp struct {
	*GroupInfo
	// 群成员
	Members []*GroupMember `json:"members" form:"members"`
}

type CreateNFTGroupReq struct {
	Name      string
	Avatar    string   `json:"avatar" form:"avatar"`
	Introduce string   `json:"introduce" form:"introduce"`
	MemberIds []string `json:"memberIds" form:"memberIds"`

	// 持有条件
	Condition *Condition `json:"condition" form:"condition"`
}

type Condition struct {
	// 持有条件，0=持有其中之一（默认），1=需全部持有
	Type int32 `json:"type" form:"type"`
	// 指定藏品
	NFT []*NFT `json:"nft" form:"nft"`
}

type NFT struct {
	Type int32  `json:"type" form:"type"`
	Name string `json:"name" form:"name"`
	ID   string `json:"id" form:"id"`
}

type CreateNFTGroupResp struct {
	*GroupInfo
	// 群成员
	Members []*GroupMember `json:"members" form:"members"`
	// 持有条件
	//Condition *Condition `json:"condition" form:"condition"`
}

type GetNFTGroupExtInfoReq struct {
	// 群 ID
	Id int64 `json:"id" form:"id"`
}

type GetNFTGroupExtInfoResp struct {
	// 持有条件
	Condition *Condition `json:"condition" form:"condition"`
}

type InviteGroupMembersReq struct {
	Id int64 `json:"id" form:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr        string        `json:"idStr"`
	Inviter      GroupMember   `json:"-"`
	NewMembers   []GroupMember `json:"-"`
	NewMemberIds []string      `json:"newMemberIds" form:"newMemberIds" binding:"required"`
}

type InviteGroupMembersResp struct {
	Id        int64  `json:"id" form:"id" example:"123821199217135616"`
	IdStr     string `json:"idStr"`
	MemberNum int32  `json:"memberNum" form:"memberNum" example:"5"`
	//Inviter    GroupMember   `json:"inviter" form:"inviter"`
	//NewMembers []GroupMember `json:"newMembers" form:"newMembers"`
}

type GetGroupInfoReq struct {
	Id int64 `json:"id" uri:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr      string `json:"idStr"`
	PersonId   string `json:"-"`
	DisPlayNum int64  `json:"-"`
}

type GetGroupInfoResp struct {
	*GroupInfo
	Members []*GroupMember `json:"members" form:"members"`
}

type GetGroupListReq struct {
	PersonId string `json:"-"`
}

type GetGroupListResp struct {
	Groups []*GroupInfo `json:"groups"`
}

type GetGroupMemberListReq struct {
	Id int64 `json:"id" uri:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr    string `json:"idStr"`
	PersonId string `json:"-"`
}

type GetGroupMemberListResp struct {
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr   string         `json:"idStr"`
	Members []*GroupMember `json:"members"`
}

type GetGroupMemberInfoReq struct {
	Id int64 `json:"id" uri:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr    string `json:"idStr"`
	MemberId string `json:"memberId" uri:"memberId" binding:"required"`
	PersonId string `json:"-"`
}

type GetGroupMemberInfoResp struct {
	*GroupMember
}

type GroupExitReq struct {
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr    string `json:"idStr"`
	PersonId string `json:"-"`
}

type GroupExitResp struct {
}

type GroupRemoveReq struct {
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr     string   `json:"idStr"`
	MemberIds []string `json:"memberIds" binding:"required"`
	PersonId  string   `json:"-"`
}

type GroupRemoveResp struct {
	// 群人数
	MemberNum int32 `json:"memberNum" form:"memberNum"`
	// 成功被踢的成员列表
	MemberIds []string `json:"memberIds"`
}

type GroupDisbandReq struct {
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr    string `json:"idStr"`
	PersonId string `json:"-"`
}

type GroupDisbandResp struct {
}

type UpdateGroupNameReq struct {
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr      string `json:"idStr"`
	PersonId   string `json:"-"`
	Name       string `json:"name"`
	PublicName string `json:"publicName"`
}

type UpdateGroupNameResp struct {
}

type UpdateGroupAvatarReq struct {
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr    string `json:"idStr"`
	PersonId string `json:"-"`
	Avatar   string `json:"avatar"`
}

type UpdateGroupAvatarResp struct {
}

type UpdateGroupMemberNameReq struct {
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr      string `json:"idStr"`
	PersonId   string `json:"-"`
	MemberName string `json:"memberName"`
}

type UpdateGroupMemberNameResp struct {
}

type UpdateGroupJoinTypeReq struct {
	// 群 ID
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr    string `json:"idStr"`
	PersonId string `json:"-"`
	// 加群方式，0=无需审批（默认），1=禁止加群，群主和管理员邀请加群, 2=普通人邀请需要审批,群主和管理员直接加群
	JoinType int32 `json:"joinType"  binding:"oneof=0 1 2"`
}

type UpdateGroupJoinTypeResp struct {
}

type UpdateGroupFriendTypeReq struct {
	// 群 ID
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr    string `json:"idStr"`
	PersonId string `json:"-"`
	// 加好友限制， 0=群内可加好友，1=群内禁止加好友
	FriendType int32 `json:"friendType"  binding:"oneof=0 1"`
}

type UpdateGroupFriendTypeResp struct {
}

type ChangeOwnerReq struct {
	// 群 ID
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr string `json:"idStr"`
	// 被转让为群主的群成员 ID
	MemberId string `json:"memberId" binding:"required"`
	PersonId string `json:"-"`
}

type ChangeOwnerResp struct {
}

// JoinGroupReq 扫二维码加群
type JoinGroupReq struct {
	Id        int64  `json:"id"`
	IdStr     string `json:"idStr"`
	InviterId string `json:"inviterId"`
	PersonId  string `json:"-"`
}

type JoinGroupResp struct {
	Id    int64  `json:"id"`
	IdStr string `json:"idStr"`
}

type SetAdminReq struct {
	// 群 ID
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr string `json:"idStr"`
	// 被设置的群成员 ID
	MemberId string `json:"memberId" binding:"required"`
	PersonId string `json:"-"`
	// 用户角色 0=群员, 1=管理员
	MemberType int32 `json:"memberType" binding:"oneof=0 1"`
}

type SetAdminResp struct {
}

type UpdateGroupMuteTypeReq struct {
	// 群 ID
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr string `json:"idStr"`
	// 禁言， 0=全员可发言， 1=全员禁言(除群主和管理员)
	MuteType int32  `json:"muteType" binding:"oneof=0 1"`
	PersonId string `json:"-"`
}

type UpdateGroupMuteTypeResp struct {
}

type UpdateGroupMemberMuteTimeReq struct {
	// 群 ID
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr string `json:"idStr"`
	// 被禁言的群员 ID
	MemberIds []string `json:"memberIds" binding:"required"`
	// 禁言持续时间, 传9223372036854775807=永久禁言, 0=解除禁言
	MuteTime int64  `json:"muteTime"`
	PersonId string `json:"-"`
}

type UpdateGroupMemberMuteTimeResp struct {
	Members []*GroupMember `json:"members"`
}

type GetMuteListReq struct {
	// 群 ID
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr    string `json:"idStr"`
	PersonId string `json:"-"`
}

type GetMuteListResp struct {
	Members []*GroupMember `json:"members"`
}

type GetGroupPubInfoReq struct {
	// 群 ID
	Id int64 `json:"id"`
	// 如果同时填了 idStr, 则优先选择 idStr
	IdStr    string `json:"idStr"`
	PersonId string `json:"-"`
}

type GetGroupPubInfoResp struct {
	*GroupInfo
}

type GetGroupInfoByConditionReq struct {
	// 查询方法 0:groupMarkId, 1:groupId
	Tp       int32  `json:"tp" binding:"oneof=0 1"`
	Query    string `json:"query" binding:"required"`
	PersonId string `json:"-"`
}

type GetGroupInfoByConditionResp struct {
	Groups []*GroupInfo `json:"groups"`
}

// ---群审批------------------

// GroupApplyInfo 群审批信息
type GroupApplyInfo struct {
	// 审批 ID
	ApplyId string `json:"applyId,omitempty"`
	// 群 ID
	GroupId string `json:"id,omitempty"`
	// 邀请人 ID, 空表示是自己主动申请的
	InviterId string `json:"inviterId,omitempty"`
	// 申请加入人 ID
	MemberId string `json:"memberId,omitempty"`
	// 申请备注
	ApplyNote string `json:"applyNote,omitempty"`
	// 审批人 ID
	OperatorId string `json:"operatorId,omitempty"`
	// 审批情况 0=待审批, 1=审批通过, 2=审批不通过, 10=审批忽略
	ApplyStatus int32 `json:"applyStatus,omitempty"`
	// 拒绝原因
	RejectReason string `json:"rejectReason,omitempty"`
	// 创建时间 ms
	CreateTime int64 `json:"createTime,omitempty"`
	// 修改时间 ms
	UpdateTime int64 `json:"updateTime,omitempty"`
}

// CreateGroupApplyReq 创建群审批请求
type CreateGroupApplyReq struct {
	// 群 ID
	Id string `json:"id,omitempty" binding:"required"`
	// 申请备注
	ApplyNote string `json:"applyNote,omitempty"`

	PersonId string `json:"-"`
}

// CreateGroupApplyResp 创建群审批响应
type CreateGroupApplyResp struct {
}

// GetGroupApplyByIdReq 查询群审批请求
type GetGroupApplyByIdReq struct {
	// 审批 ID
	ApplyId string `json:"applyId" binding:"required"`
	// 群 ID
	Id string `json:"id" binding:"required"`

	PersonId string `json:"-"`
}

// GetGroupApplysReq 查询群审批列表请求
type GetGroupApplysReq struct {
	// 群 ID
	Id string `json:"id" binding:"required"`
	// 每页记录数
	Count int32 `json:"count" binding:"required"`
	// 当前审批记录数量
	Offset int32 `json:"offset"`

	PersonId string `json:"-"`
}

// GetGroupApplysResp 查询群审批响应
type GetGroupApplysResp struct {
	GroupApplys []*GroupApplyInfo `json:"applys"`
}

// AcceptGroupApplyReq 接受加群审批请求
type AcceptGroupApplyReq struct {
	// 审批 ID
	ApplyId string `json:"applyId" binding:"required"`
	// 群 ID
	Id string `json:"id" binding:"required"`

	PersonId string `json:"-"`
}

type AcceptGroupApplyResp struct {
}

// RejectGroupApplyReq 拒绝加群审批请求
type RejectGroupApplyReq struct {
	// 审批 ID
	ApplyId string `json:"applyId" binding:"required"`
	// 群 ID
	Id string `json:"id" binding:"required"`
	// 拒绝原因
	RejectReason string `json:"rejectReason"`

	PersonId string `json:"-"`
}

type RejectGroupApplyResp struct {
}

type IsWhitelistResp struct {
	Exist bool `json:"exist" form:"exist"`
}
