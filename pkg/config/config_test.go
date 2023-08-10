package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAppConfig_GetApiUrl(t *testing.T) {
	type fields struct {
		apiUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr error
	}{
		{
			name: "api url is present in env",
			fields: fields{
				apiUrl: "API_URL",
			},
			want:    "API_URL",
			wantErr: nil,
		},
		{
			name: "api url is not set in env",
			fields: fields{
				apiUrl: "",
			},
			want:    "",
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = os.Setenv(ApiURLKey, tt.fields.apiUrl)
			a := NewAppConfig()

			got, err := a.GetApiUrl()

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
