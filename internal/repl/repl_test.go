package repl

import (
	"testing"
)

func TestVaildText(t *testing.T) {
	tests := []struct {
		input        string
		expectedText string
	}{
		{"\thello\t", "hello"},
		{" hello world ", "hello world"},
		{"\t hoge fuga \t", "hoge fuga"},
		{"\t\n", ""},
		{"    ", ""},
	}

	for _, tt := range tests {
		text, _ := validText(tt.input)

		if tt.expectedText != text {
			t.Fatalf("validText() got=%q, expectedText=%q", text, tt.expectedText)
		}
	}
}
