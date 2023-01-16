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

