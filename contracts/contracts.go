package contracts

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type ContractsResponse struct {
	Version   int    `json:"Version"`
	Subject   string `json:"Subject"`
	Contracts []struct {
		ContractDefinitionID string `json:"ContractDefinitionID"`
		ContractProgression  struct {
			TotalProgressionEarned        int `json:"TotalProgressionEarned"`
			TotalProgressionEarnedVersion int `json:"TotalProgressionEarnedVersion"`
			HighestRewardedLevel          map[string]struct {
				Amount  int `json:"Amount"`
				Version int `json:"Version"`
			} `json:"HighestRewardedLevel"`
		} `json:"ContractProgression"`
		ProgressionLevelReached     int `json:"ProgressionLevelReached"`
		ProgressionTowardsNextLevel int `json:"ProgressionTowardsNextLevel"`
	} `json:"Contracts"`
	ProcessedMatches []struct {
		ID        string `json:"ID"`
		StartTime int    `json:"StartTime"`
		XPGrants  struct {
			GamePlayed  int         `json:"GamePlayed"`
			GameWon     int         `json:"GameWon"`
			RoundPlayed int         `json:"RoundPlayed"`
			RoundWon    int         `json:"RoundWon"`
			Missions    interface{} `json:"Missions"`
			Modifier    struct {
				Value    int    `json:"Value"`
				Name     string `json:"Name"`
				BaseOnly bool   `json:"BaseOnly"`
			} `json:"Modifier"`
			NumAFKRounds int `json:"NumAFKRounds"`
		} `json:"XPGrants"`
		RewardGrants          interface{} `json:"RewardGrants"`
		MissionDeltas         interface{} `json:"MissionDeltas"`
		ContractDeltas        interface{} `json:"ContractDeltas"`
		CouldProgressMissions bool        `json:"CouldProgressMissions"`
	} `json:"ProcessedMatches"`
	ActiveSpecialContract string `json:"ActiveSpecialContract"`
	Missions              []struct {
		ID             string         `json:"ID"`
		Objectives     map[string]int `json:"Objectives"`
		Complete       bool           `json:"Complete"`
		ExpirationTime string         `json:"ExpirationTime"`
	} `json:"Missions"`
	MissionMetadata struct {
		NPECompleted     bool   `json:"NPECompleted"`
		WeeklyCheckpoint string `json:"WeeklyCheckpoint"`
		WeeklyRefillTime string `json:"WeeklyRefillTime"`
	} `json:"MissionMetadata"`
}

// Get contract details including agents, battlepass, missions, and recent games
func GetContracts(a *valorant.Auth, puuid string) (*ContractsResponse, error) {
	resp, err := utils.GetRequest("/contracts/v1/contracts/"+puuid, "pd", a)
	if err != nil {
		return nil, err
	}
	respBody := ContractsResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
