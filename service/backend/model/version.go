package model

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/txchat/dtalk/pkg/util"
)

type Description []string

func (desc *Description) ToString() string {
	b, err := json.Marshal(desc)
	if err != nil {
		return ""
	}
	return string(b)
}

func ConvertDescription(str string) (Description, error) {
	var desc Description
	err := json.Unmarshal([]byte(str), &desc)
	if err != nil {
		return nil, err
	}
	return desc, nil
}

type VersionForm struct {
	Id          int64       `json:"id"`
	Platform    string      `json:"platform"`
	Status      int32       `json:"status"`
	DeviceType  string      `json:"deviceType"`
	VersionName string      `json:"versionName"`
	VersionCode int64       `json:"versionCode"`
	Url         string      `json:"url"`
	Force       bool        `json:"force"`
	Description Description `json:"description"`
	OpeUser     string      `json:"opeUser"`
	Md5         string      `json:"md5"`
	Size        int64       `json:"size"`
	UpdateTime  int64       `json:"updateTime"`
	CreateTime  int64       `json:"createTime"`
}

func ConvertVersionForm(record *map[string]string) (*VersionForm, error) {
	description, err := ConvertDescription((*record)["description"])
	if err != nil {
		return nil, err
	}
	return &VersionForm{
		Id:          util.ToInt64((*record)["id"]),
		Platform:    (*record)["platform"],
		Status:      util.ToInt32((*record)["state"]),
		DeviceType:  (*record)["device_type"],
		VersionName: (*record)["version_name"],
		VersionCode: util.ToInt64((*record)["version_code"]),
		Url:         (*record)["download_url"],
		Force:       util.ToBool((*record)["force_update"]),
		Description: description,
		OpeUser:     (*record)["ope_user"],
		Md5:         (*record)["md5"],
		Size:        util.ToInt64((*record)["size"]),
		UpdateTime:  util.ToInt64((*record)["update_time"]),
		CreateTime:  util.ToInt64((*record)["create_time"]),
	}, nil
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserInfo struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
