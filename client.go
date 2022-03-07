package vegapunk

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	defaultDomain = "https://api.openai.com"
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
	url := fmt.Sprintf("%s%s", c.domain, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(`Authorization`, fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Add(`Content-Type`, `application/json`)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}

func (c client) post(endpoint string, payload string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.domain, endpoint)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(`Authorization`, fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Add(`Content-Type`, `application/json`)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}
