package saver

import "github.com/AryanshMahato/go-csv-proj/pkg/model"

type Saver interface {
	SaveUsers([]model.User) error
}
