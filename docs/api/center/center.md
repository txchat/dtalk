### 1. "获取默认节点"

1. route definition

- Url: /disc/nodes
- Method: POST
- Request: `GetNodesReq`
- Response: `GetNodesResp`

2. request definition



```golang
type GetNodesReq struct {
}
```


3. response definition



```golang
type GetNodesResp struct {
	Servers []*ChatNode `json:"servers"`
	Nodes []*ContractNode `json:"nodes"`
}
```

### 2. "获取启用模块"

1. route definition

- Url: /app/modules/all
- Method: POST
- Request: `GetModulesReq`
- Response: `GetModulesResp`

2. request definition



```golang
type GetModulesReq struct {
}
```


3. response definition



```golang
type GetModulesResp struct {
	Modules []Module `json:"modules"`
}
```

### 3. "后台用户登录"

1. route definition

- Url: /backend/user/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginResp`

2. request definition



```golang
type LoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginResp struct {
	UserInfo UserInfo `json:"userInfo"`
}

type UserInfo struct {
	UserName string `json:"userName"`
	Token string `json:"token"`
}
```

### 4. "查询手机号是否绑定账号"

1. route definition

- Url: /backup/phone-query
- Method: POST
- Request: `QueryPhoneReq`
- Response: `QueryPhoneResp`

2. request definition



```golang
type QueryPhoneReq struct {
	Area string `json:"area,optional"`
	Phone string `json:"phone"`
}
```


3. response definition



```golang
type QueryPhoneResp struct {
	Exists bool `json:"exists"`
}
```

### 5. "查询邮箱是否绑定账号"

1. route definition

- Url: /backup/email-query
- Method: POST
- Request: `QueryEmailReq`
- Response: `QueryEmailResp`

2. request definition



```golang
type QueryEmailReq struct {
	Email string `json:"email"`
}
```


3. response definition



```golang
type QueryEmailResp struct {
	Exists bool `json:"exists"`
}
```

### 6. "手机号找回备份"

1. route definition

- Url: /backup/phone-retrieve
- Method: POST
- Request: `PhoneRetrieveReq`
- Response: `PhoneRetrieveResp`

2. request definition



```golang
type PhoneRetrieveReq struct {
	Area string `json:"area,optional"`
	Phone string `json:"phone"`
	Code string `json:"code"`
}
```


3. response definition



```golang
type PhoneRetrieveResp struct {
	Address string `json:"address"`
	Area string `json:"area"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Mnemonic string `json:"mnemonic"`
	PrivateKey string `json:"private_key"`
	UpdateTime int64 `json:"update_time"`
	CreateTime int64 `json:"create_time"`
}

type AddressInfo struct {
	Address string `json:"address"`
	Area string `json:"area"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Mnemonic string `json:"mnemonic"`
	PrivateKey string `json:"private_key"`
	UpdateTime int64 `json:"update_time"`
	CreateTime int64 `json:"create_time"`
}
```

### 7. "邮箱找回备份"

1. route definition

- Url: /backup/email-retrieve
- Method: POST
- Request: `EmailRetrieveReq`
- Response: `EmailRetrieveResp`

2. request definition



```golang
type EmailRetrieveReq struct {
	Email string `json:"email"`
	Code string `json:"code"`
}
```


3. response definition



```golang
type EmailRetrieveResp struct {
	Address string `json:"address"`
	Area string `json:"area"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Mnemonic string `json:"mnemonic"`
	PrivateKey string `json:"private_key"`
	UpdateTime int64 `json:"update_time"`
	CreateTime int64 `json:"create_time"`
}

type AddressInfo struct {
	Address string `json:"address"`
	Area string `json:"area"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Mnemonic string `json:"mnemonic"`
	PrivateKey string `json:"private_key"`
	UpdateTime int64 `json:"update_time"`
	CreateTime int64 `json:"create_time"`
}
```

### 8. "手机号导出备份"

1. route definition

- Url: /backup/phone-export
- Method: POST
- Request: `PhoneExportReq`
- Response: `PhoneExportResp`

2. request definition



```golang
type PhoneExportReq struct {
	Area string `json:"area,optional"`
	Phone string `json:"phone"`
	Code string `json:"code"`
	Address string `json:"address"`
}
```


3. response definition



```golang
type PhoneExportResp struct {
}
```

### 9. "邮箱导出备份"

1. route definition

- Url: /backup/email-export
- Method: POST
- Request: `EmailExportReq`
- Response: `EmailExportResp`

2. request definition



```golang
type EmailExportReq struct {
	Email string `json:"email"`
	Code string `json:"code"`
	Address string `json:"address"`
}
```


3. response definition



```golang
type EmailExportResp struct {
}
```

### 10. "手机号绑定账号"

1. route definition

- Url: /backup/phone-binding
- Method: POST
- Request: `PhoneBindingReq`
- Response: `PhoneBindingResp`

2. request definition



```golang
type PhoneBindingReq struct {
	Area string `json:"area,optional"`
	Phone string `json:"phone"`
	Code string `json:"code"`
	Mnemonic string `json:"mnemonic"`
}
```


3. response definition



```golang
type PhoneBindingResp struct {
	Address string `json:"address"`
}
```

### 11. "邮箱绑定账号"

1. route definition

- Url: /backup/email-binding
- Method: POST
- Request: `EmailBindingReq`
- Response: `EmailBindingResp`

2. request definition



```golang
type EmailBindingReq struct {
	Email string `json:"email"`
	Code string `json:"code"`
	Mnemonic string `json:"mnemonic"`
}
```


3. response definition



```golang
type EmailBindingResp struct {
	Address string `json:"address"`
}
```

### 12. "手机号关联账号"

1. route definition

- Url: /backup/phone-relate
- Method: POST
- Request: `PhoneRelateReq`
- Response: `PhoneRelateResp`

2. request definition



```golang
type PhoneRelateReq struct {
	Area string `json:"area,optional"`
	Phone string `json:"phone"`
	Mnemonic string `json:"mnemonic"`
}
```


3. response definition



```golang
type PhoneRelateResp struct {
	Address string `json:"address"`
}
```

### 13. "地址找回备份"

1. route definition

- Url: /backup/address-retrieve
- Method: POST
- Request: `AddressRetrieveReq`
- Response: `AddressRetrieveResp`

2. request definition



```golang
type AddressRetrieveReq struct {
}
```


3. response definition



```golang
type AddressRetrieveResp struct {
	Address string `json:"address"`
	Area string `json:"area"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Mnemonic string `json:"mnemonic"`
	PrivateKey string `json:"private_key"`
	UpdateTime int64 `json:"update_time"`
	CreateTime int64 `json:"create_time"`
}

type AddressInfo struct {
	Address string `json:"address"`
	Area string `json:"area"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Mnemonic string `json:"mnemonic"`
	PrivateKey string `json:"private_key"`
	UpdateTime int64 `json:"update_time"`
	CreateTime int64 `json:"create_time"`
}
```

### 14. "编辑助记词"

1. route definition

- Url: /backup/edit-mnemonic
- Method: POST
- Request: `EditMnemonicReq`
- Response: `EditMnemonicResp`

2. request definition



```golang
type EditMnemonicReq struct {
	Mnemonic string `json:"mnemonic"`
}
```


3. response definition



```golang
type EditMnemonicResp struct {
}
```

### 15. "获取地址"

1. route definition

- Url: /backup/get-address
- Method: POST
- Request: `GetAddressReq`
- Response: `GetAddressResp`

2. request definition



```golang
type GetAddressReq struct {
	Query string `json:"query"`
}
```


3. response definition



```golang
type GetAddressResp struct {
	Address string `json:"address"`
}
```

### 16. "检查版本"

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

### 17. "创建新版本"

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

### 18. "更新版本"

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

### 19. "更改版本状态"

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

### 20. "版本列表"

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

