package deepl

import (
	"net/url"
	"strings"
)

type TranslateList struct {
	Translations []Translation `json:"translations"`
}

type Translation struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}

// DeepL API docs: https://www.deepl.com/ja/docs-api/translating-text
func (c *Client) Translate(text, sourceLang, targetLang string) (*TranslateList, error) {
	values := url.Values{}
	values.Add("text", text)
	values.Add("source_lang", sourceLang)
	values.Add("target_lang", targetLang)

	req, err := c.newRequest("POST", "/translate", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	translateList := new(TranslateList)
	if err := decodeBody(resp, translateList); err != nil {
		return nil, err
	}

	return translateList, nil
}

func (t *TranslateList) TranslateText() string {
	return t.Translations[0].Text
}
