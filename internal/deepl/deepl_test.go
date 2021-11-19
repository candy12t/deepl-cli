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

func testNoErr(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
