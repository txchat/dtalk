package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	xcrypt "github.com/txchat/dtalk/pkg/crypt"
)

const DefaultExpireTimeInterval = time.Second * 120

type Signatory struct {
	appId              string
	timestamp          int64
	expireTimeInterval time.Duration
	crypt              xcrypt.Encrypt
}

func NewSignatory(crypt xcrypt.Encrypt, appId string, timestamp int64) *Signatory {
	return &Signatory{
		crypt:              crypt,
		appId:              appId,
		timestamp:          timestamp,
		expireTimeInterval: DefaultExpireTimeInterval,
	}
}

func NewSignatoryFromMetadata(crypt xcrypt.Encrypt, metadata string) (*Signatory, error) {
	msg := strings.SplitN(metadata, "*", -1)
	if len(msg) < 2 {
		return nil, fmt.Errorf("metadata parse feilds error need 2 got %d", len(msg))
	}
	timestamp, err := strconv.ParseInt(msg[0], 10, 64)
	if err != nil {
		return nil, err
	}
	appId := msg[1]

	return &Signatory{
		crypt:              crypt,
		appId:              appId,
		timestamp:          timestamp,
		expireTimeInterval: DefaultExpireTimeInterval,
	}, nil
}

func (t *Signatory) getMetadata() string {
	return fmt.Sprintf("%d*%s", t.timestamp, t.appId)
}

func (t *Signatory) doSignature(privKey []byte) string {
	metadata := t.getMetadata()
	// sha356 encoding metadata
	msg256 := sha256.Sum256([]byte(metadata))
	// secp256k1(eth format) sign
	sig, _ := t.crypt.Sign(msg256[:], privKey)
	// base64 encoding signature
	token := base64.StdEncoding.EncodeToString(sig)
	return token
}

func (t *Signatory) isExpire() bool {
	return time.Now().After(time.UnixMilli(t.timestamp).Add(t.expireTimeInterval))
}

func (t *Signatory) match(signature string, pubKey []byte) (bool, error) {
	// base64 decoding signature
	sig, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	msg256 := sha256.Sum256([]byte(t.getMetadata()))
	// secp256k1(eth format) verify
	return t.crypt.Verify(msg256[:], sig, pubKey)
}
