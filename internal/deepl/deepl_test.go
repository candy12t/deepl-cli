package deepl

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testProAuthKey  = "test-pro-auth-key"
	testFreeAuthKey = "test-free-auth-key:fx"
)

func TestNewClient(t *testing.T) {
	c := NewClient(testProAuthKey)
	assert.Equal(t, "https://api.deepl.com/v2", c.BaseURL.String())
	c2 := NewClient(testProAuthKey)
	assert.NotSame(t, c.HTTPClient, c2.HTTPClient)
}

func TestNewClientURLbyAccountType(t *testing.T) {
	type args struct {
		authKey string
	}
	tests := []struct {
		name    string
		args    args
		wantURL string
	}{
		{
			name: "deepl client for pro account",
			args: args{
				authKey: testProAuthKey,
			},
			wantURL: "https://api.deepl.com/v2",
		},
		{
			name: "deepl client for free account",
			args: args{
				authKey: testFreeAuthKey,
			},
			wantURL: "https://api-free.deepl.com/v2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.args.authKey)
			assert.Equal(t, tt.wantURL, c.BaseURL.String())
		})
	}
}

func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := NewClient(testProAuthKey)
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
	data, err := io.ReadAll(r.Body)
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
