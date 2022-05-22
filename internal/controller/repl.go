package controller

import (
	"bufio"
	"fmt"
	"io"

	"github.com/candy12t/deepl-cli/internal/entity"
	"github.com/candy12t/deepl-cli/internal/usecase"
)

const PROMPT = ">> "

type Repl struct {
	uc *usecase.Translation
	*entity.Languages
	inStream  io.Reader
	outStream io.Writer
}

func NewRepl(uc *usecase.Translation, sourceLanguage, targetLanguage string, inStream io.Reader, outStream io.Writer) *Repl {
	return &Repl{
		uc: uc,
		Languages: &entity.Languages{
			SourceLanguage: sourceLanguage,
			TargetLanguage: targetLanguage,
		},
		inStream:  inStream,
		outStream: outStream,
	}
}

func (r *Repl) Apply() {
	fmt.Fprintf(r.outStream, "Translate text from %s to %s\n", r.SourceLanguage, r.TargetLanguage)
	scanner := bufio.NewScanner(r.inStream)

	for {
		fmt.Fprint(r.outStream, PROMPT)
		if !scanner.Scan() {
			return
		}
		text := scanner.Text()

		translation, err := entity.NewTranslation(text, r.Languages)
		if err != nil {
			fmt.Fprintln(r.outStream, err)
			continue
		}

		resultTranlation, err := r.uc.Translate(translation)
		if err != nil {
			fmt.Fprintln(r.outStream, err)
			break
		}

		fmt.Fprintln(r.outStream, resultTranlation.TranslatedText)
	}
}
