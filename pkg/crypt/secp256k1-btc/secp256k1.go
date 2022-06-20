package secp256K1

import (
	"github.com/btcsuite/btcd/btcec"
)

const SignatureLength = 64 + 1 // 64 bytes ECDSA signature + 1 byte recovery id

type bitcoin struct {
}

func (t *bitcoin) ethRecoverCompact(curve *btcec.KoblitzCurve, signature,
	hash []byte) (*btcec.PublicKey, bool, error) {
	// Convert to btcec input format with 'recovery id' v at the beginning.
	btcsig := make([]byte, SignatureLength)
	btcsig[0] = signature[64] + 27
	copy(btcsig[1:], signature)

	return btcec.RecoverCompact(btcec.S256(), btcsig, hash)
}

func (t *bitcoin) Sign(msg []byte, privkey []byte) ([]byte, error) {
	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), privkey)
	sig, err := btcec.SignCompact(btcec.S256(), priv, msg, false)
	if err != nil {
		return nil, err
	}
	// Convert to Ethereum signature format with 'recovery id' v at the end.
	v := sig[0] - 27
	copy(sig, sig[1:])
	sig[64] = v
	return sig, nil
}

func (t *bitcoin) Verify(msg []byte, signature []byte, pubkey []byte) (bool, error) {
	pub, err := btcec.ParsePubKey(pubkey, btcec.S256())
	if err != nil {
		return false, err
	}
	//recoverdPubKey, _, err := btcec.RecoverCompact(btcec.S256(), signature, msg)
	recoverdPubKey, _, err := t.ethRecoverCompact(btcec.S256(), signature, msg)
	if err != nil {
		return false, err
	}
	return pub.IsEqual(recoverdPubKey), nil
}
