package auth

import (
	"time"

	"github.com/txchat/im/pkg/auth"
)

const Name = "dtalk"

func init() {
	auth.Register(Name, NewAuth)
	auth.RegisterErrorDecoder(Name, DecodingErrorReject)
}

func NewAuth(url string, timeout time.Duration) auth.Auth {
	return &talkClient{url: url, timeout: timeout}
}
