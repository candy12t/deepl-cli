package deepl

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"reflect"
	"testing"
)

func TestTranslate(t *testing.T) {

	t.Run("success translate", func(t *testing.T) {
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

		want := &Translate{Translations: []Translation{{DetectedSourceLanguage: "EN", Text: "こんにちわ"}}}

		ctx := context.Background()
		got, err := client.Translate(ctx, "hello", "EN", "JA")

		testErr(t, err, nil)
		testTranslate(t, got, want)
	})

	t.Run("failed translate because unspecified target language", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()

		mux.HandleFunc(TRANSLATE, func(w http.ResponseWriter, r *http.Request) {
			testHeader(t, r, "Content-Type", "application/x-www-form-urlencoded")
			testHeader(t, r, "Authorization", fmt.Sprintf("DeepL-Auth-Key %s", testAuthKey))
			testMethod(t, r, "POST")
			testBody(t, r, "source_lang=EN&target_lang=&text=hello")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message":"\"Value for 'target_lang' not supported.\""}`))
		})

		u := *client.BaseURL
		u.Path = path.Join(client.BaseURL.Path, TRANSLATE)
		want := HTTPError{StatusCode: http.StatusBadRequest, RequestURL: u.String(), Message: `"Value for 'target_lang' not supported."`}

		ctx := context.Background()
		got, err := client.Translate(ctx, "hello", "EN", "")

		testErr(t, err, want)
		testTranslate(t, got, nil)
	})

	t.Run("failed translate because incorrect DeepL AuthKey", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()

		mux.HandleFunc(TRANSLATE, func(w http.ResponseWriter, r *http.Request) {
			testHeader(t, r, "Content-Type", "application/x-www-form-urlencoded")
			testHeader(t, r, "Authorization", fmt.Sprintf("DeepL-Auth-Key %s", testAuthKey))
			testMethod(t, r, "POST")
			testBody(t, r, "source_lang=EN&target_lang=JA&text=hello")

			w.WriteHeader(http.StatusForbidden)
		})

		u := *client.BaseURL
		u.Path = path.Join(client.BaseURL.Path, TRANSLATE)
		want := HTTPError{StatusCode: http.StatusForbidden, RequestURL: u.String(), Message: "403 Forbidden"}

		ctx := context.Background()
		got, err := client.Translate(ctx, "hello", "EN", "JA")

		testErr(t, err, want)
		testTranslate(t, got, nil)
	})
}

func testTranslate(t *testing.T, got, want *Translate) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("translated text is %s, want %s", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an translate response.")
	}
}
