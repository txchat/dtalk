### 1. "用户登录"

1. route definition

- Url: /app/user/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginResp`

2. request definition



```golang
type LoginReq struct {
	ConnType int32 `json:"connType"`
}
```


3. response definition



```golang
type LoginResp struct {
	Address string `json:"address"`
}
```

