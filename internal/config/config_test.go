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
		wantConfig      *Config
		wantBaseURL     string
		wantErr         error
	}{
		{
			name:            "parse config yml file, accoutn plan is `free`",
			inputConfigFile: "free.yaml",
			wantConfig: &Config{
				Account: Account{
					AuthKey:     "test-auth-key",
					AccountPlan: "free",
				},
				DefaultLang: DefaultLang{
					SourceLang: "EN",
					TargetLang: "JA",
				},
			},
			wantBaseURL: "https://api-free.deepl.com/v2",
			wantErr:     nil,
		},
		{
			name:            "parse config yml file, accoutn plan is `pro`",
			inputConfigFile: "pro.yaml",
			wantConfig: &Config{
				Account: Account{
					AuthKey:     "test-auth-key",
					AccountPlan: "pro",
				},
				DefaultLang: DefaultLang{
					SourceLang: "EN",
					TargetLang: "JA",
				},
			},
			wantBaseURL: "https://api.deepl.com/v2",
			wantErr:     nil,
		},
		{
			name:            "can not read config file",
			inputConfigFile: "not_read.yaml",
			wantConfig:      nil,
			wantErr:         ErrNotReadFile,
		},
		{
			name:            "can not unmarshal config file",
			inputConfigFile: "not_unmarshal.yaml",
			wantConfig:      nil,
			wantErr:         ErrNotUnmarshal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configPath := filepath.Join("testdata", tt.inputConfigFile)
			gotConfig, gotErr := ParseConfig(configPath)

			test.AssertError(t, gotErr, tt.wantErr)
			assertConfig(t, gotConfig, tt.wantConfig)
		})
	}
}

func assertConfig(t *testing.T, got, want *Config) {
	t.Helper()
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expexted got Config")
	}

	assertAuthKey(t, got.AuthKey(), want.AuthKey())
	assertDefaltLangs(t, got, want)
	assertBaseURL(t, got.BaseURL(), want.BaseURL())

}

func assertAuthKey(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got AuthKey is %q, want %q", got, want)
	}
}

func assertDefaltLangs(t *testing.T, got, want *Config) {
	gotSourceLang, gotTargetLang := got.DefaultLangs()
	wantSourceLang, wantTargetLang := want.DefaultLangs()

	if gotSourceLang != wantSourceLang {
		t.Errorf("got SourceLang is %q, want %q", gotSourceLang, wantSourceLang)
	}

	if gotTargetLang != wantTargetLang {
		t.Errorf("got TargetLang is %q, want %q", gotTargetLang, wantTargetLang)
	}
}

func assertBaseURL(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got BaseURL is %q, want %q", got, want)
	}
}
