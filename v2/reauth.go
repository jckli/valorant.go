package val

import (
	"net/http"
)

func Reauthenticate(auth *AuthBody) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://auth.riotgames.com/authorize?redirect_uri=https%3A%2F%2Fplayvalorant.com%2Fopt_in&client_id=play-valorant-web-prod&response_type=token%20id_token&nonce=1", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", auth.Cookies)
	req.Header.Set("Referer", req.URL.Host)
	return client.Do(req)
}
