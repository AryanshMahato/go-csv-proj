package getter

import (
	"github.com/AryanshMahato/go-csv-proj/pkg/constants"
	"github.com/AryanshMahato/go-csv-proj/pkg/model"
	"github.com/gocarina/gocsv"
	"os"
	"path"
)

type FileGetter struct {
	path string
}

func NewFileGetter(path string) *FileGetter {
	return &FileGetter{path: path}
}

func (f *FileGetter) GetUsers() (users []model.User, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	csvFilePath := path.Join(cwd, f.path, constants.UserCsvFileName)

	csvFile, err := os.OpenFile(csvFilePath, os.O_RDONLY, 0755)
	defer func(csvFile *os.File) {
		_ = csvFile.Close()
	}(csvFile)

	err = gocsv.UnmarshalFile(csvFile, &users)
	if err != nil {
		return nil, err
	}

	return users, err
}
