package deepl

import (
	"fmt"
	"net/http"
	"path"
	"testing"

	"github.com/candy12t/deepl-cli/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {

	t.Run("success translate", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()

		mux.HandleFunc("/translate", func(w http.ResponseWriter, r *http.Request) {
			testHeader(t, r, "Content-Type", "application/x-www-form-urlencoded")
			testHeader(t, r, "Authorization", fmt.Sprintf("DeepL-Auth-Key %s", testProAuthKey))
			testMethod(t, r, "POST")
			testBody(t, r, "source_lang=EN&target_lang=JA&text=hello")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"translations":[{"detected_source_language":"EN","text":"こんにちは"}]}`))
		})

		want := &model.TranslateText{
			OriginalText:   "hello",
			TranslateText:  "こんにちは",
			SourceLanguage: "EN",
			TargetLanguage: "JA",
		}

		got, err := client.TranslateText("hello", "EN", "JA")

		if assert.NoError(t, err) {
			assert.Equal(t, got, want)
		}
	})

	t.Run("failed translate because unspecified target language", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()

		mux.HandleFunc("/translate", func(w http.ResponseWriter, r *http.Request) {
			testHeader(t, r, "Content-Type", "application/x-www-form-urlencoded")
			testHeader(t, r, "Authorization", fmt.Sprintf("DeepL-Auth-Key %s", testProAuthKey))
			testMethod(t, r, "POST")
			testBody(t, r, "source_lang=EN&target_lang=&text=hello")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message":"\"Value for 'target_lang' not supported.\""}`))
		})

		u := *client.BaseURL
		u.Path = path.Join(client.BaseURL.Path, "/translate")
		want := HTTPError{StatusCode: http.StatusBadRequest, RequestURL: u.String(), Message: `"Value for 'target_lang' not supported."`}

		_, err := client.TranslateText("hello", "EN", "")

		assert.EqualError(t, err, want.Error())
	})

	t.Run("failed translate because incorrect DeepL AuthKey", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()

		mux.HandleFunc("/translate", func(w http.ResponseWriter, r *http.Request) {
			testHeader(t, r, "Content-Type", "application/x-www-form-urlencoded")
			testHeader(t, r, "Authorization", fmt.Sprintf("DeepL-Auth-Key %s", testProAuthKey))
			testMethod(t, r, "POST")
			testBody(t, r, "source_lang=EN&target_lang=JA&text=hello")

			w.WriteHeader(http.StatusForbidden)
		})

		u := *client.BaseURL
		u.Path = path.Join(client.BaseURL.Path, "/translate")
		want := HTTPError{StatusCode: http.StatusForbidden, RequestURL: u.String(), Message: "403 Forbidden"}

		_, err := client.TranslateText("hello", "EN", "JA")

		assert.EqualError(t, err, want.Error())
	})
}
