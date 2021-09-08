package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Account struct {
		AuthKey     string `yaml:"auth_key"`
		AccountPlan string `yaml:"account_plan"`
	} `yaml:"account"`
	DefaultTranslateLang struct {
		SourceLang string `yaml:"source_lang"`
		TargetLang string `yaml:"target_lang"`
	} `yaml:"default_translate_lang"`
}

var config Config

func ConfigDir() string {
	d, _ := os.UserHomeDir()
	return filepath.Join(d, ".config", "deepl-cli")
}

func ConfigFile() string {
	return filepath.Join(ConfigDir(), "config.yaml")
}

func ParseConfig() error {
	data, err := ioutil.ReadFile(ConfigFile())
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}
	return nil
}

func DefaultLang() (string, string) {
	return setDefaultLang(config.DefaultTranslateLang.SourceLang, "EN"), setDefaultLang(config.DefaultTranslateLang.TargetLang, "JA")
}

func AccountPlan() string { return config.Account.AccountPlan }
func AuthKey() string     { return config.Account.AuthKey }

func setDefaultLang(oldStr, newStr string) string {
	if oldStr == "" {
		return newStr
	}
	return oldStr
}
