package group

type Group struct {
	id int64
	owner string
}

func (g *Group) GetID() int64{
	return g.id
}

func (g *Group) GetOwner() string{
	return g.owner
}


type Manager struct {
	gmm *GMManager
	dbExec DBExec
	signalHub SignalHub
	noticeHub 	NoticeHub
}

func (mg *Manager) ChangeOwner(operator, newOwner *Member, group *Group) error{
	// check operator permission
	if operator.id != group.GetOwner() {
		return ErrPermissionDenied
	}
	// exec change info
	err := mg.dbExec.ChangeGroupOwner(group.id, operator.id, newOwner.id)
	if err != nil {
		return err
	}

	// unset mute
	err = mg.gmm.UnMute(operator, newOwner)
	if err != nil {
		return err
	}

	// send signal
	err = mg.signalHub.ChangeMemberRoleType(group.id, newOwner.id, Owner)
	if err != nil {
		return err
	}
	err = mg.signalHub.ChangeMemberRoleType(group.id, operator.id, Normal)
	if err != nil {
		return err
	}
	// send notify
	err = mg.noticeHub.UpdateGroupOwner(group.id, operator.id, newOwner.id)
	if err != nil {
		return err
	}
	return nil
}