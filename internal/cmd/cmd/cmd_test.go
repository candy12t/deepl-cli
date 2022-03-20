package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/candy12t/deepl-cli/internal/build"
	"github.com/candy12t/deepl-cli/internal/config"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantOut      string
		wantErr      string
		wantExitCode exitCode
	}{
		{
			name:         "--version",
			args:         strings.Split("deepl-cli --version", " "),
			wantOut:      fmt.Sprintf("deepl-cli version %s\n", build.Version),
			wantErr:      "",
			wantExitCode: exitOK,
		},
		{
			name:         "-v",
			args:         strings.Split("deepl-cli -v", " "),
			wantOut:      fmt.Sprintf("deepl-cli version %s\n", build.Version),
			wantErr:      "",
			wantExitCode: exitOK,
		},
		{
			name:         "unknown command",
			args:         strings.Split("deepl-cli hoge", " "),
			wantOut:      fmt.Sprintf("unknown command %q for \"deepl-cli\"\n", "hoge"),
			wantErr:      "",
			wantExitCode: exitOK,
		},
		{
			name:         "repl",
			args:         strings.Split("deepl-cli repl", " "),
			wantOut:      fmt.Sprintf("Translate text from %s to %s\n>> ", "EN", "JA"),
			wantErr:      "",
			wantExitCode: exitOK,
		},
		{
			name:         "repl",
			args:         strings.Split("deepl-cli repl -s JA -t EN", " "),
			wantOut:      fmt.Sprintf("Translate text from %s to %s\n>> ", "JA", "EN"),
			wantErr:      "",
			wantExitCode: exitOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &config.DeepLCLIConfig{
				Auth: config.Auth{
					AuthKey: "test-auth-key",
				},
				DefaultLanguage: config.DefaultLanguage{
					SourceLanguage: "EN",
					TargetLanguage: "JA",
				},
			}
			inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
			cli := NewCLI(inStream, outStream, errStream, conf)
			code := cli.Run(tt.args)

			if outStream.String() != tt.wantOut {
				t.Errorf("got %q, want %q", outStream.String(), tt.wantOut)
			}

			if errStream.String() != tt.wantErr {
				t.Errorf("got %q, want %q", errStream.String(), tt.wantErr)
			}

			if code != tt.wantExitCode {
				t.Errorf("got %d, want %d", code, tt.wantExitCode)
			}
		})
	}
}
