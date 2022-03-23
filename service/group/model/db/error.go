package db

import "errors"

var (
	ErrGroupNotExist       = errors.New("the group is not exist")
	ErrGroupMemberNotExist = errors.New("the member is not exist")
)
