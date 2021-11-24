package config

import "fmt"

const (
	baseHost   = "deepl.com"
	apiVersion = "v2"
	pro        = "api"
	free       = "api-free"
)

func (c *Config) BaseURL() string {
	if c.isPro() {
		return fmt.Sprintf("https://%s.%s/%s", pro, baseHost, apiVersion)
	} else {
		return fmt.Sprintf("https://%s.%s/%s", free, baseHost, apiVersion)
	}
}

func (c *Config) isPro() bool {
	return c.Account.AccountPlan == "pro"
}
