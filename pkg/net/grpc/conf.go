package grpc

import "time"

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
	Dial     time.Duration
	Timeout  time.Duration
}
