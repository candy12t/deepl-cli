package repl

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/candy12t/deepl-cli/internal/deepl"
)

func TestRepl(t *testing.T) {
	out := new(bytes.Buffer)
	in := bytes.NewBufferString("hello")
	want := "こんにちわ\n"
	client, _ := deepl.NewMockClient("https://api.deepl.com/v2", "test-auth-key")
	Repl(client, "EN", "JA", in, out)
	fmt.Println(out.String())
	if out.String() != want {
		t.Fatalf("repl output %s, want %s", out.String(), want)
	}
}
