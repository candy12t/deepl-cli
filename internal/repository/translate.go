//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/$GOFILE
package repository

import "github.com/candy12t/deepl-cli/internal/model"

type Translator interface {
	TranslateText(text, sourceLang, targetLang string) (*model.TranslateText, error)
}
