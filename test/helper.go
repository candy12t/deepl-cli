package test

import (
	"testing"
)

func AssertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error is %q, want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}
