package usecase

import (
	"github.com/candy12t/deepl-cli/internal/adapter/usecase"
	"github.com/candy12t/deepl-cli/internal/adapter/webapi"
	"github.com/candy12t/deepl-cli/internal/entity"
)

type DeepL struct {
	webapi webapi.DeepL
}

var _ usecase.DeepL = &DeepL{}

func NewTranslation(webapi webapi.DeepL) *DeepL {
	return &DeepL{
		webapi: webapi,
	}
}

func (d *DeepL) Translate(original, source, target string) (string, error) {
	t, err := d.webapi.Translate(
		&entity.Translation{
			Original: original,
			Source:   source,
			Target:   target,
		},
	)
	if err != nil {
		return "", err
	}
	return t.Text, nil
}
