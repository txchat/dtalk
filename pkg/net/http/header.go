package http

import (
	"context"
	"net/http"
)

type customHeaderKey struct{}

type Custom struct {
	Signature  string
	UUID       string
	Device     string
	DeviceName string
	Version    string
	UID        string
}

func ConvertCustom(r *http.Request) *Custom {
	return &Custom{
		Signature:  r.Header.Get(HeaderSignature),
		UUID:       r.Header.Get(HeaderUUID),
		Device:     r.Header.Get(HeaderDeviceType),
		DeviceName: r.Header.Get(HeaderDeviceName),
		Version:    r.Header.Get(HeaderVersion),
	}
}

func (c *Custom) SetWithRequest(r *http.Request) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), customHeaderKey{}, c))
}

func FromContext(ctx context.Context) (*Custom, bool) {
	bh, ok := ctx.Value(customHeaderKey{}).(*Custom)
	return bh, ok
}
