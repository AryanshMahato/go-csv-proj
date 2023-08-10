package getter

import (
	"github.com/AryanshMahato/go-csv-proj/pkg/model"
)

type FileGetter struct {
	path string
}

func NewFileGetter(path string) *FileGetter {
	return &FileGetter{path: path}
}

func (f *FileGetter) GetUsers() ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}
