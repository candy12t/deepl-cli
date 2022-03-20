package main

import (
	"os"

	"github.com/candy12t/deepl-cli/internal/cmd/cmd"
	"github.com/candy12t/deepl-cli/internal/config"
)

func main() {
	conf := config.NewDeepLCLIConfig()
	cli := cmd.NewCLI(os.Stdin, os.Stdout, os.Stderr, conf)
	code := cli.Run(os.Args)
	os.Exit(int(code))
}
