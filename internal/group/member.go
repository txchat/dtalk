package group

type RoleType int

const (
	Normal  RoleType = 0
	Manager RoleType = 1
	Owner   RoleType = 2
	Out     RoleType = 10
)

const (
	UnMute      = 0
	MuteForever = 9223372036854775807
)

type Members []*Member

func (ms Members) ToArray() []string {
	members := []*Member(ms)
	mid := make([]string, len(members))
	for i, member := range members {
		mid[i] = member.Id()
	}
	return mid
}

type Member struct {
	id       string
	nickname string
	group    *Group
	role     RoleType
	muteTime int64
}

func (m *Member) SetRole(role RoleType) {
	m.role = role
}

func (m *Member) Id() string {
	return m.id
}

func (m *Member) Nickname() string {
	return m.nickname
}

func (m *Member) Group() *Group {
	return m.group
}

func (m *Member) Role() RoleType {
	return m.role
}

func (m *Member) MuteTime() int64 {
	return m.muteTime
}

func (m *Member) AdminOrOwner() bool {
	return m.role == Manager || m.role == Owner
}
