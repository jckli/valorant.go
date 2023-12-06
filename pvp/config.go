package pvp

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type ConfigResponse struct {
	LastApplication string            `json:"LastApplication"`
	Collapsed       map[string]string `json:"Collapsed"`
	Version         int               `json:"Version"`
}

func GetConfig(a *valorant.Auth) (*ConfigResponse, error) {
	resp, err := utils.GetRequest("/v1/config/"+a.Region, "pd", a)
	if err != nil {
		return nil, err
	}

	respBody := ConfigResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
