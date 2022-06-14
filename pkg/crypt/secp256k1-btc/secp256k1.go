package secp256K1

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec"
)

type bitcoin struct {
}

func (t *bitcoin) Sign(msg []byte, privkey []byte) []byte {
	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), privkey)
	sig, err := btcec.SignCompact(btcec.S256(), priv, msg, true)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return sig
}

func (t *bitcoin) Verify(msg []byte, sig []byte, pubkey []byte) int {
	pub, err := btcec.ParsePubKey(pubkey, btcec.S256())
	if err != nil {
		fmt.Println(err)
		return 0
	}
	recoverdPubKey, _, err := btcec.RecoverCompact(btcec.S256(), sig, msg)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if pub.IsEqual(recoverdPubKey) {
		return 1
	}
	return 0
}
