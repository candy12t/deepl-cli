//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/$GOFILE
package webapi

import "github.com/candy12t/deepl-cli/internal/entity"

type Translater interface {
	Translate(*entity.Translation) (*entity.Translation, error)
}
