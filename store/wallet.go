package store

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type WalletResponse struct {
	Balances map[string]float64
}

func GetWallet(a *valorant.Auth, puuid string) (*WalletResponse, error) {
	resp, err := utils.GetRequest("/store/v1/wallet/"+puuid, "pd", a)
	if err != nil {
		return nil, err
	}
	respBody := WalletResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
