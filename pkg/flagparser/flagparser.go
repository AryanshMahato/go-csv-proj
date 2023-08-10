package flagparser

import "strings"

type FlagParser interface {
	GetTags() []string
}

type AppFlagParser struct {
	tagFlag string
}

func NewAppFlagParser(tagFlag string) *AppFlagParser {
	return &AppFlagParser{tagFlag}
}

// GetTags returns all tags from the tagFlag
func (a *AppFlagParser) GetTags() (tags []string) {
	if a.tagFlag == "" {
		return tags
	}

	flags := strings.Split(a.tagFlag, ",")
	if len(flags) > 0 {
		for _, flag := range flags {
			tags = append(tags, flag)
		}
	}

	return
}
