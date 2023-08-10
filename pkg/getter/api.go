package getter

import (
	"github.com/AryanshMahato/go-csv-proj/pkg/config"
	"github.com/AryanshMahato/go-csv-proj/pkg/model"
	"io"
	"net/http"
)

type ApiGetter struct {
	client *http.Client
	config config.Config
}

func NewApiGetter(client *http.Client, config config.Config) *ApiGetter {
	return &ApiGetter{client: client, config: config}
}

// GetUsers returns all users from the API
func (a *ApiGetter) GetUsers() ([]model.User, error) {
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
