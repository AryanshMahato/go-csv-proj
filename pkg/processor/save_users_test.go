package processor

import (
	"errors"
	"github.com/AryanshMahato/go-csv-proj/mocks"
	"github.com/AryanshMahato/go-csv-proj/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSaveUsersProcessor_Process(t *testing.T) {
	type mockValues struct {
		getUsersResponse struct {
			users []model.User
			err   error
		}
		saveUsersResponse struct {
			err error
		}
	}

	tests := []struct {
		name       string
		wantErr    error
		mockValues mockValues
	}{
		{
			name: "process successful",
			mockValues: mockValues{
				getUsersResponse: struct {
					users []model.User
					err   error
				}{[]model.User{}, nil},
				saveUsersResponse: struct{ err error }{nil},
			},
			wantErr: nil,
		},
		{
			name: "get users failed",
			mockValues: mockValues{
				getUsersResponse: struct {
					users []model.User
					err   error
				}{nil, errors.New("get users failed")},
				saveUsersResponse: struct{ err error }{nil},
			},
			wantErr: errors.New("get users failed"),
		},
		{
			name: "save users failed",
			mockValues: mockValues{
				getUsersResponse: struct {
					users []model.User
					err   error
				}{nil, nil},
				saveUsersResponse: struct{ err error }{errors.New("save users failed")},
			},
			wantErr: errors.New("save users failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getter := &mocks.Getter{}
			saver := &mocks.Saver{}

			getter.On("GetUsers").Return(tt.mockValues.getUsersResponse.users, tt.mockValues.getUsersResponse.err).Times(1)
			saver.On("SaveUsers", mock.Anything).Return(tt.mockValues.saveUsersResponse.err).Maybe()

			p := NewSaveUsersProcessor(getter, saver)
			err := p.Process()

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
