package phpserverclient

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Email      string
	Phone      string
	Ticket     string
	BusinessId string
	CodeType   string
	Code       string
}

type Client struct {
	appKey     string
	secretKey  string
	serviceURL string
	msg        string
}

func NewClient(url, appKey, secretKey, msg string) *Client {
	return &Client{
		appKey:     appKey,
		secretKey:  secretKey,
		serviceURL: url,
		msg:        msg,
	}
}

func (c *Client) SendSMS(cfg *Config) (*SendResult, error) {
	values := map[string]string{
		"mobile":     cfg.Phone,
		"codetype":   cfg.CodeType,
		"param":      c.msg,
		"ticket":     cfg.Ticket,
		"businessId": cfg.BusinessId,
	}
	reqURL := c.serviceURL + "/send/sms2"
	return c.send(reqURL, values)
}

func (c *Client) SendEmail(cfg *Config) (*SendResult, error) {
	values := map[string]string{
		"email":      cfg.Email,
		"codetype":   cfg.CodeType,
		"param":      c.msg,
		"ticket":     cfg.Ticket,
		"businessId": cfg.BusinessId,
	}
	reqURL := c.serviceURL + "/send/email2"
	return c.send(reqURL, values)
}

func (c *Client) ValidateSMS(cfg *Config) error {
	values := map[string]string{
		"t":        "sms",
		"codetype": cfg.CodeType,
		"code":     cfg.Code,
		"guide":    "0",
		"mobile":   cfg.Phone,
	}
	reqURL := c.serviceURL + "/validate/code"
	return c.validate(reqURL, values)
}

func (c *Client) ValidateEmail(cfg *Config) error {
	values := map[string]string{
		"t":        "email",
		"codetype": cfg.CodeType,
		"code":     cfg.Code,
		"guide":    "0",
		"email":    cfg.Email,
	}
	reqURL := c.serviceURL + "/validate/code"
	return c.validate(reqURL, values)
}

func (c *Client) send(reqURL string, param map[string]string) (*SendResult, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	reqMethod := "POST"
	strParams := MapToSortURLEncode(param)

	sign := sginature(c.appKey, param, c.secretKey, timestamp)

	headerMap := map[string]string{
		"Content-Type":     "application/x-www-form-urlencoded",
		"FZM-Ca-Timestamp": timestamp,
		"FZM-Ca-AppKey":    c.appKey,
		"FZM-Ca-Signature": sign,
	}

	req, err := http.NewRequest(reqMethod, reqURL, strings.NewReader(strParams))
	if err != nil {
		return nil, err
	}

	for k, v := range headerMap {
		req.Header.Add(k, v)
	}

	cli := http.Client{
		Timeout: HTTPReqTimeout,
	}

	resp, err := cli.Do(req)
	if resp != nil {
		defer func() {
			err = resp.Body.Close()
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

	sCode, err := ParseInterface(result["code"], "string")
	if nil != err {
		return nil, err
	}

	sError, err := ParseInterface(result["error"], "string")
	if nil != err {
		return nil, err
	}

	sMessage, err := ParseInterface(result["message"], "string")
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

func (c *Client) validate(reqURL string, param map[string]string) error {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	reqMethod := "POST"
	strParams := MapToSortURLEncode(param)

	sign := sginature(c.appKey, param, c.secretKey, timestamp)

	headerMap := map[string]string{
		"Content-Type":     "application/x-www-form-urlencoded",
		"FZM-Ca-Timestamp": timestamp,
		"FZM-Ca-AppKey":    c.appKey,
		"FZM-Ca-Signature": sign,
	}

	req, err := http.NewRequest(reqMethod, reqURL, strings.NewReader(strParams))
	if err != nil {
		return err
	}

	for k, v := range headerMap {
		req.Header.Add(k, v)
	}

	cli := http.Client{
		Timeout: HTTPReqTimeout,
	}

	resp, err := cli.Do(req)
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

	sCode, err := ParseInterface(result["code"], "string")
	if nil != err {
		return err
	}

	sError, err := ParseInterface(result["error"], "string")
	if nil != err {
		return err
	}

	sMessage, err := ParseInterface(result["message"], "string")
	if nil != err {
		return err
	}

	if "200" != sCode.(string) || "succ" != sError.(string) || "succ" != sMessage.(string) {
		return &Error{Code: sCode.(string), Err: sError.(string), Message: sMessage.(string)}
	}

	return nil
}

func sginature(appKey string, req map[string]string, secretKey string, time string) string {
	signParams := MapToSortURLEncode(req)
	signParams = appKey + signParams + secretKey + time
	h := md5.New()
	h.Write([]byte(signParams))
	cipgerStr := h.Sum(nil)
	sign := strings.ToUpper(hex.EncodeToString(cipgerStr))

	return sign
}

func MapToSortURLEncode(paramsMap map[string]string) string {
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

func ParseInterface(orign interface{}, ty string) (interface{}, error) {
	var result interface{}

	switch ty {
	case "":
		return nil, fmt.Errorf("invalid ty")
	case "int":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = tOrign
		case uint:
			result = int(tOrign)
		case int32:
			result = int(tOrign)
		case uint32:
			result = int(tOrign)
		case int64:
			result = int(tOrign)
		case uint64:
			result = int(tOrign)
		case float32:
			result = int(tOrign)
		case float64:
			result = int(tOrign)
		case string:
			tm, err := strconv.ParseInt(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = int(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "uint":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = uint(tOrign)
		case uint:
			result = tOrign
		case int32:
			result = uint(tOrign)
		case uint32:
			result = uint(tOrign)
		case int64:
			result = uint(tOrign)
		case uint64:
			result = uint(tOrign)
		case float32:
			result = uint(tOrign)
		case float64:
			result = uint(tOrign)
		case string:
			tm, err := strconv.ParseUint(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = uint(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "int32":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = int32(tOrign)
		case uint:
			result = int32(tOrign)
		case int32:
			result = tOrign
		case uint32:
			result = int32(tOrign)
		case int64:
			result = int32(tOrign)
		case uint64:
			result = int32(tOrign)
		case float32:
			result = int32(tOrign)
		case float64:
			result = int32(tOrign)
		case string:
			tm, err := strconv.ParseInt(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = int32(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "uint32":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = uint32(tOrign)
		case uint:
			result = uint32(tOrign)
		case int32:
			result = uint32(tOrign)
		case uint32:
			result = tOrign
		case int64:
			result = uint32(tOrign)
		case uint64:
			result = uint32(tOrign)
		case float32:
			result = uint32(tOrign)
		case float64:
			result = uint32(tOrign)
		case string:
			tm, err := strconv.ParseUint(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = uint32(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "int64":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = int64(tOrign)
		case uint:
			result = int64(tOrign)
		case int32:
			result = int64(tOrign)
		case uint32:
			result = int64(tOrign)
		case int64:
			result = tOrign
		case uint64:
			result = int64(tOrign)
		case float32:
			result = int64(tOrign)
		case float64:
			result = int64(tOrign)
		case string:
			tm, err := strconv.ParseInt(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = tm
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "uint64":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = uint64(tOrign)
		case uint:
			result = uint64(tOrign)
		case int32:
			result = uint64(tOrign)
		case uint32:
			result = uint64(tOrign)
		case int64:
			result = uint64(tOrign)
		case uint64:
			result = tOrign
		case float32:
			result = uint64(tOrign)
		case float64:
			result = uint64(tOrign)
		case string:
			tm, err := strconv.ParseUint(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = tm
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "float32":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = float32(tOrign)
		case uint:
			result = float32(tOrign)
		case int32:
			result = float32(tOrign)
		case uint32:
			result = float32(tOrign)
		case int64:
			result = float32(tOrign)
		case uint64:
			result = float32(tOrign)
		case float32:
			result = tOrign
		case float64:
			result = float32(tOrign)
		case string:
			tm, err := strconv.ParseFloat(tOrign, 32)
			if nil != err {
				return nil, err
			}
			result = float32(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "float64":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = float64(tOrign)
		case uint:
			result = float64(tOrign)
		case int32:
			result = float64(tOrign)
		case uint32:
			result = float64(tOrign)
		case int64:
			result = float64(tOrign)
		case uint64:
			result = float64(tOrign)
		case float32:
			result = float64(tOrign)
		case float64:
			result = tOrign
		case string:
			tm, err := strconv.ParseFloat(tOrign, 64)
			if nil != err {
				return nil, err
			}
			result = tm
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "string":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = fmt.Sprint(uint64(tOrign))
		case uint:
			result = fmt.Sprint(uint64(tOrign))
		case int32:
			result = fmt.Sprint(uint64(tOrign))
		case uint32:
			result = fmt.Sprint(uint64(tOrign))
		case int64:
			result = fmt.Sprint(uint64(tOrign))
		case uint64:
			result = fmt.Sprint(tOrign)
		case float32:
			// 这种只适合整数转字符串的情形, 也就是id那种情况, 带小数的转换不支持
			if float64(tOrign) > float64(uint64(tOrign)) {
				return nil, fmt.Errorf("not support the condition, float64(tOrign) > float64(uint64(tOrign))")
			}

			result = fmt.Sprint(uint64(tOrign))
		case float64:
			// 这种只适合整数转字符串的情形, 也就是id那种情况, 带小数的转换不支持
			if tOrign > float64(uint64(tOrign)) {
				return nil, fmt.Errorf("not support the condition, tOrign > float64(uint64(tOrign))")
			}
			result = fmt.Sprint(uint64(tOrign))
		case string:
			result = tOrign
		case bool:
			result = fmt.Sprint(tOrign)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "bool":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case bool:
			result = tOrign
		case string:
			tm, err := strconv.ParseBool(tOrign)
			if nil != err {
				return nil, err
			}
			result = tm
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	default:
		return nil, fmt.Errorf("unknow ty")
	}

	return result, nil
}
