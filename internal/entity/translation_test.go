package entity

import (
	"testing"
)

func TestNewTranslationErr(t *testing.T) {
	original := ""
	_, err := NewTranslation(original, &Languages{})
	got := err.Error()
	want := "Error: text length is 0"
	if got != want {
		t.Fatalf("got error is %q, want error is %q\n", got, want)
	}
}
