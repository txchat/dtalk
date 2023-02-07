package dao

import "github.com/txchat/dtalk/app/services/backup/internal/model"

type BackupRepository interface {
	QueryBind(queryType int32, queryCase string) (*model.AddrBackup, error)
	QueryRelate(queryType int32, queryCase string) (*model.AddrRelate, error)
	UpdateAddrBackup(tp int32, r *model.AddrBackup) error
	UpdateAddrRelate(tp int32, r *model.AddrRelate) error
	UpdateMnemonic(r *model.AddrBackup) error
}
