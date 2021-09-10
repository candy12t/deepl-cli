package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/candy12t/deepl-cli/internal/config"
)

func Request(endpoint string, method string, body io.Reader, data interface{}, header map[string]string) error {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil
	}

	req.Header.Set("Authorization", fmt.Sprintf("DeepL-Auth-Key %s", config.AuthKey()))
	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return HandleResponse(resp, data)
}

func HandleResponse(resp *http.Response, data interface{}) error {
	success := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !success {
		return HandleHTTPError(resp)
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}
	return nil
}
