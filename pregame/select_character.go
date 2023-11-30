package pregame

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type SelectCharacterResponse struct {
	ID      string `json:"ID"`
	Version int    `json:"Version"`
	Teams   []struct {
		TeamID  string `json:"TeamID"`
		Players []struct {
			Subject                 string `json:"Subject"`
			CharacterID             string `json:"CharacterID"`
			CharacterSelectionState string `json:"CharacterSelectionState"`
			PregamePlayerState      string `json:"PregamePlayerState"`
			CompetitiveTier         int    `json:"CompetitiveTier"`
			PlayerIdentity          struct {
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
			IsCaptain bool `json:"IsCaptain"`
		} `json:"Players"`
	} `json:"Teams"`
	AllyTeam             interface{}   `json:"AllyTeam"`
	EnemyTeam            interface{}   `json:"EnemyTeam"`
	ObserverSubjects     []interface{} `json:"ObserverSubjects"`
	MatchCoaches         []interface{} `json:"MatchCoaches"`
	EnemyTeamSize        int           `json:"EnemyTeamSize"`
	EnemyTeamLockCount   int           `json:"EnemyTeamLockCount"`
	PregameState         string        `json:"PregameState"`
	LastUpdated          string        `json:"LastUpdated"`
	MapID                string        `json:"MapID"`
	MapSelectPool        []interface{} `json:"MapSelectPool"`
	BannedMapIDs         []interface{} `json:"BannedMapIDs"`
	CastedVotes          interface{}   `json:"CastedVotes"`
	MapSelectSteps       []interface{} `json:"MapSelectSteps"`
	MapSelectStep        int           `json:"MapSelectStep"`
	Team1                string        `json:"Team1"`
	GamePodID            string        `json:"GamePodID"`
	Mode                 string        `json:"Mode"`
	VoiceSessionID       string        `json:"VoiceSessionID"`
	MUCName              string        `json:"MUCName"`
	QueueID              string        `json:"QueueID"`
	ProvisioningFlow     string        `json:"ProvisioningFlow"`
	IsRanked             bool          `json:"IsRanked"`
	PhaseTimeRemainingNS int           `json:"PhaseTimeRemainingNS"`
	StepTimeRemainingNS  int           `json:"StepTimeRemainingNS"`
	AltModesFlagADA      bool          `json:"altModesFlagADA"`
	TournamentMetadata   interface{}   `json:"TournamentMetadata"`
	RosterMetadata       interface{}   `json:"RosterMetadata"`
}

func SelectCharacter(a *valorant.Auth, matchId string, characterId string) (*SelectCharacterResponse, error) {
	resp, err := utils.PostRequest("/pregame/v1/matches/"+matchId+"/select/"+characterId, "glz", a)
	if err != nil {
		return nil, err
	}
	respBody := SelectCharacterResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
