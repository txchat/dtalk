### 1. "移交群主"

1. route definition

- Url: /group/app/change-owner
- Method: POST
- Request: `ChangeOwnerReq`
- Response: `ChangeOwnerResp`

2. request definition



```golang
type ChangeOwnerReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	MemberId string `json:"memberId"`
}
```


3. response definition



```golang
type ChangeOwnerResp struct {
}
```

### 2. "创建群"

1. route definition

- Url: /group/app/create-group
- Method: POST
- Request: `CreateGroupReq`
- Response: `CreateGroupResp`

2. request definition



```golang
type CreateGroupReq struct {
	Name string `json:"name" form:"name"`
	Avatar string `json:"avatar" form:"avatar"`
	Introduce string `json:"introduce" form:"introduce"`
	MemberIds []string `json:"memberIds" form:"memberIds"`
}
```


3. response definition



```golang
type CreateGroupResp struct {
	Id int64 `json:"id" form:"id"`
	IdStr string `json:"idStr"`
	MarkId string `json:"markId" form:"markId"`
	Name string `json:"name" form:"name"` //群成员可见的群名称（加密）
	PublicName string `json:"publicName"` //对外公开群名称
	Avatar string `json:"avatar" form:"avatar"`
	Introduce string `json:"introduce" form:"introduce"`
	Owner *GroupMember `json:"owner" form:"owner"`
	Person *GroupMember `json:"person" form:"person"` //个人群内的信息
	MemberNum int32 `json:"memberNum" form:"memberNum"`
	Maximum int32 `json:"maximum" form:"maximum"`
	Status int32 `json:"status,options=0|1|2" form:"status"` //群状态:0-&gt;正常, 1-&gt;封禁, 2-&gt;解散
	CreateTime int64 `json:"createTime" form:"createTime"`
	JoinType int32 `json:"joinType,options=0|1|2" form:"joinType"` //加群方式:0-&gt;无需审批, 1-&gt;仅群主和管理员邀请加群, 2-&gt;普通成员邀请需审批
	MuteType int32 `json:"muteType" form:"muteType"` //0-&gt;全员可发言, 1-&gt;全员禁言(除群主和管理员)
	FriendType int32 `json:"friendType"` //0-&gt;群内可加好友, 1-&gt;群内禁止加好友
	MuteNum int32 `json:"muteNum"` //禁言人数
	AdminNum int32 `json:"adminNum"` //管理员人数
	AESKey string `json:"key"`
	GroupType int32 `json:"groupType"` //群类型：0-&gt;普通群, 1-&gt;企业群, 2-&gt;部门群
	Members []*GroupMember `json:"members" form:"members"`
}

type GroupInfo struct {
	Id int64 `json:"id" form:"id"`
	IdStr string `json:"idStr"`
	MarkId string `json:"markId" form:"markId"`
	Name string `json:"name" form:"name"` //群成员可见的群名称（加密）
	PublicName string `json:"publicName"` //对外公开群名称
	Avatar string `json:"avatar" form:"avatar"`
	Introduce string `json:"introduce" form:"introduce"`
	Owner *GroupMember `json:"owner" form:"owner"`
	Person *GroupMember `json:"person" form:"person"` //个人群内的信息
	MemberNum int32 `json:"memberNum" form:"memberNum"`
	Maximum int32 `json:"maximum" form:"maximum"`
	Status int32 `json:"status,options=0|1|2" form:"status"` //群状态:0-&gt;正常, 1-&gt;封禁, 2-&gt;解散
	CreateTime int64 `json:"createTime" form:"createTime"`
	JoinType int32 `json:"joinType,options=0|1|2" form:"joinType"` //加群方式:0-&gt;无需审批, 1-&gt;仅群主和管理员邀请加群, 2-&gt;普通成员邀请需审批
	MuteType int32 `json:"muteType" form:"muteType"` //0-&gt;全员可发言, 1-&gt;全员禁言(除群主和管理员)
	FriendType int32 `json:"friendType"` //0-&gt;群内可加好友, 1-&gt;群内禁止加好友
	MuteNum int32 `json:"muteNum"` //禁言人数
	AdminNum int32 `json:"adminNum"` //管理员人数
	AESKey string `json:"key"`
	GroupType int32 `json:"groupType"` //群类型：0-&gt;普通群, 1-&gt;企业群, 2-&gt;部门群
}
```

### 3. "获取群信息"

1. route definition

- Url: /group/app/group-info
- Method: POST
- Request: `GetGroupInfoReq`
- Response: `GetGroupInfoResp`

2. request definition



```golang
type GetGroupInfoReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
}
```


3. response definition



```golang
type GetGroupInfoResp struct {
	Id int64 `json:"id" form:"id"`
	IdStr string `json:"idStr"`
	MarkId string `json:"markId" form:"markId"`
	Name string `json:"name" form:"name"` //群成员可见的群名称（加密）
	PublicName string `json:"publicName"` //对外公开群名称
	Avatar string `json:"avatar" form:"avatar"`
	Introduce string `json:"introduce" form:"introduce"`
	Owner *GroupMember `json:"owner" form:"owner"`
	Person *GroupMember `json:"person" form:"person"` //个人群内的信息
	MemberNum int32 `json:"memberNum" form:"memberNum"`
	Maximum int32 `json:"maximum" form:"maximum"`
	Status int32 `json:"status,options=0|1|2" form:"status"` //群状态:0-&gt;正常, 1-&gt;封禁, 2-&gt;解散
	CreateTime int64 `json:"createTime" form:"createTime"`
	JoinType int32 `json:"joinType,options=0|1|2" form:"joinType"` //加群方式:0-&gt;无需审批, 1-&gt;仅群主和管理员邀请加群, 2-&gt;普通成员邀请需审批
	MuteType int32 `json:"muteType" form:"muteType"` //0-&gt;全员可发言, 1-&gt;全员禁言(除群主和管理员)
	FriendType int32 `json:"friendType"` //0-&gt;群内可加好友, 1-&gt;群内禁止加好友
	MuteNum int32 `json:"muteNum"` //禁言人数
	AdminNum int32 `json:"adminNum"` //管理员人数
	AESKey string `json:"key"`
	GroupType int32 `json:"groupType"` //群类型：0-&gt;普通群, 1-&gt;企业群, 2-&gt;部门群
	Members []*GroupMember `json:"members" form:"members"`
}

type GroupInfo struct {
	Id int64 `json:"id" form:"id"`
	IdStr string `json:"idStr"`
	MarkId string `json:"markId" form:"markId"`
	Name string `json:"name" form:"name"` //群成员可见的群名称（加密）
	PublicName string `json:"publicName"` //对外公开群名称
	Avatar string `json:"avatar" form:"avatar"`
	Introduce string `json:"introduce" form:"introduce"`
	Owner *GroupMember `json:"owner" form:"owner"`
	Person *GroupMember `json:"person" form:"person"` //个人群内的信息
	MemberNum int32 `json:"memberNum" form:"memberNum"`
	Maximum int32 `json:"maximum" form:"maximum"`
	Status int32 `json:"status,options=0|1|2" form:"status"` //群状态:0-&gt;正常, 1-&gt;封禁, 2-&gt;解散
	CreateTime int64 `json:"createTime" form:"createTime"`
	JoinType int32 `json:"joinType,options=0|1|2" form:"joinType"` //加群方式:0-&gt;无需审批, 1-&gt;仅群主和管理员邀请加群, 2-&gt;普通成员邀请需审批
	MuteType int32 `json:"muteType" form:"muteType"` //0-&gt;全员可发言, 1-&gt;全员禁言(除群主和管理员)
	FriendType int32 `json:"friendType"` //0-&gt;群内可加好友, 1-&gt;群内禁止加好友
	MuteNum int32 `json:"muteNum"` //禁言人数
	AdminNum int32 `json:"adminNum"` //管理员人数
	AESKey string `json:"key"`
	GroupType int32 `json:"groupType"` //群类型：0-&gt;普通群, 1-&gt;企业群, 2-&gt;部门群
}
```

### 4. "获取入群列表"

1. route definition

- Url: /group/app/group-list
- Method: POST
- Request: `GetGroupListReq`
- Response: `GetGroupListResp`

2. request definition



```golang
type GetGroupListReq struct {
}
```


3. response definition



```golang
type GetGroupListResp struct {
	Groups []*GroupInfo `json:"groups"`
}
```

### 5. "获取群成员信息"

1. route definition

- Url: /group/app/group-member-info
- Method: POST
- Request: `GetGroupMemberInfoReq`
- Response: `GetGroupMemberInfoResp`

2. request definition



```golang
type GetGroupMemberInfoReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	MemberId string `json:"memberId" uri:"memberId"`
}
```


3. response definition



```golang
type GetGroupMemberInfoResp struct {
	MemberId string `json:"memberId" form:"memberId"`
	MemberName string `json:"memberName" form:"memberName"`
	MemberType int32 `json:"memberType" form:"memberType"` //用户角色:0-&gt;群员, 1-&gt;管理员, 2-&gt;群主, 10-&gt;退群
	MemberMuteTime int64 `json:"memberMuteTime"` // 禁言截止时间: 9223372036854775807-&gt;永久禁言
}

type GroupMember struct {
	MemberId string `json:"memberId" form:"memberId"`
	MemberName string `json:"memberName" form:"memberName"`
	MemberType int32 `json:"memberType" form:"memberType"` //用户角色:0-&gt;群员, 1-&gt;管理员, 2-&gt;群主, 10-&gt;退群
	MemberMuteTime int64 `json:"memberMuteTime"` // 禁言截止时间: 9223372036854775807-&gt;永久禁言
}
```

### 6. "获取群成员列表"

1. route definition

- Url: /group/app/group-member-list
- Method: POST
- Request: `GetGroupMemberListReq`
- Response: `GetGroupMemberListResp`

2. request definition



```golang
type GetGroupMemberListReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
}
```


3. response definition



```golang
type GetGroupMemberListResp struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	Members []*GroupMember `json:"members"`
}
```

### 7. "获取群禁言列表"

1. route definition

- Url: /group/app/mute-list
- Method: POST
- Request: `GetMuteListReq`
- Response: `GetMuteListResp`

2. request definition



```golang
type GetMuteListReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
}
```


3. response definition



```golang
type GetMuteListResp struct {
	Members []*GroupMember `json:"members"`
}
```

### 8. "获取群公开信息"

1. route definition

- Url: /group/app/group-pub-info
- Method: POST
- Request: `GetGroupPubInfoReq`
- Response: `GetGroupPubInfoResp`

2. request definition



```golang
type GetGroupPubInfoReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
}
```


3. response definition



```golang
type GetGroupPubInfoResp struct {
	Id int64 `json:"id" form:"id"`
	IdStr string `json:"idStr"`
	MarkId string `json:"markId" form:"markId"`
	Name string `json:"name" form:"name"` //群成员可见的群名称（加密）
	PublicName string `json:"publicName"` //对外公开群名称
	Avatar string `json:"avatar" form:"avatar"`
	Introduce string `json:"introduce" form:"introduce"`
	Owner *GroupMember `json:"owner" form:"owner"`
	Person *GroupMember `json:"person" form:"person"` //个人群内的信息
	MemberNum int32 `json:"memberNum" form:"memberNum"`
	Maximum int32 `json:"maximum" form:"maximum"`
	Status int32 `json:"status,options=0|1|2" form:"status"` //群状态:0-&gt;正常, 1-&gt;封禁, 2-&gt;解散
	CreateTime int64 `json:"createTime" form:"createTime"`
	JoinType int32 `json:"joinType,options=0|1|2" form:"joinType"` //加群方式:0-&gt;无需审批, 1-&gt;仅群主和管理员邀请加群, 2-&gt;普通成员邀请需审批
	MuteType int32 `json:"muteType" form:"muteType"` //0-&gt;全员可发言, 1-&gt;全员禁言(除群主和管理员)
	FriendType int32 `json:"friendType"` //0-&gt;群内可加好友, 1-&gt;群内禁止加好友
	MuteNum int32 `json:"muteNum"` //禁言人数
	AdminNum int32 `json:"adminNum"` //管理员人数
	AESKey string `json:"key"`
	GroupType int32 `json:"groupType"` //群类型：0-&gt;普通群, 1-&gt;企业群, 2-&gt;部门群
}

type GroupInfo struct {
	Id int64 `json:"id" form:"id"`
	IdStr string `json:"idStr"`
	MarkId string `json:"markId" form:"markId"`
	Name string `json:"name" form:"name"` //群成员可见的群名称（加密）
	PublicName string `json:"publicName"` //对外公开群名称
	Avatar string `json:"avatar" form:"avatar"`
	Introduce string `json:"introduce" form:"introduce"`
	Owner *GroupMember `json:"owner" form:"owner"`
	Person *GroupMember `json:"person" form:"person"` //个人群内的信息
	MemberNum int32 `json:"memberNum" form:"memberNum"`
	Maximum int32 `json:"maximum" form:"maximum"`
	Status int32 `json:"status,options=0|1|2" form:"status"` //群状态:0-&gt;正常, 1-&gt;封禁, 2-&gt;解散
	CreateTime int64 `json:"createTime" form:"createTime"`
	JoinType int32 `json:"joinType,options=0|1|2" form:"joinType"` //加群方式:0-&gt;无需审批, 1-&gt;仅群主和管理员邀请加群, 2-&gt;普通成员邀请需审批
	MuteType int32 `json:"muteType" form:"muteType"` //0-&gt;全员可发言, 1-&gt;全员禁言(除群主和管理员)
	FriendType int32 `json:"friendType"` //0-&gt;群内可加好友, 1-&gt;群内禁止加好友
	MuteNum int32 `json:"muteNum"` //禁言人数
	AdminNum int32 `json:"adminNum"` //管理员人数
	AESKey string `json:"key"`
	GroupType int32 `json:"groupType"` //群类型：0-&gt;普通群, 1-&gt;企业群, 2-&gt;部门群
}
```

### 9. "解散群"

1. route definition

- Url: /group/app/group-disband
- Method: POST
- Request: `GroupDisbandReq`
- Response: `GroupDisbandResp`

2. request definition



```golang
type GroupDisbandReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
}
```


3. response definition



```golang
type GroupDisbandResp struct {
}
```

### 10. "退出群"

1. route definition

- Url: /group/app/group-exit
- Method: POST
- Request: `GroupExitReq`
- Response: `GroupExitResp`

2. request definition



```golang
type GroupExitReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
}
```


3. response definition



```golang
type GroupExitResp struct {
}
```

### 11. "将成员移除群"

1. route definition

- Url: /group/app/group-remove
- Method: POST
- Request: `GroupRemoveReq`
- Response: `GroupRemoveResp`

2. request definition



```golang
type GroupRemoveReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	MemberIds []string `json:"memberIds"`
}
```


3. response definition



```golang
type GroupRemoveResp struct {
	MemberNum int32 `json:"memberNum" form:"memberNum"`
	MemberIds []string `json:"memberIds"`
}
```

### 12. "邀请新成员"

1. route definition

- Url: /group/app/invite-group-members
- Method: POST
- Request: `InviteGroupMembersReq`
- Response: `InviteGroupMembersResp`

2. request definition



```golang
type InviteGroupMembersReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	NewMemberIds []string `json:"newMemberIds" form:"newMemberIds"`
}
```


3. response definition



```golang
type InviteGroupMembersResp struct {
	Id int64 `json:"id,optional" form:"id,optional" example:"123821199217135616"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	MemberNum int32 `json:"memberNum" form:"memberNum" example:"5"`
}
```

### 13. "主动加群"

1. route definition

- Url: /group/app/join-group
- Method: POST
- Request: `JoinGroupReq`
- Response: `JoinGroupResp`

2. request definition



```golang
type JoinGroupReq struct {
	Id int64 `json:"id"`
	IdStr string `json:"idStr"`
	InviterId string `json:"inviterId"`
}
```


3. response definition



```golang
type JoinGroupResp struct {
	Id int64 `json:"id"`
	IdStr string `json:"idStr"`
}
```

### 14. "设置群管理员"

1. route definition

- Url: /group/app/member/type
- Method: POST
- Request: `SetAdminReq`
- Response: `SetAdminResp`

2. request definition



```golang
type SetAdminReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	MemberId string `json:"memberId"`
	MemberType int32 `json:"memberType" enums:"0,1"` //用户角色:0-&gt;群员, 1-&gt;管理员
}
```


3. response definition



```golang
type SetAdminResp struct {
}
```

### 15. "更新群头像"

1. route definition

- Url: /group/app/avatar
- Method: POST
- Request: `UpdateGroupAvatarReq`
- Response: `UpdateGroupAvatarResp`

2. request definition



```golang
type UpdateGroupAvatarReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	Avatar string `json:"avatar"`
}
```


3. response definition



```golang
type UpdateGroupAvatarResp struct {
}
```

### 16. "更新群加好友权限"

1. route definition

- Url: /group/app/friendType
- Method: POST
- Request: `UpdateGroupFriendTypeReq`
- Response: `UpdateGroupFriendTypeResp`

2. request definition



```golang
type UpdateGroupFriendTypeReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	FriendType int32 `json:"friendType,options=0|1"` //0-&gt;群内可加好友, 1-&gt;群内禁止加好友
}
```


3. response definition



```golang
type UpdateGroupFriendTypeResp struct {
}
```

### 17. "更新群入群权限"

1. route definition

- Url: /group/app/joinType
- Method: POST
- Request: `UpdateGroupJoinTypeReq`
- Response: `UpdateGroupJoinTypeResp`

2. request definition



```golang
type UpdateGroupJoinTypeReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	JoinType int32 `json:"joinType,options=0|1|2"` //加群方式:0-&gt;无需审批, 1-&gt;仅群主和管理员邀请加群, 2-&gt;普通成员邀请需审批
}
```


3. response definition



```golang
type UpdateGroupJoinTypeResp struct {
}
```

### 18. "禁言或解禁成员"

1. route definition

- Url: /group/app/member/muteTime
- Method: POST
- Request: `UpdateGroupMemberMuteTimeReq`
- Response: `UpdateGroupMemberMuteTimeResp`

2. request definition



```golang
type UpdateGroupMemberMuteTimeReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	MemberIds []string `json:"memberIds"`
	MuteTime int64 `json:"muteTime"` //禁言截止时间: 0-&gt;解除禁言, 9223372036854775807-&gt;永久禁言
}
```


3. response definition



```golang
type UpdateGroupMemberMuteTimeResp struct {
	Members []*GroupMember `json:"members"`
}
```

### 19. "更新群成员名称"

1. route definition

- Url: /group/app/member/name
- Method: POST
- Request: `UpdateGroupMemberNameReq`
- Response: `UpdateGroupMemberNameResp`

2. request definition



```golang
type UpdateGroupMemberNameReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	MemberName string `json:"memberName"`
}
```


3. response definition



```golang
type UpdateGroupMemberNameResp struct {
}
```

### 20. "更新群禁言类型"

1. route definition

- Url: /group/app/muteType
- Method: POST
- Request: `UpdateGroupMuteTypeReq`
- Response: `UpdateGroupMuteTypeResp`

2. request definition



```golang
type UpdateGroupMuteTypeReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	MuteType int32 `json:"muteType,options=0|1"` //0-&gt;全员可发言, 1-&gt;全员禁言(除群主和管理员)
}
```


3. response definition



```golang
type UpdateGroupMuteTypeResp struct {
}
```

### 21. "更新群名"

1. route definition

- Url: /group/app/name
- Method: POST
- Request: `UpdateGroupNameReq`
- Response: `UpdateGroupNameResp`

2. request definition



```golang
type UpdateGroupNameReq struct {
	Id int64 `json:"id,optional"`
	IdStr string `json:"idStr,optional"` // 群 ID, 如果同时填了 idStr, 则优先选择 idStr
	Name string `json:"name"`
	PublicName string `json:"publicName"`
}
```


3. response definition



```golang
type UpdateGroupNameResp struct {
}
```

