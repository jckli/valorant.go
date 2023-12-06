package pvp

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type NameServiceResponse []struct {
	DisplayName string `json:"DisplayName"`
	Subject     string `json:"Subject"`
	GameName    string `json:"GameName"`
	TagLine     string `json:"TagLine"`
}

type NameServiceBody []string

func PutNameService(a *valorant.Auth, puuids []string) (*NameServiceResponse, error) {
	resp, err := utils.PutBodyRequest("/name-service/v2/players", "pd", a, NameServiceBody(puuids))
	if err != nil {
		return nil, err
	}

	respBody := NameServiceResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
