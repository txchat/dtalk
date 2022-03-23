package auth

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type ApiRequest struct {
	token    string
	metadata string
	pubKey   []byte
}

func NewApiRequest(token, metadata string, pubKey []byte) *ApiRequest {
	return &ApiRequest{
		token:    token,
		metadata: metadata,
		pubKey:   pubKey,
	}
}

func NewApiRequestFromSig(sig string) (*ApiRequest, error) {
	ss := strings.SplitN(sig, "#", -1)
	if len(ss) < 3 {
		return nil, fmt.Errorf("signature parse feilds error need 3 got %d", len(ss))
	}
	pubKey, err := hex.DecodeString(ss[2])
	if err != nil {
		return nil, err
	}
	return &ApiRequest{
		token:    ss[0],
		metadata: ss[1],
		pubKey:   pubKey,
	}, nil
}

// get signature
func (t *ApiRequest) getSig() string {
	return fmt.Sprintf("%s#%s#%s", t.token, t.metadata, hex.EncodeToString(t.pubKey))
}

func (t *ApiRequest) getToken() string {
	return t.token
}

func (t *ApiRequest) getMetadata() string {
	return t.metadata
}

func (t *ApiRequest) getPublicKey() []byte {
	return t.pubKey
}
