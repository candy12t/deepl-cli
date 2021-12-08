package config

import (
	"path/filepath"
	"testing"

	"github.com/candy12t/deepl-cli/test"
	"gopkg.in/yaml.v2"
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

func WriteConfigTest(t *testing.T) {
	tests := []struct {
		name        string
		inputConfig Config
		args        struct {
			filename string
			data     func(Config) []byte
		}
		wantErr error
	}{
		{
			name: "write config file",
			inputConfig: Config{
				Account:     Account{AuthKey: "test-auth-key", AccountPlan: "free"},
				DefaultLang: DefaultLang{SourceLang: "EN", TargetLang: "JA"},
			},
			args: struct {
				filename string
				data     func(Config) []byte
			}{
				filename: filepath.Join(t.TempDir(), "config.yaml"),
				data: func(config Config) []byte {
					data, _ := yaml.Marshal(config)
					return data
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteConfigFile(tt.args.filename, tt.args.data(tt.inputConfig))
			test.AssertError(t, err, tt.wantErr)
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
