package pvp

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type PlayerMMRResponse struct {
	Version                     int    `json:"Version"`
	Subject                     string `json:"Subject"`
	NewPlayerExperienceFinished bool   `json:"NewPlayerExperienceFinished"`
	QueueSkills                 map[string]struct {
		TotalGamesNeededForRating         int `json:"TotalGamesNeededForRating"`
		TotalGamesNeededForLeaderboard    int `json:"TotalGamesNeededForLeaderboard"`
		CurrentSeasonGamesNeededForRating int `json:"CurrentSeasonGamesNeededForRating"`
		SeasonalInfoBySeasonID            map[string]struct {
			SeasonID                   string         `json:"SeasonID"`
			NumberOfWins               int            `json:"NumberOfWins"`
			NumberOfWinsWithPlacements int            `json:"NumberOfWinsWithPlacements"`
			NumberOfGames              int            `json:"NumberOfGames"`
			Rank                       int            `json:"Rank"`
			CapstoneWins               int            `json:"CapstoneWins"`
			LeaderboardRank            int            `json:"LeaderboardRank"`
			CompetitiveTier            int            `json:"CompetitiveTier"`
			RankedRating               int            `json:"RankedRating"`
			WinsByTier                 map[string]int `json:"WinsByTier"`
			GamesNeededForRating       int            `json:"GamesNeededForRating"`
			TotalWinsNeededForRank     int            `json:"TotalWinsNeededForRank"`
		} `json:"SeasonalInfoBySeasonID"`
	} `json:"QueueSkills"`
	LatestCompetitiveUpdate struct {
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
	} `json:"LatestCompetitiveUpdate"`
	IsLeaderboardAnonymized bool `json:"IsLeaderboardAnonymized"`
	IsActRankBadgeHidden    bool `json:"IsActRankBadgeHidden"`
}

func GetPlayerMmr(a *valorant.Auth, puuid string) (*PlayerMMRResponse, error) {
	resp, err := utils.GetRequest("/mmr/v1/players/"+puuid, "pd", a)
	if err != nil {
		return nil, err
	}

	respBody := PlayerMMRResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
