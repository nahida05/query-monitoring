package config

import (
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "12345")
	os.Setenv("DB_NAME", "postgres")

	os.Setenv("SERVICE_PORT", "8080")

	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "set all env variables",
			want: &Config{
				DB: db{
					Host:     "localhost",
					Port:     "5432",
					User:     "postgres",
					Password: "12345",
					Name:     "postgres",
				},
				Server: server{
					Port: "8080",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
