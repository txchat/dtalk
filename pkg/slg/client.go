package slg

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/goinggo/mapstructure"
	"github.com/txchat/dtalk/pkg/net/http"
)

type HTTPClientConfig struct {
	Host string
	Salt string
}

var ErrResponseParams = errors.New("返回参数格式错误")

type groupPermissionVerificationConditions struct {
	CategoryId string `json:"categoryId"`
}

type groupPermissionVerificationParams struct {
	Conditions []*groupPermissionVerificationConditions `json:"conditions"`
	Type       int32                                    `json:"type"`
	UID        string                                   `json:"uid"`
}

type groupPermissionVerificationReq struct {
	Params []*groupPermissionVerificationParams `json:"params"`
	Sign   string                               `json:"sign"`
}

type basicResp struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type groupPermissionVerificationDataResp struct {
	UID  string `json:"uid"`
	Flag bool   `json:"flag"`
}

func convertGroupPermissionVerificationConditions(conditions []string) []*groupPermissionVerificationConditions {
	cons := make([]*groupPermissionVerificationConditions, len(conditions))
	for i, condition := range conditions {
		cons[i] = &groupPermissionVerificationConditions{
			CategoryId: condition,
		}
	}
	return cons
}

func convertGroupPermissionVerificationParams(conditions []*UserCondition) []*groupPermissionVerificationParams {
	params := make([]*groupPermissionVerificationParams, len(conditions))
	for i, condition := range conditions {
		params[i] = &groupPermissionVerificationParams{
			UID:        condition.UID,
			Type:       condition.HandleType,
			Conditions: convertGroupPermissionVerificationConditions(condition.Conditions),
		}
	}
	return params
}

func convertGroupPermission(data interface{}) (GroupPermission, error) {
	paramTypes := reflect.TypeOf(data)
	paramValues := reflect.ValueOf(data)

	switch paramTypes.Kind() {
	case reflect.Ptr:
		paramTypes = paramTypes.Elem()
		paramValues = paramValues.Elem()
	}

	groupPermission := make(map[string]bool)
	switch paramTypes.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < paramValues.Len(); i++ {
			var o = groupPermissionVerificationDataResp{}
			err := mapstructure.Decode(paramValues.Index(i).Interface(), &o)
			if err != nil {
				return nil, ErrResponseParams
			}
			groupPermission[o.UID] = o.Flag
		}
	default:
		return nil, ErrResponseParams
	}
	return groupPermission, nil
}

type HTTPClient struct {
	host string
	salt string
}

func NewHTTPClient(host, salt string) *HTTPClient {
	return &HTTPClient{
		host: host,
		salt: salt,
	}
}

func (c *HTTPClient) GroupPermissionVerification(conditions []*UserCondition) (GroupPermission, error) {
	if len(conditions) <= 0 {
		return make(map[string]bool), nil
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	reqBody := &groupPermissionVerificationReq{
		Params: convertGroupPermissionVerificationParams(conditions),
		Sign:   "",
	}

	unSignedParamsReqData, err := json.Marshal(reqBody.Params)
	if err != nil {
		return nil, err
	}

	// signature
	reqBody.Sign = c.sign(unSignedParamsReqData)

	signedReqData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	respData, err := http.HTTPRequest("POST", c.host+"/root/mall/open/chat/group-permission-verification", headers, bytes.NewReader(signedReqData))
	if err != nil {
		return nil, err
	}

	var resp basicResp
	err = json.Unmarshal(respData, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != "00000" {
		return nil, errors.New(resp.Msg)
	}

	rlt, err := convertGroupPermission(resp.Data)
	if err != nil {
		return nil, err
	}
	return rlt, nil
}

func (c *HTTPClient) sign(data []byte) string {
	// format signature
	sig := fmt.Sprintf("salt=%s&params=%s&salt=%s", c.salt, string(data), c.salt)
	// md5 encoding
	md5Data := md5.Sum([]byte(sig))
	return hex.EncodeToString(md5Data[:])
}
