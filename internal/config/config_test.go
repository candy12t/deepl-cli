package config

import (
	"path/filepath"
	"testing"
)

func TestParseConfig(t *testing.T) {
	t.Run("parse config yml file", func(t *testing.T) {
		wantConfig := Config{
			Account: Account{
				AuthKey:     "test-auth-key",
				AccountPlan: "free",
			},
			DefaultLang: DefaultLang{
				SourceLang: "EN",
				TargetLang: "JA",
			},
		}

		configPath := filepath.Join("testdata", "config.yaml")
		err := ParseConfig(configPath)
		gotConfig := CachedConfig()

		assertNoError(t, err)
		assertConfig(t, gotConfig, wantConfig)

		cleanCachedConfig(t)
	})

	t.Run("can not read config file", func(t *testing.T) {
		wantConfig := Config{}

		configPath := filepath.Join("testdata", "not_read.yaml")
		err := ParseConfig(configPath)
		gotConfig := CachedConfig()

		assertReadFileError(t, err, ErrNotReadFile)
		assertConfig(t, gotConfig, wantConfig)

		cleanCachedConfig(t)
	})

	t.Run("can not unmarshal config file", func(t *testing.T) {
		wantConfig := Config{}

		configPath := filepath.Join("testdata", "not_unmarshal.yaml")
		err := ParseConfig(configPath)
		gotConfig := CachedConfig()

		assertUnmarshalError(t, err, ErrNotUnmarshal)
		assertConfig(t, gotConfig, wantConfig)

		cleanCachedConfig(t)
	})
}

func assertConfig(t *testing.T, got, want Config) {
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("got error but didn't want one")
	}
}

func assertReadFileError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
	if got == nil || got == ErrNotUnmarshal {
		t.Fatal("expected to get error")
	}
}

func assertUnmarshalError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
	if got == nil || got == ErrNotReadFile {
		t.Fatal("expected to get error")
	}
}

func cleanCachedConfig(t *testing.T) {
	t.Cleanup(func() {
		config = Config{}
	})
}
