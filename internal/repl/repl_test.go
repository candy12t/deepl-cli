package repl

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/candy12t/deepl-cli/internal/deepl"
)

func TestRepl(t *testing.T) {
	tests := []struct {
		name       string
		input      *bytes.Buffer
		targetLang string
		want       string
	}{
		{
			name:       "success repl",
			input:      bytes.NewBufferString("hello"),
			targetLang: "JA",
			want:       fmt.Sprintf("%sこんにちわ\n%s", PROMPT, PROMPT),
		},
		{
			name:       "failed repl because input text length is 0",
			input:      bytes.NewBufferString("\n"),
			targetLang: "JA",
			want:       fmt.Sprintf("%sError: text length is 0\n%s", PROMPT, PROMPT),
		},
		{
			name:       "failed repl because not specified target language",
			input:      bytes.NewBufferString("hello"),
			targetLang: "",
			want:       fmt.Sprintf(`%sHTTP 400: "Value for 'target_lang' not supported." (https://api-free.deepl.com/v2/translate)`, PROMPT),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := new(bytes.Buffer)

			client, _ := deepl.NewMockClient("https://api.deepl.com/v2", "test-auth-key")
			Repl(client, "EN", tt.targetLang, tt.input, out)

			got := out.String()
			if got != tt.want {
				t.Fatalf("repl output %q, want %q", got, tt.want)
			}
		})
	}
}
