package error

const (
	// CodeOK 请求成功业务状态码
	CodeOK = 0
	// MsgOK 请求成功消息
	MsgOK = "OK"

	// CodeUnexpected 意外错误业务状态码
	CodeUnexpected = -1000
	// MsgUnexpected 意外错误业务消息
	MsgUnexpected = "服务器繁忙，请稍后重试"

	grpcMaxCode = 17
)

var (
	NoErr         = NewError(CodeOK, MsgOK)
	ErrUnexpected = NewError(CodeUnexpected, MsgUnexpected)

	ErrNotFound            = NewError(-1007, "无对应信息")
	ErrOutOfRange          = NewError(-1008, "超出指定范围")
	ErrExec                = NewError(-1009, "执行失败")
	ErrSignatureInvalid    = NewError(-1010, "无效签名")
	ErrSignatureExpired    = NewError(-1011, "签名过期")
	ErrInvalidParams       = NewError(-1012, "请求参数错误")
	ErrInvalidURL          = NewError(-1013, "请求URL非法")
	ErrInvalidHeader       = NewError(-1014, "请求头部错误")
	ErrUnsupportedDevice   = NewError(-1015, "不支持该设备类型")
	ErrReconnectRejected   = NewError(-1016, "设备重连不被允许")
	ErrFeaturesUnSupported = NewError(-2000, "未支持该功能")

	ErrGroupCreateFailed           = NewError(-10000, "群创建失败")
	ErrGroupStatusBlock            = NewError(-10001, "群正在被封禁中")
	ErrGroupStatusDisBand          = NewError(-10002, "群已被解散")
	ErrGroupMemberLimit            = NewError(-10003, "超出群人数上限")
	ErrGroupInviteMemberFailed     = NewError(-10004, "邀请群成员失败")
	ErrGroupInvitePermissionDenied = NewError(-10005, "邀请群成员权限不足")
	ErrGroupMemberTypeOther        = NewError(-10006, "你已不在本群中")
	ErrGroupExit                   = NewError(-10007, "退群失败")
	ErrGroupRemove                 = NewError(-10008, "踢人失败")
	ErrGroupDisband                = NewError(-10009, "解散群失败")
	ErrGroupInviteNoMembers        = NewError(-10010, "被邀请人已经都在本群中")
	ErrGroupRemoveNoMembers        = NewError(-10011, "没有可以踢出群的人")
	ErrGroupAdminDeny              = NewError(-10012, "需要管理员权限")
	ErrGroupOwnerDeny              = NewError(-10013, "需要群主权限")
	ErrGroupChangeOwner            = NewError(-10014, "转让群主失败")
	ErrGroupPersonNotExist         = NewError(-10015, "你已不在本群中")
	ErrGroupMemberNotExist         = NewError(-10016, "对方不在本群中")
	ErrGroupChangeOwnerSelf        = NewError(-10017, "不能把群转让给自己")
	ErrGroupSetAdmin               = NewError(-10018, "设置管理员失败")
	ErrGroupAdminNumLimit          = NewError(-10019, "管理员数量已满")
	ErrGroupHigherPermission       = NewError(-10020, "需要更高权限")
	ErrGroupOwnerExit              = NewError(-10021, "群主不能主动退群")
	ErrGroupInviteMemberExist      = NewError(-10022, "被邀请加入的用户已经是群成员")
	ErrGroupOwnerDisband           = NewError(-10023, "只有群主可以解散群")
	ErrGroupOwnerChange            = NewError(-10024, "只有群主可以转让群")
	ErrGroupOwnerSetAdmin          = NewError(-10025, "只有群主可以设置管理员")
	ErrGroupNotExist               = NewError(-10026, "该群号不存在")
	ErrGroupMutePermission         = NewError(-10027, "群主或管理员不能被禁言")
	ErrGroupApplyUsed              = NewError(-10028, "该申请已处理")
	ErrGroupMemberExist            = NewError(-10029, "你已在本群中")
	ErrGroupApplyNotExist          = NewError(-10030, "该申请不存在")

	ErrCallUserBusy = NewError(-10101, "对方在忙,请稍后再试")

	ErrOssFileTooSmall = NewError(-10201, "上传分片文件太小")
	ErrOssFileTooBig   = NewError(-10202, "上传文件太大")
	ErrOssKeyIllegal   = NewError(-10203, "文件路径非法")

	ErrCdkOutOfStock      = NewError(-10301, "兑换码数量不足")
	ErrCdkOrderError      = NewError(-10302, "订单号错误")
	ErrCdkStatusNotFrozen = NewError(-10303, "兑换码状态错误")
	ErrCdkCoinNameErr     = NewError(-10304, "该票券暂时不支持兑换")
	ErrCdkCoinNameExist   = NewError(-10305, "同名票券已存在")
	ErrCdkMaxNumberErr    = NewError(-10306, "该优惠券已达兑换上限, 请兑换其他优惠券")

	ErrSendMsmCodeError               = NewError(-10400, "验证码发送失败")
	ErrCodeExpired                    = NewError(-10401, "验证码已经过期或者已使用")
	ErrCodeError                      = NewError(-10402, "验证失败")
	ErrExportAddressPhoneInconsistent = NewError(-10400, "账号与绑定手机不一致")
	ErrExportAddressEmailInconsistent = NewError(-10401, "账号与绑定邮箱不一致")
	ErrUserAccountOrPWD               = NewError(-10402, "用户名或密码错误")
)

const (
	CodeInnerError      = -1000
	QueryFailed         = -1001
	ExecFailed          = -1002
	SignatureInvalid    = -1010
	SignatureExpired    = -1011
	ParamsError         = -1012
	RPCFailed           = -1013
	TokenError          = -1014
	DeviceTypeError     = -1015
	ReconnectNotAllowed = -1016

	FeaturesUnSupported = -2000
	VerifyCodeSendError = -4006
	VerifyCodeError     = -4007
	VerifyCodeExpired   = -4008
	QueryNotExist       = -4009
	SendMsgFailed       = -5001

	ExportAddressPhoneInconsistent = -4010
	ExportAddressEmailInconsistent = -4011

	GroupCreateFailed           = -10000
	GroupStatusBlock            = -10001
	GroupStatusDisBand          = -10002
	GroupMemberLimit            = -10003
	GroupInviteMemberFailed     = -10004
	GroupInvitePermissionDenied = -10005
	GroupMemberTypeOther        = -10006
	GroupExit                   = -10007
	GroupRemove                 = -10008
	GroupDisband                = -10009
	GroupInviteNoMembers        = -10010
	GroupRemoveNoMembers        = -10011
	GroupAdminDeny              = -10012
	GroupOwnerDeny              = -10013
	GroupChangeOwner            = -10014
	GroupPersonNotExist         = -10015
	GroupMemberNotExist         = -10016
	GroupChangeOwnerSelf        = -10017
	GroupSetAdmin               = -10018
	GroupAdminNumLimit          = -10019
	GroupNotExist               = -10020
	GroupHigherPermission       = -10021
	GroupOwnerExit              = -10022
	GroupInviteMemberExist      = -10023
	GroupOwnerDisband           = -10024
	GroupOwnerChange            = -10025
	GroupOwnerSetAdmin          = -10026
	GroupMutePermission         = -10027
	GroupApplyUsed              = -10028
	GroupMemberExist            = -10029
	GroupApplyNotExist          = -10030

	CallUserBusy = -10101

	OssFileTooSmall = -10201
	OssFileTooBig   = -10202
	OssKeyIllegal   = -10203

	CdkOutOfStock      = -10301
	CdkOrderError      = -10302
	CdkStatusNotFrozen = -10303
	CdkCoinNameErr     = -10304
	CdkCoinNameExist   = -10305
	CdkMaxNumberErr    = -10306
)

var errorMsg = map[int]string{
	CodeOK:              "操作成功",
	CodeInnerError:      "访问失败",
	QueryFailed:         "查询失败",
	ExecFailed:          "修改失败",
	SignatureInvalid:    "无效签名",
	SignatureExpired:    "签名过期",
	ParamsError:         "请求参数错误",
	RPCFailed:           "调用失败",
	VerifyCodeError:     "验证失败",
	VerifyCodeExpired:   "验证码已经过期或者已使用",
	VerifyCodeSendError: "验证码发送失败",
	QueryNotExist:       "该手机号或邮箱不存在",
	SendMsgFailed:       "消息发送失败",
	FeaturesUnSupported: "未支持该功能",
	TokenError:          "token错误",
	DeviceTypeError:     "获取设备类型失败",
	ReconnectNotAllowed: "设备重连不被允许",

	ExportAddressPhoneInconsistent: "账号与绑定手机不一致",
	ExportAddressEmailInconsistent: "账号与绑定邮箱不一致",

	GroupCreateFailed:           "群创建失败",
	GroupStatusBlock:            "群正在被封禁中",
	GroupStatusDisBand:          "群已被解散",
	GroupMemberLimit:            "超出群人数上限",
	GroupInviteMemberFailed:     "邀请群成员失败",
	GroupInvitePermissionDenied: "邀请群成员权限不足",
	GroupMemberTypeOther:        "你已不在本群中",
	GroupExit:                   "退群失败",
	GroupRemove:                 "踢人失败",
	GroupDisband:                "解散群失败",
	GroupInviteNoMembers:        "被邀请人已经都在本群中",
	GroupRemoveNoMembers:        "没有可以踢出群的人",
	GroupAdminDeny:              "需要管理员权限",
	GroupOwnerDeny:              "需要群主权限",
	GroupChangeOwner:            "转让群主失败",
	GroupPersonNotExist:         "你已不在本群中",
	GroupMemberNotExist:         "对方不在本群中",
	GroupChangeOwnerSelf:        "不能把群转让给自己",
	GroupSetAdmin:               "设置管理员失败",
	GroupAdminNumLimit:          "管理员数量已满",
	GroupHigherPermission:       "需要更高权限",
	GroupOwnerExit:              "群主不能主动退群",
	GroupInviteMemberExist:      "被邀请加入的用户已经是群成员",
	GroupOwnerDisband:           "只有群主可以解散群",
	GroupOwnerChange:            "只有群主可以转让群",
	GroupOwnerSetAdmin:          "只有群主可以设置管理员",
	GroupNotExist:               "该群号不存在",
	GroupMutePermission:         "群主或管理员不能被禁言",
	GroupApplyUsed:              "该申请已处理",
	GroupMemberExist:            "你已在本群中",
	GroupApplyNotExist:          "该申请不存在",

	CallUserBusy: "对方在忙,请稍后再试",

	OssFileTooSmall: "上传分片文件太小",
	OssFileTooBig:   "上传文件太大",
	OssKeyIllegal:   "文件路径非法",

	CdkOutOfStock:      "兑换码数量不足",
	CdkOrderError:      "订单号错误",
	CdkStatusNotFrozen: "兑换码状态错误",
	CdkCoinNameErr:     "该票券暂时不支持兑换",
	CdkCoinNameExist:   "同名票券已存在",
	CdkMaxNumberErr:    "该优惠券已达兑换上限, 请兑换其他优惠券",
}
