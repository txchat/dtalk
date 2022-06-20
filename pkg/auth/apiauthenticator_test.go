package auth

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userAddress = "1AqutxNoVTtcWiVYpBtvficAgea1dYTddR"

func TestDefaultApuAuthenticator(t *testing.T) {
	pubKey, err := hex.DecodeString(publicKey)
	assert.Nil(t, err)

	privKey, err := hex.DecodeString(privateKey)
	assert.Nil(t, err)

	client := NewDefaultApuAuthenticator()
	sig := client.Request("dtalk", pubKey, privKey)

	server := NewDefaultApuAuthenticator()
	uid, err := server.Auth(sig)
	assert.Nil(t, err)
	assert.Equal(t, userAddress, uid)
}

func TestDefaultApuAuthenticatorRequest(t *testing.T) {
	pubKey, err := hex.DecodeString("02b2dcf40123a5364a4bc9fd717db92122f90321a6771a47bc922100c9852c8b68")
	if err != nil {
		t.Errorf("decode pubKey:%v", err)
		return
	}
	privKey, err := hex.DecodeString("bfae31775aeefb2eb01f604e2a4cf6d6c4cb4c072ddfbde03235252bd2765e06")
	if err != nil {
		t.Errorf("decode privKey:%v", err)
		return
	}

	client := NewDefaultApuAuthenticator()
	sig := client.Request("dtalk", pubKey, privKey)
	t.Logf("sig:%s", sig)
}

func TestDefaultApiAuthenticatorAuthInvalid(t *testing.T) {
	sig := "UXmp9mdD/iQlkAXqAHHzWaFwTNOxv6OxyqqSoJo8RnD6kJmh1dCF3f8GKL70/LCZKucvwAgUWbj9rTVXeq3gA=#1648784849078376*dtalk#02b2dcf40123a5364a4bc9fd717db92122f90321a6771a47bc922100c9852c8b68"
	server := NewDefaultApuAuthenticator()
	uid, err := server.Auth(sig)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("uid:%s", uid)
	t.Log("success")
}

func TestDefaultApuAuthenticatorAuthExpire(t *testing.T) {
	sig := "G0UXmp9mdD/iQlkAXqAHHzWaFwTNOxv6OxyqqSoJo8RnD6kJmh1dCF3f8GKL70/LCZKucvwAgUWbj9rTVXeq3gA=#1648784849078376*dtalk#02b2dcf40123a5364a4bc9fd717db92122f90321a6771a47bc922100c9852c8b68"
	server := NewDefaultApuAuthenticator()
	uid, err := server.Auth(sig)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("uid:%s", uid)
	t.Log("success")
}
