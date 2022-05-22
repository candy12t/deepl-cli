package entity

import (
	"errors"
	"strings"
)

type Translation struct {
	Languages
	OriginalText   string
	TranslatedText string
}

type Languages struct {
	SourceLanguage string
	TargetLanguage string
}

var ErrTextLength = errors.New("Error: text length is 0")

func NewTranslation(originalText string, languages *Languages) (*Translation, error) {
	trimedspace := strings.TrimSpace(originalText)
	if len(trimedspace) == 0 {
		return nil, ErrTextLength
	}

	translation := &Translation{
		Languages:    *languages,
		OriginalText: originalText,
	}
	return translation, nil
}
