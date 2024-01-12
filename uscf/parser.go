package uscf

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	reForm   = regexp.MustCompile(`(?is)<form.*?>(.*?)</form>`)
	reNFound = regexp.MustCompile(`<td colspan=7>Players found: (\d+)</td>`)
	reTR     = regexp.MustCompile(`(?s)<tr>(.*?)</tr>`)
	reTD     = regexp.MustCompile(`(?s)<td.*?>(.*?)</td>`)
	reA      = regexp.MustCompile(`(?s)<a.*?>(.*?)</a>`)
)

// CreateURL makes the URL for accessing the USCF player search
func CreateURL(baseURL, user, state string) string {
	sb := strings.Builder{}
	sb.WriteString(baseURL)
	sb.WriteString("?name=")
	sb.WriteString(strings.ReplaceAll(user, " ", "+"))
	sb.WriteString("&state=")
	sb.WriteString(state)
	sb.WriteString("&rating=R")
	sb.WriteString("&ratingmin=")
	sb.WriteString("&ratingmax=")
	sb.WriteString("&order=N")
	sb.WriteString("&mode=Find")
	return sb.String()
}

// GetRatings returns the ratings rows found in this HTML
func GetRatings(html string) ([]USCFRating, error) {

	form, err := getForm(html)
	if err != nil {
		return nil, err
	}

	ratings := make([]USCFRating, 0)

	// Get the <tr> elements in the form
	trs := getTRs(form)
	if trs == nil {
		return nil, errors.New("no <tr>...</tr> elements found in form")
	}

	// The first <tr> has the number found
	n := getNumberFound(trs[0])
	if n == 0 {
		return nil, nil
	}

	// The second <tr> has the column headings
	ntr := 1

	// The next n <tr>s have the ratings
	for i := 0; i < n; i++ {
		ntr++
		tr := trs[ntr]
		cells := getTDs(tr)
		rating, err := parseUSCFRating(cells)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, *rating)
	}
	return ratings, nil
}

// getForm finds the results form in the html
func getForm(html string) (string, error) {
	m := reForm.FindStringSubmatch(html)
	if m == nil {
		return "", errors.New("no <form> found")
	}
	output := strings.TrimSpace(m[1])
	return output, nil
}

// getNumberFound returns the number of players found
func getNumberFound(s string) int {
	m := reNFound.FindStringSubmatch(s)
	if m == nil {
		return 0
	}
	digits := m[1]
	n, _ := strconv.Atoi(digits)
	return n
}

// getTDs takes a <tr>...</tr> string and returns the individual
// <td>...</td> elements it contains as a slice of strings
func getTDs(tr string) []string {
	tds := make([]string, 0)
	m := reTD.FindAllStringSubmatch(tr, -1)
	if m == nil {
		return nil
	}
	for _, match := range m {
		s := match[1]
		m := reA.FindStringSubmatch(s)
		if m != nil {
			s = m[1]
		}
		tds = append(tds, s)
	}
	return tds
}

// getTRs extracts the <tr> elements from the form
func getTRs(form string) []string {
	trs := make([]string, 0)
	m := reTR.FindAllStringSubmatch(form, -1)
	if m == nil {
		return nil
	}
	for _, match := range m {
		trs = append(trs, match[1])
	}
	return trs
}
