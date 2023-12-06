package pvp

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type MatchHistoryResponse struct {
	Subject    string `json:"Subject"`
	BeginIndex int    `json:"BeginIndex"`
	EndIndex   int    `json:"EndIndex"`
	Total      int    `json:"Total"`
	History    []struct {
		MatchID       string `json:"MatchID"`
		GameStartTime int    `json:"GameStartTime"`
		QueueID       string `json:"QueueID"`
	} `json:"History"`
}

func GetMatchHistory(a *valorant.Auth, puuid string, opts ...func(*PvpOptionalParams)) (*MatchHistoryResponse, error) {
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
	url := "/match-history/v1/history/" + puuid
	if eP != "" {
		url += "?" + eP
	}

	resp, err := utils.GetRequest(url, "pd", a)
	if err != nil {
		return nil, err
	}

	respBody := MatchHistoryResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}
