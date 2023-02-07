package group

import "errors"

var (
	ErrPermissionDenied     = errors.New("permission denied")
	ErrGroupMaxMembersLimit = errors.New("group max members limit")
)
