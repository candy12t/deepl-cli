package setup

import (
	"bytes"
	"strings"
	"testing"

	"github.com/candy12t/deepl-cli/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  config.DeepLCLIConfig
	}{
		{
			name:  "setup",
			input: []string{"test-auth-key", "EN", "JA"},
			want: config.DeepLCLIConfig{
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
			assert.Equal(t, tt.want, *got)
		})
	}
}
