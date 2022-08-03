package valorant

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// GET CoreGame_FetchAllChatMUCToken
func Coregame_fetchAllChatMucToken(match_id string) (interface{}, error) {
	url := "/core-game/v1/matches/" + match_id + "/allchatmuctoken"
	resp, err := fetchGet(url, "glz")
	if resp.StatusCode == 404 {
		return "", fmt.Errorf("match not found")
	} else if resp.StatusCode == 500 {
		return "", fmt.Errorf("server error")
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body := new(interface{})
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
}

// GET CoreGame_FetchTeamChatMUCToken
func Coregame_fetchTeamChatMucToken(match_id string) (interface{}, error) {
	url := "/core-game/v1/matches/" + match_id + "/teamchatmuctoken"
	resp, err := fetchGet(url, "glz")
	if resp.StatusCode == 404 {
		return "", fmt.Errorf("match not found")
	} else if resp.StatusCode == 500 {
		return "", fmt.Errorf("server error")
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body := new(interface{})
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
}

// GET CoreGame_FetchMatch
func Coregame_fetchMatch(match_id string) (*CoregameFetchMatchResp, error) {
	url := "/core-game/v1/matches/" + match_id
	resp, err := fetchGet(url, "glz")
	if resp.Status == "404" {
		return nil, fmt.Errorf("match not found")
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body := new(CoregameFetchMatchResp)
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
}

// GET CoreGame_FetchMatchLoadouts
func Coregame_fetchMatchLoadouts(match_id string) (*CoregameFetchMatchLoadoutsResp, error) {
	url := "/core-game/v1/matches/" + match_id + "/loadouts"
	resp, err := fetchGet(url, "glz")
	if resp.Status == "404" {
		return nil, fmt.Errorf("match not found")
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body := new(CoregameFetchMatchLoadoutsResp)
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
}

// GET CoreGame_FetchPlayer
func Coregame_fetchPlayer(puuid string) (string, error) {
	url := "/core-game/v1/players/" + puuid
	resp, err := fetchGet(url, "glz")
	if resp.StatusCode == 404 {
		return "", fmt.Errorf("player not in a game")
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body := new(CoregameFetchPlayerResp)
	json.NewDecoder(resp.Body).Decode(body)

	return body.MatchId, nil
}

// POST CoreGame_DisassociatePlayer
func Coregame_disassociatePlayer(puuid, match_id string) (*CoregameFetchPlayerResp, error) {
	url := "/core-game/v1/players/" + puuid + "/disassociate/" + match_id
	resp, err := fetchP(http.MethodPost, url, "glz", nil)
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("player not in a game")
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body := new(CoregameFetchPlayerResp)
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
	
}