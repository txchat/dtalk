package auth

import (
	"errors"
	"fmt"

	"github.com/txchat/dtalk/pkg/address"
	xcrypt "github.com/txchat/dtalk/pkg/crypt"

	//secp256k1_ethereum "github.com/txchat/dtalk/pkg/crypt/secp256k1-ethereum"
	"time"

	secp256k1_haltingstate "github.com/txchat/dtalk/pkg/crypt/secp256k1-haltingstate"
)

var ERR_SIGNATUREEXPIRED = errors.New("signature expired")
var ERR_SIGNATUREINVALID = func(e error) SIGNATUREINVALIDERR {
	return SIGNATUREINVALIDERR{content: "signature invalid: %w", err: e}
}
var ERR_UIDINVALID = errors.New("uid invalid")

type SIGNATUREINVALIDERR struct {
	content string
	err     error
}

func (e SIGNATUREINVALIDERR) Error() string {
	return fmt.Sprintf("%s:%s", e.content, e.err)
}

func (e SIGNATUREINVALIDERR) Unwrap() error { return e.err }

type ApiAuthenticator interface {
	Request(appId string, pubKey, privKey []byte) string
	Auth(sig string) (string, error)
}

type defaultApiAuthenticator struct {
	crypt xcrypt.Encrypt
}

func NewDefaultApiAuthenticator() *defaultApiAuthenticator {
	driver, err := xcrypt.Load(secp256k1_haltingstate.Name)
	if err != nil {
		panic(err)
	}
	return &defaultApiAuthenticator{
		crypt: driver,
	}
}

func NewDefaultApiAuthenticatorAsDriver(driver xcrypt.Encrypt) *defaultApiAuthenticator {
	return &defaultApiAuthenticator{
		crypt: driver,
	}
}

func (d *defaultApiAuthenticator) Request(appId string, pubKey, privKey []byte) string {
	signatory := NewSignatory(d.crypt, appId, time.Now().UnixMilli())

	apiRequest := NewApiRequest(signatory.DoSignature(privKey), signatory.GetMetadata(), pubKey)
	return apiRequest.GetToken()
}

func (d *defaultApiAuthenticator) Auth(token string) (string, error) {
	apiRequest, err := NewApiRequestFromToken(token)
	if err != nil {
		return "", err
	}
	signatory, err := NewSignatoryFromMetadata(d.crypt, apiRequest.GetMetadata())
	if err != nil {
		return "", err
	}
	if isMatch, err := signatory.Match(apiRequest.GetSignature(), apiRequest.GetPublicKey()); !isMatch {
		return "", ERR_SIGNATUREINVALID(err)
	}
	if signatory.IsExpire() {
		return "", ERR_SIGNATUREEXPIRED
	}
	uid := address.PublicKeyToAddress(address.NormalVer, apiRequest.GetPublicKey())
	if uid == "" {
		return "", ERR_UIDINVALID
	}
	return uid, nil
}
