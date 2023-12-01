package store

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type ItemTypeEnum int

const (
	ItemTypeAgents = iota
	ItemTypeContracts
	ItemTypeSprays
	ItemTypeBuddies
	ItemTypeCards
	ItemTypeSkins
	ItemTypeSkinVariants
	ItemTypeTitles
)

func WithItemType(itemType ItemTypeEnum) string {
	switch itemType {
	case ItemTypeAgents:
		return "01bb38e1-da47-4e6a-9b3d-945fe4655707"
	case ItemTypeContracts:
		return "f85cb6f7-33e5-4dc8-b609-ec7212301948"
	case ItemTypeSprays:
		return "d5f120f8-ff8c-4aac-92ea-f2b5acbe9475"
	case ItemTypeBuddies:
		return "dd3bf334-87f3-40bd-b043-682a57a8dc3a"
	case ItemTypeCards:
		return "3f296c07-64c3-494c-923b-fe692a4fa1bd"
	case ItemTypeSkins:
		return "e7c63390-eda7-46e0-bb7a-a6abdacd2433"
	case ItemTypeSkinVariants:
		return "3ad1b2b2-acdb-4524-852f-954a76ddae0a"
	case ItemTypeTitles:
		return "de7caa6b-adf7-4588-bbd1-143831e786c6"
	}
	return ""
}

type OwnedItemsResponse struct {
	ItemTypeID   string `json:"ItemTypeID"`
	Entitlements []struct {
		TypeID     string `json:"TypeID"`
		ItemID     string `json:"ItemID"`
		InstanceID string `json:"InstanceID,omitempty"`
	} `json:"Entitlements"`
}

func GetOwnedItems(a *valorant.Auth, puuid string, itemTypeId string) (*OwnedItemsResponse, error) {
	endpoint := "/store/v1/entitlements/" + puuid + "/" + itemTypeId
	resp, err := utils.GetRequest(endpoint, "pd", a)
	if err != nil {
		return nil, err
	}
	respBody := OwnedItemsResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil

}
