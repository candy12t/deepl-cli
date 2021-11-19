package deepl

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const defaultBaseURL = "https://api.deepl.com/v2"
const testAuthKey = "test-auth-key"

func TestNewClient(t *testing.T) {
	t.Run("success new deepl client", func(t *testing.T) {
		c, err := NewClient(defaultBaseURL, testAuthKey)
		testNoErr(t, err)

		if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
			t.Errorf("NewClient BaseURL is %v, want %v", got, want)
		}
	})

	t.Run("failed new deepl client because missing deepl authkey", func(t *testing.T) {
		_, err := NewClient(defaultBaseURL, "")
		testErr(t, err)
		testMissingAuthKeyErr(t, err)
	})

	t.Run("failed new deepl client because don't parse url", func(t *testing.T) {
		_, err := NewClient("%", testAuthKey)
		testErr(t, err)
		testURLParseErr(t, err)
	})
}

func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client, _ := NewClient(defaultBaseURL, testAuthKey)
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

func testMissingAuthKeyErr(t *testing.T, err error) {
	t.Helper()
	if err != ErrMissingAuthKey {
		t.Fatalf("got err is %v, want %v", err, ErrMissingAuthKey)
	}
}

func testNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func testErr(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}
