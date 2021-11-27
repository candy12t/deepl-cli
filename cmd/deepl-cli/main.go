package main

import (
	"fmt"
	"io"
	"os"

	"github.com/candy12t/deepl-cli/internal/config"
	"github.com/candy12t/deepl-cli/internal/deepl"
	"github.com/candy12t/deepl-cli/internal/repl"
	"github.com/urfave/cli/v2"
)

var version = "DEV"

type CLI struct {
	outStream, errStream io.Writer
	config               *config.Config
}

type exitCode int

const (
	exitOK exitCode = iota
	exitErr
)

func main() {
	cfg := config.LoadConfig()
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr, config: cfg}
	code := cli.run(os.Args)
	os.Exit(int(code))
}

func (c *CLI) run(args []string) exitCode {

	defaultSourceLang, defaultTargetLang := c.config.DefaultLangs()
	var sourceLang, targetLang string

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show deepl-cli version",
	}

	app := &cli.App{
		Name:                 "deepl-cli",
		Usage:                "unofficial DeepL command line tool",
		Version:              version,
		Writer:               c.outStream,
		ErrWriter:            c.errStream,
		EnableBashCompletion: true,
		HideHelp:             true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "help",
				Aliases: []string{"h"},
				Usage:   "Show help for command",
			},
		},
		CommandNotFound: func(ctx *cli.Context, command string) {
			fmt.Fprintf(ctx.App.Writer, "unknown command %q for %q\n", command, "deepl-cli")
		},
		Commands: []*cli.Command{
			{
				Name:  "repl",
				Usage: "repl",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "source",
						Aliases:     []string{"s"},
						Value:       defaultSourceLang,
						DefaultText: defaultSourceLang,
						Destination: &sourceLang,
					},
					&cli.StringFlag{
						Name:        "target",
						Aliases:     []string{"t"},
						Value:       defaultTargetLang,
						DefaultText: defaultTargetLang,
						Destination: &targetLang,
					},
				},
				Action: func(ctx *cli.Context) error {
					fmt.Fprintf(ctx.App.Writer, "Translate text from %s to %s\n", sourceLang, targetLang)
					client, err := deepl.NewClient(c.config.BaseURL(), c.config.AuthKey())
					if err != nil {
						return err
					}
					repl.Repl(client, sourceLang, targetLang, ctx.App.Reader, ctx.App.Writer)
					return nil
				},
			},
		},
	}

	if err := app.Run(args); err != nil {
		fmt.Fprintln(c.errStream, err)
		return exitErr
	}

	return exitOK
}
