package validation

import (
	"fmt"
	"strings"
)

var ErrTextLength = fmt.Errorf("Error: text length is 0")

func ValidText(text string) (string, error) {
	if err := validTextLength(trimSpace(text)); err != nil {
		return "", err
	}
	return trimSpace(text), nil
}

func trimSpace(text string) string {
	return strings.TrimSpace(text)
}

func validTextLength(text string) error {
	if len(text) == 0 {
		return ErrTextLength
	}
	return nil
}
