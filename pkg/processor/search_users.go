package processor

import "github.com/AryanshMahato/go-csv-proj/pkg/model"

type SearchUsersProcessor struct {
	tags    []string
	csvPath string
}

func NewSearchUsersProcessor(tags []string, csvPath string) *SearchUsersProcessor {
	return &SearchUsersProcessor{tags: tags, csvPath: csvPath}
}

func (p *SearchUsersProcessor) Process() error {
	return nil
}

func (p *SearchUsersProcessor) GetUsers() ([]model.User, error) {
	return []model.User{}, nil
}
