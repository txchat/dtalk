package dao

import (
	"github.com/txchat/dtalk/pkg/oss"
)

type OssRepository interface {
	SaveAssumeRole(appId, ossType string, cfg *oss.Config, data *oss.AssumeRoleResp) error
	GetAssumeRole(appId, ossType string, cfg *oss.Config) (*oss.AssumeRoleResp, error)
}
