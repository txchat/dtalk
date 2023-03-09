package secp256K1

import (
	"bytes"

	"github.com/ethereum/go-ethereum/crypto"
)

type ethereum struct {
}

func (t *ethereum) Sign(msg []byte, privkey []byte) ([]byte, error) {
	priv, err := crypto.ToECDSA(privkey)
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(msg, priv)
	//ret , err := secp256k1.Sign(msg, privkey)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (t *ethereum) Verify(msg []byte, sig []byte, pubkey []byte) (bool, error) {
	recoveredPub, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false, err
	}

	ecdsaPubKey, err := crypto.DecompressPubkey(pubkey)
	if err != nil {
		return false, err
	}

	return bytes.Equal(crypto.FromECDSAPub(ecdsaPubKey), crypto.FromECDSAPub(recoveredPub)), nil
}
