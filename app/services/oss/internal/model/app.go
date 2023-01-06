package model

import (
	"github.com/txchat/dtalk/pkg/oss"
)

const (
	OssAliyun    = "aliyun"
	OssHuaweiuyn = "huaweiyun"
	OssMinio     = "minio"
)

type AppOssEngine struct {
	ossEngines map[string]*Engines
}

func NewAppOssManager() *AppOssEngine {
	return &AppOssEngine{
		ossEngines: make(map[string]*Engines),
	}
}

func (m *AppOssEngine) Init(appId string, ossType string, engine oss.Oss) {
	a, ok := m.ossEngines[appId]
	if !ok {
		a = NewEngines(ossType)
		m.ossEngines[appId] = a
	}
	a.Register(ossType, engine)
}

func (m *AppOssEngine) GetEngine(appId string, ossType string) (oss.Oss, error) {
	a, ok := m.ossEngines[appId]
	if !ok {
		return nil, ErrEngineUnsupported
	}
	if ossType == "" {
		return a.GetDefaultEngine(), nil
	}
	return a.GetEngine(ossType), nil
}

type Engines struct {
	defaultType string
	engines     map[string]oss.Oss
}

func NewEngines(defaultType string) *Engines {
	return &Engines{
		defaultType: defaultType,
		engines:     make(map[string]oss.Oss),
	}
}

func (e *Engines) Register(ossType string, engine oss.Oss) {
	e.engines[ossType] = engine
}

func (e *Engines) GetEngine(ossType string) oss.Oss {
	return e.engines[ossType]
}

func (e *Engines) GetDefaultEngine() oss.Oss {
	return e.engines[e.defaultType]
}
