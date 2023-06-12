package address

import (
	"encoding/hex"
	"testing"
)

func Test_Encoding(t *testing.T) {
	in, err := hex.DecodeString("02cec0b297406fc5298e9fe829c9f6f96fa176a1bd1a7c55fa0d345a6f49b09d25")
	if err != nil {
		t.Error(err)
		return
	}
	addr := PublicKeyToAddress(NormalVer, in)
	t.Log(addr)
	if err := CheckAddress(NormalVer, addr); err != nil {
		t.Error(err)
		return
	}
	t.Log("check success")
}

func Test_CheckAddress(t *testing.T) {
	addr := "1JoFzozbxvst22c2K7MBYwQGjCaMZbC5Qm"
	if err := CheckAddress(NormalVer, addr); err != nil {
		t.Error(err)
		return
	}
	t.Log("check success")
}
