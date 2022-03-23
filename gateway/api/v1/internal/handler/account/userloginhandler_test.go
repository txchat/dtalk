package account

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/txchat/dtalk/pkg/api"
)

func performRequest(r http.Handler, method, path string, headers map[string]string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewReader(body))
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAddressLoginConnect(t *testing.T) {
	reqHeader := map[string]string{
		"Content-type":    "application/json",
		"FZM-SIGNATURE":   "MOCK",
		"FZM-UUID":        "123",
		"FZM-DEVICE":      "macOS",
		"FZM-DEVICE-NAME": "test mac",
	}
	reqBody := gin.H{
		"connType": 0,
	}
	respBody := gin.H{
		"result":  0,
		"message": "",
		"data": gin.H{
			"address": "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
		},
	}

	reqData, err := json.Marshal(reqBody)
	if err != nil {
		t.Fail()
	}
	//serverCtx := svc.NewServiceContext(*config.Conf)

	router := gin.Default()
	router.Use(api.AuthMiddleWare(), api.RespMiddleWare())
	router.POST("/", AddressLogin(nil))

	w := performRequest(router, "POST", "/", reqHeader, reqData)

	assert.Equal(t, http.StatusOK, w.Code)

	//
	var response struct {
		Result  int    `json:"result"`
		Message string `json:"message"`
		Data    struct {
			Address string `json:"address"`
		} `json:"data"`
	}
	t.Log(w.Body.String())
	err = json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, respBody["result"], response.Result)
	assert.Equal(t, respBody["data"].(gin.H)["address"], response.Data.Address)
}

func TestAddressLoginReconnectReject(t *testing.T) {
	reqHeader := map[string]string{
		"Content-type":    "application/json",
		"FZM-SIGNATURE":   "MOCK",
		"FZM-UUID":        "123",
		"FZM-DEVICE":      "Android",
		"FZM-DEVICE-NAME": "test",
	}
	reqBody := gin.H{
		"connType": 1,
	}
	respBody := gin.H{
		"result":  -1016,
		"message": "设备重连不被允许",
		"data": gin.H{
			"code":    -1016,
			"service": "",
			"message": gin.H{
				"UUid":       "456",
				"Device":     0,
				"DeviceName": "xiaomi",
				"Datetime":   int64(1),
			},
		},
	}

	reqData, err := json.Marshal(reqBody)
	if err != nil {
		t.Fail()
	}
	//serverCtx := svc.NewServiceContext(*config.Conf)

	router := gin.Default()
	router.Use(api.AuthMiddleWare(), api.RespMiddleWare())
	router.POST("/", AddressLogin(nil))

	w := performRequest(router, "POST", "/", reqHeader, reqData)

	assert.Equal(t, http.StatusOK, w.Code)

	//
	var response struct {
		Result  int    `json:"result"`
		Message string `json:"message"`
		Data    struct {
			Code    int    `json:"code"`
			Service string `json:"service"`
			Message struct {
				UUid       string `json:"uuid"`
				Device     int    `json:"device"`
				DeviceName string `json:"deviceName"`
				Datetime   int64  `json:"datetime"`
			} `json:"message"`
		} `json:"data"`
	}
	t.Log(w.Body.String())
	err = json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, respBody["result"], response.Result)
	assert.Equal(t, respBody["data"].(gin.H)["message"].(gin.H)["UUid"], response.Data.Message.UUid)
	assert.Equal(t, respBody["data"].(gin.H)["message"].(gin.H)["Device"], response.Data.Message.Device)
	assert.Equal(t, respBody["data"].(gin.H)["message"].(gin.H)["DeviceName"], response.Data.Message.DeviceName)
	assert.Equal(t, respBody["data"].(gin.H)["message"].(gin.H)["Datetime"], response.Data.Message.Datetime)
}

func TestAddressLoginReconnectAllowed(t *testing.T) {
	reqHeader := map[string]string{
		"Content-type":    "application/json",
		"FZM-SIGNATURE":   "MOCK",
		"FZM-UUID":        "456",
		"FZM-DEVICE":      "Android",
		"FZM-DEVICE-NAME": "test",
	}
	reqBody := gin.H{
		"connType": 1,
	}
	respBody := gin.H{
		"result":  0,
		"message": "",
		"data": gin.H{
			"address": "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
		},
	}

	reqData, err := json.Marshal(reqBody)
	if err != nil {
		t.Fail()
	}
	//serverCtx := svc.NewServiceContext(*config.Conf)

	router := gin.Default()
	router.Use(api.AuthMiddleWare(), api.RespMiddleWare())
	router.POST("/", AddressLogin(nil))

	w := performRequest(router, "POST", "/", reqHeader, reqData)

	assert.Equal(t, http.StatusOK, w.Code)

	//
	var response struct {
		Result  int    `json:"result"`
		Message string `json:"message"`
		Data    struct {
			Address string `json:"address"`
		} `json:"data"`
	}
	t.Log(w.Body.String())
	err = json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, respBody["result"], response.Result)
	assert.Equal(t, respBody["data"].(gin.H)["address"], response.Data.Address)
}
