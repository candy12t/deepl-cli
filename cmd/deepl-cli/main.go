package main

import (
	"os"

	"github.com/candy12t/deepl-cli/internal/cmd"
	"github.com/candy12t/deepl-cli/internal/config"
)

func main() {
	conf := config.NewDeepLCLIConfig()
	cli := cmd.NewCLI(os.Stdin, os.Stdout, os.Stderr)
	code := cli.Run(os.Args, conf)
	os.Exit(int(code))
}
