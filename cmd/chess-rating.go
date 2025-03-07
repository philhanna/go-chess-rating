package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	rating "github.com/philhanna/chess-rating"
	"github.com/philhanna/chess-rating/lichess"
	"github.com/philhanna/chess-rating/uscf"
)

var (
	optType  string
	optUser  string
	optState string
)

var config *rating.Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Usage = func() {
		text := `Usage: chess-rating [OPTIONS]

Prints the chess rating of the user.

config.yaml contains the URL from which the rating is extracted and the
default userID contained in the URL (if necessary).  See README.md
for details.

options:
  -t, --type          Rating type, one of [lichess, uscf] (default=lichess)
  -u, --user          User ID
  -s, --state         State code (only for USCF)

git repository: https://github.com/philhanna/chess-rating
`
		fmt.Fprint(os.Stderr, text)
	}
	flag.StringVar(&optType, "type", "lichess", "Rating type (default=lichess)")
	flag.StringVar(&optType, "t", "lichess", "Rating type (default=lichess)")
	flag.StringVar(&optUser, "user", "", "User ID (default in configuration file)")
	flag.StringVar(&optUser, "u", "", "User ID (default in configuration file)")
	flag.StringVar(&optState, "state", "", "State code (only used for USCF)")
	flag.StringVar(&optState, "s", "", "State code (only used for USCF)")

	flag.Parse()

	// Get the rating

	switch optType {
	case "lichess", "l", "":
		doLichess()
	case "uscf", "u":
		doUSCF()
	default:
		log.Fatalf("%q is not a valid rating type\n", optType)
	}
}

func init() {
	var err error
	config, err = rating.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func doLichess() {

	// Load the configuration data

	if optUser == "" {
		defaultUser := config.Lichess.DefaultUser
		optUser = defaultUser
	}

	// Get the URL for this user in lichess

	url := lichess.GetURL(optUser)
	if url == "" {
		log.Fatalf("could not get URL for user %s\n", optUser)
	}

	// Read the HTML pointed to by the URL

	html, err := rating.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Get the rating

	rating := lichess.Parse(*html)

	// Print the rating

	if rating == nil {
		fmt.Printf("Lichess rating for %s is unknown\n", optUser)
	} else {
		fmt.Printf("%v\n", rating)
	}
}

func doUSCF() {

	// USCF requires a state code
	if optState == "" {
		defaultState := config.USCF.DefaultState
		optState = defaultState
		if optState == "" {
			optState = "ANY"
		}
	}

	// Load the default USCF user
	if optUser == "" {
		defaultUser := config.USCF.DefaultUser
		optUser = defaultUser
	}

	// Get the URL
	url := uscf.CreateURL(config.USCF.URL, optUser, optState)

	// Read the HTML pointed to by the URL

	html, err := rating.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Get the rating

	ratings, err := uscf.GetRatings(*html)
	if err != nil {
		log.Fatal(err)
	}

	// Print the rating(s)
	switch len(ratings) {
	case 0:
		fmt.Printf("No USCF ratings found for %q\n", optUser)
	default:
		for _, rating := range ratings {
			fmt.Printf("%v\n", rating)
		}
	}

}
