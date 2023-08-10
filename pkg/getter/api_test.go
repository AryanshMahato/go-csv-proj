package getter

import (
	"github.com/AryanshMahato/go-csv-proj/mocks"
	"github.com/AryanshMahato/go-csv-proj/pkg/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppApi_GetUsers(t *testing.T) {
	type fields struct {
		config *mocks.Config
	}

	type testMock struct {
		apiResponse []byte
	}

	tests := []struct {
		name   string
		fields fields
		testMock
		want    []model.User
		wantErr bool
	}{
		{
			name: "no users found",
			fields: fields{
				config: &mocks.Config{},
			},
			testMock: testMock{
				apiResponse: []byte(`[]`),
			},
			want: make([]model.User, 0),
		},
		{
			name: "users found",
			fields: fields{
				config: &mocks.Config{},
			},
			testMock: testMock{
				apiResponse: []byte(`[{"_id":"64d39b0582ec3cff5fc7f24e","index":0,"guid":"03ee84da-5a54-493f-8438-60bad7ab6e2a","isActive":true,"balance":"$2,633.92","tags":["pariatur","qui","ea","culpa","laboris","laboris","minim"],"friends":[{"id":0,"name":"KochValdez"},{"id":1,"name":"KramerBush"},{"id":2,"name":"TownsendChurch"}]}]`),
			},
			want: []model.User{
				{
					ID:       "64d39b0582ec3cff5fc7f24e",
					Index:    0,
					GUID:     "03ee84da-5a54-493f-8438-60bad7ab6e2a",
					IsActive: true,
					Balance:  "$2,633.92",
					Tags:     []string{"pariatur", "qui", "ea", "culpa", "laboris", "laboris", "minim"},
					Friends:  []model.Friend{{ID: 0, Name: "KochValdez"}, {ID: 1, Name: "KramerBush"}, {ID: 2, Name: "TownsendChurch"}},
				},
			},
		},
		{
			name: "users found",
			fields: fields{
				config: &mocks.Config{},
			},
			testMock: testMock{
				apiResponse: []byte(`[`),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write(tt.testMock.apiResponse)
			}))

			tt.fields.config.On("GetApiUrl").Return(testServer.URL, nil).Once()

			a := NewApiGetter(testServer.Client(), tt.fields.config)

			got, err := a.GetUsers()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
