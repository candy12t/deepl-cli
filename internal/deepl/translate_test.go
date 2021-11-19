package deepl

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTranslate(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(TRANSLATE, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, "Content-Type", "application/x-www-form-urlencoded")
		testHeader(t, r, "Authorization", fmt.Sprintf("DeepL-Auth-Key %s", testAuthKey))
		testMethod(t, r, "POST")
		testBody(t, r, "source_lang=EN&target_lang=JA&text=hello")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"translations":[{"detected_source_language":"EN","text":"こんにちわ"}]}`))
	})

	ctx := context.Background()
	got, err := client.Translate(ctx, "hello", "EN", "JA")
	testNoErr(t, err)

	want := &Translate{Translations: []Translation{{DetectedSourceLanguage: "EN", Text: "こんにちわ"}}}
	testTranslate(t, got, want)
}

func testTranslate(t *testing.T, got, want *Translate) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("translated text is %s, want %s", got, want)
	}
}
