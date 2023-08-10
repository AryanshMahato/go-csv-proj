package saver

import "github.com/AryanshMahato/go-csv-proj/pkg/model"

type FileSaver struct {
	path string
}

func NewFileSaver(path string) *FileSaver {
	return &FileSaver{path: path}
}

func (f *FileSaver) SaveUsers(users []model.User) error {
	//TODO implement me
	panic("implement me")
}
