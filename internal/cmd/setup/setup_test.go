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
		want  *config.Config
	}{
		{
			name:  "setup",
			input: []string{"test-auth-key", "free", "EN", "JA"},
			want: &config.Config{
				Account: config.Account{
					AuthKey:     "test-auth-key",
					AccountPlan: "free",
				},
				DefaultLang: config.DefaultLang{
					SourceLang: "EN",
					TargetLang: "JA",
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
