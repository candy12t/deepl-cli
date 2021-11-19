package repl

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/candy12t/deepl-cli/internal/config"
	"github.com/candy12t/deepl-cli/internal/deepl"
)

const PROMPT = ">> "

func Repl(sourceLang, targetLang string) {
	scanner := bufio.NewScanner(os.Stdin)
	client, err := deepl.NewClient(config.BaseURL(), config.AuthKey())
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Printf(PROMPT)
		if scanned := scanner.Scan(); !scanned {
			return
		}

		text := scanner.Text()
		if validedText, err := validText(text); err != nil {
			fmt.Println(err)
		} else {
			ctx := context.Background()
			t, err := client.Translate(ctx, validedText, sourceLang, targetLang)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(t.TranslateText())
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
