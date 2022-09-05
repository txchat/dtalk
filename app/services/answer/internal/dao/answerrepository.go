package dao

import "github.com/txchat/dtalk/app/services/answer/internal/model"

type AnswerRepository interface {
	AddRecordSeqIndex(uid string, m *model.MsgIndex) error
	GetRecordSeqIndex(uid, seq string) (*model.MsgIndex, error)
}
