package secp256K1

import (
	"encoding/hex"
	"testing"
)

func Test_ChatSign(t *testing.T) {
	var bty haltingstate
	msg := []byte("1624503307409*aaaaaaaaaaaaaaaaaa")
	privateKey, err := hex.DecodeString("5befd5ad6b2195fd44a5353e009b0ad3ca1de32c412f45f9524649ff1eb1f21c")
	if err != nil {
		t.Error(err)
		return
	}
	sig := bty.Sign(msg, privateKey)
	t.Log(hex.EncodeToString(sig))
}

func Test_Verify(t *testing.T) {
	var bty haltingstate
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
	ret := bty.Verify(msg, sig, publicKey)
	t.Log(ret)
}

func Test_Verify2(t *testing.T) {
	var bty haltingstate
	msg := []byte("1624521238050*aaaaaaaaaaaaaaaaaa")
	publicKey, err := hex.DecodeString("0341655c5af9badc5199f8bcd4efffc5fdef6f571e36a8e314a6ea2656af82d2ed")
	if err != nil {
		t.Error(err)
		return
	}
	sig, err := hex.DecodeString("30af6bf3a149a1e7f82651c71a02c470857cdf0d5821be8cc478584e2855982d5a91487067135d7077d97d621777957cd6567c51ba890c410142897de040d3e501")
	if err != nil {
		t.Error(err)
		return
	}
	ret := bty.Verify(msg, sig, publicKey)
	t.Log(ret)
}
