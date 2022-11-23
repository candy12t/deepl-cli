package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	input := `
credential:
  deepl_auth_key: "test-auth-key"
default_language:
  source: "JA"
  target: "EN"
`

	conf, err := parseDeepLCLIConfigData([]byte(input))
	if assert.NoError(t, err) {
		assert.Equal(t, "test-auth-key", conf.Credential.DeepLAuthKey)
		assert.Equal(t, "JA", conf.DefaultLanguage.Source)
		assert.Equal(t, "EN", conf.DefaultLanguage.Target)
	}
}

func TestWriteConfig(t *testing.T) {
	conf := DeepLCLIConfig{
		Credential: Credential{
			DeepLAuthKey: "test-auth-key",
		},
		DefaultLanguage: DefaultLanguage{
			Source: "JA",
			Target: "EN",
		},
	}
	filename := filepath.Join(t.TempDir(), "config.yaml")
	err := conf.writeDeepLCLIConfig(filename)
	assert.NoError(t, err)
	parseDeepLCLIConfigFile(filename)
}
