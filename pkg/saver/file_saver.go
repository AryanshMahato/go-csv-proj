package saver

import (
	"fmt"
	"github.com/AryanshMahato/go-csv-proj/pkg/constants"
	"github.com/AryanshMahato/go-csv-proj/pkg/model"
	"github.com/gocarina/gocsv"
	"os"
	"path"
)

type FileSaver struct {
	path string
}

func NewFileSaver(path string) *FileSaver {
	return &FileSaver{path: path}
}

func (f *FileSaver) SaveUsers(users []model.User) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	csvFilePath := path.Join(cwd, f.path, constants.UserCsvFileName)

	csvFile, err := os.OpenFile(csvFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	err = gocsv.MarshalFile(&users, csvFile)
	if err != nil {
		return err
	}

	fmt.Printf("Users saved successfully at: %s\n", csvFilePath)

	return nil
}
