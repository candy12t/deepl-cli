package deepl

import (
	"context"
	"net/url"
	"strings"
)

const TRANSLATE = "/translate"

type Translate struct {
	Translations []Translation `json:"translations"`
}

type Translation struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}

func (c *Client) Translate(ctx context.Context, text, sourceLang, targetLang string) (*Translate, error) {
	values := url.Values{
		"text":        {text},
		"source_lang": {sourceLang},
		"target_lang": {targetLang},
	}
	req, err := c.newRequest(ctx, "POST", TRANSLATE, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	translate := new(Translate)
	if err := decodeBody(resp, translate); err != nil {
		return nil, err
	}

	return translate, nil
}

func (t *Translate) TranslateText() string {
	return t.Translations[0].Text
}
