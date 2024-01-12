package uscf

import (
	"fmt"
	"strings"
)

type USCFRating struct {
	ID      string
	Rating  string
	State   string
	ExpDate string
	Name    string
}

// String returns a string representation of this player's rating
func (r USCFRating) String() string {
	return fmt.Sprintf("ID=%s,Rating=%s,State=%s,ExpDate=%s,Name=%q",
		r.ID, r.Rating, r.State, r.ExpDate, r.Name)
}

// parseUSCFRating creates a ratings object from an array of strings
func parseUSCFRating(cells []string) (*USCFRating, error) {
	if len(cells) != 10 {
		errmsg := fmt.Errorf("expected 10 cells, got %d", len(cells))
		return nil, errmsg
	}
	r := new(USCFRating)
	r.ID = trim(cells[0])
	r.Rating = trim(cells[1])
	r.State = trim(cells[7])
	r.ExpDate = trim(cells[8])
	r.Name = strings.TrimSpace(cells[9]) // Expecting trailing spaces
	return r, nil
}

// trim chops off anything afte
func trim(s string) string {
	tokens := strings.Split(s, " ")
	return tokens[0]
}
