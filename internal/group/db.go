package group

type DBExec interface {
	SetGroupMemberMuteInfo(m *Member) error
	ChangeGroupOwner(groupId int64, ownerId, memberId string) error
}
