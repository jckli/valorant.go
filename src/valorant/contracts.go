package valorant

import (
	"encoding/json"
)

func Fetch_contract_definitions() (*ContractDefinitionsResp, error) {
	resp, err := fetchGet("/contract-definitions/v3/item-upgrades", "pd")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body := new(ContractDefinitionsResp)
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
}