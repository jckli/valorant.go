package val

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
)

func buildUrl(ep_type, region string) string {
	shard := region
	if (region == "latam") || (region == "br") {
		shard = "na"
	}
	if (region == "") {
		region = "na"
	}
	if ep_type == "pd" {
		return "https://pd." + region +".a.pvp.net"
	} else if ep_type == "glz" {
		return "https://glz-" + region + "-1." + shard + ".a.pvp.net"
	} else if ep_type == "shared" {
		return "https://shared." + shard + ".a.pvp.net"
	} else {
		return ""
	}
}

func FetchGet(endpoint, ep_type string, auth *AuthBody) (*http.Response, error) {
	endpoints := map[string]bool {
		"pd": true,
		"glz": true,
		"shared":true,
	}
	if endpoints[ep_type] {
		client := &http.Client{}
		url := buildUrl(ep_type, auth.Region)
		req, err := http.NewRequest(http.MethodGet, url + endpoint, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Cookie", auth.Cookies)
		req.Header.Set("Authorization", "Bearer " + auth.AccessToken)
		req.Header.Set("X-Riot-Entitlements-JWT", auth.Token)
		req.Header.Set("X-Riot-ClientVersion", auth.Version)
		req.Header.Set("X-Riot-ClientPlatform", "ew0KCSJwbGF0Zm9ybVR5cGUiOiAiUEMiLA0KCSJwbGF0Zm9ybU9TIjogIldpbmRvd3MiLA0KCSJwbGF0Zm9ybU9TVmVyc2lvbiI6ICIxMC4wLjE5MDQyLjEuMjU2LjY0Yml0IiwNCgkicGxhdGZvcm1DaGlwc2V0IjogIlVua25vd24iDQp9")
		req.Header.Set("Referer", req.URL.Host)
		return client.Do(req)
		
	} else {
		return nil, fmt.Errorf("invalid endpoint type")
	}
}

func FetchP(method, endpoint, ep_type string, body interface{}, auth *AuthBody) (*http.Response, error) {
	endpoints := map[string]bool {
		"pd": true,
		"glz": true,
	}
	if endpoints[ep_type] {
		client := &http.Client{}
		url := buildUrl(ep_type, auth.Region)
		json_data, _ := json.Marshal(body)
		req, err := http.NewRequest(method, url + endpoint, bytes.NewBuffer(json_data))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Cookie", auth.Cookies)
		req.Header.Set("Authorization", "Bearer " + auth.AccessToken)
		req.Header.Set("X-Riot-Entitlements-JWT", auth.Token)
		req.Header.Set("X-Riot-ClientVersion", auth.Version)
		req.Header.Set("X-Riot-ClientPlatform", "ew0KCSJwbGF0Zm9ybVR5cGUiOiAiUEMiLA0KCSJwbGF0Zm9ybU9TIjogIldpbmRvd3MiLA0KCSJwbGF0Zm9ybU9TVmVyc2lvbiI6ICIxMC4wLjE5MDQyLjEuMjU2LjY0Yml0IiwNCgkicGxhdGZvcm1DaGlwc2V0IjogIlVua25vd24iDQp9")
		req.Header.Set("Referer", req.URL.Host)
		return client.Do(req)
	} else {
		return nil, fmt.Errorf("invalid endpoint type")
	}
}