package game

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type CurrentGamePlayerResponse struct {
	Subject string `json:"Subject"`
	MatchID string `json:"MatchID"`
	Version int    `json:"Version"`
}

func GetPlayerGame(a *valorant.Auth, puuid string) (*CurrentGamePlayerResponse, error) {
	resp, err := utils.GetRequest("/core-game/v1/players/"+puuid, "glz", a)
	if err != nil {
		return nil, err
	}
	respBody := CurrentGamePlayerResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
