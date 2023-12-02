package pvp

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type PlayerLoadoutResponse struct {
	Subject string `json:"Subject"`
	Version int    `json:"Version"`
	Guns    []struct {
		ID              string        `json:"ID"`
		CharmInstanceID string        `json:"CharmInstanceID"`
		CharmID         string        `json:"CharmID"`
		CharmLevelID    string        `json:"CharmLevelID"`
		SkinID          string        `json:"SkinID"`
		SkinLevelID     string        `json:"SkinLevelID"`
		ChromaID        string        `json:"ChromaID"`
		Attachments     []interface{} `json:"Attachments"`
	} `json:"Guns"`
	Sprays []struct {
		EquipSlotID  string      `json:"EquipSlotID"`
		SprayID      string      `json:"SprayID"`
		SprayLevelID interface{} `json:"SprayLevelID"`
	} `json:"Sprays"`
	Identity struct {
		PlayerCardID           string `json:"PlayerCardID"`
		PlayerTitleID          string `json:"PlayerTitleID"`
		AccountLevel           int    `json:"AccountLevel"`
		PreferredLevelBorderID string `json:"PreferredLevelBorderID"`
		HideAccountLevel       bool   `json:"HideAccountLevel"`
	} `json:"Identity"`
	Incognito bool `json:"Incognito"`
}

// Only works on own puuid, so no puuid parameter.
func GetPlayerLoadout(a *valorant.Auth) (*PlayerLoadoutResponse, error) {
	resp, err := utils.GetRequest("/personalization/v2/players/"+a.UserInfo.UserId+"/playerloadout", "pd", a)
	if err != nil {
		return nil, err
	}

	respBody := PlayerLoadoutResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
