package processor

import (
	"github.com/AryanshMahato/go-csv-proj/pkg/api"
	"github.com/AryanshMahato/go-csv-proj/pkg/saver"
)

type GetUsersProcessor struct {
	api   api.Api
	saver saver.Saver
}

func NewGetUsersProcessor(api api.Api, saver saver.Saver) *GetUsersProcessor {
	return &GetUsersProcessor{api: api, saver: saver}
}

func (p *GetUsersProcessor) Process() error {
	users, err := p.api.GetUsers()
	if err != nil {
		return err
	}

	err = p.saver.SaveUsers(users)
	if err != nil {
		return err
	}

	return nil
}
