package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Account     Account     `yaml:"account"`
	DefaultLang DefaultLang `yaml:"default_lang"`
	BaseURL     string
}

type Account struct {
	AuthKey     string `yaml:"auth_key"`
	AccountPlan string `yaml:"account_plan"`
}

type DefaultLang struct {
	SourceLang string `yaml:"source_lang"`
	TargetLang string `yaml:"target_lang"`
}

var config Config

func CachedConfig() Config {
	return config
}

func ConfigFile() string {
	d, _ := os.UserHomeDir()
	return filepath.Join(d, ".config", "deepl-cli", "config.yaml")
}

func ParseConfig(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return ErrNotReadFile
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return ErrNotUnmarshal
	}
	config.BaseURL = baseURL()

	return nil
}

func DefaultLangs() (string, string) {
	return config.DefaultLang.SourceLang, config.DefaultLang.TargetLang
}
func AuthKey() string { return CachedConfig().Account.AuthKey }
func BaseURL() string { return CachedConfig().BaseURL }
