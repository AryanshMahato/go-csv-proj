package config

import (
	"os"
	"strings"
)

const (
	ApiURLKey = "API_URL"
)

type Config interface {
	GetApiUrls() ([]string, error)
}

type AppConfig struct {
	apiUrls []string
}

func NewAppConfig() *AppConfig {
	apiUrls := os.Getenv(ApiURLKey)
	parsedApiUrls := strings.Split(apiUrls, ",")
	var correctApiUrls []string
	for _, url := range parsedApiUrls {
		if url != "" {
			correctApiUrls = append(correctApiUrls, url)
		}
	}

	return &AppConfig{
		apiUrls: correctApiUrls,
	}
}

// GetApiUrls return apiUrl from the config(env)
func (a *AppConfig) GetApiUrls() ([]string, error) {
	if len(a.apiUrls) == 0 {
		return nil, ErrNotFound
	}

	return a.apiUrls, nil
}
