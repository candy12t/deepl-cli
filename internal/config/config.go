package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	ErrNotReadFile  = ConfigErr("can not read config file")
	ErrNotUnmarshal = ConfigErr("can not unmarshal config data")
)

type ConfigErr string

func (c ConfigErr) Error() string {
	return string(c)
}

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

var config Config

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
	return nil
}

func CachedConfig() Config {
	return config
}

func DefaultLangs() (string, string) {
	return config.DefaultLang.SourceLang, config.DefaultLang.TargetLang
}

func AccountPlan() string { return CachedConfig().Account.AccountPlan }
func AuthKey() string     { return CachedConfig().Account.AuthKey }
