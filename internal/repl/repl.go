package repl

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/candy12t/deepl-cli/internal/config"
	"github.com/candy12t/deepl-cli/internal/deepl"
	"github.com/candy12t/deepl-cli/internal/validation"
)

const PROMPT = ">> "

func Repl(sourceLang, targetLang string) {
	scanner := bufio.NewScanner(os.Stdin)

	client, err := deepl.NewClient(config.BaseURL(), config.AuthKey())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for {
		fmt.Printf(PROMPT)
		if scanned := scanner.Scan(); !scanned {
			return
		}

		text := scanner.Text()
		validedText, err := validation.ValidText(text)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		ctx := context.Background()
		t, err := client.Translate(ctx, validedText, sourceLang, targetLang)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println(t.TranslateText())
	}
}
