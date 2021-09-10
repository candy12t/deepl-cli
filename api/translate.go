package api

import (
	"net/url"
	"strings"

	"github.com/candy12t/deepl-cli/internal/host"
)

type TranslateResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

func Translate(text, sourceLang, TargetLang string) (*TranslateResponse, error) {
	values := buildValues(text, sourceLang, TargetLang)

	translate := &TranslateResponse{}
	err := Request(
		host.TranslateEndpoint(),
		"POST",
		strings.NewReader(values.Encode()),
		translate,
		map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
	)
	if err != nil {
		return nil, err
	}

	return translate, nil
}

func (tr *TranslateResponse) TranslatedText() string {
	return tr.Translations[0].Text
}

func buildValues(text, sourceLang, TargetLang string) *url.Values {
	return &url.Values{
		"text":        {text},
		"source_lang": {sourceLang},
		"target_lang": {TargetLang},
	}
}
