package deepl

import (
	"github.com/candy12t/deepl-cli/internal/adapter/webapi"
	"github.com/candy12t/deepl-cli/internal/entity"
)

var _ webapi.DeepL = &Client{}

func (c *Client) Translate(t *entity.Translation) (*entity.Translation, error) {
	translateList, err := c.translate(t.Original, t.Source, t.Target)
	if err != nil {
		return nil, err
	}
	t.Text = translateList.Translations[0].Text
	return t, nil
}
