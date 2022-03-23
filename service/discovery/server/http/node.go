package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
)

//get all nodes
func Nodes(c *gin.Context) {
	cNodes, err := svc.CNodes()
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.QueryFailed))
		return
	}
	dNodes, err := svc.DNodes()
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.QueryFailed))
		return
	}
	ret := make(map[string]interface{})
	ret["servers"] = cNodes
	ret["nodes"] = dNodes
	c.Set(api.ReqResult, ret)
	c.Set(api.ReqError, nil)
}
