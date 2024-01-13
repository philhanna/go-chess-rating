package lichess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRating_String(t *testing.T) {
	type fields struct {
		UltraBullet    string
		Bullet         string
		Blitz          string
		Rapid          string
		Classical      string
		Correspondence string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Empty",
		},
		{
			name: "Only classical",
			fields: fields{
				Classical: "2995",
			},
			want: "Classical=2995",
		},
		{
			name: "All fields",
			fields: fields{
				UltraBullet:    "12",
				Bullet:         "34",
				Blitz:          "56",
				Rapid:          "78",
				Classical:      "910",
				Correspondence: "1112",
			},
			want: "Classical=910,Bullet=34,Blitz=56,Rapid=78,UltraBullet=12,Correspondence=1112",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rating{
				UltraBullet:    tt.fields.UltraBullet,
				Bullet:         tt.fields.Bullet,
				Blitz:          tt.fields.Blitz,
				Rapid:          tt.fields.Rapid,
				Classical:      tt.fields.Classical,
				Correspondence: tt.fields.Correspondence,
			}
			want := tt.want
			have := r.String()
			assert.Equal(t, want, have)
		})
	}
}
