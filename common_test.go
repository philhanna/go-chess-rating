package rating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHTML(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{"empty", "", true},
		{"example.com", "http://www.example.com", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pHTML, err := GetHTML(tt.url)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.NotNil(t, pHTML)
			assert.Greater(t, len(*pHTML), 0)
		})
	}
}
