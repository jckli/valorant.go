package pvp

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type AccountXPResponse struct {
	Version  int    `json:"Version"`
	Subject  string `json:"Subject"`
	Progress struct {
		Level int `json:"Level"`
		XP    int `json:"XP"`
	} `json:"Progress"`
	History []struct {
		ID            string `json:"ID"`
		MatchStart    string `json:"MatchStart"`
		StartProgress struct {
			Level int `json:"Level"`
			XP    int `json:"XP"`
		} `json:"StartProgress"`
		EndProgress struct {
			Level int `json:"Level"`
			XP    int `json:"XP"`
		} `json:"EndProgress"`
		XPDelta   int `json:"XPDelta"`
		XPSources []struct {
			ID     string `json:"ID"`
			Amount int    `json:"Amount"`
		} `json:"XPSources"`
		XPMultipliers []interface{} `json:"XPMultipliers"`
	} `json:"History"`
	LastTimeGrantedFirstWin   string `json:"LastTimeGrantedFirstWin"`
	NextTimeFirstWinAvailable string `json:"NextTimeFirstWinAvailable"`
}

func GetAccountXP(a *valorant.Auth, puuid string) (*AccountXPResponse, error) {
	resp, err := utils.GetRequest("/account-xp/v1/players/"+puuid, "pd", a)
	if err != nil {
		return nil, err
	}

	respBody := AccountXPResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
