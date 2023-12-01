package store

import (
	"encoding/json"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
)

type StorefrontResponse struct {
	FeaturedBundle struct {
		Bundle struct {
			ID             string
			DataAssetID    string
			CurrencyID     string
			Items          []ItemBundle
			BundleDuration int `json:"BundleRemainingDurationInSeconds"`
		}
		Bundles []struct {
			ID          string
			DataAssetID string
			CurrencyID  string
			Items       []ItemBundle
		}
		BundleRemainingDurationInSeconds int
	}
	SkinsPanelLayout struct {
		SingleItemOffers                           []string
		SingleItemStoreOffers                      []SingleItemStoreOffer
		SingleItemOffersRemainingDurationInSeconds int
	}
	UpgradeCurrencyStore struct {
		UpgradeCurrencyOffers []struct {
			OfferID          string
			StorefrontItemID string
			Offer            SingleItemStoreOffer
		}
	}
	BonusStore struct {
		BonusStoreOffers []struct {
			BonusOfferID    string
			Offer           SingleItemStoreOffer
			DiscountPercent float64
			DiscountCosts   map[string]float64
			IsSeen          bool
		}
		BonusStoreRemainingDurationInSeconds int
	}
}

type ItemBundle struct {
	Item struct {
		ItemTypeID string
		ItemID     string
		Quantity   int
	}
	BasePrice       float64
	CurrencyID      string
	DiscountPercent float64
	DiscountedPrice float64
	IsPromoItem     bool
}

type SingleItemStoreOffer struct {
	OfferID          string
	IsDirectPurchase bool
	StartDate        string
	Cost             map[string]float64
	Rewards          []struct {
		ItemTypeID string
		ItemID     string
		Quantity   int
	}
}

func GetStorefront(a *valorant.Auth, puuid string) (*StorefrontResponse, error) {
	resp, err := utils.GetRequest("/store/v2/storefront/"+puuid, "pd", a)
	if err != nil {
		return nil, err
	}
	respBody := StorefrontResponse{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
