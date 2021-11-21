package deepl

import (
	"context"
	"fmt"
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
	if len(targetLang) == 0 {
		return nil, fmt.Errorf(`HTTP 400: "Value for 'target_lang' not supported." (https://api-free.deepl.com/v2/translate)`)
	}
	return &Translate{Translations: []Translation{{DetectedSourceLanguage: "EN", Text: "こんにちわ"}}}, nil
}
