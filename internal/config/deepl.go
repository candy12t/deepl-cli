package config

import (
	"io/ioutil"
	"log"

	_ "github.com/candy12t/deepl-cli/statik"
	"github.com/rakyll/statik/fs"
	"gopkg.in/yaml.v2"
)

type DeeplConfig struct {
	AuthKey  string `yaml:"auth_key"`
	Endpoint string `yaml:"endpoint"`
}

var deeplConfigData DeeplConfig

func InitDeeplConfig() {
	data := file2byte("/deepl.yaml")
	if err := yaml.Unmarshal(data, &deeplConfigData); err != nil {
		log.Fatal(err)
	}
}

func DeeplConfigData() DeeplConfig {
	return deeplConfigData
}

func file2byte(filename string) []byte {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	// Access individual files by their paths.
	r, err := statikFS.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return contents
}
