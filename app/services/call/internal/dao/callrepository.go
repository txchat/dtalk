package dao

import "github.com/txchat/dtalk/app/services/call/internal/model"

type CallRepository interface {
	GetSession(traceId int64) (*model.Session, error)
	SaveSession(session model.Session) error
}
