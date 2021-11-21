package repl

import (
	"bufio"
	"context"
	"io"

	"github.com/candy12t/deepl-cli/internal/deepl"
	"github.com/candy12t/deepl-cli/internal/validation"
)

const PROMPT = ">> "

func Repl(client deepl.Clienter, sourceLang, targetLang string, in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		io.WriteString(out, PROMPT)
		if !scanner.Scan() {
			return
		}

		text := scanner.Text()
		validedText, err := validation.ValidText(text)
		if err != nil {
			io.WriteString(out, err.Error()+"\n")
			continue
		}

		ctx := context.Background()
		t, err := client.Translate(ctx, validedText, sourceLang, targetLang)
		if err != nil {
			io.WriteString(out, err.Error())
			return
		}
		io.WriteString(out, t.TranslateText()+"\n")
	}
}
