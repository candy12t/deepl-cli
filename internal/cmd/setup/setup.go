package setup

import (
	"fmt"
	"io"

	"github.com/candy12t/deepl-cli/internal/config"
)

var languageList = []string{"BG", "CS", "DA", "DE", "EL", "EN", "ES", "ET", "FI", "FR", "HU", "IT", "JA", "LT", "LV", "NL", "PL", "PT", "RO", "RU", "SK", "SL", "SV", "ZH"}

func Setup(inStream io.Reader, outStream io.Writer) error {
	cfg := PromptSetup(inStream, outStream)

	if err := cfg.Write(); err != nil {
		return err
	}

	return nil
}

// TODO: validation
func PromptSetup(inStream io.Reader, outStream io.Writer) *config.Config {
	var authKey string
	fmt.Fprintf(outStream, "Paste your DeepL auth key >> ")
	fmt.Fscanf(inStream, "%s\n", &authKey)

	var accountPlan string
	fmt.Fprintf(outStream, "your DeepL account plan \"free\" or \"pro\" >> ")
	fmt.Fscanf(inStream, "%s\n", &accountPlan)

	var sourceLanguage string
	fmt.Fprintf(outStream, "set default source language >> ")
	fmt.Fscanf(inStream, "%s\n", &sourceLanguage)

	var targetLanguage string
	fmt.Fprintf(outStream, "set default target language >> ")
	fmt.Fscanf(inStream, "%s\n", &targetLanguage)

	return &config.Config{
		Account: config.Account{
			AuthKey:     string(authKey),
			AccountPlan: accountPlan,
		},
		DefaultLang: config.DefaultLang{
			SourceLang: sourceLanguage,
			TargetLang: targetLanguage,
		},
	}
}
