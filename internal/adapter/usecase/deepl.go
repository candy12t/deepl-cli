//go:generate mockgen -source=$GOFILE -destination=./mock_$GOPACKAGE/$GOFILE
package usecase

type DeepL interface {
	Translate(original, source, target string) (string, error)
}
