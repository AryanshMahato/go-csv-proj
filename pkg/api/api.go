package api

import (
	"github.com/AryanshMahato/go-csv-proj/pkg/config"
	"github.com/AryanshMahato/go-csv-proj/pkg/model"
	"io"
	"net/http"
)

type Api interface {
	GetUsers() ([]model.User, error)
}

type AppApi struct {
	client *http.Client
	config config.Config
}

func NewAppApi(client *http.Client, config config.Config) *AppApi {
	return &AppApi{client: client, config: config}
}

// GetUsers returns all users from the API
func (a *AppApi) GetUsers() ([]model.User, error) {
	url, err := a.config.GetApiUrl()
	if err != nil {
		return nil, err
	}

	response, err := a.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return model.UsersFromJson(bodyBytes)
}
