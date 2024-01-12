package lichess

import (
	rating "github.com/philhanna/chess-rating"
	"regexp"
	"strconv"
	"strings"
)

// Regular expressions
var (
	reH3 = regexp.MustCompile(`(?s)<h3>(.*?)</h3><rating><strong>(\d*).*?</strong>`)
)

// GetURL returns the full URL from the configuration.  It can be called
// with an option string parameter indicating the user name to be
// substituted into the URL.  If no user is specified, it calls
// GetUser() to get the user name from the configuration file.
//
// To load the configuration, GetURL calls rating.LoadConfig(). The
// configuration data can be overridden for unit testing.  See
// rating.LoadConfig() for details.
func GetURL(users ...string) string {
	config, err := rating.LoadConfig()
	if err != nil {
		return ""
	}
	var user string
	if len(users) > 0 {
		user = users[0]
	} else {
		user = GetUser(config)
	}
	url := strings.ReplaceAll(config.Lichess.URL, "{{user}}", user)
	return url
}

// GetUser returns the default user from the configuration
func GetUser(config *rating.Config) string {
	user := config.Lichess.DefaultUser
	return user
}

// Parse accepts the HTML body of the lichess page and extracts the
// classical rating from it.
func Parse(html string) int {
	m := reH3.FindAllStringSubmatch(html, -1)
	for _, match := range m {
		ratingType := match[1]
		if ratingType == "Classical" {
			ratingValue := match[2]
			rating, _ := strconv.Atoi(ratingValue)
			return rating
		}
	}
	return -1
}
