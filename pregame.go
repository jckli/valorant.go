package valorant

import (
	"net/http"
	"fmt"
	"encoding/json"
)

// GET Pregame_GetPlayer
func Pregame_fetchPlayer(puuid string) (string, error) {
	url := "/pregame/v1/players/" + puuid
	resp, err := fetchGet(url, "glz")
	if resp.StatusCode == 404 {
		return "", fmt.Errorf("player not found")
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body := new(PregameFetchPlayerResp)
	json.NewDecoder(resp.Body).Decode(body)

	return body.MatchId, nil
}

// GET Pregame_GetMatch
func Pregame_fetchMatch(matchId string) (*PregameFetchMatchResp, error) {
	url := "/pregame/v1/matches/" + matchId
	resp, err := fetchGet(url, "glz")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body := new(PregameFetchMatchResp)
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
}

// POST Pregame_QuitMatch
func Pregame_quitMatch(matchId string) (error) {
	url := "/pregame/v1/matches/" + matchId + "/quit"
	resp, err := fetchP(http.MethodPost, url, "glz", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	return nil
}