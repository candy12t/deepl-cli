package cmd

import (
	"github.com/candy12t/deepl-cli/internal/controller"
	"github.com/urfave/cli/v2"
)

func NewCmdConfigure() *cli.Command {
	return &cli.Command{
		Name:   "configure",
		Usage:  "Configure deepl-cli options.",
		Action: controller.ConfigureAction,
	}
}
