//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/$GOFILE

package deepl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const (
	FreeHost             = "https://api-free.deepl.com"
	ProHost              = "https://api.deepl.com"
	APIVersion           = "v2"
	EndpointDetermineKey = ":fx"
)

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	AuthKey    string
}

func NewClient(authKey string) *Client {
	baseURL, _ := url.Parse(defaultHost(authKey))
	baseURL.Path = path.Join(baseURL.Path, APIVersion)

	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
		AuthKey:    authKey,
	}
}

func defaultHost(authKey string) string {
	if !strings.HasSuffix(authKey, EndpointDetermineKey) {
		return ProHost
	}
	return FreeHost
}

func (c *Client) newRequest(method, _path string, body io.Reader) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, _path)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("DeepL-Auth-Key %s", c.AuthKey))

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	success := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !success {
		return HandleHTTPError(resp)
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
