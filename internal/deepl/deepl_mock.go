package deepl

import (
	"context"
	"net/url"
)

type MockClient struct {
	BaseURL *url.URL
	AuthKey string
}

func NewMockClient(rawBaseURL, authKey string) (*MockClient, error) {
	if len(authKey) == 0 {
		return nil, ErrMissingAuthKey
	}

	baseURL, err := url.ParseRequestURI(rawBaseURL)
	if err != nil {
		return nil, err
	}

	return &MockClient{
		BaseURL: baseURL,
		AuthKey: authKey,
	}, nil
}

func (mc *MockClient) Translate(ctx context.Context, text, sourceLang, targetLang string) (*Translate, error) {
	return &Translate{Translations: []Translation{{DetectedSourceLanguage: "EN", Text: "こんにちわ"}}}, nil
}
