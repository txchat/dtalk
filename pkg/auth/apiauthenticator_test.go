package auth

import (
	"encoding/hex"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	xcrypt "github.com/txchat/dtalk/pkg/crypt"
	secp256k1_ethereum "github.com/txchat/dtalk/pkg/crypt/secp256k1-ethereum"
	secp256k1_haltingstate "github.com/txchat/dtalk/pkg/crypt/secp256k1-haltingstate"
)

var userAddress = "1AqutxNoVTtcWiVYpBtvficAgea1dYTddR"

func TestDefaultAuthAndVerify(t *testing.T) {
	pubKey, err := hex.DecodeString(publicKey)
	assert.Nil(t, err)

	privKey, err := hex.DecodeString(privateKey)
	assert.Nil(t, err)

	client := NewDefaultApiAuthenticator()
	sig := client.Request("dtalk", pubKey, privKey)

	server := NewDefaultApiAuthenticator()
	uid, err := server.Auth(sig)
	assert.Nil(t, err)
	assert.Equal(t, userAddress, uid)
}

func TestHaltAuthAndEthVerify(t *testing.T) {
	pubKey, err := hex.DecodeString(publicKey)
	assert.Nil(t, err)

	privKey, err := hex.DecodeString(privateKey)
	assert.Nil(t, err)

	haltDriver, err := xcrypt.Load(secp256k1_haltingstate.Name)
	if err != nil {
		panic(err)
	}
	ethereumDriver, err := xcrypt.Load(secp256k1_ethereum.Name)
	if err != nil {
		panic(err)
	}
	client := NewDefaultApiAuthenticatorAsDriver(haltDriver)
	sig := client.Request("dtalk", pubKey, privKey)

	server := NewDefaultApiAuthenticatorAsDriver(ethereumDriver)
	uid, err := server.Auth(sig)
	assert.Nil(t, err)
	assert.Equal(t, userAddress, uid)
}

func TestEthAuthAndHaltVerify(t *testing.T) {
	pubKey, err := hex.DecodeString(publicKey)
	assert.Nil(t, err)

	privKey, err := hex.DecodeString(privateKey)
	assert.Nil(t, err)

	haltDriver, err := xcrypt.Load(secp256k1_haltingstate.Name)
	if err != nil {
		panic(err)
	}
	ethereumDriver, err := xcrypt.Load(secp256k1_ethereum.Name)
	if err != nil {
		panic(err)
	}
	client := NewDefaultApiAuthenticatorAsDriver(ethereumDriver)
	sig := client.Request("dtalk", pubKey, privKey)

	server := NewDefaultApiAuthenticatorAsDriver(haltDriver)
	uid, err := server.Auth(sig)
	assert.Nil(t, err)
	assert.Equal(t, userAddress, uid)
}

func TestDefaultVerifyInvalid(t *testing.T) {
	sig := "#1655716335340*dtalk#022db0e08669b30c5dab8c564b428db4944912144088943ec9b690a9046bc8f78b"
	server := NewDefaultApiAuthenticator()
	uid, err := server.Auth(sig)
	assert.ErrorAs(t, err, &SIGNATUREINVALIDERR{})
	assert.Empty(t, uid)
}

func TestDefaultAuthExpire(t *testing.T) {
	sig := "zrrLQ9FLnpON9s3erkMJ+sug5oviPYcOR04/w4ucC5dVbkjqvuZIFUuGgaKi+5XmDzj2FvxrDIom9dt2Ons6fAE=#1655716335340*dtalk#022db0e08669b30c5dab8c564b428db4944912144088943ec9b690a9046bc8f78b"
	server := NewDefaultApiAuthenticator()
	uid, err := server.Auth(sig)
	assert.ErrorIs(t, err, ERR_SIGNATUREEXPIRED)
	assert.Empty(t, uid)
}

func TestError(t *testing.T) {
	sig := "#1655716335340*dtalk#022db0e08669b30c5dab8c564b428db4944912144088943ec9b690a9046bc8f78b"
	server := NewDefaultApiAuthenticator()
	uid, err := server.Auth(sig)
	assert.Empty(t, uid)

	assert.True(t, errors.As(err, &SIGNATUREINVALIDERR{}))
	assert.ErrorAs(t, err, &SIGNATUREINVALIDERR{})
	// once unWarp not nil
	assert.NotNil(t, errors.Unwrap(err))
	// twice unWarp is nil
	assert.Nil(t, errors.Unwrap(errors.Unwrap(err)))
}
