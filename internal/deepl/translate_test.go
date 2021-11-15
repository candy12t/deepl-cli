package deepl

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/candy12t/deepl-cli/test"
)

func TestTranslate(t *testing.T) {
	t.Run("success translate", func(t *testing.T) {
		client, server, err := deeplMockServer(t, "test-auth-key", "success-body", "success-header")
		if err != nil {
			t.Fatal(err)
		}
		defer server.Close()

		ctx := context.Background()
		got, err := client.Translate(ctx, "hello", "EN", "JA")
		if err != nil {
			t.Fatal(err)
		}
		want := "こんにちわ"

		if got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func deeplMockServer(t *testing.T, mockAuthKey string, mockResponseBody, mockResponseHeader string) (*Client, *httptest.Server, error) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := filepath.Join(test.ProjectDirPath(), "test", "testdata", "deepl", mockResponseHeader)
		headerBytes, err := ioutil.ReadFile(header)
		h := strings.Split(string(headerBytes), "\n")
		for _, line := range h {
			if strings.Contains(line, "HTTP") {
				statusCode, err := strconv.Atoi(strings.Split(line, " ")[1])
				if err != nil {
					t.Fatal(err)
				}
				w.WriteHeader(statusCode)
			}

			if strings.Contains(line, "content-type") {
				contentType := strings.Split(line, " ")[1]
				w.Header().Add("content-type", contentType)
			}
		}

		body := filepath.Join(test.ProjectDirPath(), "test", "testdata", "deepl", mockResponseBody)
		bodyBytes, err := ioutil.ReadFile(body)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(bodyBytes)
	}))

	serverURL, err := url.ParseRequestURI(server.URL)
	if err != nil {
		return nil, nil, err
	}

	return &Client{
		BaseURL:    serverURL,
		HTTPClient: server.Client(),
		AuthKey:    mockAuthKey,
	}, server, nil
}
