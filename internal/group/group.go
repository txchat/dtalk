package group

type TypeOfGroup int
type JoinGroupPermission int
type MuteTypeOfGroup int
type FriendshipOfGroupPermission int

const (
	NormalGroup     TypeOfGroup = 0
	EnterpriseGroup TypeOfGroup = 1
	DepartmentGroup TypeOfGroup = 2
)

const (
	AnybodyCanJoinGroup   JoinGroupPermission = 0 // 无需审批（默认）
	JustManagerCanInvite  JoinGroupPermission = 1 // 禁止加群，群主和管理员邀请加群
	NormalMemberCanInvite JoinGroupPermission = 2 // 普通人邀请需要审批,群主和管理员直接加群
)

const (
	NotLimited            MuteTypeOfGroup = 0 // 全员可发言
	AllMutedExceptManager MuteTypeOfGroup = 1 // 全员禁言(除群主和管理员)
)

const (
	AllowedGroupFriendship FriendshipOfGroupPermission = 0 // 群内可加好友
	DeniedGroupFriendship  FriendshipOfGroupPermission = 1 // 群内禁止加好友
)

type Group struct {
	id         int64
	name       string
	avatar     string
	markId     string
	owner      string
	maxMembers int
	createTime int64
	//permission
	joinPermission       JoinGroupPermission
	mutePermission       MuteTypeOfGroup
	friendshipPermission FriendshipOfGroupPermission
	aesKey               string

	members []*Member
}

func (g *Group) Members() []*Member {
	return g.members
}

func (g *Group) JoinPermission() JoinGroupPermission {
	return g.joinPermission
}

func (g *Group) MutePermission() MuteTypeOfGroup {
	return g.mutePermission
}

func (g *Group) FriendshipPermission() FriendshipOfGroupPermission {
	return g.friendshipPermission
}

func (g *Group) Id() int64 {
	return g.id
}

func (g *Group) Name() string {
	return g.name
}

func (g *Group) Avatar() string {
	return g.avatar
}

func (g *Group) MarkId() string {
	return g.markId
}

func (g *Group) Owner() string {
	return g.owner
}

func (g *Group) MaxMembers() int {
	return g.maxMembers
}

func (g *Group) CreateTime() int64 {
	return g.createTime
}

func (g *Group) AesKey() string {
	return g.aesKey
}

func NewGroup() *Group {
	return &Group{
		id:      0,
		markId:  "",
		owner:   "",
		members: make([]*Member, 0),
	}
}

func (g *Group) MemberCount() int {
	return len(g.members)
}

func (g *Group) Invite(id, nickname string) {
	role := Normal
	if id == g.owner {
		role = Owner
	}
	g.members = append(g.members, &Member{
		id:       id,
		nickname: nickname,
		group:    g,
		role:     role,
		muteTime: 0,
	})
	return
}

func (g *Group) ChangeOwner(operator, newOwner *Member, mmg *GMManager) error {
	// check operator permission
	if operator.id != g.owner {
		return ErrPermissionDenied
	}

	// unset mute
	err := mmg.UnMute(operator, newOwner)
	if err != nil {
		return err
	}

	operator.SetRole(Normal)
	newOwner.SetRole(Owner)

	g.owner = newOwner.Id()
	return nil
}
