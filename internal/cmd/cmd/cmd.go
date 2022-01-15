package cmd

import (
	"fmt"
	"io"

	"github.com/candy12t/deepl-cli/internal/build"
	"github.com/candy12t/deepl-cli/internal/cmd/repl"
	"github.com/candy12t/deepl-cli/internal/cmd/setup"
	"github.com/candy12t/deepl-cli/internal/config"
	"github.com/candy12t/deepl-cli/internal/deepl"
	"github.com/urfave/cli/v2"
)

type exitCode int

const (
	exitOK exitCode = iota
	exitErr
)

type CLI struct {
	InStream  io.Reader
	OutStream io.Writer
	ErrStream io.Writer

	cfg *config.Config
}

func NewCLI(inStream io.Reader, outStream, errStream io.Writer, cfg *config.Config) *CLI {
	return &CLI{
		InStream:  inStream,
		OutStream: outStream,
		ErrStream: errStream,
		cfg:       cfg,
	}
}

func (c *CLI) Run(args []string) exitCode {
	defaultSourceLang, defaultTargetLang := c.cfg.DefaultLangs()

	var sourceLang, targetLang string

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show deepl-cli version",
	}

	app := &cli.App{
		Name:                 "deepl-cli",
		Usage:                "unofficial DeepL command line tool",
		Version:              build.Version,
		Reader:               c.InStream,
		Writer:               c.OutStream,
		ErrWriter:            c.ErrStream,
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
				Name:  "setup",
				Usage: "Setup for using this cli",
				Action: func(ctx *cli.Context) error {
					err := setup.Setup(ctx.App.Reader, ctx.App.Writer)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "repl",
				Usage: "Translate with REPL",
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
					if err := c.checkAuthKey(); err != nil {
						return err
					}
					fmt.Fprintf(ctx.App.Writer, "Translate text from %s to %s\n", sourceLang, targetLang)
					client, err := deepl.NewClient(c.cfg.GetAuthKey())
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
		fmt.Fprintln(c.ErrStream, err)
		return exitErr
	}

	return exitOK
}

func (c *CLI) checkAuthKey() error {
	if len(c.cfg.GetAuthKey()) == 0 {
		return fmt.Errorf("To setup, please run `deepl-cli setup`.")
	}
	return nil
}
