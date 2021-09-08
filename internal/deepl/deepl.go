package deepl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/candy12t/deepl-cli/internal/config"
	"github.com/candy12t/deepl-cli/internal/host"
)

type DeeplRequest struct {
	AuthKey    string
	Text       string
	SourceLang string
	TargetLang string
}

type DeeplResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

func New(text, source, target string) *DeeplRequest {
	dr := &DeeplRequest{
		AuthKey:    config.AuthKey(),
		Text:       text,
		SourceLang: source,
		TargetLang: target,
	}

	return dr
}

func (dr *DeeplRequest) Post() ([]byte, error) {
	values := url.Values{
		"auth_key":    {dr.AuthKey},
		"text":        {dr.Text},
		"source_lang": {dr.SourceLang},
		"target_lang": {dr.TargetLang},
	}
	resp, err := http.PostForm(host.TranslateEndpoint(), values)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return body, nil
}

func (dr *DeeplRequest) Data() *DeeplResponse {
	var deeplResp DeeplResponse
	body, err := dr.Post()
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &deeplResp); err != nil {
		log.Fatal(err)
	}

	return &deeplResp
}

func (dresp *DeeplResponse) TranslatedText() string {
	return dresp.Translations[0].Text
}
