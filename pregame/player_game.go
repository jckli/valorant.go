package pregame

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type PregamePlayerResponse struct {
	Subject string `json:"Subject"`
	MatchID string `json:"MatchID"`
	Version int    `json:"Version"`
}

func GetPlayerPregame(a *valorant.Auth, puuid string) (*PregamePlayerResponse, error) {
	resp, err := utils.GetRequest("/pregame/v1/players/"+puuid, "glz", a)
	if err != nil {
		return nil, err
	}
	respBody := PregamePlayerResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
