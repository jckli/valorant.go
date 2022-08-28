# valorant.go

an API wrapper for VALORANT's client/unofficial API

## getting started

### installing

this assumes you have a working go envirenment, if not please see [this page](https://golang.org/doc/install) first

```sh
go get github.com/jckli/valorant.go
```

## usage

import the package into your project and then run auth

example:
```go
import "github.com/jckli/valorant.go"

func main() {
	puuid := valorant.Authentication("username", "password")
	matchid, _ := valorant.Pregame_fetchPlayer(a)
	pregame, _ := valorant.Pregame_fetchMatch(c)
}
```


## notes

- dont use this to make anything that is against the tos