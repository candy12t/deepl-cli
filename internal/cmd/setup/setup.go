package setup

import (
	"fmt"
	"io"

	"github.com/candy12t/deepl-cli/internal/config"
)

const (
	authKeyQuestion        = "your DeepL auth key"
	sourceLanguageQuestion = "set default `source` language"
	targetLanguageQuestion = "set default `target` language"
)

var languageList = []string{"BG", "CS", "DA", "DE", "EL", "EN", "ES", "ET", "FI", "FR", "HU", "IT", "JA", "LT", "LV", "NL", "PL", "PT", "RO", "RU", "SK", "SL", "SV", "ZH"}

func Setup(inStream io.Reader, outStream io.Writer) error {
	cfg := PromptSetup(inStream, outStream)

	if err := cfg.Write(); err != nil {
		return err
	}

	return nil
}

func PromptSetup(inStream io.Reader, outStream io.Writer) *config.Config {

	authKey := promptForLine(inStream, outStream, authKeyQuestion)
	sourceLanguage := promptForSelect(inStream, outStream, sourceLanguageQuestion, languageList)
	targetLanguage := promptForSelect(inStream, outStream, targetLanguageQuestion, languageList)

	return &config.Config{
		AuthKey: authKey,
		DefaultLang: config.DefaultLang{
			SourceLang: sourceLanguage,
			TargetLang: targetLanguage,
		},
	}
}

func promptForLine(inStream io.Reader, outStream io.Writer, msg string) string {
	var answer string

	fmt.Fprintln(outStream, msg)
	fmt.Fprintf(outStream, ">> ")
	fmt.Fscanf(inStream, "%s\n", &answer)

	return answer
}

func promptForSelect(inStream io.Reader, outStream io.Writer, msg string, selectList []string) string {
	var answer string

	fmt.Fprintln(outStream, msg)
	fmt.Fprintln(outStream, joinSelectList(selectList))
	fmt.Fprintf(outStream, ">> ")
	fmt.Fscanf(inStream, "%s\n", &answer)

	if ok := isValueInStringSlice(answer, selectList); !ok {
		fmt.Fprintln(outStream, "##############################################")
		fmt.Fprintln(outStream, "### no correct input value!!! try again!!! ###")
		fmt.Fprintln(outStream, "##############################################")
		return promptForSelect(inStream, outStream, msg, selectList)
	}

	return answer
}

func joinSelectList(selectList []string) string {
	var str string
	for _, v := range selectList {
		str += fmt.Sprintf("%q ", v)
	}
	return str
}

func isValueInStringSlice(value string, strList []string) bool {
	for _, v := range strList {
		if v == value {
			return true
		}
	}
	return false
}
