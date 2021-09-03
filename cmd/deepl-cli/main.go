package main

import (
	"flag"
	"fmt"

	"github.com/candy12t/deepl-cli/internal/config"
	"github.com/candy12t/deepl-cli/internal/repl"
)

const (
	defaultSourceLang = "EN"
	defaultTargetLang = "JA"
)

func main() {
	var sourceLang, targetLang string
	flag.StringVar(&sourceLang, "source", defaultSourceLang, "Language of the source text.")
	flag.StringVar(&sourceLang, "s", defaultSourceLang, "Language of the source text.")
	flag.StringVar(&targetLang, "target", defaultTargetLang, "Language of the text to be translated.")
	flag.StringVar(&targetLang, "t", defaultTargetLang, "Language of the text to be translated.")
	flag.Parse()
	fmt.Printf("Translate text from %s to %s\n", sourceLang, targetLang)

	config.InitDeeplConfig()
	repl.Repl(sourceLang, targetLang)
}
