package game

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type CurrentGameLoadoutsResponse struct {
	Loadouts []struct {
		CharacterID string `json:"CharacterID"`
		Loadout     struct {
			Sprays struct {
				SpraySelection []struct {
					SocketID string `json:"SocketID"`
					SprayID  string `json:"SprayID"`
					LevelID  string `json:"LevelID"`
				} `json:"SpraySelection"`
			} `json:"Sprays"`
			Items map[string]struct {
				ID      string `json:"ID"`
				TypeID  string `json:"TypeID"`
				Sockets map[string]struct {
					ID   string `json:"ID"`
					Item struct {
						ID     string `json:"ID"`
						TypeID string `json:"TypeID"`
					} `json:"Item"`
				} `json:"Sockets"`
			} `json:"Items"`
		} `json:"Loadout"`
	} `json:"Loadouts"`
}

// Get the current game loadout info for all players in the match
func GetGameLoadouts(a *valorant.Auth, matchId string) (*CurrentGameLoadoutsResponse, error) {
	resp, err := utils.GetRequest("/core-game/v1/matches/"+matchId+"/loadouts", "glz", a)
	if err != nil {
		return nil, err
	}
	respBody := CurrentGameLoadoutsResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
