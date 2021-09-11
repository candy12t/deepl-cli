package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/candy12t/deepl-cli/api"
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
		if validedText, err := validText(text); err != nil {
			fmt.Println(err)
		} else {
			tr, err := api.Translate(validedText, sourceLang, targetLang)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(tr.TranslatedText())
		}
	}
}

func validText(text string) (string, error) {
	validedText := strings.TrimSpace(text)
	if len(validedText) == 0 {
		return "", fmt.Errorf("Error: text length is 0")
	}
	return validedText, nil
}
