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
}

type Account struct {
	AuthKey     string `yaml:"auth_key"`
	AccountPlan string `yaml:"account_plan"`
}

type DefaultLang struct {
	SourceLang string `yaml:"source_lang"`
	TargetLang string `yaml:"target_lang"`
}

func ConfigFile() string {
	d, _ := os.UserHomeDir()
	return filepath.Join(d, ".config", "deepl-cli", "config.yaml")
}

func ParseConfig(filepath string) (*Config, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, ErrNotReadFile
	}

	config := new(Config)
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, ErrNotUnmarshal
	}

	return config, nil
}

func (c *Config) DefaultLangs() (string, string) {
	return c.DefaultLang.SourceLang, c.DefaultLang.TargetLang
}
func (c *Config) AuthKey() string {
	return c.Account.AuthKey
}
