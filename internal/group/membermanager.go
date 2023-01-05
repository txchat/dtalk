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
