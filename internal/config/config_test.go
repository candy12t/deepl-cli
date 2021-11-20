package config

import (
	"path/filepath"
	"testing"

	"github.com/candy12t/deepl-cli/test"
)

func TestParseConfig(t *testing.T) {
	t.Run("parse config yml file, accoutn plan is `free`", func(t *testing.T) {
		wantConfig := Config{
			Account: Account{
				AuthKey:     "test-auth-key",
				AccountPlan: "free",
			},
			DefaultLang: DefaultLang{
				SourceLang: "EN",
				TargetLang: "JA",
			},
			BaseURL: "https://api-free.deepl.com/v2",
		}

		configPath := filepath.Join(test.ProjectDirPath(), "test", "testdata", "config", "free.yaml")
		err := ParseConfig(configPath)
		gotConfig := CachedConfig()

		test.AssertError(t, err, nil)
		assertConfig(t, gotConfig, wantConfig)

		cleanCachedConfig(t)
	})

	t.Run("parse config yml file, accoutn plan is `pro`", func(t *testing.T) {
		wantConfig := Config{
			Account: Account{
				AuthKey:     "test-auth-key",
				AccountPlan: "pro",
			},
			DefaultLang: DefaultLang{
				SourceLang: "EN",
				TargetLang: "JA",
			},
			BaseURL: "https://api.deepl.com/v2",
		}

		configPath := filepath.Join(test.ProjectDirPath(), "test", "testdata", "config", "pro.yaml")
		err := ParseConfig(configPath)
		gotConfig := CachedConfig()

		test.AssertError(t, err, nil)
		assertConfig(t, gotConfig, wantConfig)

		cleanCachedConfig(t)
	})

	t.Run("can not read config file", func(t *testing.T) {
		wantConfig := Config{}

		configPath := filepath.Join(test.ProjectDirPath(), "test", "testdata", "config", "not_read.yaml")
		err := ParseConfig(configPath)
		gotConfig := CachedConfig()

		test.AssertError(t, err, ErrNotReadFile)
		assertConfig(t, gotConfig, wantConfig)

		cleanCachedConfig(t)
	})

	t.Run("can not unmarshal config file", func(t *testing.T) {
		wantConfig := Config{}

		configPath := filepath.Join(test.ProjectDirPath(), "test", "testdata", "config", "not_unmarshal.yaml")
		err := ParseConfig(configPath)
		gotConfig := CachedConfig()

		test.AssertError(t, err, ErrNotUnmarshal)
		assertConfig(t, gotConfig, wantConfig)

		cleanCachedConfig(t)
	})
}

func assertConfig(t *testing.T, got, want Config) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func cleanCachedConfig(t *testing.T) {
	t.Cleanup(func() {
		config = Config{}
	})
}
