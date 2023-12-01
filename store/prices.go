package store

import (
	"encoding/json"
	"fmt"
	"github.com/jckli/valorant.go"
	"github.com/jckli/valorant.go/utils"
	"github.com/valyala/fasthttp"
)

type PricesResponse struct {
	Offers []struct {
		OfferID          string             `json:"OfferID"`
		IsDirectPurchase bool               `json:"IsDirectPurchase"`
		StartDate        string             `json:"StartDate"`
		Cost             map[string]float64 `json:"Cost"`
		Rewards          []struct {
			ItemTypeID string `json:"ItemTypeID"`
			ItemID     string `json:"ItemID"`
			Quantity   int    `json:"Quantity"`
		} `json:"Rewards"`
	} `json:"Offers"`
}

var (
	defaultHeaders = map[string]string{
		"Content-Type":  "application/json",
		"Cache-Control": "no-cache",
		"User-Agent":    valorant.GenerateRandomStringURLSafe(111),
	}
)

func GetPrices(a *valorant.Auth) (*PricesResponse, error) {
	url := utils.BuildUrl("pd", a.Region)
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.Header.SetRequestURI(url + "/store/v1/offers")
	for k, v := range defaultHeaders {
		req.Header.Set(k, v)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))
	req.Header.Set("X-Riot-Entitlements-JWT", a.Token)
	req.Header.SetBytesV("Referer", req.URI().Host())
	resp := fasthttp.AcquireResponse()
	for {
		_ = a.Client.Do(req, resp)
		if resp.StatusCode() == fasthttp.StatusMovedPermanently {
			req.SetRequestURI(url + string(resp.Header.Peek("Location")))
		} else {
			break
		}
	}

	respBody := PricesResponse{}
	err := json.Unmarshal(resp.Body(), &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil

}
