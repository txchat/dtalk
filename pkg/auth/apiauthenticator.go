package auth

import (
	"time"
)

type ApiAuthenticator interface {
	Request(appId string, pubKey, privKey []byte) string
	Auth(sig string) bool
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

func (d *defaultApuAuthenticator) Auth(sig string) bool {
	ar, err := NewApiRequestFromSig(sig)
	if err != nil {
		return false
	}
	authT, err := NewAuthTokenFromMetadata(ar.getMetadata())
	if err != nil {
		return false
	}
	return !authT.isExpire() && authT.match(ar.getToken(), ar.getPublicKey())
}
