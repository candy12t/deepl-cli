package controller

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/candy12t/deepl-cli/internal/adapter/usecase"
)

const PROMPT = ">> "

type Repl struct {
	uc        usecase.DeepL
	source    string
	target    string
	inStream  io.Reader
	outStream io.Writer
}

func NewRepl(uc usecase.DeepL, source, target string, inStream io.Reader, outStream io.Writer) *Repl {
	return &Repl{
		uc:        uc,
		source:    source,
		target:    target,
		inStream:  inStream,
		outStream: outStream,
	}
}

func (r *Repl) Apply() {
	fmt.Fprintf(r.outStream, "Translate text from %s to %s\n", r.source, r.target)
	scanner := bufio.NewScanner(r.inStream)

	for {
		fmt.Fprint(r.outStream, PROMPT)
		if !scanner.Scan() {
			return
		}
		original := scanner.Text()
		if !isValid(original) {
			fmt.Fprintln(r.outStream, ErrTextLength)
			continue
		}

		text, err := r.uc.Translate(original, r.source, r.target)
		if err != nil {
			fmt.Fprintln(r.outStream, err)
			break
		}

		fmt.Fprintln(r.outStream, text)
	}
}

var ErrTextLength = errors.New("Error: text length is 0")

func isValid(text string) bool {
	trimedspace := strings.TrimSpace(text)
	if len(trimedspace) == 0 {
		return false
	}
	return true
}
