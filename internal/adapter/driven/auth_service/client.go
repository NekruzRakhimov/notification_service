package auth_service

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	getAllEmailsEndpoint = "/internal/emails"
)

type Client struct {
	url        string
	httpClient *resty.Client
}

func New(url string) *Client {
	return &Client{
		url:        url,
		httpClient: resty.New(),
	}
}

func (c *Client) GetAllEmails() ([]string, error) {
	rawResp, err := c.httpClient.R().Get(fmt.Sprintf("%s%s", c.url, getAllEmailsEndpoint))
	if err != nil {
		return nil, err
	}

	var resp EmailsResponse
	if err = json.Unmarshal(rawResp.Body(), &resp); err != nil {
		return nil, err
	}

	return resp.Emails, nil
}
