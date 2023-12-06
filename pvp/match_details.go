package pvp

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type MatchDetailsResponse struct {
	MatchInfo struct {
		MatchId                     string      `json:"matchId"`
		MapId                       string      `json:"mapId"`
		GamePodId                   string      `json:"gamePodId"`
		GameLoopZone                string      `json:"gameLoopZone"`
		GameServerAddress           string      `json:"gameServerAddress"`
		GameVersion                 string      `json:"gameVersion"`
		GameLengthMillis            interface{} `json:"gameLengthMillis"`
		GameStartMillis             int         `json:"gameStartMillis"`
		ProvisioningFlowID          string      `json:"provisioningFlowID"`
		IsCompleted                 bool        `json:"isCompleted"`
		CustomGameName              string      `json:"customGameName"`
		ForcePostProcessing         bool        `json:"forcePostProcessing"`
		QueueID                     string      `json:"queueID"`
		GameMode                    string      `json:"gameMode"`
		IsRanked                    bool        `json:"isRanked"`
		IsMatchSampled              bool        `json:"isMatchSampled"`
		SeasonId                    string      `json:"seasonId"`
		CompletionState             string      `json:"completionState"`
		PlatformType                string      `json:"platformType"`
		PremierMatchInfo            interface{} `json:"premierMatchInfo"`
		PartyRRPenalties            interface{} `json:"partyRRPenalties"`
		ShouldMatchDisablePenalties bool        `json:"shouldMatchDisablePenalties"`
	} `json:"matchInfo"`
	Players []struct {
		Subject      string `json:"subject"`
		GameName     string `json:"gameName"`
		TagLine      string `json:"tagLine"`
		PlatformInfo struct {
			PlatformType      string `json:"platformType"`
			PlatformOS        string `json:"platformOS"`
			PlatformOSVersion string `json:"platformOSVersion"`
			PlatformChipset   string `json:"platformChipset"`
		} `json:"platformInfo"`
		TeamId                     string      `json:"teamId"`
		PartyId                    string      `json:"partyId"`
		CharacterId                string      `json:"characterId"`
		Stats                      interface{} `json:"stats"`
		RoundDamage                interface{} `json:"roundDamage"`
		CompetitiveTier            int         `json:"competitiveTier"`
		IsObserver                 bool        `json:"isObserver"`
		PlayerCard                 string      `json:"playerCard"`
		PlayerTitle                string      `json:"playerTitle"`
		PreferredLevelBorder       interface{} `json:"preferredLevelBorder"`
		AccountLevel               int         `json:"accountLevel"`
		SessionPlaytimeMinutes     interface{} `json:"sessionPlaytimeMinutes"`
		XpModifications            interface{} `json:"xpModifications"`
		BehaviorFactors            interface{} `json:"behaviorFactors"`
		NewPlayerExperienceDetails interface{} `json:"newPlayerExperienceDetails"`
	} `json:"players"`
	Bots    []interface{} `json:"bots"`
	Coaches []struct {
		Subject string `json:"subject"`
		TeamId  string `json:"teamId"`
	} `json:"coaches"`
	Teams        interface{} `json:"teams"`
	RoundResults interface{} `json:"roundResults"`
	Kills        interface{} `json:"kills"`
}

func GetMatchDetails(a *valorant.Auth, matchID string) (*MatchDetailsResponse, error) {
	resp, err := utils.GetRequest("/match-details/v1/matches/"+matchID, "pd", a)
	if err != nil {
		return nil, err
	}

	respBody := MatchDetailsResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
