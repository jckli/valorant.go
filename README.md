# valorant.go

> **disclaimer:** this wrapper is still a big work in progress and doesn't cover every endpoint yet. use at own risk.

an golang API wrapper for VALORANT's client/unofficial API written with fasthttp

## getting started

### installing

this assumes you have a working go envirenment, if not please see [this page](https://golang.org/doc/install) first

```sh
go get github.com/jckli/valorant.go
```

## usage

import the package into your project and then login to a valorant account.

example:

```go
package main

import (
    "github.com/jckli/valorant.go"
    "github.com/jckli/valorant.go/game"
)

func main() {
	client, err := valorant.New("username", "password")
    if err != nil {
        panic(err)
    }
    playerGame, _ := game.GetPlayerGame(client, client.UserInfo.UserId)
    gameInfo, _ := game.GetGameInfo(client, playerGame.MatchID)
}
```

## notes

- dont use this to make anything that is against the tos
