package contracts

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type ItemUpgradesResponse struct {
	Definitions []struct {
		ID   string `json:"ID"`
		Item struct {
			ItemTypeID string `json:"ItemTypeID"`
			ItemID     string `json:"ItemID"`
		} `json:"Item"`
		RequiredEntitlement struct {
			ItemTypeID string `json:"ItemTypeID"`
			ItemID     string `json:"ItemID"`
		} `json:"RequiredEntitlement"`
		ProgressionSchedule struct {
			Name                     string `json:"Name"`
			ProgressionCurrencyID    string `json:"ProgressionCurrencyID"`
			ProgressionDeltaPerLevel []int  `json:"ProgressionDeltaPerLevel"`
		} `json:"ProgressionSchedule"`
		RewardSchedule struct {
			ID              string      `json:"ID"`
			Name            string      `json:"Name"`
			Prerequisites   interface{} `json:"Prerequisites"`
			RewardsPerLevel []struct {
				EntitlementRewards []struct {
					Amount     int    `json:"Amount"`
					ItemTypeID string `json:"ItemTypeID"`
					ItemID     string `json:"ItemID"`
				} `json:"EntitlementRewards"`
				WalletRewards  interface{} `json:"WalletRewards"`
				CounterRewards interface{} `json:"CounterRewards"`
			} `json:"RewardsPerLevel"`
		} `json:"RewardSchedule"`
		Sidegrades []struct {
			SidegradeID string `json:"SidegradeID"`
			Options     []struct {
				OptionID string `json:"OptionID"`
				Cost     struct {
					WalletCosts []struct {
						CurrencyID     string `json:"CurrencyID"`
						AmountToDeduct int    `json:"AmountToDeduct"`
					} `json:"WalletCosts"`
				} `json:"Cost"`
				Rewards []struct {
					Amount     int    `json:"Amount"`
					ItemTypeID string `json:"ItemTypeID"`
					ItemID     string `json:"ItemID"`
				} `json:"Rewards"`
			} `json:"Options"`
			Prerequisites struct {
				RequiredEntitlements []struct {
					ItemTypeID string `json:"ItemTypeID"`
					ItemID     string `json:"ItemID"`
				} `json:"RequiredEntitlements"`
			} `json:"Prerequisites"`
		} `json:"Sidegrades"`
	} `json:"Definitions"`
}

func GetItemUpgrades(a *valorant.Auth) (*ItemUpgradesResponse, error) {
	resp, err := utils.GetRequest("/contract-definitions/v3/item-upgrades", "pd", a)
	if err != nil {
		return nil, err
	}

	respBody := ItemUpgradesResponse{}
	if err := json.Unmarshal(resp, &respBody); err != nil {
		return nil, err
	}

	return &respBody, nil
}
