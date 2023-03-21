package httpreq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/zeromicro/go-zero/rest/httpc"

	"chatim/shared/errorx"
)

const (
	contentType     = "Content-Type"
	applicationJson = "application/json"
)

type (
	Client struct {
		svc httpc.Service
	}
	errResp struct {
		Code int    `json:"code"`
		Desc string `json:"desc,omitempty"`
	}
)

// NewDefaultClient new http client
func NewDefaultClient(opts ...httpc.Option) *Client {
	res := &Client{}
	res.svc = httpc.NewService("default", opts...)
	return res
}

// DoWithJson data will marshal to json
func (c *Client) DoWithJson(method, reqUrl string, data, val interface{}) error {
	err := c.validate(val)
	if err != nil {
		return err
	}
	var body io.Reader
	if data != nil {
		reqBody, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body = bytes.NewReader(reqBody)
	}
	req, err := http.NewRequest(method, reqUrl, body)
	if err != nil {
		return err
	}
	req.Header.Set(contentType, applicationJson)

	resp, err := c.svc.DoRequest(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return processResponse(resp, val)
}

// Do data will marshal json tag to json; path to url path; form to form url
func (c *Client) Do(ctx context.Context, method, url string, data interface{}, val interface{}) error {
	err := c.validate(val)
	if err != nil {
		return err
	}

	resp, err := c.svc.Do(ctx, method, url, data)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return processResponse(resp, val)
}

// DoRequest support other req like form content req
func (c *Client) DoRequest(r *http.Request, val interface{}) error {
	err := c.validate(val)
	if err != nil {
		return err
	}

	resp, err := c.svc.DoRequest(r)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return processResponse(resp, val)
}

func (c *Client) validate(val interface{}) error {
	if val != nil {
		v := reflect.ValueOf(val)
		if !v.IsValid() || v.Kind() != reflect.Ptr || v.IsNil() {
			return fmt.Errorf("error: not a valid pointer: %v", v)
		}
	}
	return nil
}

func processResponse(resp *http.Response, val interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var errResp errResp
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return errorx.NewStatCodeError(resp.StatusCode, 1,
				fmt.Sprintf("unmarshal %s failed: %s", string(body), err.Error()))
		}
		return errorx.NewStatCodeError(resp.StatusCode, errResp.Code, errResp.Desc)
	}

	if val != nil {
		if err := json.Unmarshal(body, val); err != nil {
			return errorx.NewStatCodeError(resp.StatusCode, 2,
				fmt.Sprintf("unmarshal %s failed: %s", string(body), err.Error()))
		}
	}

	return nil
}
