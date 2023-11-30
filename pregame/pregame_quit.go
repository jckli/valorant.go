package pregame

import (
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

func PostPregameQuit(a *valorant.Auth, matchId string) (ok bool, err error) {
	_, err = utils.PostRequest("/pregame/v1/matches/"+matchId+"/quit", "glz", a)
	if err != nil {
		return false, err
	}
	return true, nil
}
