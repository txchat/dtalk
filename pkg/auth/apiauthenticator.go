package auth

import (
	"errors"
	"fmt"
	"github.com/txchat/dtalk/pkg/address"
	xcrypt "github.com/txchat/dtalk/pkg/crypt"
	//secp256k1_ethereum "github.com/txchat/dtalk/pkg/crypt/secp256k1-ethereum"
	secp256k1_haltingstate "github.com/txchat/dtalk/pkg/crypt/secp256k1-haltingstate"
	"time"
)

var ERR_SIGNATUREEXPIRED = errors.New("signature expired")
var ERR_SIGNATUREINVALID = errors.New("signature invalid")
var ERR_UIDINVALID = errors.New("uid invalid")

type ApiAuthenticator interface {
	Request(appId string, pubKey, privKey []byte) string
	Auth(sig string) (string, error)
}

type defaultApuAuthenticator struct {
	crypt xcrypt.Encrypt
}

func NewDefaultApuAuthenticator() *defaultApuAuthenticator {
	driver, err := xcrypt.Load(secp256k1_haltingstate.Name)
	if err != nil {
		panic(err)
	}
	return &defaultApuAuthenticator{
		crypt: driver,
	}
}

func (d *defaultApuAuthenticator) Request(appId string, pubKey, privKey []byte) string {
	authT := NewAuthToken(d.crypt, appId, time.Now().UnixNano()/1000)

	ar := NewApiRequest(authT.getToken(privKey), authT.getMetadata(), pubKey)
	return ar.getSig()
}

func (d *defaultApuAuthenticator) Auth(sig string) (string, error) {
	ar, err := NewApiRequestFromSig(sig)
	if err != nil {
		return "", err
	}
	authT, err := NewAuthTokenFromMetadata(d.crypt, ar.getMetadata())
	if err != nil {
		return "", err
	}
	if isMatch, err := authT.match(ar.getToken(), ar.getPublicKey()); !isMatch {
		return "", fmt.Errorf("%s, %s", ERR_SIGNATUREINVALID.Error(), err.Error())
	}
	if authT.isExpire() {
		return "", ERR_SIGNATUREEXPIRED
	}
	uid := address.PublicKeyToAddress(address.NormalVer, ar.getPublicKey())
	if uid == "" {
		return "", ERR_UIDINVALID
	}
	return uid, nil
}
