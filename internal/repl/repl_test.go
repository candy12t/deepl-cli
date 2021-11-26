package repl

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/candy12t/deepl-cli/internal/deepl"
	"github.com/candy12t/deepl-cli/internal/deepl/mock_deepl"
	"github.com/golang/mock/gomock"
)

func TestRepl(t *testing.T) {
	type args struct {
		inputText  string
		sourceLang string
		targetLang string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(m *mock_deepl.MockAPIClient)
		want          string
	}{
		{
			name: "success repl",
			args: args{
				inputText:  "hello",
				sourceLang: "EN",
				targetLang: "JA",
			},
			prepareMockFn: func(m *mock_deepl.MockAPIClient) {
				ctx := context.Background()
				m.EXPECT().
					Translate(ctx, "hello", "EN", "JA").
					Return(&deepl.Translate{Translations: []deepl.Translation{{DetectedSourceLanguage: "EN", Text: "こんにちわ"}}}, nil)
			},
			want: fmt.Sprintf("%sこんにちわ\n%s", PROMPT, PROMPT),
		},
		{
			name: "failed repl because input text length is 0",
			args: args{
				inputText:  "\n",
				sourceLang: "EN",
				targetLang: "JA",
			},
			prepareMockFn: func(m *mock_deepl.MockAPIClient) {},
			want:          fmt.Sprintf("%sError: text length is 0\n%s", PROMPT, PROMPT),
		},
		{
			name: "failed repl because not specified target language",
			args: args{
				inputText:  "hello",
				sourceLang: "EN",
				targetLang: "",
			},
			prepareMockFn: func(m *mock_deepl.MockAPIClient) {
				ctx := context.Background()
				m.EXPECT().
					Translate(ctx, "hello", "EN", "").
					Return(nil, fmt.Errorf(`HTTP 400: "Value for 'target_lang' not supported." (https://api-free.deepl.com/v2/translate)`))
			},
			want: fmt.Sprintf("%sHTTP 400: \"Value for 'target_lang' not supported.\" (https://api-free.deepl.com/v2/translate)\n", PROMPT),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mock_deepl.NewMockAPIClient(ctrl)

			tt.prepareMockFn(mockClient)

			out := new(bytes.Buffer)
			Repl(mockClient, tt.args.sourceLang, tt.args.targetLang, bytes.NewBufferString(tt.args.inputText), out)

			got := out.String()
			if got != tt.want {
				t.Fatalf("repl output %q, want %q", got, tt.want)
			}
		})
	}
}
