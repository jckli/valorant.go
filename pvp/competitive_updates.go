package pvp

import (
	"encoding/json"
	"fmt"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
	"net/url"
)

type CompetitiveUpdatesResponse struct {
	Version int    `json:"Version"`
	Subject string `json:"Subject"`
	Matches []struct {
		MatchID                      string `json:"MatchID"`
		MapID                        string `json:"MapID"`
		SeasonID                     string `json:"SeasonID"`
		MatchStartTime               int    `json:"MatchStartTime"`
		TierAfterUpdate              int    `json:"TierAfterUpdate"`
		TierBeforeUpdate             int    `json:"TierBeforeUpdate"`
		RankedRatingAfterUpdate      int    `json:"RankedRatingAfterUpdate"`
		RankedRatingBeforeUpdate     int    `json:"RankedRatingBeforeUpdate"`
		RankedRatingEarned           int    `json:"RankedRatingEarned"`
		RankedRatingPerformanceBonus int    `json:"RankedRatingPerformanceBonus"`
		CompetitiveMovement          string `json:"CompetitiveMovement"`
		AFKPenalty                   int    `json:"AFKPenalty"`
	} `json:"Matches"`
}

func GetCompetitiveUpdates(a *valorant.Auth, puuid string, opts ...func(*PvpOptionalParams)) (*CompetitiveUpdatesResponse, error) {
	params := PvpOptionalParams{}
	for _, opt := range opts {
		opt(&params)
	}

	p := url.Values{}
	if params.startIndex >= 0 {
		p.Add("startIndex", fmt.Sprint(params.startIndex))
	}
	if params.endIndex != 0 {
		p.Add("endIndex", fmt.Sprint(params.endIndex))
	}
	if params.queue != "" {
		p.Add("queue", params.queue)
	}
	eP := p.Encode()
	url := "/mmr/v1/players/" + puuid + "/competitiveupdates"
	if eP != "" {
		url += "?" + eP
	}

	resp, err := utils.GetRequest(url, "pd", a)
	if err != nil {
		return nil, err
	}

	respBody := CompetitiveUpdatesResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}
