package game

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type CurrentGameMatchResponse struct {
	MatchID           string `json:"MatchID"`
	Version           int    `json:"Version"`
	State             string `json:"State"`
	MapID             string `json:"MapID"`
	ModeID            string `json:"ModeID"`
	ProvisioningFlow  string `json:"ProvisioningFlow"`
	GamePodID         string `json:"GamePodID"`
	AllMUCName        string `json:"AllMUCName"`
	TeamMUCName       string `json:"TeamMUCName"`
	TeamVoiceID       string `json:"TeamVoiceID"`
	IsReconnectable   bool   `json:"IsReconnectable"`
	ConnectionDetails struct {
		GameServerHosts        []string `json:"GameServerHosts"`
		GameServerHost         string   `json:"GameServerHost"`
		GameServerPort         int      `json:"GameServerPort"`
		GameServerObfuscatedIP int      `json:"GameServerObfuscatedIP"`
		GameClientHash         int      `json:"GameClientHash"`
		PlayerKey              string   `json:"PlayerKey"`
	} `json:"ConnectionDetails"`
	PostGameDetails interface{} `json:"PostGameDetails"`
	Players         []struct {
		Subject        string `json:"Subject"`
		TeamID         string `json:"TeamID"`
		CharacterID    string `json:"CharacterID"`
		PlayerIdentity struct {
			Subject                string `json:"Subject"`
			PlayerCardID           string `json:"PlayerCardID"`
			PlayerTitleID          string `json:"PlayerTitleID"`
			AccountLevel           int    `json:"AccountLevel"`
			PreferredLevelBorderID string `json:"PreferredLevelBorderID"`
			Incognito              bool   `json:"Incognito"`
			HideAccountLevel       bool   `json:"HideAccountLevel"`
		} `json:"PlayerIdentity"`
		SeasonalBadgeInfo struct {
			SeasonID        string      `json:"SeasonID"`
			NumberOfWins    int         `json:"NumberOfWins"`
			WinsByTier      interface{} `json:"WinsByTier"`
			Rank            int         `json:"Rank"`
			LeaderboardRank int         `json:"LeaderboardRank"`
		} `json:"SeasonalBadgeInfo"`
		IsCoach      bool `json:"IsCoach"`
		IsAssociated bool `json:"IsAssociated"`
	} `json:"Players"`
	MatchmakingData interface{} `json:"MatchmakingData"`
}

// Get the current game match info
func GetGameInfo(a *valorant.Auth, matchId string) (*CurrentGameMatchResponse, error) {
	resp, err := utils.GetRequest("/core-game/v1/matches/"+matchId, "glz", a)
	if err != nil {
		return nil, err
	}

	respBody := CurrentGameMatchResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
