package uscf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseUSCFRating(t *testing.T) {
	tests := []struct {
		name    string
		cells   []string
		want    *USCFRating
		wantErr bool
	}{
		{
			name:    "Empty",
			wantErr: true,
		},
		{
			name: "Good",
			cells: []string{
				"128",
				"1840",
				"xxxx",
				"xxx",
				"1616",
				"Unrated",
				"xxxx",
				"NC",
				"1953-12-04",
				"ME, MYSELF   ",
			},
			want: &USCFRating{
				ID:      "128",
				Rating:  "1840",
				State:   "NC",
				ExpDate: "1953-12-04",
				Name:    "ME, MYSELF",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have, err := parseUSCFRating(tt.cells)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, want, have)
		})
	}
}

func Test_trim(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Has trailing spaces", args{"1531/ &nbsp;&nbsp;"}, "1531/"},
		{"Only 1 token", args{"1531"}, "1531"},
		{"empty", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trim(tt.args.s); got != tt.want {
				t.Errorf("trim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUSCFRating_String(t *testing.T) {
	type fields struct {
		ID      string
		Rating  string
		State   string
		ExpDate string
		Name    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Good",
			fields: fields{
				ID:      "123",
				Rating:  "456",
				State:   "OR",
				ExpDate: "12/31/1999",
				Name:    "JOHN DOE",
			},
			want: `ID=123,Rating=456,State=OR,ExpDate=12/31/1999,Name="JOHN DOE"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := USCFRating{
				ID:      tt.fields.ID,
				Rating:  tt.fields.Rating,
				State:   tt.fields.State,
				ExpDate: tt.fields.ExpDate,
				Name:    tt.fields.Name,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("USCFRating.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
