package auth

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	xcrypt "github.com/txchat/dtalk/pkg/crypt"
	//secp256k1_ethereum "github.com/txchat/dtalk/pkg/crypt/secp256k1-ethereum"
	secp256k1_haltingstate "github.com/txchat/dtalk/pkg/crypt/secp256k1-haltingstate"
)

var (
	privateKey = "fa884fd1b47d9e9e8dc19e47dc1a794a524ce5d4ee1b82ec92b1ffc1f109c2b6"
	publicKey  = "022db0e08669b30c5dab8c564b428db4944912144088943ec9b690a9046bc8f78b"
)

func TestCreateAuthAsClient(t *testing.T) {
	driver, err := xcrypt.Load(secp256k1_haltingstate.Name)
	assert.Nil(t, err)

	pubKey, err := hex.DecodeString(publicKey)
	assert.Nil(t, err)

	privKey, err := hex.DecodeString(privateKey)
	assert.Nil(t, err)

	authT := NewAuthToken(driver, "dtalk", time.Now().UnixNano()/1000)

	ar := NewApiRequest(authT.getToken(privKey), authT.getMetadata(), pubKey)
	t.Log(ar.getSig())
}

func TestCreateAuthAsServer(t *testing.T) {
	driver, err := xcrypt.Load(secp256k1_haltingstate.Name)
	assert.Nil(t, err)

	sig := "cIwU/HCDoaqRb9GYeXoLpv2L/Qojuvn2SNgbbkLheD9Lh37AY3iDpawH9uUtDj0j8pp/i0LTiKQNzWw9d0UsFQA=#164594983600*dtalk#02b2dcf40123a5364a4bc9fd717db92122f90321a6771a47bc922100c9852c8b68"

	ar, err := NewApiRequestFromSig(sig)
	assert.Nil(t, err)

	authT, err := NewAuthTokenFromMetadata(driver, ar.getMetadata())
	assert.Nil(t, err)

	isMatch, err := authT.match(ar.getToken(), ar.getPublicKey())
	assert.Nil(t, err)
	assert.Equal(t, true, isMatch)
	assert.Equal(t, true, authT.isExpire())
}
