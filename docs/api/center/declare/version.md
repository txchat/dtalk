### 1. "检查版本"

1. route definition

- Url: /app/version/check
- Method: POST
- Request: `VersionCheckReq`
- Response: `VersionCheckResp`

2. request definition



```golang
type VersionCheckReq struct {
	VersionCode int64 `json:"versionCode"`
	DeviceType string `json:"deviceType,optional"`
}
```


3. response definition



```golang
type VersionCheckResp struct {
	Id int64 `json:"id"`
	Platform string `json:"platform"`
	Status int32 `json:"status"`
	DeviceType string `json:"deviceType"`
	VersionName string `json:"versionName"`
	VersionCode int64 `json:"versionCode"`
	Url string `json:"url"`
	Force bool `json:"force"`
	Description []string `json:"description"`
	OpeUser string `json:"opeUser"`
	Md5 string `json:"md5"`
	Size int64 `json:"size"`
	UpdateTime int64 `json:"updateTime"`
	CreateTime int64 `json:"createTime"`
}

type VersionInfo struct {
	Id int64 `json:"id"`
	Platform string `json:"platform"`
	Status int32 `json:"status"`
	DeviceType string `json:"deviceType"`
	VersionName string `json:"versionName"`
	VersionCode int64 `json:"versionCode"`
	Url string `json:"url"`
	Force bool `json:"force"`
	Description []string `json:"description"`
	OpeUser string `json:"opeUser"`
	Md5 string `json:"md5"`
	Size int64 `json:"size"`
	UpdateTime int64 `json:"updateTime"`
	CreateTime int64 `json:"createTime"`
}
```

### 2. "创建新版本"

1. route definition

- Url: /backend/version/create
- Method: POST
- Request: `CreateVersionReq`
- Response: `CreateVersionResp`

2. request definition



```golang
type CreateVersionReq struct {
	Platform string `json:"platform"`
	Description []string `json:"description"`
	Force bool `json:"force"`
	Url string `json:"url"`
	VersionCode int64 `json:"versionCode"`
	VersionName string `json:"versionName"`
	DeviceType string `json:"deviceType"`
	OpeUser string `json:"opeUser"`
	Md5 string `json:"md5"`
	Size int64 `json:"size"`
}
```


3. response definition



```golang
type CreateVersionResp struct {
	Version VersionInfo `json:"version"`
}

type VersionInfo struct {
	Id int64 `json:"id"`
	Platform string `json:"platform"`
	Status int32 `json:"status"`
	DeviceType string `json:"deviceType"`
	VersionName string `json:"versionName"`
	VersionCode int64 `json:"versionCode"`
	Url string `json:"url"`
	Force bool `json:"force"`
	Description []string `json:"description"`
	OpeUser string `json:"opeUser"`
	Md5 string `json:"md5"`
	Size int64 `json:"size"`
	UpdateTime int64 `json:"updateTime"`
	CreateTime int64 `json:"createTime"`
}
```

### 3. "更新版本"

1. route definition

- Url: /backend/version/update
- Method: PUT
- Request: `UpdateVersionReq`
- Response: `UpdateVersionResp`

2. request definition



```golang
type UpdateVersionReq struct {
	Description []string `json:"description"`
	Force bool `json:"force"`
	Url string `json:"url"`
	VersionCode int64 `json:"versionCode"`
	VersionName string `json:"versionName"`
	Id int64 `json:"id"`
	OpeUser string `json:"opeUser"`
	Md5 string `json:"md5"`
	Size int64 `json:"size"`
}
```


3. response definition



```golang
type UpdateVersionResp struct {
	Version VersionInfo `json:"version"`
}

type VersionInfo struct {
	Id int64 `json:"id"`
	Platform string `json:"platform"`
	Status int32 `json:"status"`
	DeviceType string `json:"deviceType"`
	VersionName string `json:"versionName"`
	VersionCode int64 `json:"versionCode"`
	Url string `json:"url"`
	Force bool `json:"force"`
	Description []string `json:"description"`
	OpeUser string `json:"opeUser"`
	Md5 string `json:"md5"`
	Size int64 `json:"size"`
	UpdateTime int64 `json:"updateTime"`
	CreateTime int64 `json:"createTime"`
}
```

### 4. "更改版本状态"

1. route definition

- Url: /backend/version/change-status
- Method: PUT
- Request: `ChangeVersionStateReq`
- Response: `ChangeVersionStateResp`

2. request definition



```golang
type ChangeVersionStateReq struct {
	Id int64 `json:"id"`
	OpeUser string `json:"opeUser,optional"`
}
```


3. response definition



```golang
type ChangeVersionStateResp struct {
}
```

### 5. "版本列表"

1. route definition

- Url: /backend/version/list
- Method: GET
- Request: `ListVersionReq`
- Response: `ListVersionResp`

2. request definition



```golang
type ListVersionReq struct {
	Page int64 `form:"page,default=0"`
	Platform string `form:"platform,default=%"`
	DeviceType string `form:"deviceType,default=%"`
}
```


3. response definition



```golang
type ListVersionResp struct {
	TotalElements int64 `json:"totalElements"`
	TotalPages int64 `json:"totalPages"`
	VersionList []VersionInfo `json:"versionList"`
}
```

