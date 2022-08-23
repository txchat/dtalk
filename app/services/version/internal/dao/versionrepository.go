package dao

import (
	"context"

	"github.com/txchat/dtalk/app/services/version/internal/model"
)

type DeviceRepository interface {
	GetVersionInfo(ctx context.Context, vid int64) (*model.VersionForm, error)
	AddVersionInfo(ctx context.Context, version *model.VersionForm) (int64, int64, error)
	UpdateVersionInfo(ctx context.Context, version *model.VersionForm) (int64, int64, error)
	ReleaseSpecificVersion(ctx context.Context, vid, updateTime int64, operator string) error
	SpecificPlatformAndDeviceTypeVersions(ctx context.Context, platform, deviceType string, page int64, size int64) ([]*model.VersionForm, error)
	SpecificPlatformAndDeviceTypeCount(ctx context.Context, platform, deviceType string) (int64, error)
	LastReleaseVersion(ctx context.Context, platform, deviceType string) (*model.VersionForm, error)
	ForceNumberBetween(ctx context.Context, platform, deviceType string, begin, end int64) (int64, error)
}
