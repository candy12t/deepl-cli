package deepl

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/candy12t/deepl-cli/test"
)

const defaultBaseURL = "https://api.deepl.com/v2"
const testAuthKey = "test-auth-key"

func TestNewClient(t *testing.T) {
	t.Run("success new deepl client", func(t *testing.T) {
		c, err := NewClient(testAuthKey)
		test.AssertError(t, err, nil)

		if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
			t.Errorf("NewClient BaseURL is %v, want %v", got, want)
		}
	})

	t.Run("failed new deepl client because missing deepl authkey", func(t *testing.T) {
		_, err := NewClient("")
		test.AssertError(t, err, ErrMissingAuthKey)
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
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	t.Helper()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("request Body is %s, want %s", got, want)
	}
}

func testURLParseErr(t *testing.T, err error) {
	t.Helper()
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}
