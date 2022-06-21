package auth

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type ApiRequest struct {
	signature string
	metadata  string
	pubKey    []byte
}

func NewApiRequest(signature, metadata string, pubKey []byte) *ApiRequest {
	return &ApiRequest{
		signature: signature,
		metadata:  metadata,
		pubKey:    pubKey,
	}
}

func NewApiRequestFromToken(token string) (*ApiRequest, error) {
	ss := strings.SplitN(token, "#", -1)
	if len(ss) < 3 {
		return nil, fmt.Errorf("token parse feilds error need 3 got %d", len(ss))
	}
	pubKey, err := hex.DecodeString(ss[2])
	if err != nil {
		return nil, err
	}
	return &ApiRequest{
		signature: ss[0],
		metadata:  ss[1],
		pubKey:    pubKey,
	}, nil
}

// get signature
func (t *ApiRequest) getToken() string {
	return fmt.Sprintf("%s#%s#%s", t.signature, t.metadata, hex.EncodeToString(t.pubKey))
}

func (t *ApiRequest) getSignature() string {
	return t.signature
}

func (t *ApiRequest) getMetadata() string {
	return t.metadata
}

func (t *ApiRequest) getPublicKey() []byte {
	return t.pubKey
}
