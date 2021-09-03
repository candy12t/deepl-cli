package repl

import (
	"fmt"

	"github.com/candy12t/deepl-cli/internal/deepl"
)

const PROMPT = ">> "

func Repl(sourceLang, targetLang string) {
	for {
		fmt.Printf(PROMPT)
		var text string
		fmt.Scanf("%s", &text)

		d := deepl.New(text, sourceLang, targetLang)
		resp := d.Data()
		fmt.Println(resp.TranslatedText())
	}
}
