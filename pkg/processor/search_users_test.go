package processor

import (
	"errors"
	"github.com/AryanshMahato/go-csv-proj/mocks"
	"github.com/AryanshMahato/go-csv-proj/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchUsersProcessor_Process(t *testing.T) {
	type mockValues struct {
		getUsersResponse struct {
			users []model.User
			err   error
		}
	}

	type fields struct {
		tags []string
	}

	tests := []struct {
		name       string
		fields     fields
		mockValues mockValues
		wantErr    error
	}{
		{
			name: "users matched",
			fields: fields{
				tags: []string{"tag1", "tag2"},
			},
			mockValues: mockValues{
				getUsersResponse: struct {
					users []model.User
					err   error
				}{
					[]model.User{
						{
							ID:       "64d39b0582ec3cff5fc7f24e",
							Index:    0,
							GUID:     "03ee84da-5a54-493f-8438-60bad7ab6e2a",
							IsActive: true,
							Balance:  "$2,633.92",
							Tags:     []string{"tag1", "tag2"},
							Friends:  []model.Friend{{ID: 0, Name: "KochValdez"}, {ID: 1, Name: "KramerBush"}, {ID: 2, Name: "TownsendChurch"}},
						},
					},
					nil,
				},
			},
			wantErr: nil,
		},
		{
			name: "partial users tag matched",
			fields: fields{
				tags: []string{"tag1", "tag2"},
			},
			mockValues: mockValues{
				getUsersResponse: struct {
					users []model.User
					err   error
				}{
					[]model.User{
						{
							ID:       "64d39b0582ec3cff5fc7f24e",
							Index:    0,
							GUID:     "03ee84da-5a54-493f-8438-60bad7ab6e2a",
							IsActive: true,
							Balance:  "$2,633.92",
							Tags:     []string{"tag1"},
							Friends:  []model.Friend{{ID: 0, Name: "KochValdez"}, {ID: 1, Name: "KramerBush"}, {ID: 2, Name: "TownsendChurch"}},
						},
					},
					nil,
				},
			},
			wantErr: ErrNoMatchingUsers,
		},
		{
			name: "users tag did not matched",
			fields: fields{
				tags: []string{"tag1", "tag2"},
			},
			mockValues: mockValues{
				getUsersResponse: struct {
					users []model.User
					err   error
				}{
					[]model.User{
						{
							ID:       "64d39b0582ec3cff5fc7f24e",
							Index:    0,
							GUID:     "03ee84da-5a54-493f-8438-60bad7ab6e2a",
							IsActive: true,
							Balance:  "$2,633.92",
							Tags:     []string{},
							Friends:  []model.Friend{{ID: 0, Name: "KochValdez"}, {ID: 1, Name: "KramerBush"}, {ID: 2, Name: "TownsendChurch"}},
						},
					},
					nil,
				},
			},
			wantErr: ErrNoMatchingUsers,
		},
		{
			name: "get users failed",
			fields: fields{
				tags: []string{"tag1", "tag2"},
			},
			mockValues: mockValues{
				getUsersResponse: struct {
					users []model.User
					err   error
				}{
					nil,
					errors.New("get users failed"),
				},
			},
			wantErr: errors.New("get users failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getter := &mocks.Getter{}

			getter.On("GetUsers").Return(tt.mockValues.getUsersResponse.users, tt.mockValues.getUsersResponse.err).Times(1)

			p := NewSearchUsersProcessor(tt.fields.tags, getter)
			err := p.Process()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestSearchUsersProcessor_isTagMatching(t *testing.T) {
	type fields struct {
		tags []string
	}
	type args struct {
		user model.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "tags matched",
			fields: fields{
				tags: []string{"tag1", "tag2"},
			},
			args: args{
				user: model.User{
					ID:       "64d39b0582ec3cff5fc7f24e",
					Index:    0,
					GUID:     "03ee84da-5a54-493f-8438-60bad7ab6e2a",
					IsActive: true,
					Balance:  "$2,633.92",
					Tags:     []string{"tag1", "tag2"},
					Friends:  []model.Friend{{ID: 0, Name: "KochValdez"}, {ID: 1, Name: "KramerBush"}, {ID: 2, Name: "TownsendChurch"}},
				},
			},
			want: true,
		},
		{
			name: "tags partial matched",
			fields: fields{
				tags: []string{"tag1", "tag2"},
			},
			args: args{
				user: model.User{
					ID:       "64d39b0582ec3cff5fc7f24e",
					Index:    0,
					GUID:     "03ee84da-5a54-493f-8438-60bad7ab6e2a",
					IsActive: true,
					Balance:  "$2,633.92",
					Tags:     []string{"tag1"},
					Friends:  []model.Friend{{ID: 0, Name: "KochValdez"}, {ID: 1, Name: "KramerBush"}, {ID: 2, Name: "TownsendChurch"}},
				},
			},
			want: false,
		},
		{
			name: "tags did not matched",
			fields: fields{
				tags: []string{"something else"},
			},
			args: args{
				user: model.User{
					ID:       "64d39b0582ec3cff5fc7f24e",
					Index:    0,
					GUID:     "03ee84da-5a54-493f-8438-60bad7ab6e2a",
					IsActive: true,
					Balance:  "$2,633.92",
					Tags:     []string{"tag1"},
					Friends:  []model.Friend{{ID: 0, Name: "KochValdez"}, {ID: 1, Name: "KramerBush"}, {ID: 2, Name: "TownsendChurch"}},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &SearchUsersProcessor{
				tags: tt.fields.tags,
			}
			got := p.isTagMatching(tt.args.user)
			assert.Equal(t, tt.want, got)
		})
	}
}
