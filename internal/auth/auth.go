package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/api/proto/auth"
	"github.com/txchat/dtalk/api/proto/signal"
	"github.com/txchat/dtalk/internal/auth/tools"
)

type EndpointRejectResp struct {
	Result  int    `json:"result"`
	Message string `json:"message"`
	Data    struct {
		Code    int    `json:"code"`
		Service string `json:"service"`
		Message struct {
			UUID       string `json:"uuid"`
			Device     int    `json:"device"`
			DeviceName string `json:"deviceName"`
			Datetime   int64  `json:"datetime"`
		} `json:"message"`
	} `json:"data"`
}

type ErrorReject signal.SignalEndpointLogin

func (e *ErrorReject) Error() string {
	return "reconnect not be allowed code -1016"
}

func (e *ErrorReject) Domain() string {
	return Name
}

func (e *ErrorReject) Encoding() (string, error) {
	data, err := proto.Marshal((*signal.SignalEndpointLogin)(e))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func errorDetail(respData *ErrorDataReconnectNotAllowed) *ErrorReject {
	return &ErrorReject{
		Uuid:       respData.Message.UUID,
		Device:     auth.Device(respData.Message.Device),
		DeviceName: respData.Message.DeviceName,
		Datetime:   respData.Message.Datetime,
	}
}

func DecodingErrorReject(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

type talkClient struct {
	url     string
	timeout time.Duration
}

func (a *talkClient) DoAuth(token string, ext []byte) (uid string, err error) {
	var (
		bytes     []byte
		strParams string
	)
	headers := map[string]string{}
	headers["FZM-SIGNATURE"] = token

	if len(ext) != 0 {
		var device auth.Login
		err = proto.Unmarshal(ext, &device)
		if err != nil {
			return "", err
		}

		headers["FZM-UUID"] = device.Uuid
		headers["FZM-DEVICE"] = device.Device.String()
		headers["FZM-DEVICE-NAME"] = device.DeviceName
		headers["Content-type"] = "application/json"
		reqBody := gin.H{
			"connType": device.ConnType,
		}
		reqData, err := json.Marshal(reqBody)
		if err != nil {
			return "", err
		}
		strParams = string(reqData)
	}

	bytes, err = tools.HTTPReq(&tools.HTTPParams{
		Method:    "POST",
		ReqURL:    a.url,
		HeaderMap: headers,
		Timeout:   a.timeout,
		StrParams: strParams,
	})
	if err != nil {
		return
	}

	var res Reply
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return
	}
	log.Debug().Interface("resp", res).Msg("auth reply")

	switch res.Result {
	case 0:
		var success SuccessData
		if err = mapstructure.Decode(res.Data, &success); err != nil {
			err = errors.New("invalid auth success data")
			return
		}
		uid = success.Address
	case -1016:
		var errNotAllowed ErrorDataReconnectNotAllowed
		if err = mapstructure.Decode(res.Data, &errNotAllowed); err != nil {
			return
		}
		err = errorDetail(&errNotAllowed)
	default:
		err = errors.New(res.Message)
	}
	log.Debug().Interface("err", err).Msg("auth reply code")
	return
}
