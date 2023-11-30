package pregame

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type PregameLoadoutsResponse struct {
	Loadouts []struct {
		Sprays struct {
			SpraySelections interface{} `json:"SpraySelections"`
		} `json:"Sprays"`
		Items interface{} `json:"Items"`
	} `json:"Loadouts"`
	LoadoutsValid bool `json:"LoadoutsValid"`
}

func GetPregameLoadouts(a *valorant.Auth, matchId string) (*PregameLoadoutsResponse, error) {
	resp, err := utils.GetRequest("/pregame/v1/matches/"+matchId+"/loadouts", "glz", a)
	if err != nil {
		return nil, err
	}
	respBody := PregameLoadoutsResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
