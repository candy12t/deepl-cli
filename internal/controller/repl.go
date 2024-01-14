package controller

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/candy12t/deepl-cli/internal/repository"
	"github.com/peterh/liner"
	"github.com/urfave/cli/v2"
)

var ErrTextLength = errors.New("Error: input text length is 0")
var historyFilePath = filepath.Join(os.TempDir(), ".deepl_cli_history")

func ReplAction(client repository.Translator) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		l := liner.NewLiner()
		defer l.Close()

		l.SetCtrlCAborts(true)

		f, err := os.Open(historyFilePath)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			if _, err := l.ReadHistory(f); err != nil {
				return err
			}
		}
		defer f.Close()

		source := ctx.String("source")
		target := ctx.String("target")
		prompt := fmt.Sprintf("deepl-cli (%s->%s)> ", source, target)
		var promptError error
		for {
			text, err := l.Prompt(prompt)
			if err != nil {
				promptError = err
				break
			}
			if text == "exit" {
				break
			}
			t, err := client.TranslateText(text, source, target)
			if err != nil {
				promptError = err
			}
			fmt.Fprintln(ctx.App.Writer, t.TranslateText)
			l.AppendHistory(text)
		}

		ff, err := os.Create(historyFilePath)
		if err != nil {
			return err
		}
		defer ff.Close()
		if _, err := l.WriteHistory(ff); err != nil {
			return err
		}

		return promptError
	}
}
