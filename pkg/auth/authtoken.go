package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	xcrypt "github.com/txchat/dtalk/pkg/crypt"
	"github.com/txchat/dtalk/pkg/util"
)

const DefaultExpireTimeInterval = time.Second * 120

type AuthToken struct {
	appId              string
	timestamp          int64
	expireTimeInterval time.Duration
	crypt              xcrypt.Encrypt
}

func NewAuthToken(crypt xcrypt.Encrypt, appId string, timestamp int64) *AuthToken {
	return &AuthToken{
		crypt:              crypt,
		appId:              appId,
		timestamp:          timestamp,
		expireTimeInterval: DefaultExpireTimeInterval,
	}
}

func NewAuthTokenFromMetadata(crypt xcrypt.Encrypt, metadata string) (*AuthToken, error) {
	msg := strings.SplitN(metadata, "*", -1)
	if len(msg) < 2 {
		return nil, fmt.Errorf("metadata parse feilds error need 2 got %d", len(msg))
	}
	timestamp, err := strconv.ParseInt(msg[0], 10, 64)
	if err != nil {
		return nil, err
	}
	appId := msg[1]

	return &AuthToken{
		crypt:              crypt,
		appId:              appId,
		timestamp:          timestamp,
		expireTimeInterval: DefaultExpireTimeInterval,
	}, nil
}

func (t *AuthToken) getMetadata() string {
	return fmt.Sprintf("%d*%s", t.timestamp, t.appId)
}

func (t *AuthToken) getToken(privKey []byte) string {
	metadata := t.getMetadata()
	// sha356 encoding metadata
	msg256 := sha256.Sum256([]byte(metadata))
	// secp256k1(eth format) sign
	sig, _ := t.crypt.Sign(msg256[:], privKey)
	// base64 encoding signature
	token := base64.StdEncoding.EncodeToString(sig)
	return token
}

func (t *AuthToken) isExpire() bool {
	return util.CheckTimeOut(t.timestamp, t.expireTimeInterval)
}

func (t *AuthToken) match(token string, pubKey []byte) (bool, error) {
	// base64 decoding token
	sig, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false, err
	}
	msg256 := sha256.Sum256([]byte(t.getMetadata()))
	// secp256k1(eth format) verify
	return t.crypt.Verify(msg256[:], sig, pubKey)
}
