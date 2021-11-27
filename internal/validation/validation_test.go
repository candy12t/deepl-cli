package validation

import (
	"testing"

	"github.com/candy12t/deepl-cli/test"
)

func TestVaildText(t *testing.T) {
	tests := []struct {
		input    string
		wantText string
		wantErr  error
	}{
		{"\thello\t", "hello", nil},
		{" hello world ", "hello world", nil},
		{"\t hoge fuga \t", "hoge fuga", nil},
		{"\t\n", "", ErrTextLength},
		{"    ", "", ErrTextLength},
	}

	for _, tt := range tests {
		text, err := ValidText(tt.input)
		test.AssertError(t, err, tt.wantErr)

		if tt.wantText != text {
			t.Fatalf("validText(%q) returned %q, want %q", tt.input, text, tt.wantText)
		}
	}
}
