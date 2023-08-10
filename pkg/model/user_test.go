package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsersFromJson(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name      string
		args      args
		wantUsers []User
		wantErr   bool
	}{
		{
			name: "empty array",
			args: args{
				body: []byte(`[]`),
			},
			wantUsers: []User{},
		},
		{
			name: "users found",
			args: args{
				body: []byte(`[{"_id":"64d39b0582ec3cff5fc7f24e","index":0,"guid":"03ee84da-5a54-493f-8438-60bad7ab6e2a","isActive":true,"balance":"$2,633.92","tags":["pariatur","qui","ea","culpa","laboris","laboris","minim"],"friends":[{"id":0,"name":"KochValdez"},{"id":1,"name":"KramerBush"},{"id":2,"name":"TownsendChurch"}]}]`),
			},
			wantUsers: []User{
				{
					ID:       "64d39b0582ec3cff5fc7f24e",
					Index:    0,
					GUID:     "03ee84da-5a54-493f-8438-60bad7ab6e2a",
					IsActive: true,
					Balance:  "$2,633.92",
					Tags:     []string{"pariatur", "qui", "ea", "culpa", "laboris", "laboris", "minim"},
					Friends:  []Friend{{ID: 0, Name: "KochValdez"}, {ID: 1, Name: "KramerBush"}, {ID: 2, Name: "TownsendChurch"}},
				},
			},
		},
		{
			name: "invalid json",
			args: args{
				body: []byte(`[`),
			},
			wantUsers: nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UsersFromJson(tt.args.body)
			assert.Equal(t, tt.wantUsers, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
