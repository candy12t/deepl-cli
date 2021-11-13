package deepl

import (
	"context"
	"net/url"
	"strings"
)

const TRANSLATE = "/translate"

type TranslateResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

func (c *Client) Translate(ctx context.Context, text, sourceLang, targetLang string) (string, error) {
	values := url.Values{
		"text":        {text},
		"source_lang": {sourceLang},
		"target_lang": {targetLang},
	}
	req, err := c.newRequest(ctx, "POST", TRANSLATE, strings.NewReader(values.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	transalted := new(TranslateResponse)
	if err := decodeBody(resp, transalted); err != nil {
		return "", err
	}

	return transalted.Translations[0].Text, nil
}
