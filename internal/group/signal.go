package group

type SignalHub interface {
	UpdateMembersMuteTime(groupId int64, members []string, muteTime int64) error
	ChangeMemberRoleType(groupId int64, memberId string, roleType RoleType) error
}