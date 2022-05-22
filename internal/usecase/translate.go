package usecase

import (
	"github.com/candy12t/deepl-cli/internal/entity"
	"github.com/candy12t/deepl-cli/internal/webapi"
)

type Translation struct {
	webapi webapi.Translater
}

func NewTranslation(webapi webapi.Translater) *Translation {
	return &Translation{
		webapi: webapi,
	}
}

func (t *Translation) Translate(translation *entity.Translation) (*entity.Translation, error) {
	return t.webapi.Translate(translation)
}
