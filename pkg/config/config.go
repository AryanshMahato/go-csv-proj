package config

import "os"

const (
	ApiURLKey = "API_URL"
)

type Config interface {
	GetApiUrl() (string, error)
}

type AppConfig struct {
	apiUrl string
}

func NewAppConfig() *AppConfig {
	apiUrl := os.Getenv(ApiURLKey)

	return &AppConfig{
		apiUrl,
	}
}

// GetApiUrl return apiUrl from the config(env)
func (a *AppConfig) GetApiUrl() (string, error) {
	if len(a.apiUrl) == 0 {
		return "", ErrNotFound
	}

	return a.apiUrl, nil
}
