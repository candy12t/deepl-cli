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

func Repl(client deepl.Clienter, sourceLang, targetLang string, in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		if !scanner.Scan() {
			return
		}

		text := scanner.Text()
		validedText, err := validation.ValidText(text)
		if err != nil {
			io.WriteString(out, err.Error())
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
