package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/candy12t/deepl-cli/internal/build"
	"github.com/candy12t/deepl-cli/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestRun_Success(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantOut      string
		wantExitCode exitCode
	}{
		{
			name:         "--version",
			args:         strings.Split("deepl-cli --version", " "),
			wantOut:      fmt.Sprintf("deepl-cli version %s\n", build.Version),
			wantExitCode: exitOK,
		},
		{
			name:         "-v",
			args:         strings.Split("deepl-cli -v", " "),
			wantOut:      fmt.Sprintf("deepl-cli version %s\n", build.Version),
			wantExitCode: exitOK,
		},
		{
			name:         "unknown command",
			args:         strings.Split("deepl-cli hoge", " "),
			wantOut:      fmt.Sprintf("unknown command %q for \"deepl-cli\"\n", "hoge"),
			wantExitCode: exitOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &config.DeepLCLIConfig{
				Credential: config.Credential{
					DeepLAuthKey: "test-auth-key",
				},
				DefaultLanguage: config.DefaultLanguage{
					Source: "EN",
					Target: "JA",
				},
			}
			reader, writer, errWriter := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
			cli := NewCLI(reader, writer, errWriter)
			code := cli.Run(tt.args, conf)

			assert.Equal(t, tt.wantOut, writer.String())
			assert.Equal(t, tt.wantExitCode, code)
		})
	}
}

func TestRun_Failed(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantOut      string
		wantExitCode exitCode
	}{
		{
			name:         "check auth key",
			args:         strings.Split("deepl-cli repl", " "),
			wantOut:      fmt.Sprintf("To get started with deepl-cli, please run: `deepl-cli configure`\n"),
			wantExitCode: exitErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &config.DeepLCLIConfig{}
			reader, writer, errWriter := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
			cli := NewCLI(reader, writer, errWriter)
			code := cli.Run(tt.args, conf)

			assert.Equal(t, tt.wantOut, errWriter.String())
			assert.Equal(t, tt.wantExitCode, code)
		})
	}
}
