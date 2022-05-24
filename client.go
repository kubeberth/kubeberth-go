package kubeberth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Config struct {
	URL string
}

func NewConfig(url string) *Config {
	return &Config{
		URL: url,
	}
}

type Client struct {
	httpClient *http.Client
	config     *Config
}

func NewClient(config *Config) *Client {
	return &Client{
		httpClient: &http.Client{},
		config: config,
	}
}

func decodeBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	return decoder.Decode(out)
}

func (c *Client) newRequest(
	ctx context.Context,
	method string,
	apiPath string,
	body io.Reader,
) (*http.Request, error) {
	u, err := url.Parse(c.config.URL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, apiPath)
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}
