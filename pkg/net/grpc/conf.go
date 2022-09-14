package grpc

import xtime "github.com/txchat/dtalk/pkg/time"

type Discovery struct {
	Address  string
	Name     string
	Password string
}

// RPCClient is RPC client config.
type RPCClient struct {
	RegAddrs string
	Schema   string
	SrvName  string
	Dial     xtime.Duration
	Timeout  xtime.Duration
}
