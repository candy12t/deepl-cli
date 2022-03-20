package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVaildText(t *testing.T) {
	tests := []struct {
		input    string
		wantText string
	}{
		{"\thello\t", "hello"},
		{" hello world ", "hello world"},
		{"\t hoge fuga \t", "hoge fuga"},
	}

	for _, tt := range tests {
		text, err := ValidText(tt.input)
		if assert.NoError(t, err) {
			assert.Equal(t, tt.wantText, text)
		}
	}
}

func TestInVaildText(t *testing.T) {
	tests := []struct {
		input   string
		wantErr error
	}{
		{"\t\n", ErrTextLength},
		{"    ", ErrTextLength},
	}

	for _, tt := range tests {
		_, err := ValidText(tt.input)
		assert.EqualError(t, err, tt.wantErr.Error())
	}
}
