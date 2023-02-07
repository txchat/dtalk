package tencentyun

import (
	"testing"
)

func TestGenUserSign(t *testing.T) {
	ts := NewTCTLSSig(1400543084, "1ed1b5e2729395c1e8b55b83be72e60139dfa36ff3fa67bc3c3285592a7b3cf6", 86400)
	userSign1, err := ts.GetUserSig("AAAABBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	if err != nil {

	}
	t.Log(userSign1)
	userSign2, err := ts.GetUserSig("AAAABBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	if err != nil {

	}
	t.Log(userSign2)
	userSign3, err := ts.GetUserSig("share_chenhongyu")
	if err != nil {

	}
	t.Log(userSign3)
	userSign4, err := ts.GetUserSig("chenhongyu")
	if err != nil {

	}
	t.Log(userSign4)
}
