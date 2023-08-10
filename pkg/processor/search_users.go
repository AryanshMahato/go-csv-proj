package processor

import (
	"fmt"
	"github.com/AryanshMahato/go-csv-proj/pkg/getter"
	"github.com/AryanshMahato/go-csv-proj/pkg/model"
	"github.com/AryanshMahato/go-csv-proj/pkg/util"
)

type SearchUsersProcessor struct {
	tags   []string
	getter getter.Getter
}

func NewSearchUsersProcessor(tags []string, getter getter.Getter) *SearchUsersProcessor {
	return &SearchUsersProcessor{tags: tags, getter: getter}
}

func (p *SearchUsersProcessor) Process() error {
	users, err := p.getter.GetUsers()
	if err != nil {
		return err
	}

	var matchedUsers []model.User
	for _, user := range users {
		if p.isTagMatching(user) {
			matchedUsers = append(matchedUsers, user)
		}
	}

	fmt.Println(matchedUsers)

	return nil
}

func (p *SearchUsersProcessor) isTagMatching(user model.User) bool {
	for _, tag := range p.tags {
		if !util.Contains(user.Tags, tag) {
			return false
		}
	}

	return true
}
