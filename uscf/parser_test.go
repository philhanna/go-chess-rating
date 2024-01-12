package uscf

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetForm(t *testing.T) {
	tests := []struct {
		name    string
		html    string
		wantErr bool
	}{
		{
			name: "Good",
			html: func() string {
				filename := filepath.Join("..", "testdata", "uscf-with-results.html")
				body, err := os.ReadFile(filename)
				assert.Nil(t, err)
				return string(body)
			}(),
			wantErr: false,
		},
		{
			name:    "Bad",
			html:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have, err := getForm(tt.html)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.True(t, strings.HasPrefix(have, "<table"))
			assert.True(t, strings.HasSuffix(have, "</table>"))
		})
	}
}

func TestGetPlayersFound(t *testing.T) {
	tests := []struct {
		name string
		form string
		want int
	}{
		{
			name: "Good",
			form: "<tr><td colspan=7>Players found: 3</td></tr>",
			want: 3,
		},
		{
			name: "Bogus",
			form: "<ttd colspan=7>Players found: 3</td></tr>",
			want: 0,
		},
		{
			name: "Zero",
			form: "<tr><td colspan=7>Players found: 0</td></tr>",
			want: 0,
		},
		{
			name: "One",
			form: `
Something
Something
<td colspan=7>Players found: 1</td></tr>
Something`,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := getNumberFound(tt.form)
			assert.Equal(t, want, have)
		})
	}
}

func Test_getTrs(t *testing.T) {
	tests := []struct {
		name    string
		form    string
		number  int
		wantErr bool
	}{
		{
			name: "Good",
			form: func() string {
				filename := filepath.Join("..", "testdata", "goodForm.html")
				body, err := os.ReadFile(filename)
				assert.Nil(t, err)
				return string(body)
			}(),
			number: 8,
		},
		{
			name:    "Bad",
			form:    `Something else`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trs := getTRs(tt.form)
			if tt.wantErr {
				assert.Nil(t, trs)
				return
			}
			assert.Equal(t, tt.number, len(trs))
		})
	}
}

func Test_getTDs(t *testing.T) {
	tests := []struct {
		name   string
		tr     string
		number int
		want   []string
	}{
		{
			name:   "Zero tds",
			tr:     `<tr></tr>`,
			number: 0,
			want:   nil,
		},
		{
			name:   "Good",
			tr:     `<tr><td first="true">abc</td><td>def</td></tr>`,
			number: 2,
			want:   []string{"abc", "def"},
		},
		{
			name: "Including <a>...</a>",
			tr: `<tr><td>abc</td><td>def</td><td>
			<a href=https://www.example.com>LAST, FIRST</a>
</td></tr>`,
			number: 3,
			want:   []string{"abc", "def", "LAST, FIRST"},
		},
		{
			name:   "Empty strings",
			tr:     `<tr><td></td><td></td></tr>`,
			number: 2,
			want:   []string{"", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tds := getTDs(tt.tr)
			if tt.want == nil {
				assert.Nil(t, tds)
				return
			}
			assert.Equal(t, tt.number, len(tds))
			assert.Equal(t, tt.want, tds)
		})
	}
}

func TestGetRatings(t *testing.T) {
	tests := []struct {
		name    string
		html    string
		want    []USCFRating
		wantErr bool
	}{
		{
			name: "Bogus results",
			html: func() string {
				filename := filepath.Join("..", "testdata", "uscf-bogus.html")
				body, err := os.ReadFile(filename)
				assert.Nil(t, err)
				return string(body)
			}(),
			wantErr: true,
		},
		{
			name: "Zero results",
			html: func() string {
				filename := filepath.Join("..", "testdata", "uscf-no-match.html")
				body, err := os.ReadFile(filename)
				assert.Nil(t, err)
				return string(body)
			}(),
		},
		{
			name:    "Empty form",
			html:    `<form>...</form>`,
			wantErr: true,
		},
		{
			name: "Good",
			html: func() string {
				filename := filepath.Join("..", "testdata", "uscf-with-results.html")
				body, err := os.ReadFile(filename)
				assert.Nil(t, err)
				return string(body)
			}(),
			want: []USCFRating{
				{"17025343", "105/7", "NC", "2020-03-31", "FARHAT, ALEXANDER JOHNATHAN"},
				{"12877028", "1830", "NC", "Life", "HANNA, JOHN"},
				{"13056143", "543/5", "NC", "2006-01-31", "JOHNSON, HANNAH L"},
			},
		},
		{
			name:    "Empty",
			html:    "",
			wantErr: true,
		},
		{
			name: "Zero results found",
			html: "<tr><td colspan=7>Players found: 0</td></tr>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have, err := GetRatings(tt.html)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, have)
				return
			}
			assert.Equal(t, len(tt.want), len(have))
			for i := 0; i < len(have); i++ {
				wantRating := tt.want[i]
				haveRating := have[i]
				assert.Equal(t, wantRating, haveRating)
			}
		})
	}
}

func TestCreateURL(t *testing.T) {
	type args struct {
		baseURL string
		user    string
		state   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Good",
			args: args{
				baseURL: "https://www.example.com",
				user:    "JOHN Q. PUBLIC",
				state:   "NC",
			},
			want: "https://www.example.com?name=JOHN+Q.+PUBLIC&state=NC&rating=R&ratingmin=&ratingmax=&order=N&mode=Find"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := CreateURL(tt.args.baseURL, tt.args.user, tt.args.state)
			assert.Equal(t, want, have)
		})
	}
}
