package rating

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfigFile(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"live", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := ReadConfigFile()
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			data := string(body)
			fmt.Print(data)
		})
	}
}

func TestLoadConfig(t *testing.T) {
	// Mock data getter for "file not found"
	fnf := func() ([]byte, error) {
		return make([]byte, 0), errors.New("file not found")
	}
	// Mock data getter for "invalid yaml"
	invalid := func() ([]byte, error) {
		data := `
		lichess
		  url www.example.com
		  defaultUser me
		`
		body := []byte(data)
		return body, nil
	}
	tests := []struct {
		name    string
		f       func() ([]byte, error)
		wantErr bool
		url     string
		user    string
	}{
		{"live", DEFAULT_DATA_GETTER, false, "https://lichess.org/@/{{user}}", "pehanna"},
		{"no file", fnf, true, "", ""},
		{"invalid yaml", invalid, true, "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DATA_GETTER = tt.f
			config, err := LoadConfig()
			DATA_GETTER = DEFAULT_DATA_GETTER
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, tt.url, config.Lichess.URL)
			assert.Equal(t, tt.user, config.Lichess.DefaultUser)
		})
	}
}
