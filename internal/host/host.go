package host

import (
	"fmt"

	"github.com/candy12t/deepl-cli/internal/config"
)

const (
	baseHost   = "deepl.com"
	apiVersion = "v2"
)

func Host() string {
	if isPro() {
		return fmt.Sprintf("%s.%s", "api", baseHost)
	} else {
		return fmt.Sprintf("%s.%s", "api-free", baseHost)
	}
}

func TranslateEndpoint() string {
	return fmt.Sprintf("https://%s/%s/translate", Host(), apiVersion)
}

func UsageEndpoint() string {
	return fmt.Sprintf("https://%s/%s/usage", Host(), apiVersion)
}

func isPro() bool {
	return config.AccountPlan() == "pro"
}
