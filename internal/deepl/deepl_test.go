package deepl

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testAuthKey = "test-auth-key"

func TestNewClient(t *testing.T) {
	t.Run("success new deepl client", func(t *testing.T) {
		c, err := NewClient(testAuthKey)
		if assert.NoError(t, err) {
			assert.Equal(t, "https://api.deepl.com/v2", c.BaseURL.String())
		}
	})

	t.Run("failed new deepl client because missing deepl authkey", func(t *testing.T) {
		_, err := NewClient("")
		assert.EqualError(t, err, ErrMissingAuthKey.Error())
	})
}

func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client, _ := NewClient(testAuthKey)
	url, _ := url.ParseRequestURI(server.URL)
	client.BaseURL = url

	return client, mux, server.Close
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	t.Helper()
	assert.Equal(t, want, r.Header.Get(header))
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	assert.Equal(t, want, r.Method)
}

func testBody(t *testing.T, r *http.Request, want string) {
	t.Helper()
	data, err := ioutil.ReadAll(r.Body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, string(data))
	}
}

func testURLParseErr(t *testing.T, err error) {
	t.Helper()
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}
