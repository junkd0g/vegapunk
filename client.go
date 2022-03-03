package vegapunk

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	defaultDomain = "api.openai.com"
)

type client struct {
	domain string
	apiKey string
	client *http.Client
}

func NewClient(domain, apiKey string) *client {
	return &client{
		domain: domain,
		apiKey: apiKey,
		client: &http.Client{},
	}
}

func NewClientDefaultDomain(apiKey string) *client {
	return &client{
		domain: "",
		apiKey: apiKey,
		client: &http.Client{},
	}
}

func (c client) get(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", defaultDomain, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, resErr := c.client.Do(req)
	if resErr != nil {
		return nil, resErr
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)

	return body, bodyErr
}
