package valorant

import (
	"encoding/json"
)

// GET ContractDefinitions_Fetch
func Contracts_contractDefiniticions_fetch() (*ContractDefinitionsResp, error) {
	resp, err := fetchGet("/contract-definitions/v3/item-upgrades", "pd")
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
	resp, err := fetchGet(url, "pd")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body := new(ContractFetchResp)
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
}
