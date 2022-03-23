package model

type VersionCreateRequest struct {
	Platform    string   `json:"platform"`
	Description []string `json:"description"`
	Force       bool     `json:"force"`
	Url         string   `json:"url"`
	VersionCode int64    `json:"versionCode"`
	VersionName string   `json:"versionName"`
	DeviceType  string   `json:"deviceType"`
	OpeUser     string   `json:"opeUser"`
	Md5         string   `json:"md5"`
	Size        int64    `json:"size"`
}

type VersionUpdateRequest struct {
	Description []string `json:"description"`
	Force       bool     `json:"force"`
	Url         string   `json:"url"`
	VersionCode int64    `json:"versionCode"`
	VersionName string   `json:"versionName"`
	Id          int64    `json:"id"`
	OpeUser     string   `json:"opeUser"`
	Md5         string   `json:"md5"`
	Size        int64    `json:"size"`
}

type VersionChangeStatusRequest struct {
	Id      int64  `json:"id"`
	OpeUser string `json:"opeUser"`
}

type VersionCheckAndUpdateRequest struct {
	VersionCode int64  `form:"versionCode" json:"versionCode"`
	DeviceType  string `json:"deviceType"`
}

type GetVersionListRequest struct {
	Page       int64  `json:"page"`
	Platform   string `json:"platform"`
	DeviceType string `json:"deviceType"`
}

type GetTokenRequest struct {
	UserName string `form:"userName" json:"userName"`
	Password string `form:"password" json:"password"`
}

type VersionCreateResponse struct {
	Version VersionForm `json:"version"`
}

type VersionUpdateResponse struct {
	Version VersionForm `json:"version"`
}

type VersionChangeStatusResponse struct {
	VersionList []VersionForm `json:"versionList"`
}

type VersionCheckAndUpdateResponse struct {
	VersionForm
}

type GetVersionListResponse struct {
	TotalElements int64         `json:"totalElements"`
	TotalPages    int64         `json:"totalPages"`
	VersionList   []VersionForm `json:"versionList"`
}

type GetTokenResponse struct {
	UserInfo UserInfoResponse `json:"userInfo"`
}

type UserInfoResponse struct {
	UserName string `json:"userName"`
	Token    string `json:"token"`
}
