### 1. "后台用户登录"

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

