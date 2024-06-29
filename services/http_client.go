package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HttpClient struct {
	baseUrl string
}

func NewHttpClient(baseUrl string) *HttpClient {
	return &HttpClient{baseUrl: baseUrl}
}

func (c *HttpClient) Get(relativeUrl string) (*http.Response, error) {
	return http.Get(c.baseUrl + relativeUrl)
}

func (c *HttpClient) Post(relativeUrl string, body interface{}) (*http.Response, error) {
	if body == nil {
		return http.Post(c.baseUrl+relativeUrl, "application/json", nil)
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	result, err := http.Post(c.baseUrl+relativeUrl, "application/json", bytes.NewBuffer(jsonData))
	return result, err
}
