package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/candy12t/deepl-cli/api"
)

const PROMPT = ">> "

func Repl(sourceLang, targetLang string) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(PROMPT)
		if scanned := scanner.Scan(); !scanned {
			return fmt.Errorf("can not scan")
		}

		text := scanner.Text()
		tr, err := api.Translate(text, sourceLang, targetLang)
		if err != nil {
			return err
		}
		out := os.Stdout
		io.WriteString(out, tr.TranslatedText())
		io.WriteString(out, "\n")
	}
}
