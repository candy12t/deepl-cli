package config

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	input := `
auth:
  auth_key: "test-auth-key"
default_language:
  source_language: "JA"
  target_language: "EN"
`

	conf, err := parseDeepLCLIConfigData([]byte(input))
	if assert.NoError(t, err) {
		assert.Equal(t, "test-auth-key", conf.Auth.AuthKey)
		assert.Equal(t, "JA", conf.DefaultLanguage.SourceLanguage)
		assert.Equal(t, "EN", conf.DefaultLanguage.TargetLanguage)
	}
}

func TestWriteConfig(t *testing.T) {
	conf := DeepLCLIConfig{
		Auth: Auth{
			AuthKey: "test-auth-key",
		},
		DefaultLanguage: DefaultLanguage{
			SourceLanguage: "JA",
			TargetLanguage: "EN",
		},
	}
	filename := filepath.Join(t.TempDir(), "config.yaml")
	err := conf.writeDeepLCLIConfig(filename)
	assert.NoError(t, err)
	data, _ := parseDeepLCLIConfigFile(filename)
	fmt.Println(data)
}
