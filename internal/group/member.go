package group

type RoleType int

const (
	Normal RoleType = 0
	Admin  RoleType = 1
	Owner  RoleType = 2
	Out    RoleType = 10
)

const (
	UnMute = 0
	MuteForever = 9223372036854775807
)

type Member struct {
	id string
	group *Group
	role RoleType
	muteTime int64
}

func (m *Member) AdminOrOwner() bool{
	return m.role == Admin || m.role == Owner
}



type GMManager struct {
	dbExec DBExec
	signalHub SignalHub
	noticeHub 	NoticeHub
}

func (gmm *GMManager) Mute(operator, target *Member, muteTime int64) error{
	if !operator.AdminOrOwner() {
		return ErrPermissionDenied
	}
	// can not mute admin
	if target.AdminOrOwner() {
		return ErrPermissionDenied
	}

	target.muteTime = muteTime
	err := gmm.dbExec.SetGroupMemberMuteInfo(target)
	if err != nil {
		return err
	}

	// send signal
	err = gmm.signalHub.UpdateMembersMuteTime(target.group.GetID(), []string{target.id}, target.muteTime)
	if err != nil {
		return err
	}
	// send notify
	err = gmm.noticeHub.UpdateMembersMuteTime(target.group.GetID(), operator.id, []string{target.id})
	if err != nil {
		return err
	}
	return nil
}

func (gmm *GMManager) UnMute(operator, target *Member) error{
	if !operator.AdminOrOwner() {
		return ErrPermissionDenied
	}

	if target.muteTime == UnMute {
		return nil
	}

	target.muteTime = UnMute
	err := gmm.dbExec.SetGroupMemberMuteInfo(target)
	if err != nil {
		return err
	}

	// send signal
	err = gmm.signalHub.UpdateMembersMuteTime(target.group.GetID(), []string{target.id}, target.muteTime)
	if err != nil {
		return err
	}
	// send notify
	err = gmm.noticeHub.UpdateMembersMuteTime(target.group.GetID(), operator.id, []string{target.id})
	if err != nil {
		return err
	}
	return nil
}