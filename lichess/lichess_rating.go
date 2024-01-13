package lichess

import (
	"fmt"
	"strings"
)

type Rating struct {
	UltraBullet    string
	Bullet         string
	Blitz          string
	Rapid          string
	Classical      string
	Correspondence string
}

// String returns a string representation of this rating
func (r Rating) String() string {
	parts := make([]string, 0)
	if r.Classical != "" {
		parts = append(parts, fmt.Sprintf("Classical=%s", r.Classical))
	}
	if r.Bullet != "" {
		parts = append(parts, fmt.Sprintf("Bullet=%s", r.Bullet))
	}
	if r.Blitz != "" {
		parts = append(parts, fmt.Sprintf("Blitz=%s", r.Blitz))
	}
	if r.Rapid != "" {
		parts = append(parts, fmt.Sprintf("Rapid=%s", r.Rapid))
	}
	if r.UltraBullet != "" {
		parts = append(parts, fmt.Sprintf("UltraBullet=%s", r.UltraBullet))
	}
	if r.Correspondence != "" {
		parts = append(parts, fmt.Sprintf("Correspondence=%s", r.Correspondence))
	}
	output := strings.Join(parts, ",")
	return output
}
