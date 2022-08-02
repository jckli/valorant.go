package valorant

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

func fetchGet(endpoint, ep_type string) (*http.Response, error) {
	endpoints := map[string]bool {
		"pd": true,
		"glz": true,
		"shared":true,
	}
	if endpoints[ep_type] {
		client := &http.Client{}
		url := buildUrl(ep_type, userRegion)
		req, err := http.NewRequest(http.MethodGet, url + endpoint, nil)
		if err != nil {
			return nil, err
		}
		req.Header = authHeaders.Clone()
		req.Header.Set("Referer", req.URL.Host)
		return client.Do(req)
	} else {
		return nil, fmt.Errorf("invalid endpoint type")
	}
}

func fetchP(method, endpoint, ep_type string, body interface{}) (*http.Response, error) {
	endpoints := map[string]bool {
		"pd": true,
		"glz": true,
	}
	if endpoints[ep_type] {
		client := &http.Client{}
		url := buildUrl(ep_type, userRegion)
		json_data, _ := json.Marshal(body)
		req, err := http.NewRequest(method, url + endpoint, bytes.NewBuffer(json_data))
		if err != nil {
			return nil, err
		}
		req.Header = authHeaders.Clone()
		req.Header.Set("Referer", req.URL.Host)
		return client.Do(req)
	} else {
		return nil, fmt.Errorf("invalid endpoint type")
	}
}