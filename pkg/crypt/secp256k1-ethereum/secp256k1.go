package secp256K1

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/crypto/secp256k1"
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

	pub, err := crypto.UnmarshalPubkey(pubkey)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if bytes.Equal(crypto.FromECDSAPub(pub), crypto.FromECDSAPub(recoveredPub)) {
		return 1
	}
	fmt.Println(hex.EncodeToString(crypto.FromECDSAPub(pub)))
	fmt.Println(hex.EncodeToString(crypto.FromECDSAPub(recoveredPub)))
	fmt.Println(hex.EncodeToString(crypto.FromECDSAPub(recoveredPub)))
	fmt.Println(crypto.PubkeyToAddress(*recoveredPub).String())
	fmt.Println(crypto.PubkeyToAddress(*recoveredPub).String())

	return 0
}
