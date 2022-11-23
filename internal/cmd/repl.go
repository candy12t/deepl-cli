package cmd

import (
	"github.com/candy12t/deepl-cli/internal/config"
	"github.com/candy12t/deepl-cli/internal/controller"
	"github.com/candy12t/deepl-cli/internal/deepl"
	"github.com/urfave/cli/v2"
)

func NewCmdRepl(conf *config.DeepLCLIConfig) *cli.Command {
	client := deepl.NewClient(conf.Credential.DeepLAuthKey)

	return &cli.Command{
		Name:  "repl",
		Usage: "Translate with REPL.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "source",
				Aliases: []string{"s"},
				Value:   conf.DefaultLanguage.Source,
			},
			&cli.StringFlag{
				Name:    "target",
				Aliases: []string{"t"},
				Value:   conf.DefaultLanguage.Target,
			},
		},
		Before: CheckAuthKeyAction(conf.Credential.DeepLAuthKey),
		Action: controller.ReplAction(client),
	}
}
