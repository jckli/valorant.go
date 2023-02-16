package valorant

import (
	"encoding/json"
	"net/http"
)

// GET ContractDefinitions_Fetch
func Contracts_contractDefiniticions_fetch() (*ContractDefinitionsResp, error) {
	resp, err := FetchGet("/contract-definitions/v3/item-upgrades", "pd")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body := new(ContractDefinitionsResp)
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
}

// GET Contracts_Fetch
func Contracts_fetch(puuid string) (*ContractFetchResp, error) {
	url := "/contracts/v1/contracts/" + puuid
	resp, err := FetchGet(url, "pd")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body := new(ContractFetchResp)
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
}

// POST Contracts_Activate
func Contracts_activate(puuid, contract_id string) (*ContractFetchResp, error) {
	url := "/contracts/v1/contracts/" + puuid + "/special/" + contract_id
	resp, err := FetchP(http.MethodPost, url, "pd", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body := new(ContractFetchResp)
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
}