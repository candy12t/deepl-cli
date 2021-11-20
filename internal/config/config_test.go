package config

import (
	"path/filepath"
	"testing"

	"github.com/candy12t/deepl-cli/test"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name            string
		inputConfigFile string
		wantConfig      Config
		wantErr         error
	}{
		{
			name:            "parse config yml file, accoutn plan is `free`",
			inputConfigFile: "free.yaml",
			wantConfig: Config{
				Account: Account{
					AuthKey:     "test-auth-key",
					AccountPlan: "free",
				},
				DefaultLang: DefaultLang{
					SourceLang: "EN",
					TargetLang: "JA",
				},
				BaseURL: "https://api-free.deepl.com/v2",
			},
			wantErr: nil,
		},
		{
			name:            "parse config yml file, accoutn plan is `pro`",
			inputConfigFile: "pro.yaml",
			wantConfig: Config{
				Account: Account{
					AuthKey:     "test-auth-key",
					AccountPlan: "pro",
				},
				DefaultLang: DefaultLang{
					SourceLang: "EN",
					TargetLang: "JA",
				},
				BaseURL: "https://api.deepl.com/v2",
			},
			wantErr: nil,
		},
		{
			name:            "can not read config file",
			inputConfigFile: "not_read.yaml",
			wantConfig:      Config{},
			wantErr:         ErrNotReadFile,
		},
		{
			name:            "can not unmarshal config file",
			inputConfigFile: "not_unmarshal.yaml",
			wantConfig:      Config{},
			wantErr:         ErrNotUnmarshal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer teardownConfig(t)

			wantConfig := tt.wantConfig
			configPath := filepath.Join(test.ProjectDirPath(), "test", "testdata", "config", tt.inputConfigFile)

			test.AssertError(t, ParseConfig(configPath), tt.wantErr)
			assertConfig(t, CachedConfig(), wantConfig)
		})
	}
}

func assertConfig(t *testing.T, got, want Config) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func teardownConfig(t *testing.T) {
	t.Cleanup(func() {
		config = Config{}
	})
}
