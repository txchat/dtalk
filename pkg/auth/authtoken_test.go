package auth

import (
	"encoding/hex"
	"testing"
	"time"
)

func TestCreateAuthAsClient(t *testing.T) {
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

	authT := NewAuthToken("dtalk", int64(time.Now().UnixNano()/1000))

	ar := NewApiRequest(authT.getToken(privKey), authT.getMetadata(), pubKey)
	t.Log(ar.getSig())
}

func TestCreateAuthAsServer(t *testing.T) {
	//sig := "mGIbusvBL/otgLJvDaBIWTVO7PQ6qEP3egqd4aepArYUJxvl6bV/eDYagQHZbeh7lBgO3Hlc8lc6eslEV03iPgA=#1645949322767597*dtalk#02b2dcf40123a5364a4bc9fd717db92122f90321a6771a47bc922100c9852c8b68"
	sig := "cIwU/HCDoaqRb9GYeXoLpv2L/Qojuvn2SNgbbkLheD9Lh37AY3iDpawH9uUtDj0j8pp/i0LTiKQNzWw9d0UsFQA=#164594983600*dtalk#02b2dcf40123a5364a4bc9fd717db92122f90321a6771a47bc922100c9852c8b68"

	ar, _ := NewApiRequestFromSig(sig)
	authT, _ := NewAuthTokenFromMetadata(ar.getMetadata())
	t.Log(hex.EncodeToString(ar.getPublicKey()))
	t.Logf("isExpire:%v\n", authT.isExpire())
	t.Logf("match:%v\n", authT.match(ar.getToken(), ar.getPublicKey()))
}
