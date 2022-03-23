package service

import (
	"encoding/hex"
	"testing"

	"github.com/txchat/dtalk/service/record/answer/config"
)

func TestService_Push(t *testing.T) {
	config.Conf = config.Default()
	svc := New(config.Conf)

	body, err := hex.DecodeString("1297011a2463653238663235312d326137612d346163312d613335352d34616462616436303866313222223134563767384141396e78523461347479737a683974344b766254585373504864712a2131424b4d35776d6b69596e4b4e6b72724839764259624e734c516b71737147563130013a1f8f8658c7f0fa4ebfb9713e76eeb1beedafc0d6356bdcffbc135d892b30d9e7409dc9aeb9a52f")
	if err != nil {
		t.Error(err)
		return
	}
	got, got1, err := svc.Push("", "14V7g8AA9nxR4a4tyszh9t4KvbTXSsPHdq", body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(got, got1)
}
