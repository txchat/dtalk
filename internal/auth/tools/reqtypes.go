package tools

import (
	"encoding/json"
	"time"
)

type HTTPParams struct {
	Method    string            `json:"Method"`
	ReqURL    string            `json:"ReqUrl"`
	StrParams string            `json:"StrParams"`
	HeaderMap map[string]string `json:"HeaderMap"`
	Timeout   time.Duration     `json:"Timeout"`
}

func HTTPParamsUnmarshal(data []byte) (*HTTPParams, error) {
	params := &HTTPParams{}
	err := json.Unmarshal(data, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
