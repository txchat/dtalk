package group

type GMManager struct {
}

func (gmm *GMManager) Mute(operator, target *Member, muteTime int64) error {
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

func (gmm *GMManager) UnMute(operator, target *Member) error {
	if !operator.AdminOrOwner() {
		return ErrPermissionDenied
	}

	if target.muteTime == UnMute {
		return nil
	}

	target.muteTime = UnMute
	return nil
}
