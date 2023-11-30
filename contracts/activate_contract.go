package contracts

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type ActivateContractResponse struct {
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
				Value               int `json:"Value"`
				BaseMultiplierValue int `json:"BaseMultiplierValue"`
				Modifiers           []struct {
					Value    int    `json:"Value"`
					Name     string `json:"Name"`
					BaseOnly bool   `json:"BaseOnly"`
				} `json:"Modifiers"`
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

// Activate a specific contract by ID
func PostActivateContract(a *valorant.Auth, puuid, contractId string) (*ActivateContractResponse, error) {
	resp, err := utils.PostRequest("/contracts/v1/contracts/"+puuid+"/special/"+contractId, "pd", a)
	if err != nil {
		return nil, err
	}
	respBody := ActivateContractResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
