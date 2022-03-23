package model

import (
	"github.com/pkg/errors"
)

var (
	ErrSessionNotExist     = errors.New("the session is not exist")
	ErrUserBusy            = errors.New("user is busy")
	ErrFeaturesUnSupported = errors.New("Features UnSupported")
)
