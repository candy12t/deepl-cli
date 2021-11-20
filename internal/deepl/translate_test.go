package deepl

import (
	"context"
	"fmt"
	"net/http"
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

		ctx := context.Background()
		got, err := client.Translate(ctx, "hello", "EN", "JA")
		testNoErr(t, err)

		want := &Translate{Translations: []Translation{{DetectedSourceLanguage: "EN", Text: "こんにちわ"}}}
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

		ctx := context.Background()
		got, err := client.Translate(ctx, "hello", "EN", "")
		if got != nil {
			t.Errorf("Expected no response, got %s", got)
		}

		want := fmt.Sprintf("HTTP %d: %s (%s)", http.StatusBadRequest, fmt.Sprintf(`"Value for 'target_lang' not supported."`), client.BaseURL.String()+TRANSLATE)
		if err.Error() != want {
			t.Errorf("got error is %s, want %s", err, want)
		}
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

		ctx := context.Background()
		got, err := client.Translate(ctx, "hello", "EN", "JA")
		if got != nil {
			t.Errorf("Expected no response, got %s", got)
		}

		want := fmt.Sprintf("HTTP %d: %s (%s)", http.StatusForbidden, fmt.Sprintf("403 Forbidden"), client.BaseURL.String()+TRANSLATE)
		if err.Error() != want {
			t.Errorf("got error is %s, want %s", err, want)
		}
	})
}

func testTranslate(t *testing.T, got, want *Translate) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("translated text is %s, want %s", got, want)
	}
}
