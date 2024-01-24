package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/parkeringskompagniet/goquickpay/pkg/httpmethod"

	"github.com/gorilla/schema"
)

type QuickpayClient struct {
	BaseUrl string
	ApiKey  string
}

func NewClient(baseUrl, apiKey string) QuickpayClient {
	return QuickpayClient{baseUrl, apiKey}
}

func (c QuickpayClient) setupRequest(method httpmethod.HTTPMethod, endpoint string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(string(method), c.BaseUrl+endpoint, body)
	if err != nil {
		return nil, errors.New("there was an error setting up base request")
	}

	encodedAPIKey := base64.StdEncoding.EncodeToString([]byte(":" + c.ApiKey))

	request.Header.Add("Accept-Version", "v10")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "Basic "+encodedAPIKey)

	return request, nil
}

func (c QuickpayClient) CallEndpoint(method httpmethod.HTTPMethod, endpoint string, data any) (*http.Response, error) {
	request, err := c.PrepareEndPoint(method, endpoint, data)
	if err != nil {
		return nil, err
	}

	return c.CallEndPointWith(request)
}

func (q QuickpayClient) PrepareEndPoint(method httpmethod.HTTPMethod, endpoint string, data any) (*http.Request, error) {
	if data == nil {
		return q.setupRequest(method, endpoint, strings.NewReader(""))
	}

	body, err := q.EncodeBody(data)
	if err != nil {
		return nil, err
	}

	return q.setupRequest(method, endpoint, body)
}

func (c QuickpayClient) CallEndPointWith(request *http.Request) (*http.Response, error) {
	client := &http.Client{}

	return client.Do(request)
}

// NOTE: as of now use gorilla/schema for handling encoding,

func (c QuickpayClient) EncodeQuery(data any) (string, error) {
	encoder := schema.NewEncoder()
	values := url.Values{}
	err := encoder.Encode(data, values)
	if err != nil {
		return "", err
	}

	return values.Encode(), nil
}

// TODO: custom parser for custom schema
func (c QuickpayClient) EncodeBody(data any) (io.Reader, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(bytes)), nil
}
