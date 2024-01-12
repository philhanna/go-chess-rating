package lichess

import (
	"io"
	"os"
	"path/filepath"
	"rating"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		html string
		want int
	}{
		{"empty", "", -1},
		{"me", func() string {
			filename := filepath.Join("..", "testdata", "lichess.html")
			fp, err := os.Open(filename)
			assert.Nil(t, err)
			defer fp.Close()
			body, err := io.ReadAll(fp)
			assert.Nil(t, err)
			html := string(body)
			return html
		}(), 1526},
		{"john", func() string {
			filename := filepath.Join("..", "testdata", "john.html")
			fp, err := os.Open(filename)
			assert.Nil(t, err)
			defer fp.Close()
			body, err := io.ReadAll(fp)
			assert.Nil(t, err)
			html := string(body)
			return html
		}(), 1895},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := Parse(tt.html)
			assert.Equal(t, want, have)
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name       string
		dataGetter func() ([]byte, error)
		wantErr    bool
		want       string
	}{
		{
			name: "me",
			dataGetter: func() ([]byte, error) {
				data := "lichess:\n  url: foo\n  defaultUser: me\n"
				return []byte(data), nil
			},
			wantErr: false,
			want:    "me",
		},
		{
			name: "bad yaml",
			dataGetter: func() ([]byte, error) {
				data := "lichess:\n  url foo\n  defaultUser: me\n"
				return []byte(data), nil
			},
			wantErr: true,
		},
		/*
			{
				name: "live",
				wantErr: false,
				want: "pehanna",
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.dataGetter != nil {
				rating.DATA_GETTER = tt.dataGetter
			}
			config, err := rating.LoadConfig()
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			want := tt.want
			have := GetUser(config)
			assert.Equal(t, want, have)
		})
	}
}

func TestGetURL(t *testing.T) {
	tests := []struct {
		name       string
		dataGetter func() ([]byte, error)
		wantErr    bool
		user       string
		want       string
	}{
		{
			name: "me",
			dataGetter: func() ([]byte, error) {
				data := "lichess:\n  url: http://www.example.com/{{user}}\n  defaultUser: me\n"
				return []byte(data), nil
			},
			want: "http://www.example.com/me",
		},
		{
			name: "other",
			dataGetter: func() ([]byte, error) {
				data := "lichess:\n  url: http://www.example.com/{{user}}\n  defaultUser: me\n"
				return []byte(data), nil
			},
			user: "other",
			want: "http://www.example.com/other",
		},
		{
			name: "bad config",
			dataGetter: func() ([]byte, error) {
				data := "lichess:\n  url http://www.example.com/{{user}}\n  defaultUser: me\n"
				return []byte(data), nil
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rating.DATA_GETTER = tt.dataGetter
			var url string
			if tt.user != "" {
				url = GetURL(tt.user)
			} else {
				url = GetURL()
			}
			if tt.wantErr {
				assert.Equal(t, "", url)
				return
			}
			assert.Equal(t, tt.want, url)
		})
	}
}
