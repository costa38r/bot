package openaiclient

import (
	"net/http"
)

// HttpClient é a interface para realizar requisições HTTP.
type HttpClient interface {
    Do(req *http.Request) (*http.Response, error)
}

// HttpClientImpl é a implementação padrão da interface HttpClient.
type HttpClientImpl struct {
    Client *http.Client
}

// Do realiza a requisição HTTP.
func (c *HttpClientImpl) Do(req *http.Request) (*http.Response, error) {
    return c.Client.Do(req)
}