package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	type args struct {
		slice []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true if item is present in slice",
			args: args{
				slice: []string{"a", "b", "c"},
				item:  "a",
			},
			want: true,
		},
		{
			name: "should return false if item is not present in slice",
			args: args{
				slice: []string{"a", "b", "c"},
				item:  "z",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.args.slice, tt.args.item)

			assert.Equal(t, tt.want, got)
		})
	}
}
