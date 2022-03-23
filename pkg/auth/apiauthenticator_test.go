package auth

import (
	"encoding/hex"
	"testing"
)

func TestDefaultApuAuthenticator(t *testing.T) {
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
	server := NewDefaultApuAuthenticator()
	if !server.Auth(sig) {
		t.Fail()
	}
	t.Log("success")
}
