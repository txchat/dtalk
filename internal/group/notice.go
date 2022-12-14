package group

type NoticeHub interface {
	UpdateMembersMuteTime(groupId int64, operator string, members []string) error
	UpdateGroupOwner(groupId int64, operator, newOwner string) error
}