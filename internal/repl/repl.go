package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/candy12t/deepl-cli/internal/deepl"
)

const PROMPT = ">> "

func Repl(sourceLang, targetLang string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(PROMPT)
		if scanned := scanner.Scan(); !scanned {
			return
		}

		text := scanner.Text()
		d := deepl.New(text, sourceLang, targetLang)
		resp := d.Data()
		out := os.Stdout
		io.WriteString(out, resp.TranslatedText())
		io.WriteString(out, "\n")
	}
}
