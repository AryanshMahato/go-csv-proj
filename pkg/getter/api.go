package getter

import (
	"errors"
	"fmt"
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
	var errorsList []error
	urls, err := a.config.GetApiUrls()
	if err != nil {
		return nil, err
	}

	for _, url := range urls {
		users, err := a.getUsers(url)
		if err != nil {
			errorsList = append(errorsList, err)
			fmt.Println(err, "for url", url)
		} else {
			return users, nil
		}
	}

	return nil, errors.Join(errorsList...)
}

func (a *ApiGetter) getUsers(url string) ([]model.User, error) {
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
