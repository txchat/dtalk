package tools

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var client = &http.Client{}

func HTTPReq(params *HTTPParams) ([]byte, error) {
	var req *http.Request
	var err error

	switch params.Method {
	case GET, DELETE:
		if "" == params.StrParams {
			req, err = http.NewRequest(params.Method, params.ReqURL, nil)
		} else {
			req, err = http.NewRequest(params.Method, params.ReqURL+"?"+params.StrParams, nil)
		}
	case POST:
		req, err = http.NewRequest(params.Method, params.ReqURL, strings.NewReader(params.StrParams))
	default:
		panic(params.Method)
	}
	if nil != err {
		return nil, err
	}

	for k, v := range params.HeaderMap {
		req.Header.Add(k, v)
	}

	if 0 != int64(params.Timeout) {
		client.Timeout = params.Timeout
	} else {
		client.Timeout = DefaultTimeout
	}

	resp, err := client.Do(req)
	if nil != resp {
		defer resp.Body.Close()
	}

	if nil != err {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(fmt.Sprintf("http request errcode: %v", resp.StatusCode))
	}

	return ioutil.ReadAll(resp.Body)
}
