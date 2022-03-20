package setup

import (
	"bytes"
	"strings"
	"testing"

	"github.com/candy12t/deepl-cli/internal/config"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  *config.DeepLCLIConfig
	}{
		{
			name:  "setup",
			input: []string{"test-auth-key", "fred", "EN", "JA"},
			want: &config.DeepLCLIConfig{
				Auth: config.Auth{
					AuthKey: "test-auth-key",
				},
				DefaultLanguage: config.DefaultLanguage{
					SourceLanguage: "EN",
					TargetLanguage: "JA",
				},
			},
		},
		{
			name:  "setup validate",
			input: []string{"test-auth-key", "hoge", "free", "EN", "JA"},
			want: &config.DeepLCLIConfig{
				Auth: config.Auth{
					AuthKey: "test-auth-key",
				},
				DefaultLanguage: config.DefaultLanguage{
					SourceLanguage: "EN",
					TargetLanguage: "JA",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := bytes.NewBufferString(strings.Join(tt.input, "\n"))
			out := new(bytes.Buffer)
			got := PromptSetup(in, out)
			if *got != *tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
