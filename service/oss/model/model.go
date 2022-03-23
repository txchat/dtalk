package model

import (
	"github.com/txchat/dtalk/pkg/oss"
)

// App 代表一个应用, 一个 app 中存在多种 oss 存储方式
type App struct {
	DefaultOssType string
	ossInv         map[string]oss.Oss
	AppId          string
}

func NewApp(appId string) *App {
	return &App{
		ossInv: make(map[string]oss.Oss),
		AppId:  appId,
	}
}

func (a *App) Register(ossType string, eng oss.Oss) {
	a.ossInv[ossType] = eng
}

func (a *App) GetInvoker(ossType string) oss.Oss {
	return a.ossInv[ossType]
}
