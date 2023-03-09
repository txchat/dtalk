package secp256K1

import (
	"fmt"

	"github.com/haltingstate/secp256k1-go"
)

type haltingstate struct {
}

func (t *haltingstate) Sign(msg []byte, privkey []byte) (sig []byte, err error) {
	defer func() {
		if rcov := recover(); rcov != nil {
			if errMsg, ok := rcov.(string); ok {
				err = fmt.Errorf(errMsg)
			}
		}
	}()
	sig = secp256k1.Sign(msg, privkey)
	return
}

func (t *haltingstate) Verify(msg []byte, sig []byte, pubkey []byte) (b bool, err error) {
	defer func() {
		if rcov := recover(); rcov != nil {
			if errMsg, ok := rcov.(string); ok {
				err = fmt.Errorf(errMsg)
			}
		}
	}()
	if secp256k1.VerifySignature(msg, sig, pubkey) == 1 {
		b = true
	}
	return
}
