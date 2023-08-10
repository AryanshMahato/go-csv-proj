package processor

import (
	"github.com/AryanshMahato/go-csv-proj/pkg/getter"
	"github.com/AryanshMahato/go-csv-proj/pkg/saver"
)

type SaveUsersProcessor struct {
	getter getter.Getter
	saver  saver.Saver
}

func NewSaveUsersProcessor(getter getter.Getter, saver saver.Saver) *SaveUsersProcessor {
	return &SaveUsersProcessor{getter: getter, saver: saver}
}

func (p *SaveUsersProcessor) Process() error {
	users, err := p.getter.GetUsers()
	if err != nil {
		return err
	}

	err = p.saver.SaveUsers(users)
	if err != nil {
		return err
	}

	return nil
}
