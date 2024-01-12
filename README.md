# Chess-rating
[![Go Report Card](https://goreportcard.com/badge/github.com/philhanna/chess-rating)][idGoReportCard]
[![PkgGoDev](https://pkg.go.dev/badge/github.com/philhanna/chess-rating)][idPkgGoDev]


Prints the chess rating of the user.

## Configuration

The URL from which the rating is extracted is found in the user
configuration file, which you need to create when installing the code.
This is a file with the name of `config.yaml`. It should be placed in
the user configuration directory, which is:
- On Windows: `%APPDATA%/chess-rating`
- On Linux: `$HOME/.config/chess-rating`
  
The configuration file format format is:
```yaml
lichess:
  url: https://lichess.org/@/{{user}}
  defaultUser: pehanna

USCF:
  url: https://www.uschess.org/datapage/player-search.php
  defaultUser: "PHIL HANNA"
  defaultState: "NC"
```
NOTE: Be sure to use spaces, not tabs.  YAML does not work with tabs.

## Links
- [Lichess player profile](https://lichess.org/@/pehanna) (substitute your user ID for pehanna)
- [USCF player search](https://www.uschess.org/datapage/player-search.php)


[idGoReportCard]: https://goreportcard.com/report/github.com/philhanna/chess-rating
[idPkgGoDev]: https://pkg.go.dev/github.com/philhanna/chess-rating
