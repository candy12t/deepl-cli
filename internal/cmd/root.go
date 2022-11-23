package cmd

import (
	"fmt"
	"io"

	"github.com/candy12t/deepl-cli/internal/build"
	"github.com/candy12t/deepl-cli/internal/config"
	"github.com/urfave/cli/v2"
)

type exitCode int

const (
	exitOK exitCode = iota
	exitErr
)

type CLI struct {
	Reader    io.Reader
	Writer    io.Writer
	ErrWriter io.Writer
}

func NewCLI(reader io.Reader, writer, errWriter io.Writer) *CLI {
	return &CLI{
		Reader:    reader,
		Writer:    writer,
		ErrWriter: errWriter,
	}
}

func (c *CLI) Run(args []string, conf *config.DeepLCLIConfig) exitCode {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show deepl-cli version",
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show help for command",
	}

	app := &cli.App{
		Name:            "deepl-cli",
		Usage:           "unofficial DeepL command line tool.",
		Version:         build.Version,
		Reader:          c.Reader,
		Writer:          c.Writer,
		ErrWriter:       c.ErrWriter,
		HideHelpCommand: true,
		CommandNotFound: func(ctx *cli.Context, command string) {
			fmt.Fprintf(ctx.App.Writer, "unknown command %q for %q\n", command, "deepl-cli")
		},
		Commands: []*cli.Command{
			NewCmdRepl(conf),
			NewCmdConfigure(),
		},
	}

	if err := app.Run(args); err != nil {
		fmt.Fprintln(app.ErrWriter, err)
		return exitErr
	}

	return exitOK
}

func CheckAuthKeyAction(authKey string) cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		if len(authKey) == 0 {
			return fmt.Errorf("To get started with deepl-cli, please run: `deepl-cli configure`")
		}
		return nil
	}
}
