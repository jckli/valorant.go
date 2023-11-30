package game

import (
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

// Quits the current game
func PostGameQuit(a *valorant.Auth, puuid, matchId string) (ok bool, err error) {
	_, err = utils.GetRequest("/core-game/v1/players/"+puuid+"/disassociate/"+matchId, "glz", a)
	if err != nil {
		return false, err
	}
	return true, nil
}
