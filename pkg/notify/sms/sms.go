package sms

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/inconshreveable/log15"
	http_tools "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/notify"
	. "github.com/txchat/dtalk/pkg/notify/sms/model"
)

type SMS struct {
	appKey     string
	secretKey  string
	serviceUrl string
	msg        string
}

func NewSMS(url, appKey, secretKey, msg string) *SMS {
	return &SMS{
		serviceUrl: url,
		appKey:     appKey,
		secretKey:  secretKey,
		msg:        msg,
	}
}

func (s *SMS) Send(param map[string]string) (interface{}, error) {
	phone := param[notify.ParamMobile]
	ticket := param[notify.ParamTicket]
	businessId := param[notify.ParamBizId]
	codeType := param[notify.ParamCodeType]

	values := map[string]string{
		"mobile":     phone,
		"codetype":   codeType,
		"param":      s.msg,
		"ticket":     ticket,
		"businessId": businessId,
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	reqMethod := "POST"
	reqUrl := s.serviceUrl + "/send/sms2"
	strParams := MapToSortUrlEncode(values)

	sign := sginature(s.appKey, values, s.secretKey, timestamp)

	headerMap := map[string]string{
		"Content-Type":     "application/x-www-form-urlencoded",
		"FZM-Ca-Timestamp": timestamp,
		"FZM-Ca-AppKey":    s.appKey,
		"FZM-Ca-Signature": sign,
	}

	req, err := http.NewRequest(reqMethod, reqUrl, strings.NewReader(strParams))
	if err != nil {
		return nil, err
	}

	for k, v := range headerMap {
		req.Header.Add(k, v)
	}

	c := http.Client{
		Timeout: HttpReqTimeout,
	}

	resp, err := c.Do(req)
	if resp != nil {
		defer func() {
			err := resp.Body.Close()
			if err != nil {
			}
		}()
	}

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tresult interface{}
	err = json.Unmarshal(body, &tresult)
	if nil != err {
		return nil, err
	}

	result, ok := tresult.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invaild tresult")
	}

	sCode, err := http_tools.ParseInterface(result["code"], "string")
	if nil != err {
		return nil, err
	}

	sError, err := http_tools.ParseInterface(result["error"], "string")
	if nil != err {
		return nil, err
	}

	sMessage, err := http_tools.ParseInterface(result["message"], "string")
	if nil != err {
		return nil, err
	}

	if "200" != sCode.(string) || "succ" != sError.(string) || "succ" != sMessage.(string) {
		//return fmt.Errorf("code : " + sCode.(string) + ", error : " + sError.(string) + ", message : " + sMessage.(string))
		return nil, &Error{Code: sCode.(string), Err: sError.(string), Message: sMessage.(string)}
	}

	data, ok := result["data"]
	if !ok {
		return nil, fmt.Errorf("no 'data' info")
	}

	info := data.(map[string]interface{})
	log15.Debug("send result", "data", info)
	isShow := int(info["isShow"].(float64))
	isValidate := int(info["isValidate"].(float64))

	var rltData map[string]interface{}
	if rltData, ok = info["data"].(map[string]interface{}); ok {
	}

	return &SendResult{
		IsShow:     isShow,
		IsValidate: isValidate,
		Data:       rltData,
	}, nil
}

func (s *SMS) ValidateCode(param map[string]string) error {
	phone := param[notify.ParamMobile]
	code := param[notify.ParamCode]
	codeType := param[notify.ParamCodeType]

	values := map[string]string{
		"t":        "sms",
		"codetype": codeType,
		"code":     code,
		"guide":    "0",
		"mobile":   phone,
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	reqMethod := "POST"
	reqUrl := s.serviceUrl + "/validate/code"
	strParams := MapToSortUrlEncode(values)

	sign := sginature(s.appKey, values, s.secretKey, timestamp)

	headerMap := map[string]string{
		"Content-Type":     "application/x-www-form-urlencoded",
		"FZM-Ca-Timestamp": timestamp,
		"FZM-Ca-AppKey":    s.appKey,
		"FZM-Ca-Signature": sign,
	}

	req, err := http.NewRequest(reqMethod, reqUrl, strings.NewReader(strParams))
	if err != nil {
		return err
	}

	for k, v := range headerMap {
		req.Header.Add(k, v)
	}

	c := http.Client{
		Timeout: HttpReqTimeout,
	}

	resp, err := c.Do(req)
	if resp != nil {
		defer func() {
			err := resp.Body.Close()
			if err != nil {
			}
		}()
	}

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var tresult interface{}
	err = json.Unmarshal(body, &tresult)
	if nil != err {
		return err
	}

	result, ok := tresult.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invaild tresult")
	}

	sCode, err := http_tools.ParseInterface(result["code"], "string")
	if nil != err {
		return err
	}

	sError, err := http_tools.ParseInterface(result["error"], "string")
	if nil != err {
		return err
	}

	sMessage, err := http_tools.ParseInterface(result["message"], "string")
	if nil != err {
		return err
	}

	if "200" != sCode.(string) || "succ" != sError.(string) || "succ" != sMessage.(string) {
		//return fmt.Errorf("code : " + sCode.(string) + ", error : " + sError.(string) + ", message : " + sMessage.(string))
		return &Error{Code: sCode.(string), Err: sError.(string), Message: sMessage.(string)}
	}

	return nil
}

func sginature(appKey string, req map[string]string, secretKey string, time string) string {
	signParams := MapToSortUrlEncode(req)
	signParams = appKey + signParams + secretKey + time
	h := md5.New()
	h.Write([]byte(signParams))
	cipgerStr := h.Sum(nil)
	sign := strings.ToUpper(hex.EncodeToString(cipgerStr))

	return sign
}

func MapToSortUrlEncode(paramsMap map[string]string) string {
	v := url.Values{}

	mapKeys := []string{}
	for k := range paramsMap {
		mapKeys = append(mapKeys, k)
	}
	sort.Strings(mapKeys)

	for k := range mapKeys {
		v.Add(mapKeys[k], paramsMap[mapKeys[k]])
	}
	body := v.Encode()
	return body
}
