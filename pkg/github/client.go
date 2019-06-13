package github

import (
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

type ValidHttpMethod string

const (
	Get  ValidHttpMethod = http.MethodGet
	Post ValidHttpMethod = http.MethodPost
)

const BaseURL = "https://api.github.com"

func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient: httpClient}
}

func (c *Client) GetPublicEvents(username string) interface {
	// TODO
}

func (c *Client) request(method ValidHttpMethod, path string) {
	request, err := http.NewRequest(
		string(method),
		path,
		http.NoBody,
	)
	if err != nil {
	}

	c.httpClient.Do(request)
}
