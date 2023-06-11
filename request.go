package matchmaker

import (
	"bytes"
	"net/http"
	"reflect"

	"github.com/0xzer/matchmaker/util"
)

func (c *Client) PostRequest(url string, payload []byte, headers interface{}) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	reqHeaders := &http.Header{}
	SetHeaders(reqHeaders, headers)
	req.Header = *reqHeaders
	//c.Logger.Info().Any("headers", req.Header).Msg("POST Request Headers")
	res, reqErr := c.http.Do(req)
	if reqErr != nil {
		return res, reqErr
	}
	if res.StatusCode == 401 {
		if !c.Authenticator.isRefreshing() {
			_, err := c.Authenticator.RefreshAuth()
			if err != nil {
				return res, err
			}
			var headers util.Headers
			headers.Build(*c.Device, c.App)
			headers.ContentType = req.Header.Get("content-type")
			res, reqErr = c.PostRequest(url, payload, headers)
			return res,reqErr
		} else {
			return res,ErrUnauthorized
		}
	}
	return res, nil
}

func (c *Client) GetRequest(url string, headers interface{}) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	reqHeaders := &http.Header{}
	SetHeaders(reqHeaders, headers)
	req.Header = *reqHeaders
	//c.Logger.Info().Any("headers", req.Header).Msg("GET Request Headers")
	res, reqErr := c.http.Do(req)
	if reqErr != nil {
		return res, reqErr
	}
	if res.StatusCode == 401 {
		if !c.Authenticator.isRefreshing() {
			_, err := c.Authenticator.RefreshAuth()
			if err != nil {
				return res, err
			}
			var headers util.Headers
			headers.Build(*c.Device, c.App)
			headers.ContentType = req.Header.Get("content-type")
			res, reqErr = c.GetRequest(url, headers)
			return res,reqErr
		} else {
			return res,ErrUnauthorized
		}
	}
	return res, nil
}

func SetHeaders(h *http.Header, headers interface{}) {
	if headers == nil {
		return
	}
	v := reflect.ValueOf(headers)
	for i := 0; i < v.NumField(); i++ {
	    field := v.Type().Field(i)
		value := v.Field(i).String()
		if !v.Field(i).IsZero() {
			h.Set(field.Tag.Get("header"), value)
		}
	}
}