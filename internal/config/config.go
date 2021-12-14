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

func LoadConfig() *Config {
	cfg, err := ParseConfig(ConfigFile())
	if err != nil {
		return &Config{
			DefaultLang: setDefaultLang(),
		}
	}

	if sourceLang, targetLang := cfg.DefaultLangs(); sourceLang == "" || targetLang == "" {
		cfg.DefaultLang = setDefaultLang()
	}

	return cfg
}

func (c *Config) Write() error {
	cfgBytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	if err := WriteConfigFile(ConfigFile(), cfgBytes); err != nil {
		return err
	}

	return nil
}

func WriteConfigFile(filename string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(filename), 0644)
	if err != nil {
		return err
	}

	cfg, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer cfg.Close()

	_, err = cfg.Write(data)

	return err
}

func (c *Config) DefaultLangs() (string, string) {
	return c.DefaultLang.SourceLang, c.DefaultLang.TargetLang
}

func (c *Config) AuthKey() string {
	return c.Account.AuthKey
}

func setDefaultLang() DefaultLang {
	return DefaultLang{
		SourceLang: "EN",
		TargetLang: "JA",
	}
}
