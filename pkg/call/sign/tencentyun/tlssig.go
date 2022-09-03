package tencentyun

import (
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
	"github.com/txchat/dtalk/pkg/call/sign"
)

// TCTLSSig 腾讯音视频签名实例
type TCTLSSig struct {
	sdkAppId  int
	secretKey string
	expire    int
}

func NewTCTLSSig(sdkAppId int, secretKey string, expire int) sign.TLSSig {
	return &TCTLSSig{
		sdkAppId:  sdkAppId,
		secretKey: secretKey,
		expire:    expire,
	}
}

func (t *TCTLSSig) GetAppId() int32 {
	return int32(t.sdkAppId)
}

//GetUserSig
//【功能说明】用于签发 TRTC 和 IM 服务中必须要使用的 UserSig 鉴权票据
//
//参数说明】
//sdkappid - 应用id
//key - 计算 usersig 用的加密密钥,控制台可获取
//userid - 用户id，限制长度为32字节，只允许包含大小写英文字母（a-zA-Z）、数字（0-9）及下划线和连词符。
//expire - UserSig 票据的过期时间，单位是秒，比如 86400 代表生成的 UserSig 票据在一天后就无法再使用了。
func (t *TCTLSSig) GetUserSig(userId string) (string, error) {
	return tencentyun.GenUserSig(t.sdkAppId, t.secretKey, userId, t.expire)
}

// GenPrivateMapKey
//【功能说明】
// 用于签发 TRTC 进房参数中可选的 PrivateMapKey 权限票据。
// PrivateMapKey 需要跟 UserSig 一起使用，但 PrivateMapKey 比 UserSig 有更强的权限控制能力：
//  - UserSig 只能控制某个 UserID 有无使用 TRTC 服务的权限，只要 UserSig 正确，其对应的 UserID 可以进出任意房间。
//  - PrivateMapKey 则是将 UserID 的权限控制的更加严格，包括能不能进入某个房间，能不能在该房间里上行音视频等等。
// 如果要开启 PrivateMapKey 严格权限位校验，需要在【实时音视频控制台】=>【应用管理】=>【应用信息】中打开“启动权限密钥”开关。
//
//【参数说明】
// sdkappid - 应用id。
// key - 计算 usersig 用的加密密钥,控制台可获取。
// userid - 用户id，限制长度为32字节，只允许包含大小写英文字母（a-zA-Z）、数字（0-9）及下划线和连词符。
// expire - PrivateMapKey 票据的过期时间，单位是秒，比如 86400 生成的 PrivateMapKey 票据在一天后就无法再使用了。
// roomid - 房间号，用于指定该 userid 可以进入的房间号
// privilegeMap - 权限位，使用了一个字节中的 8 个比特位，分别代表八个具体的功能权限开关：
//  - 第 1 位：0000 0001 = 1，创建房间的权限
//  - 第 2 位：0000 0010 = 2，加入房间的权限
//  - 第 3 位：0000 0100 = 4，发送语音的权限
//  - 第 4 位：0000 1000 = 8，接收语音的权限
//  - 第 5 位：0001 0000 = 16，发送视频的权限
//  - 第 6 位：0010 0000 = 32，接收视频的权限
//  - 第 7 位：0100 0000 = 64，发送辅路（也就是屏幕分享）视频的权限
//  - 第 8 位：1000 0000 = 200，接收辅路（也就是屏幕分享）视频的权限
//  - privilegeMap == 1111 1111 == 255 代表该 userid 在该 roomid 房间内的所有功能权限。
//  - privilegeMap == 0010 1010 == 42  代表该 userid 拥有加入房间和接收音视频数据的权限，但不具备其他权限。
func (t *TCTLSSig) GenPrivateMapKey(userId string, roomId int32, privilegeMap int32) (string, error) {
	return tencentyun.GenPrivateMapKey(t.sdkAppId, t.secretKey, userId, t.expire, uint32(roomId), uint32(privilegeMap))
}
