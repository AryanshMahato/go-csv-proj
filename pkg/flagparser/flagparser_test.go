package flagparser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppFlagParser_GetTags(t *testing.T) {
	type fields struct {
		tagFlag string
	}

	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "should return empty array when tagFlag is empty",
			fields: fields{
				tagFlag: "",
			},
			want: []string(nil),
		},
		{
			name: "should return array of tags when tagFlag is not empty",
			fields: fields{
				tagFlag: "tag1,tag2",
			},
			want: []string{"tag1", "tag2"},
		},
		{
			name: "should return one tag when tagFlag has one tag",
			fields: fields{
				tagFlag: "tag1",
			},
			want: []string{"tag1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagParser := NewAppFlagParser(tt.fields.tagFlag)

			got := flagParser.GetTags()
			assert.Equal(t, tt.want, got)
		})
	}
}
