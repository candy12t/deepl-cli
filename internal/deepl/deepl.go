package deepl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/candy12t/deepl-cli/internal/repository"
)

const (
	FreeHost   = "https://api-free.deepl.com"
	ProHost    = "https://api.deepl.com"
	APIVersion = "v2"

	AccountPlanIdentificationKey = ":fx"
)

var _ repository.Translator = &Client{}

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
	if isPro(authKey) {
		return ProHost
	}
	return FreeHost
}

func isPro(authKey string) bool {
	return !strings.HasSuffix(authKey, AccountPlanIdentificationKey)
}

func (c *Client) NewRequest(method, _path string, body io.Reader) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, _path)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("DeepL-Auth-Key %s", c.AuthKey))

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	success := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !success {
		return HandleHTTPError(resp)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}
