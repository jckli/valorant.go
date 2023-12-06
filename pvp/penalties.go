package pvp

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type PenaltiesResponse struct {
	Subject   string        `json:"Subject"`
	Penalties []interface{} `json:"Penalties"`
	Version   int           `json:"Version"`
}

func GetPenalties(a *valorant.Auth) (*PenaltiesResponse, error) {
	resp, err := utils.GetRequest("/restrictions/v3/penalties", "pd", a)
	if err != nil {
		return nil, err
	}

	respBody := PenaltiesResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
