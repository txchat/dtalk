package auth

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type APIRequest struct {
	signature string
	metadata  string
	pubKey    []byte
}

func NewAPIRequest(signature, metadata string, pubKey []byte) *APIRequest {
	return &APIRequest{
		signature: signature,
		metadata:  metadata,
		pubKey:    pubKey,
	}
}

func NewAPIRequestFromToken(token string) (*APIRequest, error) {
	ss := strings.SplitN(token, "#", -1)
	if len(ss) < 3 {
		return nil, fmt.Errorf("token parse feilds error need 3 got %d", len(ss))
	}
	pubKey, err := hex.DecodeString(ss[2])
	if err != nil {
		return nil, err
	}
	return &APIRequest{
		signature: ss[0],
		metadata:  ss[1],
		pubKey:    pubKey,
	}, nil
}

func (t *APIRequest) GetToken() string {
	return fmt.Sprintf("%s#%s#%s", t.signature, t.metadata, hex.EncodeToString(t.pubKey))
}

func (t *APIRequest) GetSignature() string {
	return t.signature
}

func (t *APIRequest) GetMetadata() string {
	return t.metadata
}

func (t *APIRequest) GetPublicKey() []byte {
	return t.pubKey
}
