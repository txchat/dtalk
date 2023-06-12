package secp256K1

import (
	"github.com/txchat/dtalk/pkg/crypt"
)

const Name = "secp256k1-bitcoin"

func init() {
	crypt.Register(Name, New())
}

func New() crypt.Encrypt {
	return &bitcoin{}
}
