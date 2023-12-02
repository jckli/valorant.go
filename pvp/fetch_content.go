package pvp

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type FetchContentResponse struct {
	DisabledIDs []interface{} `json:"DisabledIDs"`
	Seasons     []struct {
		ID        string `json:"ID"`
		Name      string `json:"Name"`
		Type      string `json:"Type"`
		StartTime string `json:"StartTime"`
		EndTime   string `json:"EndTime"`
		IsActive  bool   `json:"IsActive"`
	} `json:"Seasons"`
	Events []struct {
		ID        string `json:"ID"`
		Name      string `json:"Name"`
		StartTime string `json:"StartTime"`
		EndTime   string `json:"EndTime"`
		IsActive  bool   `json:"IsActive"`
	} `json:"Events"`
}

func GetFetchContent(a *valorant.Auth) (*FetchContentResponse, error) {
	resp, err := utils.GetRequest("/content-service/v3/content", "shared", a)
	if err != nil {
		return nil, err
	}

	respBody := FetchContentResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
