package controller

import (
	"errors"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/candy12t/deepl-cli/internal/repository"
	"github.com/urfave/cli/v2"
)

var ErrTextLength = errors.New("Error: input text length is 0")

func ReplAction(client repository.Translator) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		sourceLanguage := ctx.String("source")
		targetLanguage := ctx.String("target")
		message := fmt.Sprintf("(%s->%s) >>", sourceLanguage, targetLanguage)

		for {
			var text string
			prompt := &survey.Input{
				Message: message,
			}
			if err := survey.AskOne(prompt, &text, survey.WithIcons(func(icons *survey.IconSet) {
				icons.Question.Text = "deepl-cli"
				icons.Question.Format = "default+hb"
			})); err != nil {
				return err
			}

			trimedSpaceText := strings.TrimSpace(text)
			if len(trimedSpaceText) == 0 {
				fmt.Fprintln(ctx.App.ErrWriter, ErrTextLength.Error())
				continue
			}

			tr, err := client.TranslateText(trimedSpaceText, sourceLanguage, targetLanguage)
			if err != nil {
				return err
			}

			fmt.Fprintln(ctx.App.Writer, tr.TranslateText)
		}
	}
}
