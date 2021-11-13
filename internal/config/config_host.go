package config

import "fmt"

const (
	baseHost   = "deepl.com"
	apiVersion = "v2"
	pro        = "api"
	free       = "api-free"
)

func baseURL() string {
	if isPro() {
		return fmt.Sprintf("https://%s.%s/%s", pro, baseHost, apiVersion)
	} else {
		return fmt.Sprintf("https://%s.%s/%s", free, baseHost, apiVersion)
	}
}

func isPro() bool {
	return CachedConfig().Account.AccountPlan == "pro"
}
