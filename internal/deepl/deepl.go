//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/$GOFILE

package deepl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type APIClient interface {
	Translate(context.Context, string, string, string) (*Translate, error)
}

var ErrMissingAuthKey = errors.New("missing DeepL authKey")

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	AuthKey    string
}

func NewClient(rawBaseURL, authKey string) (*Client, error) {
	if len(authKey) == 0 {
		return nil, ErrMissingAuthKey
	}

	baseURL, err := url.ParseRequestURI(rawBaseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
		AuthKey:    authKey,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method, _path string, body io.Reader) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, _path)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
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
