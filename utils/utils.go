package utils

import (
	"encoding/json"
	"fmt"
	"github.com/jckli/valorant.go"
	"github.com/valyala/fasthttp"
)

var (
	defaultHeaders = map[string]string{
		"Content-Type":          "application/json",
		"Cache-Control":         "no-cache",
		"User-Agent":            valorant.GenerateRandomStringURLSafe(111),
		"X-Riot-ClientPlatform": "ew0KCSJwbGF0Zm9ybVR5cGUiOiAiUEMiLA0KCSJwbGF0Zm9ybU9TIjogIldpbmRvd3MiLA0KCSJwbGF0Zm9ybU9TVmVyc2lvbiI6ICIxMC4wLjE5MDQyLjEuMjU2LjY0Yml0IiwNCgkicGxhdGZvcm1DaGlwc2V0IjogIlVua25vd24iDQp9",
	}
)

func BuildUrl(ep_type, region string) string {
	if region == "" {
		region = "na"
	}
	shard := region
	if (region == "latam") || (region == "br") {
		shard = "na"
	}
	if ep_type == "pd" {
		return "https://pd." + shard + ".a.pvp.net"
	} else if ep_type == "glz" {
		return "https://glz-" + region + "-1." + shard + ".a.pvp.net"
	} else if ep_type == "shared" {
		return "https://shared." + shard + ".a.pvp.net"
	} else {
		return ""
	}
}

func GetRequest(endpoint, ep_type string, a *valorant.Auth) ([]byte, error) {
	endpoints := map[string]bool{
		"pd":     true,
		"glz":    true,
		"shared": true,
	}
	if endpoints[ep_type] {
		url := BuildUrl(ep_type, a.Region)
		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("GET")
		req.Header.SetRequestURI(url + endpoint)
		for k, v := range defaultHeaders {
			req.Header.Set(k, v)
		}
		req.Header.Set("Cookie", a.CookieJar)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))
		req.Header.Set("X-Riot-Entitlements-JWT", a.Token)
		req.Header.Set("X-Riot-ClientVersion", a.Version)
		req.Header.SetBytesV("Referer", req.URI().Host())
		resp := fasthttp.AcquireResponse()
		err := a.Client.Do(req, resp)
		if err != nil {
			return nil, err
		}
		fmt.Println(resp.StatusCode())
		fmt.Println(resp.Body())

		return resp.Body(), nil
	} else {
		return nil, fmt.Errorf("invalid endpoint type")
	}
}

func PostRequest(endpoint, ep_type string, a *valorant.Auth) ([]byte, error) {
	endpoints := map[string]bool{
		"pd":  true,
		"glz": true,
	}
	if endpoints[ep_type] {
		url := BuildUrl(ep_type, a.Region)
		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("POST")
		req.Header.SetRequestURI(url + endpoint)
		for k, v := range defaultHeaders {
			req.Header.Set(k, v)
		}
		req.Header.Set("Cookie", a.CookieJar)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))
		req.Header.Set("X-Riot-Entitlements-JWT", a.Token)
		req.Header.Set("X-Riot-ClientVersion", a.Version)
		req.Header.SetBytesV("Referer", req.URI().Host())
		resp := fasthttp.AcquireResponse()
		err := a.Client.Do(req, resp)
		if err != nil {
			return nil, err
		}
		return resp.Body(), nil
	} else {
		return nil, fmt.Errorf("invalid endpoint type")
	}
}

func PutRequest(endpoint, ep_type string, a *valorant.Auth) ([]byte, error) {
	endpoints := map[string]bool{
		"pd":  true,
		"glz": true,
	}
	if endpoints[ep_type] {
		url := BuildUrl(ep_type, a.Region)
		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("PUT")
		req.Header.SetRequestURI(url + endpoint)
		for k, v := range defaultHeaders {
			req.Header.Set(k, v)
		}
		req.Header.Set("Cookie", a.CookieJar)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))
		req.Header.Set("X-Riot-Entitlements-JWT", a.Token)
		req.Header.Set("X-Riot-ClientVersion", a.Version)
		req.Header.SetBytesV("Referer", req.URI().Host())
		resp := fasthttp.AcquireResponse()
		err := a.Client.Do(req, resp)
		if err != nil {
			return nil, err
		}
		return resp.Body(), nil
	} else {
		return nil, fmt.Errorf("invalid endpoint type")
	}
}

func PutBodyRequest(endpoint, ep_type string, a *valorant.Auth, body interface{}) ([]byte, error) {
	endpoints := map[string]bool{
		"pd":  true,
		"glz": true,
	}
	if endpoints[ep_type] {
		url := BuildUrl(ep_type, a.Region)
		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("PUT")
		req.Header.SetRequestURI(url + endpoint)
		for k, v := range defaultHeaders {
			req.Header.Set(k, v)
		}
		req.Header.Set("Cookie", a.CookieJar)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))
		req.Header.Set("X-Riot-Entitlements-JWT", a.Token)
		req.Header.Set("X-Riot-ClientVersion", a.Version)
		req.Header.SetBytesV("Referer", req.URI().Host())
		if body != nil {
			bodyBytes, _ := json.Marshal(body)
			req.SetBody(bodyBytes)
		}
		resp := fasthttp.AcquireResponse()
		err := a.Client.Do(req, resp)
		if err != nil {
			return nil, err
		}
		return resp.Body(), nil
	} else {
		return nil, fmt.Errorf("invalid endpoint type")
	}
}
