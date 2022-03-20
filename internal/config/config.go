package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type DeepLCLIConfig struct {
	Auth            Auth            `yaml:"auth"`
	DefaultLanguage DefaultLanguage `yaml:"default_language"`
}

type Auth struct {
	AuthKey string `yaml:"auth_key"`
}

type DefaultLanguage struct {
	SourceLanguage string `yaml:"source_language"`
	TargetLanguage string `yaml:"target_language"`
}

func configFile() string {
	d, _ := os.UserHomeDir()
	return filepath.Join(d, ".config", "deepl-cli", "config.yaml")
}

func NewDeepLCLIConfig() *DeepLCLIConfig {
	conf, err := parseDeepLCLIConfigFile(configFile())
	if err != nil {
		return &DeepLCLIConfig{
			DefaultLanguage: setDefaultLanguage(),
		}
	}
	if len(conf.DefaultLanguage.SourceLanguage) == 0 || len(conf.DefaultLanguage.TargetLanguage) == 0 {
		conf.DefaultLanguage = setDefaultLanguage()
	}
	return conf
}

func parseDeepLCLIConfigFile(filename string) (*DeepLCLIConfig, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return parseDeepLCLIConfigData(data)
}

func parseDeepLCLIConfigData(data []byte) (*DeepLCLIConfig, error) {
	conf := new(DeepLCLIConfig)
	err := yaml.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func setDefaultLanguage() DefaultLanguage {
	return DefaultLanguage{
		SourceLanguage: "JA",
		TargetLanguage: "EN",
	}
}

func (c *DeepLCLIConfig) WriteDeepLCLIConfig() error {
	return c.writeDeepLCLIConfig(configFile())
}

func (c *DeepLCLIConfig) writeDeepLCLIConfig(filename string) error {
	conf, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	if err := writeConfigFile(filename, conf); err != nil {
		return err
	}

	return nil
}

func writeConfigFile(filename string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(filename), 0644)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)

	return err
}
