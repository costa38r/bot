package openaiclient

import (
	"net/http"
)

type HttpClient interface {
    Do(req *http.Request) (*http.Response, error)
}

type HttpClientImpl struct {
    Client *http.Client
}

func (c *HttpClientImpl) Do(req *http.Request) (*http.Response, error) {
    return c.Client.Do(req)
}

