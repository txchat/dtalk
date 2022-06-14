package secp256K1

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

type ethereum struct {
}

func (t *ethereum) Sign(msg []byte, privkey []byte) []byte {
	priv, err := crypto.ToECDSA(privkey)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	sig, err := crypto.Sign(msg, priv)
	//ret , err := secp256k1.Sign(msg, privkey)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return sig
}

func (t *ethereum) Verify(msg []byte, sig []byte, pubkey []byte) int {
	recoveredPub, err := crypto.SigToPub(msg, sig)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	ecdsaPubKey, err := crypto.DecompressPubkey(pubkey)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	if bytes.Equal(crypto.FromECDSAPub(ecdsaPubKey), crypto.FromECDSAPub(recoveredPub)) {
		return 1
	}

	return 0
}
