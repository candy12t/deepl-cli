package deepl

import (
	"net/url"
	"strings"

	"github.com/candy12t/deepl-cli/internal/model"
)

type TranslationResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

func (c *Client) TranslateText(text, sourceLang, targetLang string) (*model.TranslateText, error) {
	tr, err := c.translateText(text, sourceLang, targetLang)
	if err != nil {
		return nil, err
	}
	tt := &model.TranslateText{
		OriginalText:   text,
		TranslateText:  tr.Translations[0].Text,
		SourceLanguage: sourceLang,
		TargetLanguage: targetLang,
	}
	return tt, nil
}

// DeepL API docs: https://www.deepl.com/ja/docs-api/translating-text
func (c *Client) translateText(text, sourceLang, targetLang string) (*TranslationResponse, error) {
	values := url.Values{}
	values.Add("text", text)
	values.Add("source_lang", sourceLang)
	values.Add("target_lang", targetLang)

	req, err := c.NewRequest("POST", "/translate", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tr := new(TranslationResponse)
	if err := c.Do(req, tr); err != nil {
		return nil, err
	}

	return tr, nil
}
