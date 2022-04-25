package secp256K1

import (
	"github.com/haltingstate/secp256k1-go"
)

type haltingstate struct {
	isCkTimeOut bool
}

func (t *haltingstate) Sign(msg []byte, privkey []byte) []byte {
	return secp256k1.Sign(msg, privkey)
}

func (t *haltingstate) Verify(msg []byte, sig []byte, pubkey []byte) int {
	return secp256k1.VerifySignature(msg, sig, pubkey)
}
