package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/haltingstate/secp256k1-go"
	"github.com/txchat/dtalk/pkg/util"
)

const DefaultExpireTimeInterval = time.Second * 120

type AuthToken struct {
	appId              string
	timestamp          int64
	expireTimeInterval time.Duration
}

func NewAuthToken(appId string, timestamp int64) *AuthToken {
	return &AuthToken{
		appId:              appId,
		timestamp:          timestamp,
		expireTimeInterval: DefaultExpireTimeInterval,
	}
}

func NewAuthTokenFromMetadata(metadata string) (*AuthToken, error) {
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
	// enc metadata
	msg256 := sha256.Sum256([]byte(metadata))
	token := base64.StdEncoding.EncodeToString(secp256k1.Sign(msg256[:], privKey))
	return token
}

func (t *AuthToken) isExpire() bool {
	return util.CheckTimeOut(t.timestamp, t.expireTimeInterval)
}

func (t *AuthToken) match(token string, pubKey []byte) bool {
	//desc msg
	sig, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false
	}
	msg256 := sha256.Sum256([]byte(t.getMetadata()))
	return util.Secp256k1Verify(msg256[:], sig, pubKey)
}
