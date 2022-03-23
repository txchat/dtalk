package http

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const HttpReqTimeout = 20 * time.Second

func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if nil != resp {
		defer func() {
			err := resp.Body.Close()
			if err != nil {
			}
		}()
	}

	if nil != err {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func HTTPPostForm(reqUrl string, headers map[string]string, payload io.Reader) ([]byte, error) {
	if headers == nil {
		headers = make(map[string]string)
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	return HTTPRequest("POST", reqUrl, headers, payload)
}

// HTTPPostJSON with json []byte
func HTTPPostJSON(reqUrl string, headers map[string]string, payload io.Reader) ([]byte, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "text/json"
	return HTTPRequest("POST", reqUrl, headers, payload)
}

func HTTPRequest(method string, reqUrl string, headers map[string]string, payload io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, reqUrl, payload)
	if err != nil {
		return nil, errors.New("make http request error: " + err.Error())
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	/*ctx := req.Context()
	nextCtx, _ := context.WithDeadline(ctx, time.Now().Add(HttpReqTimeout))
	req.WithContext(nextCtx)*/

	client := &http.Client{
		Timeout: HttpReqTimeout,
	}
	resp, err := client.Do(req)
	if resp != nil {
		defer func() {
			err := resp.Body.Close()
			if err != nil {
			}
		}()
	}

	if err != nil {
		return nil, errors.New("do http request error: " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("ready http response error: " + err.Error())
	}

	if resp.StatusCode != 200 {
		return body, fmt.Errorf(fmt.Sprintf("HttpStatusCode:%d, Desc:%s", resp.StatusCode, string(body)))
	}

	return body, nil
}
