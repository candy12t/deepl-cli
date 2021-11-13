package deepl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type HTTPError struct {
	StatusCode int
	RequestURL *url.URL
	Message    string
}

func (err HTTPError) Error() string {
	if err.Message != "" {
		return fmt.Sprintf("HTTP %d: %s (%s)", err.StatusCode, err.Message, err.RequestURL)
	}
	return fmt.Sprintf("HTTP %d: (%s)", err.StatusCode, err.RequestURL)
}

func HandleHTTPError(resp *http.Response) error {
	httpError := HTTPError{
		StatusCode: resp.StatusCode,
		RequestURL: resp.Request.URL,
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "json") {
		httpError.Message = resp.Status
		return httpError
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		httpError.Message = err.Error()
		return httpError
	}

	var parseBody struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(body, &parseBody); err != nil {
		return httpError
	}
	httpError.Message = parseBody.Message

	return httpError
}
