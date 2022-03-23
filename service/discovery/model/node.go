package model

//聊天服务器节点
type CNode struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

//合约节点
type DNode struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
