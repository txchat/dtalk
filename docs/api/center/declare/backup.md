### 1. "查询手机号是否绑定账号"

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

### 2. "查询邮箱是否绑定账号"

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

### 3. "手机号找回备份"

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

### 4. "邮箱找回备份"

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

### 5. "手机号导出备份"

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

### 6. "邮箱导出备份"

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

### 7. "手机号绑定账号"

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

### 8. "邮箱绑定账号"

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

### 9. "手机号关联账号"

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

### 10. "地址找回备份"

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

### 11. "编辑助记词"

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

### 12. "获取地址"

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

