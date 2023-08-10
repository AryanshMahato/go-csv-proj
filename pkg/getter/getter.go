package getter

import "github.com/AryanshMahato/go-csv-proj/pkg/model"

type Getter interface {
	GetUsers() ([]model.User, error)
}
