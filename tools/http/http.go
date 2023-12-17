package http

import (
	"errors"
	"net/http"
)

type HttpHeader struct {
	Path          string `json:"path"`
	ContentLength string `json:"content_length"`
	Cookie        string `json:"cookie"`
}

func Post(request *http.Request, client *http.Client, contentType string) (*http.Response, error) {
	var (
		err      error
		response = &http.Response{}
		errMsg   string
	)
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	response, err = client.Do(request)
	if err != nil || response == nil || response.StatusCode != 200 {
		if err != nil {
			errMsg = err.Error()
		} else if response != nil {
			errMsg = response.Status
		} else {
			errMsg = "请求的接口发生未知错误"
		}
		return response, errors.New(errMsg)
	}
	return response, err
}

func Get(request *http.Request, client *http.Client) (*http.Response, error) {
	var (
		err      error
		response = &http.Response{}
		errMsg   string
	)
	response, err = client.Do(request)
	if err != nil || response == nil || response.StatusCode != 200 {
		if err != nil {
			errMsg = err.Error()
		} else if response != nil {
			errMsg = response.Status
		} else {
			errMsg = "请求的接口发生未知错误"
		}
		return response, errors.New(errMsg)
	}
	return response, err
}

func Get1(url string) (*http.Response, error) {
	var (
		err      error
		response = &http.Response{}
		errMsg   string
	)
	response, err = http.Get(url)
	if err != nil || response == nil || response.StatusCode != 200 {
		if err != nil {
			errMsg = err.Error()
		} else if response != nil {
			errMsg = response.Status
		} else {
			errMsg = "请求的接口发生未知错误"
		}
		return response, errors.New(errMsg)
	}
	return response, nil
}
