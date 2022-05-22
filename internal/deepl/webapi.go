package deepl

import (
	"github.com/candy12t/deepl-cli/internal/entity"
	"github.com/candy12t/deepl-cli/internal/webapi"
)

var _ webapi.Translater = &Client{}

func (c *Client) Translate(translation *entity.Translation) (*entity.Translation, error) {
	translateList, err := c.translate(translation.OriginalText, translation.SourceLanguage, translation.TargetLanguage)
	if err != nil {
		return nil, err
	}
	translation.TranslatedText = translateList.Translations[0].Text
	return translation, nil
}
