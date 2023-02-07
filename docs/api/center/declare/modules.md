### 1. "获取启用模块"

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

