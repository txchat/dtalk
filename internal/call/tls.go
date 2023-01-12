package call

type TLSSig interface {
	GetAppId() int32
	GetUserSig(userId string) (string, error)
	GenPrivateMapKey(userId string, roomId int32, privilegeMap int32) (string, error)
}
