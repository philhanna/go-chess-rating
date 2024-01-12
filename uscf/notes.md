## Design notes

The search feature uses an HTTP GET to the player search php page.
Parameters are:

- `name` char(50) -- The player name
- `state` char(2) -- The 2-character postal abbreviation of the state
- `rating` char(1) -- "R"
- `ratingmin` char(0) -- empty
- `ratingmax` char(0) -- empty
- `order` char(1) -- "N"
- `mode` char(4) -- "Find"

### Links
- [USCF player search](https://www.uschess.org/datapage/player-search.php)
