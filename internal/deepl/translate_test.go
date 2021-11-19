package deepl

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const defaultBaseURL = "https://api.deepl.com/v2"

func TestTranslate(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	body := `{"translations":[{"detected_source_language":"EN","text":"こんにちわ"}]}`
	want := "こんにちわ"

	mux.HandleFunc(TRANSLATE, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, "Content-Type", "application/x-www-form-urlencoded")
		testHeader(t, r, "Authorization", "DeepL-Auth-Key test-auth-key")
		testMethod(t, r, "POST")
		testBody(t, r, "source_lang=EN&target_lang=JA&text=hello")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	})

	ctx := context.Background()
	got, err := client.Translate(ctx, "hello", "EN", "JA")
	assertNoErr(t, err)
	assertTranslatedText(t, got, want)
}

func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client, _ := NewClient(defaultBaseURL, "test-auth-key")
	url, _ := url.ParseRequestURI(server.URL)
	client.BaseURL = url

	return client, mux, server.Close
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

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	t.Helper()
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func assertTranslatedText(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("translated text is %s, want %s", got, want)
	}
}

func assertNoErr(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
