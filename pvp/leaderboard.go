package pvp

import (
	"encoding/json"
	"fmt"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
	"net/url"
)

type LeaderboardResponse struct {
	Deployment string `json:"Deployment"`
	QueueID    string `json:"QueueID"`
	SeasonID   string `json:"SeasonID"`
	Players    []struct {
		PlayerCardID    string `json:"PlayerCardID"`
		TitleID         string `json:"TitleID"`
		IsBanned        bool   `json:"IsBanned"`
		IsAnonymized    bool   `json:"IsAnonymized"`
		Puuid           string `json:"puuid"`
		GameName        string `json:"gameName"`
		TagLine         string `json:"tagLine"`
		LeaderboardRank int    `json:"leaderboardRank"`
		RankedRating    int    `json:"rankedRating"`
		NumberOfWins    int    `json:"numberOfWins"`
		CompetitiveTier int    `json:"competitiveTier"`
	} `json:"Players"`
	TotalPlayers          int `json:"totalPlayers"`
	ImmortalStartingPage  int `json:"immortalStartingPage"`
	ImmortalStartingIndex int `json:"immortalStartingIndex"`
	TopTierRRThreshold    int `json:"topTierRRThreshold"`
	TopTierStartingPage   int `json:"topTierStartingPage"`
	TierDetails           map[string]struct {
		RankedRatingThreshold int `json:"rankedRatingThreshold"`
		StartingPage          int `json:"startingPage"`
		StartingIndex         int `json:"startingIndex"`
	} `json:"tierDetails"`
	StartIndex int    `json:"startIndex"`
	Query      string `json:"query"`
}

func GetLeaderboard(a *valorant.Auth, seasonId string, opts ...func(*PvpOptionalParams)) (*LeaderboardResponse, error) {
	params := PvpOptionalParams{}
	for _, opt := range opts {
		opt(&params)
	}

	p := url.Values{}
	if params.startIndex >= 0 {
		p.Add("startIndex", fmt.Sprint(params.startIndex))
	}
	if params.size != 0 {
		p.Add("size", fmt.Sprint(params.size))
	} else {
		p.Add("size", "510")
	}
	if params.queue != "" {
		p.Add("queue", params.queue)
	}
	eP := p.Encode()
	url := "/mmr/v1/leaderboards/affinity/na/queue/competitive/season/" + seasonId
	if eP != "" {
		url += "?" + eP
	}

	fmt.Println(url)

	resp, err := utils.GetRequest(url, "pd", a)
	if err != nil {
		return nil, err
	}

	respBody := LeaderboardResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}
