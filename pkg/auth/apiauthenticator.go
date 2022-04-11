package auth

import (
	"errors"
	"github.com/txchat/dtalk/pkg/address"
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
}

func NewDefaultApuAuthenticator() *defaultApuAuthenticator {
	return &defaultApuAuthenticator{}
}

func (d *defaultApuAuthenticator) Request(appId string, pubKey, privKey []byte) string {
	authT := NewAuthToken(appId, time.Now().UnixNano()/1000)

	ar := NewApiRequest(authT.getToken(privKey), authT.getMetadata(), pubKey)
	return ar.getSig()
}

func (d *defaultApuAuthenticator) Auth(sig string) (string, error) {
	ar, err := NewApiRequestFromSig(sig)
	if err != nil {
		return "", err
	}
	authT, err := NewAuthTokenFromMetadata(ar.getMetadata())
	if err != nil {
		return "", err
	}
	if !authT.match(ar.getToken(), ar.getPublicKey()) {
		return "", ERR_SIGNATUREINVALID
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
