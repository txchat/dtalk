package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/txchat/dtalk/pkg/address"
	xcrypt "github.com/txchat/dtalk/pkg/crypt"

	//secp256k1_ethereum "github.com/txchat/dtalk/pkg/crypt/secp256k1-ethereum"
	secp256k1_haltingstate "github.com/txchat/dtalk/pkg/crypt/secp256k1-haltingstate"
)

var ErrSignatureExpired = errors.New("signature expired")
var ErrSignatureInvalid = func(e error) SignatureInvalidError {
	return SignatureInvalidError{content: "signature invalid: %w", err: e}
}
var ErrUIDInvalid = errors.New("uid invalid")

type SignatureInvalidError struct {
	content string
	err     error
}

func (e SignatureInvalidError) Error() string {
	return fmt.Sprintf("%s:%s", e.content, e.err)
}

func (e SignatureInvalidError) Unwrap() error { return e.err }

type APIAuthenticator interface {
	Request(appId string, pubKey, privKey []byte) string
	Auth(sig string) (string, error)
}

type DefaultAPIAuthenticator struct {
	crypt xcrypt.Encrypt
}

func NewDefaultAPIAuthenticator() *DefaultAPIAuthenticator {
	driver, err := xcrypt.Load(secp256k1_haltingstate.Name)
	if err != nil {
		panic(err)
	}
	return &DefaultAPIAuthenticator{
		crypt: driver,
	}
}

func NewDefaultAPIAuthenticatorAsDriver(driver xcrypt.Encrypt) *DefaultAPIAuthenticator {
	return &DefaultAPIAuthenticator{
		crypt: driver,
	}
}

func (d *DefaultAPIAuthenticator) Request(appId string, pubKey, privKey []byte) string {
	signatory := NewSignatory(d.crypt, appId, time.Now().UnixMilli())

	apiRequest := NewAPIRequest(signatory.DoSignature(privKey), signatory.GetMetadata(), pubKey)
	return apiRequest.GetToken()
}

func (d *DefaultAPIAuthenticator) Auth(token string) (string, error) {
	apiRequest, err := NewAPIRequestFromToken(token)
	if err != nil {
		return "", err
	}
	signatory, err := NewSignatoryFromMetadata(d.crypt, apiRequest.GetMetadata())
	if err != nil {
		return "", err
	}
	if isMatch, err := signatory.Match(apiRequest.GetSignature(), apiRequest.GetPublicKey()); !isMatch {
		return "", ErrSignatureInvalid(err)
	}
	if signatory.IsExpire() {
		return "", ErrSignatureExpired
	}
	uid := address.PublicKeyToAddress(address.NormalVer, apiRequest.GetPublicKey())
	if uid == "" {
		return "", ErrUIDInvalid
	}
	return uid, nil
}
