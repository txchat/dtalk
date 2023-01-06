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
	ErrPermissionDenied    = NewError(-1019, "权限不足")
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
	ErrCallTimeout  = NewError(-10102, "通话已过期")

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

	ErrSendMsgFailed = NewError(-10500, "发送失败")
)
