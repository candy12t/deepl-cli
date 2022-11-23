package controller

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/candy12t/deepl-cli/internal/config"
	"github.com/urfave/cli/v2"
)

var languages = []string{"BG", "CS", "DA", "DE", "EL", "EN", "ES", "ET", "FI", "FR", "HU", "IT", "JA", "LT", "LV", "NL", "PL", "PT", "RO", "RU", "SK", "SL", "SV", "ZH"}

var questions = []*survey.Question{
	{
		Name: "deepl_auth_key",
		Prompt: &survey.Password{
			Message: "Paste your auth key",
		},
	},
	{
		Name: "source_language",
		Prompt: &survey.Select{
			Message: "Choose a default source language:",
			Options: languages,
			Default: "JA",
		},
	},
	{
		Name: "target_language",
		Prompt: &survey.Select{
			Message: "Choose a default target language:",
			Options: languages,
			Default: "EN",
		},
	},
}

func ConfigureAction(ctx *cli.Context) error {
	answers := struct {
		DeepLAuthKey   string `survey:"deepl_auth_key"`
		SourceLanguage string `survey:"source_language"`
		TargetLanguage string `survey:"target_language"`
	}{}

	if err := survey.Ask(questions, &answers); err != nil {
		return err
	}

	conf := config.DeepLCLIConfig{
		Credential: config.Credential{
			DeepLAuthKey: answers.DeepLAuthKey,
		},
		DefaultLanguage: config.DefaultLanguage{
			Source: answers.SourceLanguage,
			Target: answers.TargetLanguage,
		},
	}

	if err := conf.WriteDeepLCLIConfig(); err != nil {
		return err
	}
	return nil
}
