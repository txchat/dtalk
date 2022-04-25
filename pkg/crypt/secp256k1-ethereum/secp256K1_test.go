package secp256K1

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func Test_ChatSign(t *testing.T) {
	var eth ethereum
	msg := []byte("1624503307409*aaaaaaaaaaaaaaaaaa")
	privateKey, err := hex.DecodeString("5befd5ad6b2195fd44a5353e009b0ad3ca1de32c412f45f9524649ff1eb1f21c")
	if err != nil {
		t.Error(err)
		return
	}
	sig := eth.Sign(msg, privateKey)
	t.Log(hex.EncodeToString(sig))
}

func Test_ChatSign2(t *testing.T) {
	var eth ethereum
	msg := []byte("1624503307409*aaaaaaaaaaaaaaaaaa")
	sa := sha256.New()
	sa.Write(msg)
	msg32 := sa.Sum(nil)
	privateKey, err := hex.DecodeString("5befd5ad6b2195fd44a5353e009b0ad3ca1de32c412f45f9524649ff1eb1f21c")
	if err != nil {
		t.Error(err)
		return
	}
	sig := eth.Sign(msg32, privateKey)
	t.Log(hex.EncodeToString(sig))
}

func Test_Verify(t *testing.T) {
	var eth ethereum
	msg := []byte("1624503307409*aaaaaaaaaaaaaaaaaa")
	publicKey, err := hex.DecodeString("03c94c689f4ae2002bc05215684617146b84e81571ab854de4b171066087fc9df6")
	if err != nil {
		t.Error(err)
		return
	}
	sig, err := hex.DecodeString("93181cf41ebc1b65ff3fe464a2e088711fb376c0ac27802311b183b2995fe3d57bb9908cf4dd3de7ce3fffbdeee9ce0e0526daa4a4b245432380cb6c90fdafbe01")
	if err != nil {
		t.Error(err)
		return
	}
	ret := eth.Verify(msg, sig, publicKey)
	t.Log(ret)
}
