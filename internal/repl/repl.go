package repl

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/candy12t/deepl-cli/internal/deepl"
	"github.com/candy12t/deepl-cli/internal/validation"
)

const PROMPT = ">> "

func Repl(client deepl.APIClient, sourceLang, targetLang string, in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		if !scanner.Scan() {
			return
		}

		text := scanner.Text()
		validedText, err := validation.ValidText(text)
		if err != nil {
			fmt.Fprintln(out, err)
			continue
		}

		ctx := context.Background()
		t, err := client.Translate(ctx, validedText, sourceLang, targetLang)
		if err != nil {
			fmt.Fprintln(out, err)
			return
		}
		fmt.Fprintln(out, t.TranslateText())
	}
}
