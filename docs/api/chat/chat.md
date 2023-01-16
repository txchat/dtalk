### 1. "发起音视频通话"

1. route definition

- Url: /app/start-call
- Method: POST
- Request: `StartCallReq`
- Response: `StartCallResp`

2. request definition



```golang
type StartCallReq struct {
	GroupId string `json:"groupId"`
	Invitees []string `json:"invitees"`
	RTCType int32 `json:"RTCType,options=1|2"` //1-&gt;音频, 2-&gt;视频
}
```


3. response definition



```golang
type StartCallResp struct {
	TraceId int64 `json:"traceId"`
	TraceIdStr string `json:"traceIdStr"`
	RTCType int32 `json:"RTCType"` //1-&gt;音频, 2-&gt;视频
	Invitees []string `json:"invitees"`
	Caller string `json:"caller"`
	CreateTime int64 `json:"createTime"`
	Timeout int32 `json:"timeout"`
	Deadline int64 `json:"deadline"`
	GroupId string `json:"groupId"` // 0-&gt;私聊, ^0-&gt;群id
}
```

### 2. "通话响应-繁忙"

1. route definition

- Url: /app/reply-busy
- Method: POST
- Request: `ReplyBusyReq`
- Response: `ReplyBusyResp`

2. request definition



```golang
type ReplyBusyReq struct {
	TraceId int64 `json:"traceId"`
	TraceIdStr string `json:"traceIdStr,optional"` // 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
}
```


3. response definition



```golang
type ReplyBusyResp struct {
}
```

### 3. "检查通话"

1. route definition

- Url: /app/check-call
- Method: POST
- Request: `CheckCallReq`
- Response: `CheckCallResp`

2. request definition



```golang
type CheckCallReq struct {
	TraceId int64 `json:"traceId"`
	TraceIdStr string `json:"traceIdStr,optional"` // 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
}
```


3. response definition



```golang
type CheckCallResp struct {
	TraceId int64 `json:"traceId"`
	TraceIdStr string `json:"traceIdStr"`
	RTCType int32 `json:"RTCType"` //1-&gt;音频, 2-&gt;视频
	Invitees []string `json:"invitees"`
	Caller string `json:"caller"`
	CreateTime int64 `json:"createTime"`
	Timeout int32 `json:"timeout"`
	Deadline int64 `json:"deadline"`
	GroupId string `json:"groupId"` // 0-&gt;私聊, ^0-&gt;群id
}
```

### 4. "处理通话"

1. route definition

- Url: /app/handle-call
- Method: POST
- Request: `HandleCallReq`
- Response: `HandleCallResp`

2. request definition



```golang
type HandleCallReq struct {
	Answer bool `json:"answer"`
	TraceId int64 `json:"traceId"`
	TraceIdStr string `json:"traceIdStr,optional"` // 如果同时填了 tracedIdStr, 则优先选择 traceIdStr
}
```


3. response definition



```golang
type HandleCallResp struct {
	RoomId int32 `json:"roomId"`
	UserSig string `json:"userSig"`
	PrivateMapKey string `json:"privateMapKey"`
	SDKAppId int32 `json:"sdkAppId"`
}
```

### 5. "撤回消息"

1. route definition

- Url: /app/record/revoke
- Method: POST
- Request: `RevokeReq`
- Response: `RevokeResp`

2. request definition



```golang
type RevokeReq struct {
	Type int `json:"type,optional,options=0|1"` //撤回类型: 0-&gt;私聊, 1-&gt;群聊
	Mid int64 `json:"logId"`
}
```


3. response definition



```golang
type RevokeResp struct {
}
```

### 6. "关注消息"

1. route definition

- Url: /app/record/focus
- Method: POST
- Request: `FocusReq`
- Response: `FocusResp`

2. request definition



```golang
type FocusReq struct {
	Type int `json:"type,optional,options=0|1"` //关注类型: 0-&gt;私聊, 1-&gt;群聊
	Mid int64 `json:"logId"`
}
```


3. response definition



```golang
type FocusResp struct {
}
```

### 7. "同步聊天记录"

1. route definition

- Url: /app/record/sync-record
- Method: POST
- Request: `SyncReq`
- Response: `SyncResp`

2. request definition



```golang
type SyncReq struct {
	MaxCount int64 `json:"count,range=[1:1000]"` // 消息数量
	StartMid int64 `json:"start,optional"` // 消息 ID
}
```


3. response definition



```golang
type SyncResp struct {
	RecordCount int `json:"record_count"` // 聊天记录数量
	Records []string `json:"records"` // 聊天记录 base64 encoding
}
```

### 8. "获取私聊消息"

1. route definition

- Url: /app/record/pri-chat-record
- Method: POST
- Request: `PrivateRecordReq`
- Response: `PrivateRecordResp`

2. request definition



```golang
type PrivateRecordReq struct {
	FromId string `json:"-"`
	TargetId string `json:"targetId"`
	RecordCount int64 `json:"count,range=[1:100]"`
	Mid string `json:"logId"`
}
```


3. response definition



```golang
type PrivateRecordResp struct {
	RecordCount int `json:"record_count"` // 聊天记录数量
	Records []*Record `json:"records"` // 聊天记录
}
```

### 9. "发送消息"

1. route definition

- Url: /record/push
- Method: POST
- Request: `PushReq`
- Response: `PushResp`

2. request definition



```golang
type PushReq struct {
	File string `form:"file"`
}
```


3. response definition



```golang
type PushResp struct {
	Mid int64 `json:"logId"`
	Datetime uint64 `json:"datetime"`
}
```

### 10. "发送消息"

1. route definition

- Url: /record/push2
- Method: POST
- Request: `PushReq`
- Response: `PushResp`

2. request definition



```golang
type PushReq struct {
	File string `form:"file"`
}
```


3. response definition



```golang
type PushResp struct {
	Mid int64 `json:"logId"`
	Datetime uint64 `json:"datetime"`
}
```

### 11. "用户登录"

1. route definition

- Url: /app/user/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginResp`

2. request definition



```golang
type LoginReq struct {
	ConnType int32 `json:"connType"`
}
```


3. response definition



```golang
type LoginResp struct {
	Address string `json:"address"`
}
```

### 12. "移交群主"

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

### 13. "创建群"

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

### 14. "获取群信息"

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

### 15. "获取入群列表"

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

### 16. "获取群成员信息"

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

### 17. "获取群成员列表"

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

### 18. "获取群禁言列表"

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

### 19. "获取群公开信息"

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

### 20. "解散群"

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

### 21. "退出群"

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

### 22. "将成员移除群"

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

### 23. "邀请新成员"

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

### 24. "主动加群"

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

### 25. "设置群管理员"

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

### 26. "更新群头像"

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

### 27. "更新群加好友权限"

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

### 28. "更新群入群权限"

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

### 29. "禁言或解禁成员"

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

### 30. "更新群成员名称"

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

### 31. "更新群禁言类型"

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

### 32. "更新群名"

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

### 33. "获取token"

1. route definition

- Url: /oss/get-token
- Method: POST
- Request: `GetTokenReq`
- Response: `GetTokenResp`

2. request definition



```golang
type GetTokenReq struct {
}
```


3. response definition



```golang
type GetTokenResp struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
	Credentials Credentials `json:"Credentials" xml:"Credentials"`
	AssumedRoleUser AssumedRoleUser `json:"AssumedRoleUser" xml:"AssumedRoleUser"`
}

type Credentials struct {
	AccessKeySecret string `json:"AccessKeySecret" xml:"AccessKeySecret"`
	Expiration string `json:"Expiration" xml:"Expiration"`
	AccessKeyId string `json:"AccessKeyId" xml:"AccessKeyId"`
	SecurityToken string `json:"SecurityToken" xml:"SecurityToken"`
}

type AssumedRoleUser struct {
	AssumedRoleId string `json:"AssumedRoleId" xml:"AssumedRoleId"`
	Arn string `json:"Arn" xml:"Arn"`
}
```

### 34. "获取华为云token"

1. route definition

- Url: /oss/get-huaweiyun-token
- Method: POST
- Request: `GetHWCloudTokenReq`
- Response: `GetHWCloudTokenResp`

2. request definition



```golang
type GetHWCloudTokenReq struct {
}
```


3. response definition



```golang
type GetHWCloudTokenResp struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
	Credentials Credentials `json:"Credentials" xml:"Credentials"`
	AssumedRoleUser AssumedRoleUser `json:"AssumedRoleUser" xml:"AssumedRoleUser"`
}

type Credentials struct {
	AccessKeySecret string `json:"AccessKeySecret" xml:"AccessKeySecret"`
	Expiration string `json:"Expiration" xml:"Expiration"`
	AccessKeyId string `json:"AccessKeyId" xml:"AccessKeyId"`
	SecurityToken string `json:"SecurityToken" xml:"SecurityToken"`
}

type AssumedRoleUser struct {
	AssumedRoleId string `json:"AssumedRoleId" xml:"AssumedRoleId"`
	Arn string `json:"Arn" xml:"Arn"`
}
```

### 35. "上传"

1. route definition

- Url: /oss/upload
- Method: POST
- Request: `UploadReq`
- Response: `UploadResp`

2. request definition



```golang
type UploadReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
	Key string `json:"key" form:"key"` // 文件名(包含路径)
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type UploadResp struct {
	Url string `json:"url"`
	Uri string `json:"uri"`
}
```

### 36. "初始化分段上传"

1. route definition

- Url: /oss/init-multipart-upload
- Method: POST
- Request: `InitMultiUploadReq`
- Response: `InitMultiUploadResp`

2. request definition



```golang
type InitMultiUploadReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
	Key string `json:"key" form:"key"` // 文件名(包含路径)
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type InitMultiUploadResp struct {
	UploadId string `json:"uploadId"` // 分段上传任务全局唯一标识
	Key string `json:"key"` // 文件名(包含路径)
}
```

### 37. "上传某一段"

1. route definition

- Url: /oss/upload-part
- Method: POST
- Request: `UploadPartReq`
- Response: `UploadPartResp`

2. request definition



```golang
type UploadPartReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
	Key string `json:"key" form:"key"` // 文件名(包含路径)
	UploadId string `json:"uploadId" form:"uploadId"` // 分段上传任务全局唯一标识
	PartNumber int32 `json:"partNumber,range=[1:10000]" form:"partNumber,range=[1:10000]"` // 分段序号, 范围是1~10000
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type UploadPartResp struct {
	ETag string `json:"ETag" form:"ETag"` // 段数据的MD5值
	PartNumber int32 `json:"partNumber,range=[1:10000]" form:"partNumber,range=[1:10000]"` //分段序号, 范围是1~10000
	UploadId string `json:"uploadId"` // 分段上传任务全局唯一标识
	Key string `json:"key"` // 文件名(包含路径)
}

type Part struct {
	ETag string `json:"ETag" form:"ETag"` // 段数据的MD5值
	PartNumber int32 `json:"partNumber,range=[1:10000]" form:"partNumber,range=[1:10000]"` //分段序号, 范围是1~10000
}
```

### 38. "完成分段上传"

1. route definition

- Url: /oss/complete-multipart-upload
- Method: POST
- Request: `CompleteMultiUploadReq`
- Response: `CompleteMultiUploadResp`

2. request definition



```golang
type CompleteMultiUploadReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
	Key string `json:"key" form:"key"` // 文件名(包含路径)
	UploadId string `json:"uploadId" form:"uploadId"` // 分段上传任务全局唯一标识
	Parts []Part `json:"parts"`
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type CompleteMultiUploadResp struct {
	Url string `json:"url"`
	Uri string `json:"uri"`
}
```

### 39. "终止分段上传"

1. route definition

- Url: /oss/abort-multipart-upload
- Method: POST
- Request: `AbortMultiUploadReq`
- Response: `AbortMultiUploadResp`

2. request definition



```golang
type AbortMultiUploadReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
	Key string `json:"key" form:"key"` // 文件名(包含路径)
	UploadId string `json:"uploadId" form:"uploadId"` // 分段上传任务全局唯一标识
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type AbortMultiUploadResp struct {
}
```

### 40. "获取主机地址"

1. route definition

- Url: /oss/get-host
- Method: POST
- Request: `GetHostReq`
- Response: `GetHostResp`

2. request definition



```golang
type GetHostReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type GetHostResp struct {
	Host string `json:"host"`
}
```

